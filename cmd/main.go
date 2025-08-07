package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rafael-andrade-se/navo/api/domain/services"
	"github.com/rafael-andrade-se/navo/api/infrastructure/database"
	"github.com/rafael-andrade-se/navo/api/infrastructure/repository"
	"github.com/rafael-andrade-se/navo/api/interfaces/handlers"
	"github.com/rafael-andrade-se/navo/api/internal/routes"
	"github.com/rafael-andrade-se/navo/api/internal/server"
	"github.com/rafael-andrade-se/navo/api/pkg/logger"
)

func main() {
	// Carregar variáveis de ambiente
	godotenv.Load(".env")

	// Inicializa logger
	logger.InitLogger()
	logger.Log.Info("Iniciando servidor NAVO...")

	// Conectar ao banco
	db, err := database.ConectarPostgres()
	if err != nil {
		logger.Log.Error("Erro ao conectar com o banco de dados", "erro", err)
		return
	}

	// Repositórios e serviços
	usuarioRepo := repository.NovoUsuarioRepositoryPG(db)
	usuarioService := services.NovoUsuarioService(usuarioRepo)

	veiculoRepo := repository.NovoVeiculoRepositoryPG(db)
	veiculoService := services.NovoVeiculoService(veiculoRepo)

	plataformaRepo := repository.NovaPlataformaRepositoryPG(db)
	plataformaService := services.NovaPlataformaService(plataformaRepo)

	tipoGastoRepo := repository.NovoTipoGastoRepositoryPG(db)
	tipoGastoService := services.NovoTipoGastoService(tipoGastoRepo)

	// Handlers
	usuarioHandler := handlers.NewUsuarioHandler(usuarioService)
	authHandler := handlers.NewAuthHandler(usuarioService)
	veiculoHandler := handlers.NewVeiculoHandler(veiculoService)
	plataformaHandler := handlers.NewPlataformaHandler(plataformaService)
	tipoGastoHandler := handlers.NewTipoGastoHandler(tipoGastoService)

	// Criar servidor
	router := gin.Default()
	server.SetupMiddleware(router)

	// Registrar rotas com injeção de handlers
	routes.RegisterRoutes(router, usuarioHandler, authHandler, veiculoHandler, plataformaHandler, tipoGastoHandler)

	// Rodar servidor
	logger.Log.Info("Servidor rodando na porta 3000 🚀")
	if err := router.Run(":3000"); err != nil {
		logger.Log.Error("Erro ao iniciar servidor", "erro", err)
	}
}
