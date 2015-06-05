// Copyright (c) 2015 John Graham-Cumming

package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"io"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	h := fnv.New64()
	for s.Scan() {
		h.Reset()
		t := s.Text()
		io.WriteString(h, t)
		fmt.Printf("%x %s\n", h.Sum(nil), t)
	}

	if err := s.Err(); err != nil {
		fmt.Printf("Error reading standard input: %s\n", err)
	}
}
