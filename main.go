package main

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

ReadLoop:
	for {
		line, err := r.ReadString(0x0A)
		if err == io.EOF {
			break ReadLoop
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(url.QueryEscape(strings.TrimRight(line, "\n")))
	}
}
