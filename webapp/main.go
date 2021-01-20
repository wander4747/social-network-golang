package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	routers := router.Generate()
	utils.LoadTemplates()

	fmt.Println("Run app!")
	log.Fatal(http.ListenAndServe(":3000", routers))
}

//ver video 7
