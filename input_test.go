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
	"testing"

	"github.com/magiconair/properties/assert"
)

type mockScanner struct {
	scannerOut   []bool
	textOut      []string
	scannerIndex int
	textIndex    int
}

func (m *mockScanner) Scan() bool {
	m.scannerIndex++
	return m.scannerOut[m.scannerIndex-1]
}

func (m *mockScanner) Text() string {
	m.textIndex++
	return m.textOut[m.textIndex-1]
}

func TestGetInputInt(t *testing.T) {
	tables := []struct {
		scannerOut []bool
		textOut    []string
		def        int
		out        int
		msg        string
	}{
		{[]bool{true, false}, []string{"10"}, 0, 10, "case 1"},
		{[]bool{true, false}, []string{""}, 10, 10, "case 2"},
		{[]bool{true, true, false}, []string{"abc", "10"}, 0, 10, "case 3"},
	}

	for _, table := range tables {
		ms := &mockScanner{scannerOut: table.scannerOut, textOut: table.textOut}

		out := GetInputInt(ms, "", table.def)
		assert.Equal(t, table.out, out, table.msg)
	}
}

func TestGetInputString(t *testing.T) {
	tables := []struct {
		scannerOut []bool
		textOut    []string
		def        string
		out        string
		msg        string
	}{
		// Case 1: no default value, got input
		{[]bool{true, false}, []string{"abc"}, "", "abc", "case 1"},
		// Case 2: default value, empty input
		{[]bool{true, false}, []string{""}, "abc", "abc", "case 2"},
		// Case 3: no default value, simulated error
		{[]bool{false}, []string{}, "", "", "case 3"},
	}

	for _, table := range tables {
		ms := &mockScanner{scannerOut: table.scannerOut, textOut: table.textOut}

		out := GetInputString(ms, "", table.def)
		assert.Equal(t, table.out, out, table.msg)
	}
}

func TestCheckYes(t *testing.T) {
	tables := []struct {
		scannerOut []bool
		textOut    []string
		defYes     bool
		out        bool
		msg        string
	}{
		// Case 1: Y
		{[]bool{true, false}, []string{"Y"}, false, true, "case 1"},
		// Case 2: y
		{[]bool{true, false}, []string{"y"}, false, true, "case 2"},
		// Case 3: yes
		{[]bool{true, false}, []string{"yes"}, false, true, "case 3"},
		// Case 4: Non-yes
		{[]bool{true, false}, []string{"N"}, false, false, "case 4"},
		// Case 5: Empty, default no
		{[]bool{true, false}, []string{""}, false, false, "case 5"},
		// Case 6: Empty, default yes
		{[]bool{true, false}, []string{""}, true, true, "case 6"},
	}

	for _, table := range tables {
		ms := &mockScanner{scannerOut: table.scannerOut, textOut: table.textOut}

		out := CheckYes(ms, "", table.defYes)
		assert.Equal(t, table.out, out, table.msg)
	}
}
