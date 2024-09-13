package utils

import (
	"fmt"

	"github.com/dlclark/regexp2"
)

func GetModelID(model string) (string, error) {

	reg := regexp2.MustCompile(`.* ID: (.*)`, regexp2.None)
	val, err := reg.FindStringMatch(model)
	if err != nil {
		return "", err
	}
	if val.GroupCount() != 2 {
		return "", fmt.Errorf("invalid string representation of the model")
	}
	return val.GroupByNumber(1).String(), nil
}
