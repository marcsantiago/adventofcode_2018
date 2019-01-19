package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/pkg/errors"
)

type state string

const (
	opp      = 'A' - 'a'
	aplhabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(errors.Wrap(err, "could not open file"))
	}
	fData := strings.TrimSpace(string(b))

	pairs := createPairs()
	lowests := len(fData)
	for _, p := range pairs {
		tmp := strings.Replace(fData, p[0], "", -1)
		tmp = strings.Replace(tmp, p[1], "", -1)

		tmp = react(tmp)
		if len(tmp) < lowests {
			lowests = len(tmp)
		}
	}
	fmt.Println(lowests)
}

func react(fData string) string {
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
	return fData
}

func createPairs() [][]string {
	var pairs [][]string
	for _, r := range aplhabet {
		pairs = append(pairs, []string{string(r), string(r - opp)})
	}
	return pairs
}
