package main

import "testing"

func TestHello(t *testing.T){

	got := Hello("Mohsen")
	want := HelloPrefix + "Mohsen"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}