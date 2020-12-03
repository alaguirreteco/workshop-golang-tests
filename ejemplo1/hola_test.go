package main

import "testing"

func TestHolaEmptyArg(t *testing.T) {

	emptyResult := hola("")

	if emptyResult != "Hola que tal" {
		t.Errorf("hola() : FAILED, expected %s, got %s'", "Hola que tal", emptyResult)
	} else {
		t.Logf("hola() : PASSED, expected %s, got %s'", "Hola que tal", emptyResult)
	}
}

func TestHolaValidArg(t *testing.T) {

	result := hola("Carlin")

	if result != "Hola Carlin" {
		t.Errorf("hola() : FAILED, expected %s, got %s'", "Hola Carlin", result)
	} else {
		t.Logf("hola() : PASSED, expected %s, got %s'", "Hola Carlin", result)
	}
}
