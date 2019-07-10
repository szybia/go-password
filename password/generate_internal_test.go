package password

import (
	"strings"
	"testing"
)

func TestNewGenerator(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    *CharSet
		expected *Generator
	}{
		// Expect defaults
		{nil, &Generator{
			lowercase: Lowercase,
			uppercase: Uppercase,
			digits:    Digits,
			symbols:   Symbols,
		}},
		// If empty, leave as empty
		{&CharSet{}, &Generator{
			lowercase: "",
			uppercase: "",
			digits:    "",
			symbols:   "",
		}},
		{&CharSet{
			Lowercase: "a",
			Uppercase: "A",
			Digits:    "1",
			Symbols:   "!",
		}, &Generator{
			lowercase: "a",
			uppercase: "A",
			digits:    "1",
			symbols:   "!",
		}},
		// Missing one value in each
		{&CharSet{
			Lowercase: "abcdefghijklmnopqrstuvwxy",
			Uppercase: "ABCDEFGHIJKLMNOPQRSTUVWXY",
			Digits:    "012345678",
			Symbols:   "~!@#$%^&*()_+`-={}|[]\\:\"<>?,.",
		}, &Generator{
			lowercase: "abcdefghijklmnopqrstuvwxy",
			uppercase: "ABCDEFGHIJKLMNOPQRSTUVWXY",
			digits:    "012345678",
			symbols:   "~!@#$%^&*()_+`-={}|[]\\:\"<>?,.",
		}},
	}

	for _, test := range tests {
		actual := NewGenerator(test.input)

		switch {
		case actual.lowercase != test.expected.lowercase:
			t.Errorf("NewGenerator(%q).lowercase: expected %q, got %q",
				actual,
				test.expected.lowercase,
				actual.lowercase)

		case actual.uppercase != test.expected.uppercase:
			t.Errorf("NewGenerator(%q).uppercase: expected %q, got %q",
				actual,
				test.expected.uppercase,
				actual.uppercase)

		case actual.digits != test.expected.digits:
			t.Errorf("NewGenerator(%q).digits: expected %q, got %q",
				actual,
				test.expected.digits,
				actual.digits)

		case actual.symbols != test.expected.symbols:
			t.Errorf("NewGenerator(%q).symbols: expected %q, got %q",
				actual,
				test.expected.symbols,
				actual.symbols)
		}
	}
}

func TestRandomNumberN(t *testing.T) {
	t.Parallel()

	tests := make([]int, N)
	for i := 0; i < N; i++ {
		tests[i] = i + 1
	}

	for _, n := range tests {
		r, err := randomNumberN(n)
		if err != nil {
			t.Errorf("randomNumberN(%v): %v", n, err)
		}
		if !(0 <= r && r < n) {
			t.Errorf("randomNumberN(%v): %v is not in the range of 0-%v",
				n, r, n-1)
		}
	}
}
func TestRandomNumber(t *testing.T) {
	t.Parallel()

	tests := make([]int, N)
	for i := 0; i < N; i++ {
		tests[i] = i
	}

	for _, n := range tests {
		r, err := randomNumber(n)
		if err != nil {
			t.Errorf("randomNumberN(%v): %v", n, err)
		}
		if !(0 <= r && r <= n) {
			t.Errorf("randomNumberN(%v): %v is not in the range of 0-%v",
				n, r, n-1)
		}
	}
}

func TestRandomElement(t *testing.T) {
	t.Parallel()

	tests := []string{
		"1",
		"12",
		"abc",
		"@#$%^&*()!",
	}

	for _, test := range tests {
		b, err := randomElement(test)
		if err != nil {
			t.Error(err)
		}
		if !strings.Contains(test, string(b)) {
			t.Errorf("randomElement(%q): %q is not in string", test, string(b))
		}
	}
}
