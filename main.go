// Package main provides ...
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diegoparra/pet-adocao-frontend/src/config"
	"github.com/diegoparra/pet-adocao-frontend/src/cookies"
	"github.com/diegoparra/pet-adocao-frontend/src/router"
	"github.com/diegoparra/pet-adocao-frontend/src/utils"
)

func main() {
	config.Load()
	cookies.Config()
	utils.LoadTemplate()
	r := router.Generate()

	fmt.Printf("Running webapp on port: %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
