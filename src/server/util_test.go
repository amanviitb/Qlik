package server

import "testing"

func Test_checkIfPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Return true if the given string is of length 0 or 1",
			args: args{s: "A"},
			want: true,
		},
		{
			name: "Return true if the given string is a palindrome of length greater than 1",
			args: args{s: "Amore, Roma"},
			want: true,
		},
		{
			name: "Return true if the given string is not a palindrome",
			args: args{s: "Test"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfPalindrome(tt.args.s); got != tt.want {
				t.Errorf("checkIfPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
