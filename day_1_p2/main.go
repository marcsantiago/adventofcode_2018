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

	var arr []int
	seen := make(map[int]struct{})
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var k int
		_, err = fmt.Sscanf(scanner.Text(), "%d", &k)
		if err != nil {
			log.Fatal(errors.Wrap(err, "could not scan line"))
		}
		arr = append(arr, k)
	}

	if scanner.Err() != nil {
		log.Fatal(errors.Wrap(scanner.Err(), "could not scan file"))
	}

	var ans int
	var total int

	found := false
	for !found {
		for _, n := range arr {
			total += n
			_, ok := seen[total]
			if ok {
				ans = total
				found = true
				break
			}
			seen[total] = struct{}{}
		}
	}

	fmt.Println(ans)

}
