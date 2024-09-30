package main

import (
	"capstone_project/internal/api"
	"capstone_project/internal/config"
	"capstone_project/internal/database"
	"capstone_project/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.NewLogger()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	db, err := database.Initialize(cfg.DatabaseURL)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to initialize database: %s", cfg.DatabaseURL)
	}

	if err := database.Migrate(db); err != nil {
		log.Fatal().Err(err).Msg("Failed to run database migrations")
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	api.SetupRoutes(router, db)

	log.Info().Msgf("Server starting on %s", cfg.ServerAddress)
	err = router.Run(cfg.ServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
