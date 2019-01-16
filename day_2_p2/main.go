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

	var arr []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

loop:
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if diffence(arr[i], arr[j]) == 1 {
				fmt.Println(whoDis(arr[i], arr[j]))
				break loop
			}
		}
	}

	if scanner.Err() != nil {
		log.Fatal(errors.Wrap(scanner.Err(), "could not scan file"))
	}

}

func diffence(a, b string) int {
	var dif int
	for i := range a {
		if a[i] != b[i] {
			dif++
		}
	}
	return dif
}

func whoDis(a, b string) string {
	var s string
	for i := range a {
		if a[i] == b[i] {
			s += string(a[i])
		}
	}
	return s
}
