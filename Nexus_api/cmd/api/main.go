package main

import (
	"fmt"

	"github.com/MatheusMikio/Nexus/config"
	"github.com/MatheusMikio/Nexus/internal/routes"
)

// @title Nexus API
// @version 1.0
// @description API para gerenciamento de usuarios, metas e tarefas.
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Digite "Bearer " seguido do token JWT.
func main() {
	err := config.Init()
	if err != nil {
		fmt.Printf("Erro ao inicializar config: %s", err.Error())
		return
	}

	db := config.GetDb()
	if err := routes.Init(db); err != nil {
		fmt.Printf("Erro ao inicializar servidor: %s", err.Error())
	}
}
