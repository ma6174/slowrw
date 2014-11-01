package slowrw

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"
	"time"
)

func TestReader(t *testing.T) {
	rawData := "Hello"
	src := strings.NewReader(rawData)
	sr := NewReader(src, time.Second)
	now := time.Now()
	data, err := ioutil.ReadAll(sr)
	timeUse := time.Since(now).Seconds()
	if err != nil {
		t.Fatal("ioutil.ReadAll Failed:", err)
	}
	if string(data) != rawData {
		t.Fatal("data not equal", string(data))
	}
	if timeUse < float64(len(rawData)) {
		t.Fatal("used time too short:", timeUse)
	}
}

func TestWriter(t *testing.T) {
	rawData := "Hello"
	pr, pw := io.Pipe()
	sw := NewWriter(pw, time.Second)
	go func() {
		data, err := ioutil.ReadAll(pr)
		if err != nil {
			t.Fatal("ioutil.ReadAll failed:", err)
		}
		if string(data) != rawData {
			t.Fatal("data not equal", string(data))
		}
	}()
	now := time.Now()
	n, err := sw.Write([]byte(rawData))
	timeUse := time.Since(now).Seconds()
	if err != nil {
		t.Fatal("sw.Write failed:", err)
	}
	if n != len(rawData) {
		t.Fatal("writed length not equal:", n)
	}
	if timeUse < float64(len(rawData)) {
		t.Fatal("used time too short:", timeUse)
	}
}
