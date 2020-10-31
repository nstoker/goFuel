package main

import (
	"testing"
)

func Test_Routing(t *testing.T) {
	r := initialiseRouting()

	if r == nil {
		t.Fatal("no routing returned")
	}
}
