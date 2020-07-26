package avail

import (
	"fmt"
	"testing"
	"time"

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
				Days: Field{
					Kind:   day,
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
				Weekdays: Field{
					Kind:   weekday,
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
	avail, err := New("* * * * * *")
	if err != nil {
		t.Error(err)
	}

	now := time.Now()

	isAvailable := avail.Able(now)
	if isAvailable != true {
		t.Error("expected true; got false")
	}
}

func ExampleAvail_Able() {
	avail, _ := New("* * * * * *")

	now := time.Now()

	fmt.Println(avail.Able(now))
	// Output: true
}
