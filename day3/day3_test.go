package main

import (
	"strconv"
	"testing"
)

func TestP1(t *testing.T) {
	answerP1 := 4361
	actualP1 := p1("example.txt")
	if actualP1 != answerP1 {
		t.Error("P1 was: " + strconv.Itoa(actualP1) + " but needs to be: " + strconv.Itoa(answerP1))
	}
}

func TestP2(t *testing.T) {
	answer := 467835
	actual := p2("example.txt")
	if actual != answer {
		t.Error("P2 was: " + strconv.Itoa(actual) + " but needs to be: " + strconv.Itoa(answer))
	}
}
