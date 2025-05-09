package main

import "testing"

func TestThatTheOutPutIsTrue(t *testing.T) {
	valueString, stringGoal := "abcde", "cdeab"
	got := check(valueString, stringGoal)
	want := true
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}

//print(check("abcde", "cdeab"))
//print(check("ab", "ba"))
//print(check("cba", "abc"))
//print(check("13456", "45613"))
