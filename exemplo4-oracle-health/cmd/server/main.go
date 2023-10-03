package main

import (
	"net/http"

	"github.com/heptiolabs/healthcheck"
	go_ora "github.com/sijms/go-ora/v2"
)

func main() {
	connStr := `(DESCRIPTION=
    (ADDRESS_LIST=
    	(LOAD_BALANCE=OFF)
        (FAILOVER=ON)
    	(address=(protocol=tcp)(host=localhost)(port=1521))
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
	defer db.Close()





	// Create a new health check handler
	h := healthcheck.NewHandler()

	// Register health checks for any dependencies
	//h.AddLivenessCheck("database", healthcheck.DatabasePingCheck("mysql", "user:password@tcp(db:3306)/mydatabase"))
	h.AddReadinessCheck("database", DatabasePingCheck(db, 1*time.Second))


	// Create an HTTP server and add the health check handler as a handler
	http.HandleFunc("/health", h.Handler)
	http.ListenAndServe(":3000", nil)
}