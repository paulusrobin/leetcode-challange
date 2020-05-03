package main

import "fmt"

type schedule struct {
	start int
	end   int
}

func getIntersect(a, b schedule) schedule {
	return schedule{
		start: max(a.start, b.start),
		end:   min(a.end, b.end),
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type MyCalendarTwo struct {
	schedules []schedule
	overlaps  []schedule
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		schedules: make([]schedule, 0),
		overlaps:  make([]schedule, 0),
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	newBooking := schedule{
		start: start,
		end:   end,
	}

	for _, overlapData := range this.overlaps {
		if overlapData.start < newBooking.end && newBooking.start < overlapData.end {
			return false
		}
	}

	for _, scheduleData := range this.schedules {
		if scheduleData.start < newBooking.end && newBooking.start < scheduleData.end {
			this.overlaps = append(this.overlaps, getIntersect(newBooking, scheduleData))
		}
	}

	this.schedules = append(this.schedules, newBooking)
	return true
}

func main() {
	var inputs = [][]int{
		{10,12},
		{50, 60},
		{10, 40},
		{5, 15},
		{5, 10},
		{25, 55},
	}

	MyCalendar := Constructor()
	for _, input := range inputs {
		fmt.Println(MyCalendar.Book(input[0], input[1]))
	}
}
