package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", s)
}

func main() {
	if os.Getenv("LOAD_ENV_FILE") == "true" {
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}
	}

	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Executando...", os.Getenv("BASE_URL"))
	log.Println("Teste env...", os.Getenv("ORACLE_DB_USER"))
	log.Fatal(http.ListenAndServe(":3000", nil))
}
