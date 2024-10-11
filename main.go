/*Main: Application Entry Point.*/

package main

import (
	"log"
	"os"

	DATABASE "github.com/dfunani/go_coin/database"
	"github.com/dfunani/go_coin/services"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	DATABASE.SetupDatabase()
	services.UpdateBlockChain()
	log.Println(services.GetBlockChain())
	os.Exit(0)
}
