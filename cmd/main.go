package main

import (
	"clean-arch-hex/internal/cache/memcache"
	"clean-arch-hex/internal/controller/server/grpc"
	"clean-arch-hex/internal/db/postgres"
	"context"
	"fmt"
	"log"
)

func main() {

	ctx := context.Background()
	const (
		dbUser     = "postgres"
		dbPassword = "root"
		dbHost     = "localhost"
		dbPort     = "5432"
		dbName     = "tests"
		dbSSLMode  = "disable"
	)
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
		"postgres",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
		dbSSLMode,
	)

	_db, err := postgres.New(ctx, dsn, 3)
	if err != nil {
		panic(err)
	}
	serv := grpc.NewServer(_db, memcache.New())
	log.Fatal(serv.Start())
}
