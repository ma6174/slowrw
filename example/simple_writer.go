package main

import (
	"log"
	"os"
	"time"

	"github.com/ma6174/slowrw"
)

func main() {
	sw := slowrw.NewWriter(os.Stdout, time.Second)
	n, err := sw.Write([]byte("Hello World!\n"))
	if err != nil {
		log.Fatal(n, err)
	}
}
