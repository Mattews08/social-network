package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/hako/branca"
	_ "github.com/jackc/pgx/stdlib"

	"github.com/Mattews08/social-network.git/internal/handler"
	"github.com/Mattews08/social-network.git/internal/service"
)

const (
	databaseURL = ""
	port        = 3000
)

func main() {
	var (
		port        = env("PORT", "3000")
		origin      = env("ORIGIN", "http://localhost:"+port)
		databaseURL = env("DATABASE_URL", "postgresql://root@127.0.0.1:26257/?sslmode=disable&database=legendshub")
		brancaKey   = env("BRANCA_KEY", "supersecretkeyyoushouldnotcommit")
	)

	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatalf("could not open db connection: %v\n", err)
		return
	}

	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatalf("could not ping to db: %v\n", err)
		return
	}

	codec := branca.NewBranca(brancaKey)
	codec.SetTTL(uint32(service.TokenLifespan.Seconds()))
	s := service.New(db, codec, origin)
	h := handler.New(s)
	log.Printf("accepting connections on port %s\n", port)
	if err = http.ListenAndServe(":"+port, h); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}

func env(key, fallBackValue string) string {
	s := os.Getenv(key)
	if s == "" {
		return fallBackValue
	}

	return s
}
