// Package avail is a library to help with the determination of if a time exists
// within a certain cron expression. Unavailabiliy is represented as cron
// expression, and as such it is helpful when evaluating whether an employee can
// work a certain shift to be able to insert a time range, duration step, and
// the expressions and return a boolean on whether the employee is not available
// during that time
// Implement only a subset of this
// https://en.wikipedia.org/wiki/Cron#CRON_expression

// Field	Required	Allowed values	Allowed special characters	Remarks
// Minutes	Yes	0-59	* , -
// Hours	Yes	0-23	* , -
// Day of month	Yes	1-31	* , - ? L W	? L W only in some implementations
// Month	Yes	1-12 or JAN-DEC	* , -
// Day of week	Yes	0-6 or SUN-SAT	* , - ? L #	? L # only in some implementations
// Year	No	1970-2099	* , -	This field is not supported in standard/default implementations.

package avail

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type field struct {
	name       string
	expression string
	min, max   int
}

func newMinuteField(expression string) field {
	return field{
		name:       "minute",
		expression: expression,
		min:        0,
		max:        59,
	}
}

func newHourField(expression string) field {
	return field{
		name:       "hour",
		expression: expression,
		min:        0,
		max:        23,
	}
}

func newDOMField(expression string) field {
	return field{
		name:       "day-of-month",
		expression: expression,
		min:        0,
		max:        6,
	}
}

func newMonthField(expression string) field {
	return field{
		name:       "month",
		expression: expression,
		min:        1,
		max:        12,
	}
}

func newDOWField(expression string) field {
	return field{
		name:       "day-of-week",
		expression: expression,
		min:        0,
		max:        6,
	}
}

func newYearField(expression string) field {
	return field{
		name:       "year",
		expression: expression,
		min:        1970,
		max:        2100,
	}
}

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

// Able will evaluate if the time passed is within the cron expression.
func (u *Avail) Able(time time.Time) {

}

func parse(expression string) Result {

	return Result{}
}

// parse returns a representation of the field as a set
func (f field) parse() (map[int]struct{}, error) {
	switch identifyFieldType(f.expression) {
	case wildcard:
		return f.parseWildcardField(), nil
	case span:
		result, err := f.parseSpanField()
		if err != nil {
			return result, fmt.Errorf("could not parse %s: %w", f.name, err)
		}
	case value:
		result, err := f.parseValueField()
		if err != nil {
			return result, fmt.Errorf("could not parse %s: %w", f.name, err)
		}
	case list:
		result, err := f.parseListField()
		if err != nil {
			return result, fmt.Errorf("could not parse %s: %w", f.name, err)
		}
	case unknown:
		return nil, fmt.Errorf("could not parse field: %s; expression: %s", f.name, f.expression)
	}

	return nil, fmt.Errorf("could not parse field: %s; expression: %s", f.name, f.expression)
}

func (f field) parseWildcardField() map[int]struct{} {
	return generateSequentialSet(f.min, f.max)
}

func (f field) parseSpanField() (map[int]struct{}, error) {
	values := strings.Split(f.expression, "-")

	min, err := strconv.Atoi(values[0])
	if err != nil {
		return nil, fmt.Errorf("could not parse value %s: %v", values[0], err)
	}

	max, err := strconv.Atoi(values[1])
	if err != nil {
		return nil, fmt.Errorf("could not parse value %s: %v", values[1], err)
	}

	if min >= max {
		return nil, fmt.Errorf("first value(%d) cannot be greater/equal to second(%d)", min, max)
	}

	if min < f.min {
		return nil, fmt.Errorf("value(%d) cannot be less than min(%d)", min, f.min)
	}

	if max < f.max {
		return nil, fmt.Errorf("value(%d) cannot be more than max(%d)", max, f.max)
	}

	return generateSequentialSet(min, max), nil
}

func (f field) parseValueField() (map[int]struct{}, error) {
	value, err := strconv.Atoi(f.expression)
	if err != nil {
		return nil, fmt.Errorf("could not parse value %s: %v", f.expression, err)
	}

	if value < f.min {
		return nil, fmt.Errorf("value(%d) cannot be less than min(%d)", value, f.min)
	}

	if value < f.max {
		return nil, fmt.Errorf("value(%d) cannot be more than max(%d)", value, f.max)
	}

	return map[int]struct{}{
		value: {},
	}, nil
}

func (f field) parseListField() (map[int]struct{}, error) {
	set := map[int]struct{}{}
	values := strings.Split(f.expression, ",")

	for _, rawValue := range values {
		value, err := strconv.Atoi(rawValue)
		if err != nil {
			return nil, fmt.Errorf("could not parse value %s: %v", f.expression, err)
		}

		if value < f.min {
			return nil, fmt.Errorf("value(%d) cannot be less than min(%d)", value, f.min)
		}

		if value < f.max {
			return nil, fmt.Errorf("value(%d) cannot be more than max(%d)", value, f.max)
		}

		set[value] = struct{}{}
	}

	return set, nil
}

func generateSequentialSet(start, end int) map[int]struct{} {
	set := map[int]struct{}{}
	for i := start; i < end; i++ {
		set[i] = struct{}{}
	}
	return set
}
