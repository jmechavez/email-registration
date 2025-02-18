package http

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
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
	router.HandleFunc("/users/{id_no}/user", uh.NewUser).Methods(http.MethodPost)

	// Set CORS options
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Change this to specific origins if needed
		handlers.AllowedMethods(
			[]string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
				http.MethodOptions,
			},
		),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	log.Fatal(http.ListenAndServe("localhost:8000", corsHandler(router)))
}

func getDbUser() *sqlx.DB {
	connStr := "user=admin password=admin123 dbname=email_dir sslmode=disable"
	userDb, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return userDb
}
