package app_tests

import "testing"

func TestMain(t *testing.T) {
	t.Run("pow sub", func(t *testing.T) {
		testCompute1(t)
	})
	t.Run("add sub", func(t *testing.T) {
		testCompute2(t)
	})
}

func testCompute1(t *testing.T) {
	// https://golang.org/doc/go1.9#test-helper
	t.Helper()
	computeTests := []struct {
		in  int8
		out int8
	}{
		{0, 0},
		{1, 1},
		{2, 4},
		{3, 9},
	}

	for _, test := range computeTests {
		s := pow(test.in)

		if s != test.out {
			t.Errorf("Compute(%d) = %d, want %d", test.in, s, test.out)
		}
	}
}

func testCompute2(t *testing.T) {
	// https://golang.org/doc/go1.9#test-helper
	t.Helper()
	computeTests := []struct {
		in1 int8
		in2 int8
		out int8
	}{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 2},
		{2, 3, 5},
	}

	for _, test := range computeTests {
		s := add(test.in1, test.in2)

		if s != test.out {
			t.Errorf("Compute(%d, %d) = %d, want %d", test.in1, test.in2, s, test.out)
		}
	}
}

func pow(num int8) int8 {
	return num * num
}

func add(x, y int8) int8 {
	return x + y
}