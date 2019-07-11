/*
Copyright © 2019 Szymon Bialkowski

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"github.com/szybia/go-password/password"
)

var noSymbols bool
var clip bool
var length int

var rootCmd = &cobra.Command{
	Use:   "randpw",
	Short: "Generate random passwords.",
	Long:  "randpw is a command-line application which generates cryptographically-secure passwords.",

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		g := password.NewGenerator(nil)

		if noSymbols {
			g = password.NewGenerator(&password.CharSet{
				Lowercase: password.Lowercase,
				Uppercase: password.Uppercase,
				Digits:    password.Digits,
			})
		}

		pass, err := g.GenerateLength(length)
		if err != nil {
			return
		}

		if clip {
			err = clipboard.WriteAll(pass)
		} else {
			fmt.Print(pass)
		}

		return
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&noSymbols, "no-symbols", "n", false, "exclude symbols from password")
	rootCmd.PersistentFlags().BoolVarP(&clip, "clip", "c", false, "copy password to clipboard")
	rootCmd.PersistentFlags().IntVarP(&length, "length", "l", 25, "specify the lenght of the password")
}
