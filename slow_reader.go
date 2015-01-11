package slowrw

import (
	"io"
	"time"
)

type SlowReader struct {
	Duration time.Duration
	Reader   io.Reader
}

func (s *SlowReader) Read(p []byte) (n int, err error) {
	t := make([]byte, 1)
	for i := 0; i < len(p); i++ {
		tn, err := s.Reader.Read(t)
		if err != nil {
			return n, err
		}
		p[i] = t[0]
		n += tn
		time.Sleep(s.Duration)
	}
	return
}

type SlowReadSeeker struct {
	*SlowReader
	Reader io.ReadSeeker
}

func (s *SlowReadSeeker) Seek(offset int64, whence int) (int64, error) {
	return s.Reader.Seek(offset, whence)
}

type SlwoReaderAt struct {
	Reader   io.ReaderAt
	Duration time.Duration
}

func (s *SlwoReaderAt) ReadAt(p []byte, off int64) (n int, err error) {
	t := make([]byte, 1)
	for i := 0; i < len(p); i++ {
		tn, err := s.Reader.ReadAt(t, off)
		if err != nil {
			return n, err
		}
		p[i] = t[0]
		n += tn
		time.Sleep(s.Duration)
	}
	return
}

func NewReader(reader io.Reader, duration time.Duration) *SlowReader {
	return &SlowReader{
		Reader:   reader,
		Duration: duration,
	}
}
func NewReadSeeker(readSeeker io.ReadSeeker, duration time.Duration) *SlowReadSeeker {
	return &SlowReadSeeker{
		SlowReader: NewReader(readSeeker, duration),
		Reader:     readSeeker,
	}
}

func NewReaderAt(readerAt io.ReaderAt, duration time.Duration) *SlwoReaderAt {
	return &SlwoReaderAt{
		Reader:   readerAt,
		Duration: duration,
	}
}
