package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

type claim struct {
	id     int
	point  point
	points []point
	w      int
	h      int
}

type point struct {
	x, y int
}

func (c *claim) generatePoints() {
	for i := 0; i < c.w; i++ {
		for j := 0; j < c.h; j++ {
			c.points = append(c.points, point{x: c.point.x + i, y: c.point.y + j})
		}
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(errors.Wrap(err, "could not open file"))
	}
	defer f.Close()

	var claims []claim
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		c := new(claim)
		// id, from left,from top: wxh
		_, err = fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &c.id, &c.point.x, &c.point.y, &c.w, &c.h)
		if err != nil {
			log.Fatal(errors.Wrap(err, "could not scan line"))
		}
		c.generatePoints()
		claims = append(claims, *c)
	}

	if scanner.Err() != nil {
		log.Fatal(errors.Wrap(scanner.Err(), "could not scan file"))
	}

	occ := make(map[point]int)
	for _, claim := range claims {
		for _, point := range claim.points {
			occ[point]++
		}
	}

	var inches int
	for k := range occ {
		if occ[k] >= 2 {
			inches++
		}
	}
	fmt.Println(inches)

}
