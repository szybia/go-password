package password

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const (
	// Lowercase is the default list of lowercase letters.
	Lowercase = "abcdefghijklmnopqrstuvwxyz"

	// Uppercase is the default list of uppercase letters.
	Uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Digits is the default list of permitted digits.
	Digits = "0123456789"

	// Symbols is the default list of symbols.
	Symbols = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

var (
	// ErrNegativeCharNum error is returned when a negative number of character types is requested.
	// E.g. impossible to return password with -1 total lowercase letters.
	ErrNegativeCharNum = errors.New("password with negative number of specific character types is impossible")

	// ErrTypeExceedsAvailable error is returned when a non-zero number of characters
	// of specific type is requested but the character set for that type is empty
	ErrTypeExceedsAvailable = errors.New("non-zero number of character type requested but character set is empty")

	// ErrLengthEmptyGenerator error is returned when a password of specific length is requested
	// but the generator does not have any characters, therefore no password can be generated
	ErrLengthEmptyGenerator = errors.New("password of length requested but generator character set is empty")

	// ErrNegativeLength error is returned when a password of negative length is requested
	ErrNegativeLength = errors.New("password with negative length is not possible")

	// defaultGenerator is used when the helper Generate functions are called.
	// This prevents the generator from being created on each call.
	defaultGenerator = NewGenerator(nil)
)

// CharSet is used for specifying a custom character set which will be used
// for creating a password generator.
type CharSet struct {
	Lowercase string
	Uppercase string
	Digits    string
	Symbols   string
}

// Generator is used for the generating of a cryptographically secure password
// with a initialized character set.
type Generator struct {
	lowercase string
	uppercase string
	digits    string
	symbols   string
	chars     string
}

// NewGenerator creates a new Generator from the specified character set.
// If zero value (nil) is provided then default values are used. Otherwise,
// the values in the CharSet are used.
func NewGenerator(c *CharSet) *Generator {
	if c == nil {
		// Only initialize to default values if
		// user didn't supply character set.
		// This way if user specifically wants
		// no digits in the character set, it's possible.
		// If digits are then requested later an error will
		// be returned rather than quietly using default digits.
		c = &CharSet{
			Lowercase: Lowercase,
			Uppercase: Uppercase,
			Digits:    Digits,
			Symbols:   Symbols,
		}
	}

	g := &Generator{
		lowercase: c.Lowercase,
		uppercase: c.Uppercase,
		digits:    c.Digits,
		symbols:   c.Symbols,
		chars:     c.Lowercase + c.Uppercase + c.Digits + c.Symbols,
	}

	return g
}

// Generate uses the generator character set to produce a cryptographically secure random password.
func (g *Generator) Generate(numLower, numUpper, numDigits, numSymbols int) (s string, err error) {

	switch {
	case numLower < 0 || numUpper < 0 || numDigits < 0 || numSymbols < 0:
		return "", ErrNegativeCharNum

	case numLower > 0 && len(g.lowercase) == 0:
		return "", ErrTypeExceedsAvailable
	case numUpper > 0 && len(g.uppercase) == 0:
		return "", ErrTypeExceedsAvailable
	case numDigits > 0 && len(g.digits) == 0:
		return "", ErrTypeExceedsAvailable
	case numSymbols > 0 && len(g.symbols) == 0:
		return "", ErrTypeExceedsAvailable
	}

	length := numLower + numUpper + numDigits + numSymbols
	if length == 0 {
		return "", nil
	}

	pass := make([]byte, 0, length)

	pass, err = generateCharsetPermute(pass, numLower, g.lowercase)
	if err != nil {
		return
	}

	pass, err = generateCharsetPermute(pass, numUpper, g.uppercase)
	if err != nil {
		return
	}

	pass, err = generateCharsetPermute(pass, numDigits, g.digits)
	if err != nil {
		return
	}

	pass, err = generateCharsetPermute(pass, numSymbols, g.symbols)
	if err != nil {
		return
	}

	return string(pass), nil
}

// GenerateLength generates a password of a specified length
// using all characters within the generators character set.
func (g *Generator) GenerateLength(length int) (string, error) {
	if length < 0 {
		return "", ErrNegativeLength
	} else if length == 0 {
		return "", nil
	} else if len(g.chars) == 0 {
		return "", ErrLengthEmptyGenerator
	}

	pass := make([]byte, 0, length)

	for i := 0; i < length; i++ {
		b, err := randomElement(g.chars)
		if err != nil {
			return "", err
		}
		pass = append(pass, b)
	}

	return string(pass), nil
}

// Generate is the package shortcut for Generator.Generate.
// It uses the default character set to generate a cryptographically secure random password.
func Generate(numLower, numUpper, numDigits, numSymbols int) (string, error) {
	return defaultGenerator.Generate(numLower, numUpper, numDigits, numSymbols)
}

// GenerateLength is the package shortcut for Generator.GenerateLength.
// It uses the default character set to generate a cryptographically secure random password.
func GenerateLength(length int) (string, error) {
	return defaultGenerator.GenerateLength(length)
}

// generateCharsetPermute generates n characters from the charset
// and appends them to the passed byte slice while also permuting it
func generateCharsetPermute(pass []byte, n int, charset string) ([]byte, error) {
	for i := 0; i < n; i++ {
		b, err := randomElement(charset)
		if err != nil {
			return []byte{}, err
		}

		pass, err = randomlySwap(pass, b)
		if err != nil {
			return []byte{}, err
		}
	}
	return pass, nil
}

// randomElement returns a random byte from the string argument
func randomElement(s string) (b byte, err error) {
	i, err := randomNumberN(len(s))
	if err != nil {
		return
	}
	return s[i], nil
}

// randomlySwap applies the Fisher–Yates/Knuth shuffle to append the byte
// to a random location within the byte slice.
func randomlySwap(b []byte, e byte) (r []byte, err error) {
	length := len(b)

	// Decide new element index so be inclusive
	j, err := randomNumber(length)
	if err != nil {
		return
	}

	if j == length {
		r = append(b, e)
	} else {
		r = append(b, b[j])
		r[j] = e
	}
	return
}

// randomNumberN generates a random int n where 0 <= n < max.
func randomNumberN(max int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()), nil
}

// randomNumber generates a random int n where 0 <= n <= max
func randomNumber(max int) (int, error) {
	return randomNumberN(max + 1)
}
