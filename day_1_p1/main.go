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

	var n int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var k int
		_, err = fmt.Sscanf(scanner.Text(), "%d", &k)
		if err != nil {
			log.Fatal(errors.Wrap(err, "could not scan line"))
		}
		n += k
	}

	if scanner.Err() != nil {
		log.Fatal(errors.Wrap(scanner.Err(), "could not scan file"))
	}

	fmt.Println(n)

}
