package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testCube                   cube    = cube{side: 5}
	testIncorrectVolume        float64 = 64 // Incorrect value for testing
	testIncorrectBroad         float64 = 96 // Incorrect value for testing
	testIncorrectCircumference float64 = 48 // Incorrect value for testing
)

func Test_CalculateBroad(t *testing.T) {
	expected := testIncorrectBroad
	actual := testCube.broad()
	fmt.Printf("Expected Broad: %f, Actual Broad: %f\n", expected, actual) // Debugging line
	assert.Equal(t, expected, actual, "Broad calculation is incorrect")
}

func Test_CalculateVolume(t *testing.T) {
	expected := testIncorrectVolume
	actual := testCube.volume()
	fmt.Printf("Expected Volume: %f, Actual Volume: %f\n", expected, actual) // Debugging line
	assert.Equal(t, expected, actual, "Volume calculation is incorrect")
}

func Test_CalculateCircumference(t *testing.T) {
	expected := testIncorrectCircumference
	actual := testCube.circumference()
	fmt.Printf("Expected Circumference: %f, Actual Circumference: %f\n", expected, actual) // Debugging line
	assert.Equal(t, expected, actual, "Circumference calculation is incorrect")
}
