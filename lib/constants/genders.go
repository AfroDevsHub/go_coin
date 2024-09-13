package constants

import "fmt"

type Gender int

const (
	MALE Gender = iota
	FEMALE
	OTHER_GENDER
)

type TupleGender struct {
	Gender string
	Symbol rune
}

var genders = map[Gender]TupleGender{
	MALE:         {"male", 'm'},
	FEMALE:       {"female", 'f'},
	OTHER_GENDER: {"other", 'o'},
}

func (g Gender) String() (TupleGender, error) {
	if value, ok := genders[g]; ok {
		return value, nil
	}
	return TupleGender{}, fmt.Errorf("invalid gender")
}
