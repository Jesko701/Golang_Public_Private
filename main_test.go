package main

import "testing"

var (
	myCube              cube    = cube{side: 4}
	actualVolume        float64 = 64
	actualBroad         float64 = 96
	actualCircumference float64 = 48
)

// * No Library Testing
func TestCalculateVolume(t *testing.T) {
	t.Logf("Volume : %.2f", myCube.volume())
	if myCube.volume() != actualVolume {
		t.Errorf("Wrong!, must the actual %.2f", actualVolume)
	}
}

func TestCalculateBroad(t *testing.T) {
	t.Logf("Broad : %.2f", myCube.broad())
	if myCube.broad() != actualBroad {
		t.Errorf("Wrong!, Broad must be actual %.2f", actualBroad)
	}
}

func TestCalculateCircumference(t *testing.T) {
	t.Logf("Circumference : %.2f", myCube.circumference())
	if myCube.circumference() != actualCircumference {
		t.Errorf("Wrong, Circumference must be actual %.2f", actualCircumference)
	}
}

// * benchmarking
func BenchmarkCalculateVolume(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myCube.broad()
	}
}
