// Package unavail is a library to help with the determination of if a time exists
// within a certain cron expression. Unavailabiliy is represented as cron
// expression, and as such it is helpful when evaluating whether an employee can
// work a certain shift to be able to insert a time range, duration step, and
// the expressions and return a boolean on whether the employee is not available
// during that time
// Implement only a subset of this
// https://en.wikipedia.org/wiki/Cron#CRON_expression
package avail

import "time"

// api should look like:
// someTime := time.Now()
// unavail := unavail.New("* * * * 8,9 2020")
// isAble := unavail.Able(sometime)

// instead of taking everything in at once should we instead have a new function
// which takes the expected cron expression and then allow the user to call the able function
// many times. This way we can drop the duration step for things

// Result represents a breakdown of a specific cron time expression
// Most fields have sets so that it's easy to check if a specific time frame
// is within the set. These sets are made with structs because empty structs are 0 bytes.
// https://dave.cheney.net/2014/03/25/the-empty-struct
type Result struct {
	seconds map[int]struct{}
}

// Avail represents both the raw cron expression and the datastructures used to represent that
// expression for easy checking
type Avail struct {
	Expression string // * * * * * * 6 fields min, hours, day of month, month, day of week, year
	Result     Result
}

// New will parse the given cron expression and allow user to check if the time given is within
func New(expression string) *Avail {
	return &Avail{}
}

// Able will evaluate if the time passed is within the cron expression. If it is NOT it will return
// true.
func (u *Avail) Able(time time.Time) {

}

func parse(expression string) Result {

	return Result{}
}

// *, 0-12, 1,2,3,4,5,6, 6

// minutesFieldHandler returns a representation of the minutes field as a data structure
func minutesFieldHandler(field string) map[int]struct{} {

	return map[int]struct{}{}
}
