package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(errors.Wrap(err, "could not open file"))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var twos, threes int
	for scanner.Scan() {
		aggregator := make(map[rune]int)
		text := scanner.Text()
		for _, r := range text {
			aggregator[r]++
		}

		var f, f2 bool
		for k := range aggregator {
			if aggregator[k] == 2 && !f {
				f = true
				twos++
			} else if aggregator[k] == 3 && !f2 {
				f2 = true
				threes++
			}

		}
	}

	if scanner.Err() != nil {
		log.Fatal(errors.Wrap(scanner.Err(), "could not scan file"))
	}

	fmt.Println(twos * threes)

}
