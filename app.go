package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //using postgres
	"github.com/stevenkie/project-test/config"

	userDelivery "github.com/stevenkie/project-test/internal/delivery/http/user"
	sessionRedisR "github.com/stevenkie/project-test/internal/repository/session/redis"
	userRepoPG "github.com/stevenkie/project-test/internal/repository/userdb/postgres"
	userUsecase "github.com/stevenkie/project-test/internal/usecase/user"
)

// main app function

func main() {
	// Init resources
	cfg := config.GetConfig()
	db, err := ConnectDb(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	redis, err := ConnectRedis(cfg)
	if err != nil {
		panic(err)
	}
	defer redis.Close()

	// Init layers
	sessionRedisRepo := sessionRedisR.InitSessionRedisRepo(redis)
	userPGRepo := userRepoPG.InitUserPGRepo(db)
	userUsecase := userUsecase.InitUserUsecase(cfg, userPGRepo, sessionRedisRepo)
	userHttpDelivery := userDelivery.InitUserHttpDelivery(userUsecase)

	// Init routes
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", userHttpDelivery.GetUserByID).Methods(http.MethodGet)
	r.HandleFunc("/user", userHttpDelivery.InsertUser).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", userHttpDelivery.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/user/{id}", userHttpDelivery.DeleteUser).Methods(http.MethodDelete)
	r.HandleFunc("/login", userHttpDelivery.Login).Methods(http.MethodPost)

	// Start server
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type"},
	})
	handler := c.Handler(r)
	port := fmt.Sprintf(":%s", cfg.Server.Port)
	fmt.Printf("Server started at %s", port)
	log.Fatal(http.ListenAndServe(port, handler))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}

// ConnectDb connect to accounting db
func ConnectDb(cfg *config.Config) (*sqlx.DB, error) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Pass, cfg.DB.Name, cfg.DB.SSLMode)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return db, err
	}
	err = db.Ping()
	if err != nil {
		return db, err
	}
	return db, nil
}

func ConnectRedis(cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Host,
	})
	err := rdb.Ping(rdb.Context()).Err()
	return rdb, err
}
