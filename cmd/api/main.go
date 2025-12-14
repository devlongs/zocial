package main

import (
	"log"

	"github.com/devlongs/zocial/internal/db"
	"github.com/devlongs/zocial/internal/env"
	"github.com/devlongs/zocial/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", "8080"),
		db: dbConfig{
			dsn:          env.GetString("DB_DSN", "postgres://admin:adminpassword@localhost/zocial?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 25),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 25),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.dsn,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	log.Print("DB connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))

}
