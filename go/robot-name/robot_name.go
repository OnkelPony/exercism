// Package robotname implements methods for generating unique name for robot
// and resetting its name to empty value.
package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var randGen = rand.New(rand.NewSource(time.Now().UnixNano()))
var allNamesCount = 26 * 26 * 10 * 10 * 10
var allNames = make([]string, allNamesCount)
var initialRun = true

// Robot has unique name
type Robot struct {
	name string
}

// Name returns robot name consisting of two characters and three digits.
func (r *Robot) Name() (string, error) {
	if len(allNames) == 0 {
		return "", errors.New("names over")
	}
	if initialRun {
		genNames()
		initialRun = false
	}
	if r.name == "" {
		r.name = allNames[len(allNames)-1]
		allNames = allNames[:len(allNames)-1]
	}
	return r.name, nil
}

// Reset resets robot name to empty string.
func (r *Robot) Reset() {
	r.name = ""
}

// genNames generates all possible names of robots.
func genNames() {
	var p, s rune
	var t int
	for i := 0; i < allNamesCount; i++ {
		t = i % 1000
		s = rune(i/1000%26) + 'A'
		p = rune(i/26000%26) + 'A'
		allNames[i] = fmt.Sprintf("%c%c%03d", p, s, t)
	}
	randGen.Shuffle(len(allNames), func(i, j int) {
		allNames[i], allNames[j] = allNames[j], allNames[i]
	})
}
