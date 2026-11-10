/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

 import "slices"

func canAttendMeetings(intervals []Interval) bool {
	//order by start then end
	slices.SortFunc(intervals, cmpInterval)

	for i, interval := range intervals {
		if i == len(intervals)-1 {
			return true
		}
		if intersect(interval, intervals[i+1]) {
			return false
		}
	}

	return true
}

func intersect(a, b Interval) bool {
	return a.end > b.end || a.start >= b.start || (a.start < b.start && a.end > b.start)
}

func cmpInterval(a, b Interval) int {
	if a.start > b.start {
		return 1
	}
	if a.start < b.start {
		return -1
	}

	if a.end < b.end {
		return -1
	}
	return 1
}

