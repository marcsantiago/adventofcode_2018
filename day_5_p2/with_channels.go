package main

// // just for kicks

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"strings"

// 	"github.com/pkg/errors"
// )

// type state string

// const (
// 	opp      = 'A' - 'a'
// 	aplhabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// )

// func main() {
// 	b, err := ioutil.ReadFile("input.txt")
// 	if err != nil {
// 		log.Fatal(errors.Wrap(err, "could not open file"))
// 	}
// 	fData := strings.TrimSpace(string(b))

// 	pairs := createPairs()
// 	ch := make(chan int, len(pairs))
// 	for _, p := range pairs {
// 		go func(p []string, ch chan int) {
// 			tmp := strings.Replace(fData, p[0], "", -1)
// 			ch <- len(react(strings.Replace(tmp, p[1], "", -1)))
// 		}(p, ch)
// 	}

// 	lowests := len(fData)
// 	for i := 0; i < len(pairs); i++ {
// 		n := <-ch
// 		if n < lowests {
// 			lowests = n
// 		}
// 	}

// 	fmt.Println(lowests)
// }

// func react(fData string) string {
// 	for {
// 		var remove []int
// 		for i, r := range fData {
// 			j := i + 1
// 			if j < len(fData) {
// 				if r+opp == rune(fData[j]) || r-opp == rune(fData[j]) {
// 					remove = append(remove, i)
// 					break
// 				}
// 			}
// 		}

// 		if len(remove) == 0 {
// 			break
// 		}

// 		i := remove[0]
// 		fData = fData[:i] + fData[i+2:]
// 	}
// 	return fData
// }

// func createPairs() [][]string {
// 	var pairs [][]string
// 	for _, r := range aplhabet {
// 		pairs = append(pairs, []string{string(r), string(r - opp)})
// 	}
// 	return pairs
// }
