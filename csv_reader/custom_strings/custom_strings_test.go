package custom_strings

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestLen(t *testing.T) {

	path := "../csv/test_len.csv"
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("Unable to load CSV file %s", err.Error())
	}

	r := csv.NewReader(f)
	testCases, err := r.ReadAll()
	if err != nil {
		t.Fatalf("Unable to read test cases: err: %s", err.Error())
	}

	for _, testC := range testCases {
		input := string(testC[0])
		expect := string(testC[1])
		if got := Len(input); got != expect {
			t.Errorf("Failed expecting %s and got %s, input %s",
				expect, got, input)
		}
	}

}
