package utils

import "testing"

/*
func TestCountDigits(t *testing.T) {
	input := 123
	expOutput := 3

	digits := CountDigits(input)
	assert.Equal(t, expOutput, digits)
}
*/

func TestCountDigits(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name       string
		args       args
		wantDigits int
	}{
		{
			name: "success",
			args: args{
				input: 123,
			},
			wantDigits: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDigits := CountDigits(tt.args.input); gotDigits != tt.wantDigits {
				t.Errorf("CountDigits() = %v, want %v", gotDigits, tt.wantDigits)
			}
		})
	}
}
