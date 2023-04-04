package utils

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/VikashChauhan51/go-sample/internal/models"
)

func GetConfig() *models.Config {
	// Read the config file
	var config models.Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatal(err)
	}

	// Print the config values
	fmt.Println("Server port:", config.Server.Port)
	fmt.Println("Server host:", config.Server.Host)
	fmt.Println("Database user:", config.Database.User)
	fmt.Println("Database pass:", config.Database.Pass)
	return &config
}
