package main

import (
	"fmt"
	"os"

	"github.com/dfunani/go_coin/lib/constants"
	"github.com/dfunani/go_coin/models"
	"github.com/dfunani/go_coin/models/warehouse"
	"github.com/dlclark/regexp2"
)

func main() {

	users := models.User{}
	user, err := users.Create_user("user@email.co.za", "funani@12")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cards := warehouse.Card{}
	card, err := cards.Create_card(constants.CardTypes.SAVINGS, "256381")
	re := regexp2.MustCompile(".* ID: (.*)", regexp2.None)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if isMatch, _ := re.MatchString(user.String()); isMatch {
		fmt.Println(user.Dict())
		fmt.Println(card.Dict())
	}
	os.Exit(0)
}
