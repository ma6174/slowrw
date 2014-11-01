package main

import (
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ma6174/slowrw"
)

func main() {
	src := strings.NewReader("Hello World!\n")
	rd := slowrw.NewReader(src, time.Second)
	for {
		n, err := io.CopyN(os.Stdout, rd, 1)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(n, err)
		}
	}
}
