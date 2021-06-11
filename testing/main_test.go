package main

import "testing"

func TestSum(t *testing.T) {
	type testCase struct {
		input1 int
		input2 int
		want   int
	}

	testCases := []testCase{testCase{12, 23, 35},
		testCase{10, 20, 30},
		testCase{1, 2, 3},
		testCase{1200, 2333, 3533},
		testCase{111, 123, 234}}

	for _, tC := range testCases {
		result := Sum(tC.input1, tC.input2)

		if tC.want != result {
			t.Errorf("Sum was incorrect, got: %d, want: %d. with input1=%d, input2=%d",
				result, tC.want, tC.input1, tC.input2)
		}
	}

}

func TestIsPalindrome(t *testing.T) {
	type testCase struct {
		input1 string
		want   bool
	}

	testCases := []testCase{testCase{"CIVIC", true},
		testCase{"MOM", true},
		testCase{"LEVEL", true},
		testCase{"", true},
		testCase{"Building", false}}

	for _, tC := range testCases {
		result := IsPalindrome(tC.input1)

		if tC.want != result {
			t.Errorf("IsPalindrome was incorrect, got: %v, want: %v. with input1=%s",
				result, tC.want, tC.input1)
		}
	}

}
