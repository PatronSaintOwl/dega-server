package main

import (
	"log"
	"net/http"

	"github.com/factly/dega-server/config"
	"github.com/factly/dega-server/service"
	coreModel "github.com/factly/dega-server/service/core/model"
	factCheckModel "github.com/factly/dega-server/service/fact-check/model"
	"github.com/joho/godotenv"

	_ "github.com/factly/dega-server/docs" // docs is generated by Swag CLI, you have to import it.
)

// @title Dega API
// @version 1.0
// @description Dega server API

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("unable to load .env")
	}

	// db setup
	config.SetupDB()

	factCheckModel.Migration()
	coreModel.Migration()

	r := service.RegisterRoutes()

	if err = http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
