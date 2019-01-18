package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
)

type state string

const opp = 'A' - 'a'

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(errors.Wrap(err, "could not open file"))
	}
	fData := string(b)

	for {
		var remove []int
		for i, r := range fData {
			j := i + 1
			if j < len(fData) {
				if r+opp == rune(fData[j]) || r-opp == rune(fData[j]) {
					remove = append(remove, i)
					break
				}
			}
		}

		if len(remove) == 0 {
			break
		}

		i := remove[0]
		fData = fData[:i] + fData[i+2:]
	}

	fmt.Println(len(fData))
}
