package main

import (
	"database/sql"
	"net/http"
	"time"
	"strings"
	"log"
	"fmt"
	"net/http/httptest"
	"net/http/httputil"

	"github.com/heptiolabs/healthcheck"
	go_ora "github.com/sijms/go-ora/v2"
)

func connectToDatabase() *sql.DB {
	connStr := `(DESCRIPTION=
    (ADDRESS_LIST=
    	(LOAD_BALANCE=OFF)
        (FAILOVER=ON)
    	(address=(protocol=tcp)(host=oracle-xe)(port=1521))
    )
    (CONNECT_DATA=
    	(SERVICE_NAME=XE)
        (SERVER=DEDICATED)
    )
    (SOURCE_ROUTE=yes)
    )`
	databaseUrl := go_ora.BuildJDBC("sato", "abc123", connStr, nil)
	db, err := sql.Open("oracle", databaseUrl)

	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select nome, cpf from usuario")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var nome string
		var cpf string
		err := rows.Scan(&nome, &cpf)
		if err != nil {
			panic(err)
		}
		fmt.Println(nome, cpf)
	}
	return db
}

func databasePingCheck(w http.ResponseWriter, r *http.Request) {

	var database *sql.DB
	database = connectToDatabase()
	defer database.Close()

	// Create a new health check handler
	h := healthcheck.NewHandler()

	// Register health checks for any dependencies
	h.AddReadinessCheck("database", healthcheck.DatabasePingCheck(database, 1*time.Second))

	fmt.Fprintf(w, "%s", dumpRequest(h, "GET", "/ready?full=1"))
}

func dumpRequest(handler http.Handler, method string, path string) string {
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		panic(err)
	}
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	dump, err := httputil.DumpResponse(rr.Result(), true)
	if err != nil {
		panic(err)
	}
	return strings.Replace(string(dump), "\r\n", "\n", -1)
}


func main() {
	// Create an HTTP server and add the health check handler as a handler
	http.HandleFunc("/health", databasePingCheck)
	// Make a request to the readiness endpoint and print the response.
    log.Println("Servidor de healthcheck em execução na porta :3000")
    if err := http.ListenAndServe(":3000", nil); err != nil {
        log.Fatal(err)
    }
}
