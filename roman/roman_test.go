package roman

import (
	"errors"
	"testing"
)

func assertResult(t *testing.T, got, want Roman) {
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	if !errors.Is(got, want) {
		t.Errorf("got %v want error %v", got, want)
	}
}

func TestRoman(t *testing.T) {
	tests := map[string]struct {
		input   Number
		output  Roman
		wantErr error
	}{
		"The number is 1": {
			input:   1,
			output:  "I",
			wantErr: nil,
		},
		"The number is 0": {
			input:   0,
			output:  "",
			wantErr: ErrInvalidNumber,
		},
		"The number is 2": {
			input:   2,
			output:  "II",
			wantErr: nil,
		},
		"The number is 3": {
			input:   3,
			output:  "III",
			wantErr: nil,
		},
		"The number is 4": {
			input:   4,
			output:  "IV",
			wantErr: nil,
		},
		"The number is 5": {
			input:   5,
			output:  "V",
			wantErr: nil,
		}, "The number is 9": {
			input:   9,
			output:  "IX",
			wantErr: nil,
		},
		"The number is 10": {
			input:   10,
			output:  "X",
			wantErr: nil,
		},
		"The number is 14": {
			input:   14,
			output:  "XIV",
			wantErr: nil,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := ConvertToRoman(tc.input)
			assertResult(t, result, tc.output)
			assertError(t, err, tc.wantErr)
		})
	}	
}
