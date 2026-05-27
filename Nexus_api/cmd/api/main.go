package main

import (
	"fmt"

	"github.com/MatheusMikio/Nexus/config"
	"github.com/MatheusMikio/Nexus/internal/routes"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Printf("Erro ao inicializar config: %s", err.Error())
		return
	}

	db := config.GetDb()
	routes.Init(db)
}
