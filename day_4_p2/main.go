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

	sleep := make(map[string]map[int]int)
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
			for i := startTime.Minute(); i < e.t.Minute(); i++ {
				if _, ok := sleep[e.id]; !ok {
					sleep[e.id] = make(map[int]int)
				} else {
					sleep[e.id][i]++
				}
			}

		}
	}

	sid, min := whichMinute(sleep)
	gid, _ := strconv.Atoi(sid)
	fmt.Println(gid * min)

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

func whichMinute(data map[string]map[int]int) (string, int) {
	var minute, count int
	var gid string

	for id := range data {
		for min := range data[id] {
			c := data[id][min]
			if c > count {
				count = c
				minute = min
				gid = id
			}
		}
	}

	return gid, minute
}
