package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// filterUnavailableEmployees
func TestFilterUnavailableEmployees(t *testing.T) {

	dateTime, err := time.Parse("01-02-2006", "06-19-1990")
	if err != nil {
		t.Error(err)
	}

	tests := map[string]struct {
		date time.Time
		pool map[string]*Employee
		want []string
	}{
		"all unavailable": {
			date: dateTime,
			pool: map[string]*Employee{
				"1": {ID: "1", Unavailabilities: []string{"* * * * * *"}},
				"2": {ID: "2", Unavailabilities: []string{"* * * * * *"}},
			},
			want: []string{},
		},
		"all available": {
			date: dateTime,
			pool: map[string]*Employee{
				// empty unavailablility list should count as available for any day
				"1": {ID: "1", Unavailabilities: []string{}},
				"2": {ID: "2", Unavailabilities: []string{"* * * * * 2000"}},
			},
			want: []string{"1", "2"},
		},
		// "all employees eligible": {},
		// "no employees eligible":  {},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			resultPool := filterUnavailableEmployees(tc.date, tc.pool)
			got := []string{}
			for id := range resultPool {
				got = append(got, id)
			}
			require.ElementsMatch(t, tc.want, got)
		})
	}
}
