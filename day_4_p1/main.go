package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type state string

// 35169
const (
	begins state = "begins"
	alseep       = "asleep"
	awake        = "awake"
)

const (
	layout = "01-02 15:04" // "[1518-11-01 00:00" // drop the year
)

type event struct {
	id    string
	t     time.Time
	state state
}

func (e event) String() string {
	return fmt.Sprintf("id: %s, month: %02d, day: %02d, time: %02d:%02d, event: %s", e.id, e.t.Month(), e.t.Day(), e.t.Hour(), e.t.Minute(), e.state)
}

type events []event

func (e events) Len() int           { return len(e) }
func (e events) Less(i, j int) bool { return e[i].t.Before(e[j].t) }
func (e events) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

type line struct {
	t   time.Time
	msg string
}

type lines []line

func (l lines) Len() int           { return len(l) }
func (l lines) Less(i, j int) bool { return l[i].t.Before(l[j].t) }
func (l lines) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(errors.Wrap(err, "could not open file"))
	}
	defer f.Close()

	var slines []line
	// sort the file by time first
	// [1518-11-01 00:00] Guard #10 begins shift
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		index := strings.Index(scanner.Text(), "]")
		t, err := time.Parse(layout, scanner.Text()[6:index])
		if err != nil {
			log.Fatal(err)
		}
		slines = append(slines, line{
			t:   t,
			msg: scanner.Text()[index+1:],
		})
	}

	if scanner.Err() != nil {
		log.Fatal(errors.Wrap(scanner.Err(), "could not scan file"))
	}
	sort.Sort(lines(slines))

	var gevents []event
	for _, line := range slines {
		e := event{
			t: line.t,
		}

		parts := strings.Fields(line.msg)
		state := getState(line.msg)
		if state == begins {
			e.id = parts[1][1:]
		} else {
			e.id = gevents[len(gevents)-1].id
		}
		e.state = state
		gevents = append(gevents, e)
	}

	sleep := make(map[string]int)
	var startTime time.Time
	for _, e := range gevents {
		if e.state == begins {
			continue
		}

		if e.state == alseep {
			startTime = e.t
			continue
		}

		if e.state == awake {
			t := e.t.Sub(startTime)
			sleep[e.id] += int(t.Minutes())
		}
	}

	guardID := whoSleepsAlot(sleep)
	hour := initHour()
	for _, e := range gevents {
		if e.id != guardID || e.state == begins {
			continue
		}

		minute := e.t.Minute()
		if e.state == awake {
			minute--
		}

		hour[minute]++
	}

	minute := whichMinute(hour)
	cgid, _ := strconv.Atoi(guardID)
	fmt.Println(cgid, minute)
	fmt.Println(cgid * minute)

}

func getState(s string) state {
	if strings.Contains(s, "falls") {
		return alseep
	} else if strings.Contains(s, "wakes") {
		return awake
	} else if strings.Contains(s, "begins") {
		return begins
	} else {
		panic("state not found")
	}
}

func whoSleepsAlot(data map[string]int) string {
	var id string
	var amount int
	for k := range data {
		if len(id) == 0 {
			id = k
			amount = data[k]
			continue
		}

		if data[k] > amount {
			id = k
			amount = data[k]
		}
	}

	return id
}

func whichMinute(data map[int]int) int {
	var minute int
	var tot int
	for k := range data {
		if data[k] > tot {
			tot = data[k]
			minute = k
		}
	}

	return minute
}

func initHour() map[int]int {
	hour := make(map[int]int)
	for i := 0; i < 60; i++ {
		hour[i] = 0
	}
	return hour
}
