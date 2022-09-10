package log

import (
	"bufio"
	"encoding/binary"
	"os"
	"sync"
)

var (
	enc = binary.BigEndian
)

const (
	lenWidth = 8 // レコードの長さを格納するために使うバイト数
)

type store struct {
	*os.File
	mu   sync.Mutex
	buf  *bufio.Writer
	size uint64
}

func newStore(f *os.File) (*store, error) {
	fi, err := os.Stat(f.Name()) // 既にデータを含むファイルからstoreを再作成する場合に便利
	if err != nil {
		return nil, err
	}

	size := uint64(fi.Size()) // length in bytes for regular files
	return &store{
		File: f,
		size: size,
		buf:  bufio.NewWriter(f),
	}, nil
}

// 書き込まれたbyte数と、書き込み前のsize、errorを返す
// lenWidth分のbyteにレコードの長さの情報、その後にレコードのbyteが保存される
func (s *store) Append(p []byte) (uint64, uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	pos := s.size
	if err := binary.Write(s.buf, enc, uint64(len(p))); err != nil {
		return 0, 0, err
	}

	// Write writes the contents of p into the buffer. It returns the number of bytes written.
	w, err := s.buf.Write(p)
	if err != nil {
		return 0, 0, err
	}

	w += lenWidth
	s.size += uint64(w)
	return uint64(w), pos, nil
}

func (s *store) Read(pos uint64) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.buf.Flush(); err != nil {
		return nil, err
	}

	size := make([]byte, lenWidth)
	// ReadAt reads len(b) bytes from the File starting at byte offset off.
	// In this implementation, `pos` is offset
	if _, err := s.File.ReadAt(size, int64(pos)); err != nil {
		return nil, err
	}

	b := make([]byte, enc.Uint64(size))
	if _, err := s.File.ReadAt(b, int64(pos+lenWidth)); err != nil {
		return nil, err
	}

	return b, nil
}

func (s *store) ReadAt(p []byte, off int64) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.buf.Flush(); err != nil {
		return 0, err
	}

	return s.File.ReadAt(p, off)
}

func (s *store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.buf.Flush(); err != nil {
		return err
	}

	return s.File.Close()
}
