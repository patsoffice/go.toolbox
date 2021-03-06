// Copyright © 2020 Patrick Lawrence <patrick.lawrence@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package toolbox

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type scanner interface {
	Scan() bool
	Text() string
}

// GetInputInt takes a message string to present to the user and a default
// value. It scans input for text and attempts to convert the input to an
// integer. If the conversion to an integer fails, the user will be prompted
// to input an integer again. If the input value is empty, the default value
// is returned. Otherwise, the converted integer is returned.
func GetInputInt(scnr scanner, msg string, def int) int {
	for {
		fmt.Printf("%s [%d]: ", msg, def)

		for scnr.Scan() {
			input := scnr.Text()
			if input == "" {
				return def
			}
			v, err := strconv.Atoi(input)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid input: %v\n", err)
			} else {
				return v
			}
		}
	}
}

// GetInputString takes a message string to present to the user and a default
// value. It scans input for text. If the input value is empty, the default
// value is returned. Otherwise, the input string is returned.
func GetInputString(scnr scanner, msg, def string) string {
	if def == "" {
		fmt.Printf("%s: ", msg)
	} else {
		fmt.Printf("%s [%s]: ", msg, def)
	}
	for scnr.Scan() {
		input := scnr.Text()
		if input == "" {
			return def
		}
		return input
	}
	return ""
}

// CheckYes takes a message string to present to the user and a boolen value
// representing the default value. It scans input for text. If the input value
// is empty, the default value is returned. If the input is either "y", "yes"
// (any case) true is returned. Otherwise, false is returned.
func CheckYes(scnr scanner, msg string, defaultYes bool) bool {
	yn := "y/N"
	if defaultYes {
		yn = "Y/n"
	}

	res := GetInputString(scnr, fmt.Sprintf("%s", msg), yn)
	if defaultYes && res == yn {
		res = "y"
	}

	if StringInSlice(strings.ToLower(res), []string{"y", "yes"}) {
		return true
	}

	return false
}
