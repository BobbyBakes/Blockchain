package test

import (
	"log"
	"os/exec"
	"reflect"
	"testing"
)

func TestNoArg(t *testing.T) {
	cmdOutput, err := exec.Command("../CopyrightedPhrases").Output()
	if err != nil {
		log.Fatal(err)
	}
	got := string(cmdOutput)
	want := `
Copyright Created
`
	if !reflect.DeepEqual(got, got) {
		t.Errorf("create = %v, want %v", got, want)
	}
}

func TestCreate(t *testing.T) {
	cmdOutput, err := exec.Command("../CopyrightedPhrases", "create").Output()
	if err != nil {
		log.Fatal(err)
	}
	got := string(cmdOutput)
	want := `
Copyright Created
`
	if !reflect.DeepEqual(got, got) {
		t.Errorf("create = %v, want %v", got, want)
	}
}
