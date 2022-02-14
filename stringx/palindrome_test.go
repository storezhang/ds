package stringx

import (
	`testing`
)

func TestPalindrome(t *testing.T) {
	tests := []struct {
		in     string
		expect bool
	}{
		{`tet`, true},
		{`test`, false},
	}

	for _, test := range tests {
		actual := Palindrome(test.in)
		if actual != test.expect {
			t.Errorf(`Palindrome(%v) = %v；期望：%v`, test.in, actual, test.expect)
		}
	}
}
