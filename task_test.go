package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestTaskRunner_ListAll(t *testing.T) {
	expected := []Task{
		{Name: "error"},
		{Name: "good-morning", Desc: "desc"},
		{Name: "hello", Desc: "desc"},
	}
	target := TaskRunner{
		Dir: "TestData",
	}
	got, err := target.ListAll()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(expected, got); diff != "" {
		t.Fatal(diff)
	}
}

func TestTaskRunner_Run(t *testing.T) {
	target := &TaskRunner{
		Dir: "TestData",
	}
	err := target.Run(Task{
		Name: "hello",
	})
	if err != nil {
		t.Fatal(err)
	}
}
