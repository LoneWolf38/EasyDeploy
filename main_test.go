package main

import (
		"os"
		"testing"
)

func TestAWScreds(t *testing.T) {
	env := os.Getenv("TF_VAR_acc")
	if env != "asd"{
		t.Errorf("Environment variables not set coz value is %s", env)
	}
}