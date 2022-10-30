package main

import (
	"github.com/reinhard18/go-postgre/initializers"
	"github.com/reinhard18/go-postgre/models"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnvInitializers()
}

func main() {
	initializers.DB.AutoMigrate(&models.Activity{})
}
