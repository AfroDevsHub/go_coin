package constants

import "fmt"

type Theme int

const (
	SYSTEM Theme = iota
	DARK
	LIGHT
	RED
	GREEN
	BLUE
)

var themes = map[Theme]string{
	SYSTEM: "Defaults to Device Settings.",
	DARK:   "DARK Theme Preferred.",
	LIGHT:  "Light Theme Preferred.",
	RED:    "Red Theme Preferred.",
	GREEN:  "Green Theme Preferred.",
	BLUE:   "Blue Theme Preferred.",
}

func (t Theme) String() (string, error) {
	if value, ok := themes[t]; ok {
		return value, nil
	}
	return "", fmt.Errorf("invalid theme")
}
