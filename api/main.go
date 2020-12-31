package main

import (
	"api/src/config"
	"api/src/router"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

func init2() {
	key := make([]byte, 64)

	if _, erro := rand.Read(key); erro != nil {
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(key)
	fmt.Println(stringBase64)
}

func main() {
	config.Load()
	fmt.Printf("Running API port %d", config.Port)

	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
