package model

import (
	"testing"
)

func TestEmployeeValidationError(t *testing.T) {
	employee := Employee{}
	err := employee.IsValid()
	if err == nil {
		t.Error("expected error got nil")
	}
}

func TestEmployeeValidation(t *testing.T) {
	tests := map[string]struct {
		employee Employee
		want     error
	}{
		"valid": {
			employee: Employee{
				Name: "someName",
			},
			want: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.employee.IsValid()

			if got != tc.want {
				t.Errorf("struct invalid, got %t; want %t", got, tc.want)
			}
		})
	}
}
