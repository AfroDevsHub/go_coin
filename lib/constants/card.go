package constants

import "fmt"

type CardType int

const (
	CHEQUE CardType = iota
	SAVINGS
	CREDIT
)

var cardtype = map[CardType]string{
	CHEQUE:  "1991",
	SAVINGS: "1992",
	CREDIT:  "1993",
}

func (c CardType) String() (string, error) {
	if value, ok := cardtype[c]; ok {
		return value, nil
	}
	return "", fmt.Errorf("invalid card type")
}
