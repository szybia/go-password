package password

import (
	"strings"
	"testing"
)

const N = 10

func TestGenerator_Generate(t *testing.T) {
	t.Parallel()

	g := NewGenerator(nil)

	t.Run("no_characters_empty_base_case", func(t *testing.T) {
		t.Parallel()

		r, err := g.Generate(0, 0, 0, 0)
		if err != nil {
			t.Error(err)
		}
		//	If no characters requested then empty string is expected
		if r != "" {
			t.Errorf("Generator.Generate(%v, %v, %v, %v): expected %q, got %q", 0, 0, 0, 0, "", r)
		}
	})

	t.Run("negative_number_of_characters", func(t *testing.T) {
		t.Parallel()

		inputs := []int{0, 0, 0, 0}

		for i := range inputs {
			inputs[i]--

			_, err := g.Generate(inputs[0], inputs[1], inputs[2], inputs[3])
			if err != ErrNegativeCharNum {
				t.Errorf("Generator.Generate(%v, %v, %v, %v): expected %v error, got %T",
					inputs[0], inputs[1], inputs[2], inputs[3], "ErrNegativeCharNum", err)
			}

			inputs[i]++
		}
	})

	t.Run("character_counts_match", func(t *testing.T) {
		for numLower := 0; numLower < N; numLower++ {
			for numUpper := 0; numUpper < N; numUpper++ {
				for numDigits := 0; numDigits < N; numDigits++ {
					for numSymbols := 0; numSymbols < N; numSymbols++ {

						pass, err := Generate(numLower, numUpper, numDigits, numSymbols)
						if err != nil {
							t.Error(err)
						}

						var countLower, countUpper, countDigits, countSymbols int

						for _, char := range pass {
							switch {
							case strings.ContainsRune(Lowercase, char):
								countLower++
							case strings.ContainsRune(Uppercase, char):
								countUpper++
							case strings.ContainsRune(Digits, char):
								countDigits++
							case strings.ContainsRune(Symbols, char):
								countSymbols++
							}
						}

						if countLower != numLower {
							t.Errorf("Generate(%v, %v, %v, %v): expected %v lowercase letters, got %v",
								numLower, numUpper, numDigits, numSymbols, numLower, countLower)
						}
						if countUpper != numUpper {
							t.Errorf("Generate(%v, %v, %v, %v): expected %v uppercase letters, got %v",
								numLower, numUpper, numDigits, numSymbols, numUpper, countUpper)
						}
						if countDigits != numDigits {
							t.Errorf("Generate(%v, %v, %v, %v): expected %v digits, got %v",
								numLower, numUpper, numDigits, numSymbols, numDigits, countDigits)
						}
						if countSymbols != numSymbols {
							t.Errorf("Generate(%v, %v, %v, %v): expected %v symbols, got %v",
								numLower, numUpper, numDigits, numSymbols, numSymbols, countSymbols)
						}

					}
				}
			}
		}
	})
}

func TestGenerator_GenerateLength(t *testing.T) {
	t.Parallel()

	t.Run("negative_length_err", func(t *testing.T) {
		t.Parallel()

		for _, i := range []int{-2, -1} {
			_, err := GenerateLength(i)
			if err == nil {
				t.Errorf("GenerateLength(%v): expected negative length error, got none", i)
			}
		}
	})

	t.Run("password_length_matches", func(t *testing.T) {
		t.Parallel()

		for i := 0; i < N; i++ {
			p, err := GenerateLength(i)
			if err != nil {
				t.Error(err)
			}

			if len(p) != i {
				t.Errorf("GenerateLength(%v): expected password of length %v, got %v", i, i, len(p))
			}
		}
	})

}
