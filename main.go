package main

import (
	"bufio"
	"errors"
	"flag"
	"os"
	"slices"
)

func main() {

	var revFlag bool
	flag.BoolVar(&revFlag, "reverse", false, "Reverse each line on output")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadBytes('\n')
	for ; err == nil; line, err = reader.ReadBytes('\n') {
		if revFlag == true {
			withoutN, _ := removeTrailingByte(line)
			slices.Reverse(withoutN)
			os.Stdout.Write(append(withoutN, '\n'))
		} else {
			os.Stdout.Write(line)
		}

	}
}

func removeTrailingByte(b []byte) ([]byte, error) {
	if len(b) == 0 {
		return nil, errors.New("Cannot remove trailing byte from empty slice")
	}
	return b[:len(b)-1], nil
}
