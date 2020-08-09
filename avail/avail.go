package avail

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// fieldType is an enum which represents different parts of a total cron expression.
// For example in the expression "0 10 15 * * *", 0 would be of type "minute".
type fieldType string

const (
	minute  fieldType = "minute"
	hour              = "hour"
	day               = "day"
	month             = "month"
	weekday           = "weekday"
	year              = "year"
)

var cronExpressionRegex = regexp.MustCompile(`^((((\d+,)+\d+|(\d+(-)\d+)|\d+|\*) ?){6})$`)

// Result represents a breakdown of a given cron time expression
type Result struct {
	Minutes  Field
	Hours    Field
	Days     Field
	Months   Field
	Weekdays Field
	Years    Field
}

// Avail represents both the raw cron expression and the datastructures used to represent that
// expression for easy checking
type Avail struct {
	Expression string // * * * * * * 6 fields - min, hours, day of month, month, day of week, year
	Result     Result
}

// New will parse the given cron expression and allow user to check if the time given is within
func New(expression string) (Avail, error) {
	isMatch := cronExpressionRegex.MatchString(expression)
	if !isMatch {
		return Avail{}, fmt.Errorf("could not parse cron expression: %s", expression)
	}

	terms := strings.Split(expression, " ")
	// we need this extra check to make sure there are the proper amount of fields because I am bad at regex
	if len(terms) != 6 {
		return Avail{}, fmt.Errorf("could not parse cron expression: %s; must have 6 terms", expression)
	}

	minutes, err := newField(minute, terms[0], 0, 59)
	hours, err := newField(hour, terms[1], 0, 23)
	day, err := newField(day, terms[2], 1, 31)
	month, err := newField(month, terms[3], 1, 12)
	weekday, err := newField(weekday, terms[4], 0, 6)
	year, err := newField(year, terms[5], 1970, 2100)
	if err != nil {
		return Avail{}, err
	}

	return Avail{
		Expression: expression,
		Result: Result{
			Minutes:  minutes,
			Hours:    hours,
			Days:     day,
			Months:   month,
			Weekdays: weekday,
			Years:    year,
		},
	}, nil
}

// Able will evaluate if the time given is within the cron expression.
func (a *Avail) Able(time time.Time) bool {
	fieldTypes := []fieldType{
		minute,
		hour,
		day,
		month,
		weekday,
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
		case day:
			if _, ok := a.Result.Days.Values[time.Day()]; !ok {
				return false
			}
		case month:
			if _, ok := a.Result.Months.Values[int(time.Month())]; !ok {
				return false
			}
		case weekday:
			if _, ok := a.Result.Weekdays.Values[int(time.Weekday())]; !ok {
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

func generateSequentialSet(start, end int) map[int]struct{} {
	set := map[int]struct{}{}
	for i := start; i < end+1; i++ {
		set[i] = struct{}{}
	}
	return set
}
