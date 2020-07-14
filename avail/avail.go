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
	"regexp"
	"strconv"
	"strings"
	"time"
)

// fieldType is an enum which represents different parts of a total cron expression.
// For example in the expression "0 10 15 * * *", 0 would be of type "minute".
type fieldType string

const (
	minute fieldType = "minute"
	hour             = "hour"
	date             = "date"
	month            = "month"
	day              = "day"
	year             = "year"
)

type Field struct {
	Kind fieldType
	// Term is a single field in a complete cron expression.
	// Ex. in the expression: "0 15 10 * * *", "15" would be a term.
	Term     string
	Min, Max int
	Values   map[int]struct{}
}

var cronExpressionRegex = regexp.MustCompile(`^((((\d+,)+\d+|(\d+(-)\d+)|\d+|\*) ?){6})$`)

func newField(kind fieldType, term string, min, max int) (Field, error) {
	newField := Field{
		Kind: kind,
		Term: term,
		Min:  min,
		Max:  max,
	}

	err := newField.parse()
	if err != nil {
		return Field{}, err
	}

	return newField, nil
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
	Minutes Field
	Hours   Field
	Dates   Field
	Months  Field
	Days    Field
	Years   Field
}

// Avail represents both the raw cron expression and the datastructures used to represent that
// expression for easy checking
type Avail struct {
	Expression string // * * * * * * 6 fields min, hours, day of month, month, day of week, year
	Result     Result
}

// New will parse the given cron expression and allow user to check if the time given is within
func New(expression string) (Avail, error) {
	isMatch := cronExpressionRegex.MatchString(expression)
	if !isMatch {
		return Avail{}, fmt.Errorf("could not parse cron expression: %s", expression)
	}

	terms := strings.Split(expression, " ")

	minutes, err := newField(minute, terms[0], 0, 59)
	hours, err := newField(hour, terms[1], 0, 23)
	date, err := newField(date, terms[2], 1, 31)
	month, err := newField(month, terms[3], 1, 12)
	day, err := newField(day, terms[4], 0, 6)
	year, err := newField(year, terms[5], 1970, 2100)
	if err != nil {
		return Avail{}, err
	}

	return Avail{
		Expression: expression,
		Result: Result{
			Minutes: minutes,
			Hours:   hours,
			Dates:   date,
			Months:  month,
			Days:    day,
			Years:   year,
		},
	}, nil
}

// Able will evaluate if the time given is within the cron expression.
func (a *Avail) Able(time time.Time) bool {
	fieldTypes := []fieldType{
		minute,
		hour,
		date,
		month,
		day,
		year,
	}

	for _, field := range fieldTypes {
		switch field {
		case minute:
			if _, ok := a.Result.Minutes.Values[time.Minute()]; !ok {
				return false
			}
		case hour:
			if _, ok := a.Result.Hours.Values[time.Hour()]; !ok {
				return false
			}
		case date:
			if _, ok := a.Result.Dates.Values[time.Day()]; !ok {
				return false
			}
		case month:
			if _, ok := a.Result.Months.Values[int(time.Month())]; !ok {
				return false
			}
		case day:
			if _, ok := a.Result.Days.Values[int(time.Weekday())]; !ok {
				return false
			}
		case year:
			if _, ok := a.Result.Years.Values[time.Year()]; !ok {
				return false
			}
		}
	}

	return true
}

// parse returns a representation of the field as a set
func (f *Field) parse() error {
	switch identifyTermType(f.Term) {
	case wildcard:
		f.Values = f.parseWildcardField()
		return nil
	case span:
		result, err := f.parseSpanField()
		if err != nil {
			return fmt.Errorf("could not parse %s: %w", f.Kind, err)
		}
		f.Values = result
		return nil
	case value:
		result, err := f.parseValueField()
		if err != nil {
			return fmt.Errorf("could not parse %s: %w", f.Kind, err)
		}
		f.Values = result
		return nil
	case list:
		result, err := f.parseListField()
		if err != nil {
			return fmt.Errorf("could not parse %s: %w", f.Kind, err)
		}
		f.Values = result
		return nil
	case unknown:
		return fmt.Errorf("could not parse field: %s; expression: %s", f.Kind, f.Term)
	}

	return fmt.Errorf("could not parse field: %s; expression: %s", f.Kind, f.Term)
}

func (f *Field) parseWildcardField() map[int]struct{} {
	return generateSequentialSet(f.Min, f.Max)
}

// TODO(clintjedwards): We should break field into one more struct
// that defines descriptors for individual field types
func (f *Field) parseSpanField() (map[int]struct{}, error) {
	values := strings.Split(f.Term, "-")

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

	if min < f.Min {
		return nil, fmt.Errorf("value(%d) cannot be less than min(%d)", min, f.Min)
	}

	if max > f.Max {
		return nil, fmt.Errorf("value(%d) cannot be more than max(%d)", max, f.Max)
	}

	return generateSequentialSet(min, max), nil
}

func (f *Field) parseValueField() (map[int]struct{}, error) {
	value, err := strconv.Atoi(f.Term)
	if err != nil {
		return nil, fmt.Errorf("could not parse value %s: %v", f.Term, err)
	}

	if value < f.Min {
		return nil, fmt.Errorf("value(%d) cannot be less than min(%d)", value, f.Min)
	}

	if value < f.Max {
		return nil, fmt.Errorf("value(%d) cannot be more than max(%d)", value, f.Max)
	}

	return map[int]struct{}{
		value: {},
	}, nil
}

func (f *Field) parseListField() (map[int]struct{}, error) {
	set := map[int]struct{}{}
	values := strings.Split(f.Term, ",")

	for _, rawValue := range values {
		value, err := strconv.Atoi(rawValue)
		if err != nil {
			return nil, fmt.Errorf("could not parse value %s: %v", f.Term, err)
		}

		if value < f.Min {
			return nil, fmt.Errorf("value(%d) cannot be less than min(%d)", value, f.Min)
		}

		if value < f.Max {
			return nil, fmt.Errorf("value(%d) cannot be more than max(%d)", value, f.Max)
		}

		set[value] = struct{}{}
	}

	return set, nil
}

func generateSequentialSet(start, end int) map[int]struct{} {
	set := map[int]struct{}{}
	for i := start; i < end+1; i++ {
		set[i] = struct{}{}
	}
	return set
}
