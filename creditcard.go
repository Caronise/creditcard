// Package creditcard provides facilities for working with
// credit cards.
package creditcard

import (
	"fmt"
	"regexp"
)

// card represents a credit card.
type card struct {
	number string
}

// matchesPattern checks that the card number matches either of
// following patterns: 0000 1111 2222 3333 or 0000111122223333
func matchesPattern(number string) bool {
	pattern := `^\d{4}\s?\d{4}\s?\d{4}\s?\d{4}$`
	// ^ assets the start of a line
	// \d{4} matches exactly four digits
	// \s? matches zero or one whitespace character
	// $ asserts the end of a line
	match, _ := regexp.MatchString(pattern, number)
	return match
}

// New takes a credit card number and returns a 'card' value
// representing that card, or an error if the number is invalid.
func New(input string) (card, error) {
	if matchesPattern(input) {
		return card{
			number: input,
		}, nil
	}

	return card{}, fmt.Errorf("invalid input: %s", input)
}

// Number returns the card number.
func (c card) Number() string {
	return c.number
}

// SetNumber takes a string representing a credit card number
// and sets the card number to that number, or returns an error
// if the number is invalid.
func (c *card) SetNumber(number string) error {
	if matchesPattern(number) {
		c.number = number
		return nil
	}
	return fmt.Errorf("invalid number %q", number)
}
