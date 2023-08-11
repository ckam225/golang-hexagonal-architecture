package main

import (
	"clean-arch-hex/internal/cache"
	"clean-arch-hex/internal/db/postgres"
	"clean-arch-hex/internal/server/rest"
	"fmt"
)

func main() {
	_db, err := postgres.New("")
	if err != nil {
		panic(err)
	}
	serv := rest.New(_db, cache.Cache{})
	serv.Start()

	fmt.Println(serv.Test())
}
