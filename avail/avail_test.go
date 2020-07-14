package avail

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		expression string
		want       Avail
	}{
		"wildcard": {"* * * * * *", Avail{
			Expression: "* * * * * *",
			Result: Result{
				Minutes: Field{
					Kind:   minute,
					Term:   "*",
					Min:    0,
					Max:    59,
					Values: generateSequentialSet(0, 59),
				},
				Hours: Field{
					Kind:   hour,
					Term:   "*",
					Min:    0,
					Max:    23,
					Values: generateSequentialSet(0, 23),
				},
				Dates: Field{
					Kind:   date,
					Term:   "*",
					Min:    1,
					Max:    31,
					Values: generateSequentialSet(1, 31),
				},
				Months: Field{
					Kind:   month,
					Term:   "*",
					Min:    1,
					Max:    12,
					Values: generateSequentialSet(1, 12),
				},
				Days: Field{
					Kind:   day,
					Term:   "*",
					Min:    0,
					Max:    6,
					Values: generateSequentialSet(0, 6),
				},
				Years: Field{
					Kind:   year,
					Term:   "*",
					Min:    1970,
					Max:    2100,
					Values: generateSequentialSet(1970, 2100),
				},
			},
		}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := New(tc.expression)
			if err != nil {
				t.Error(err)
			}

			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Errorf("result is different than expected(-want +got):\n%s", diff)
			}
		})
	}
}

func TestParseWildcard(t *testing.T) {
	want := Field{
		Kind:   minute,
		Term:   "*",
		Min:    0,
		Max:    59,
		Values: generateSequentialSet(0, 59),
	}
	got, err := newField(minute, "*", 0, 59)
	if err != nil {
		t.Error(err)
	}

	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Errorf("result is different than expected(-want +got):\n%s", diff)
	}
}

func TestParseSpan(t *testing.T) {
	want := Field{
		Kind:   hour,
		Term:   "4-14",
		Min:    0,
		Max:    23,
		Values: generateSequentialSet(4, 14),
	}
	got, err := newField(hour, "4-14", 0, 23)
	if err != nil {
		t.Error(err)
	}

	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Errorf("result is different than expected(-want +got):\n%s", diff)
	}
}

func TestAble(t *testing.T) {

}

// func TestIdentifyTermType(t *testing.T) {
// 	tests := map[string]struct {
// 		input string
// 		want  termType
// 	}{
// 		"span":     {"1-12", span},
// 		"wildcard": {"*", wildcard},
// 		"list":     {"1,2,3,4,5,6", list},
// 		"value":    {"45", value},
// 		"unknown":  {"233)#!", unknown},
// 	}

// 	for name, tc := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			got := identifyTermType(tc.input)
// 			if got != tc.want {
// 				t.Errorf("incorrect field type identified for %s; got %s, want %s", tc.input, got, tc.want)
// 			}
// 		})
// 	}
// }

// func New(expression string) (*Avail, error) {
// 	func (a *Avail) Able(time time.Time) bool {
