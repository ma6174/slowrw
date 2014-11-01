package slowrw

import (
	"io"
	"time"
)

type SlowWriter struct {
	Duration time.Duration
	Writer   io.Writer
}

func (s *SlowWriter) Write(data []byte) (n int, err error) {
	for i := 0; i < len(data); i++ {
		nt, err := s.Writer.Write([]byte{data[i]})
		if err != nil {
			return n, err
		}
		n += nt
		time.Sleep(s.Duration)
	}
	return
}

func NewWriter(writer io.Writer, duration time.Duration) *SlowWriter {
	return &SlowWriter{
		Writer:   writer,
		Duration: duration,
	}
}
