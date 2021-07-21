package utils

import "testing"

func TestIsPalinDrome(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name             string
		args             args
		wantIsPalinDrome bool
	}{
		{
			name: "when number is palindrome",
			args: args{
				input: 12321,
			},
			wantIsPalinDrome: true,
		},
		{
			name: "when number is not palindrome",
			args: args{
				input: 123,
			},
			wantIsPalinDrome: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsPalinDrome := IsPalinDrome(tt.args.input); gotIsPalinDrome != tt.wantIsPalinDrome {
				t.Errorf("IsPalinDrome() = %v, want %v", gotIsPalinDrome, tt.wantIsPalinDrome)
			}
		})
	}
}
