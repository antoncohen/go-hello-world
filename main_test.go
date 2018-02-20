package main

import (
	"os"
	"testing"
)

func TestGetEnvDefault(t *testing.T) {
	defaultVal := "THE_DEFAULT"
	value := getEnv("RANDOM_GLKBJCUBBLWUBXNMZU", defaultVal)
	if value != defaultVal {
		t.Errorf("getEnv incorrect, got: %s, want: %s.", value, defaultVal)
	}
}

func TestGetEnvSet(t *testing.T) {
	envVar := "RANDOM_BMMCLKEMMSKLMKLSFD"
	defaultVal := "THE_DEFAULT"
	correctVal := "THE_SET_VAL"

	os.Setenv(envVar, correctVal)
	value := getEnv(envVar, defaultVal)
	if value != correctVal {
		t.Errorf("getEnv incorrect, got: %s, want: %s.", value, correctVal)
	}
}

func TestGetEnvEmpty(t *testing.T) {
	envVar := "RANDOM_EJKDJKSHSKDJDKLSJ"
	defaultVal := "THE_DEFAULT"
	correctVal := ""

	os.Setenv(envVar, correctVal)
	value := getEnv(envVar, defaultVal)
	if value != correctVal {
		t.Errorf("getEnv incorrect, got: %s, want: %s.", value, correctVal)
	}
}
