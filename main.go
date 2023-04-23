// name    :	main.go
// version :	0.0.1
// date    :	20230423
// author  :	Leam Hall
// desc    :    Playing with the wc command, in Go

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func count(r io.Reader) (wl, ww, wc int) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		wl++
		s := scanner.Text()
		sf := []string(strings.Fields(s))
		ww += len(sf)
		for i := 0; i < len(sf); i++ {
			word := strings.Trim(sf[i], " \n")
			wc += len(word)
		}
	}
	return wl, ww, wc
}

func main() {
	wl, ww, wc := count(os.Stdin)
	fmt.Printf("%d lines, %d words, %d characters\n", wl, ww, wc)
}
