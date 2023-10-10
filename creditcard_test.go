package creditcard_test

import (
	"creditcard"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input       string
		errExpected bool
	}

	testCases := []testCase{
		{input: "", errExpected: true},
		{input: "X", errExpected: true},
		{input: "AAAABBBBCCCCDDDD", errExpected: true},
		{input: "AAAA BBBB CCCC DDDD", errExpected: true},
		{input: "111122223333444D", errExpected: true},
		{input: "1111 2222 3333 444D", errExpected: true},
		{input: "1111222233334444", errExpected: false},
		{input: "1111 2222 3333 4444", errExpected: false},
	}

	for _, tc := range testCases {
		_, err := creditcard.New(tc.input)

		if tc.errExpected && err == nil {
			t.Error("expected an error creating new creditcard but got none")
		} else if !tc.errExpected && err != nil {
			t.Errorf("Got error creating new creditcard: %v", err)
		}
	}
}

func TestNumber(t *testing.T) {
	t.Parallel()

	input := "1234 1234 1234 1234"
	c, err := creditcard.New(input)
	if err != nil {
		t.Errorf("error creating new card: %v", err)
	}

	num := c.Number()
	if num != input {
		t.Errorf("wanted: %s, but got: %s", input, num)
	}
}

func TestSetNumber(t *testing.T) {
	t.Parallel()

	input := "1000 2000 3000 4000"
	c, err := creditcard.New(input)
	if err != nil {
		t.Errorf("error creating new card: %v", err)
	}

	number := "9999 8888 7777 6666"
	err = c.SetNumber(number)
	if err != nil {
		t.Errorf("error setting card number: %v", err)
	}

	// fmt.Printf("Card number is now: %q", c.Number())

	if number != c.Number() {
		t.Errorf("card number mismatch. Wanted: %q, but got %q", number, c.Number())
	}
}
