package constants

import "fmt"

type Gender int

const (
	MALE Gender = iota
	FEMALE
	OTHER_GENDER
)

type tupleGender struct {
	gender string
	symbol rune
}

var genders = map[Gender]tupleGender{
	MALE:         {"male", 'm'},
	FEMALE:       {"female", 'f'},
	OTHER_GENDER: {"other", 'o'},
}

func (g Gender) String() (*tupleGender, error) {
	if value, ok := genders[g]; ok {
		return &value, nil
	}
	return nil, fmt.Errorf("invalid gender")
}
