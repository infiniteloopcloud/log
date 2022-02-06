package log

import "fmt"

var p out = stdOut{}

type out interface {
	send(s string, passedLevel uint8)
}

type blank struct{}

func (blank) send(s string, passedLevel uint8) {}

type stdOut struct{}

func (stdOut) send(s string, passedLevel uint8) {
	if level >= passedLevel {
		//nolint:forbidigo
		fmt.Println(s)
	}
}
