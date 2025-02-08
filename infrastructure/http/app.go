package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/jmechavez/email-registration/infrastructure/db"
	"github.com/jmechavez/email-registration/internal/ports/service"
)

func Start() {
	router := mux.NewRouter()

	dbUser := getDbUser()
	userRepositoryDb := db.NewUserRepoDb(dbUser)
	uh := UserHandlers{
		service.NewUserService(userRepositoryDb),
	}

	router.HandleFunc("/users", uh.GetAllUsersEmail).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getDbUser() *sqlx.DB {
	connStr := "user=admin password=admin123 dbname=email_dir sslmode=disable"
	userDb, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return userDb
}
