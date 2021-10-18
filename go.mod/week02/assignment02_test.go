package main

import (
	"fmt"
	"testing"
	"log"
)

func arrayStuff(t *testing.T) {
	t.Parallel()

	//creating exp array
	exp := [4]string{"John", "Paul", "George", "Ringo"}
	//prints exp array
	fmt.Println(exp)
	//creates empty act variable
	var act [4]string
	//iterates through the exp variable and copy its values into the act variable
	for i, x := range exp {
		act[i] = x
	}
	//itereates through the act variable and assets that its contents match that of the exp variable
	for i, a := range act {
		x := exp[i]
		if a != x {
			t.Errorf("expected value: %s, got: %s", x, a)
		}
	}
}

func sliceStuff(t *testing.T) {
	t.Parallel()

	exp := []string{"John", "Paul", "George", "Ringo"}
	fmt.Println(exp)
	var act []string

	for i, x := range exp {
		act[i] = x
	}

	if len(act) != len(exp){
		t.Errorf("expected slice length: %d, got: %d", len(exp), len(act))
	}

	for i, a := range act {
		x := exp[i]
		if a != x {
			t.Errorf("expected: %s, got: %s", x, a)
		}
	}
}

func mapStuff(t *testing.T) {
	t.Parallel()

	exp := map[string]string{

	exp["John"] == "guitar",
	exp["Paul"] == "bass",
	exp["George"] == "guitar",
	exp["Ringo"] == "drums",
	}
	//fmt.Println(exp)

	act := map[string][string]{}

	for i, a := range exp{
		act[i] = a
	}
	if len(act) != len(exp){
		t.Errorf("expected map length: %d, got: %d", len(exp), len(act))
	}
}
