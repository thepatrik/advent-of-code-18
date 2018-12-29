package reposerecord

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type action int

const (
	beginShift action = iota
	fallAsleep
	wakeUp
)

func getMap(strslice []string) map[int]map[int]int {
	sort.Strings(strslice)
	guard, prevkey := -1, -1
	m := make(map[int]map[int]int) //map[guardID][minute(0-1400)]count

	for _, s := range strslice {
		id, action, hour, minute := parseEntry(s, guard)
		guard = id
		key := hour*60 + minute

		switch action {
		case wakeUp:
			if _, ok := m[guard]; !ok {
				m[guard] = make(map[int]int)
			}
			for i := prevkey; i < key; i++ {
				m[guard][i]++
			}
		}
		prevkey = key
	}
	return m
}

func parseEntry(s string, guard int) (int, action, int, int) {
	i := strings.Index(s, "] ")
	txt := s[i+2:]
	var id int
	n, _ := fmt.Sscanf(txt, "Guard #%d begins shift", &id)
	if n == 0 {
		// New guard present
		id = guard
	}

	action := beginShift
	switch txt {
	case "falls asleep":
		action = fallAsleep
	case "wakes up":
		action = wakeUp
	}
	date := date(s)
	return id, action, date.Hour(), date.Minute()
}

func date(s string) time.Time {
	var y, m, d, hr, mn int
	_, _ = fmt.Sscanf(s, "[%d-%d-%d %d:%d]", &y, &m, &d, &hr, &mn)
	return time.Date(y, time.Month(m), d, hr, mn, 0, 0, time.UTC)
}
