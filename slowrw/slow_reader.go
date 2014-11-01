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

func NewReader(Reader io.Reader, duration time.Duration) *SlowReader {
	return &SlowReader{
		Reader:   Reader,
		Duration: duration,
	}
}
