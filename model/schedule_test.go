package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseTimeSlot(t *testing.T) {
	tests := map[string]struct {
		start string
		end   string
		want  []string
	}{
		"simple": {
			start: "0300",
			end:   "0500",
			want:  []string{"0300", "0330", "0400", "0430", "0500"},
		},
		"rollover": {
			start: "1100",
			end:   "1430",
			want:  []string{"1100", "1130", "1200", "1230", "1300", "1330", "1400", "1430"},
		},
		"all": {
			start: "0000",
			end:   "2330",
			want: []string{"0000", "0030", "0100", "0130", "0200", "0230", "0300", "0330", "0400",
				"0430", "0500", "0530", "0600", "0630", "0700", "0730", "0800", "0830", "0900", "0930",
				"1000", "1030", "1100", "1130", "1200", "1230", "1300", "1330", "1400", "1430", "1500",
				"1530", "1600", "1630", "1700", "1730", "1800", "1830", "1900", "1930", "2000", "2030",
				"2100", "2130", "2200", "2230", "2300", "2330"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := parseTimeSlots(tc.start, tc.end)
			if err != nil {
				t.Errorf("could not parse time slots: %v", err)
			}
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Errorf("parse time slot returned diff (-want +got):\n%s", diff)
			}
		})
	}

}
