package api

import (
	"capstone_project/internal/api/handlers"
	"capstone_project/internal/middleware"
	"capstone_project/internal/repository"
	"capstone_project/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"os"
)

// @title Capstone Project API
// @version 1.0
// @description This is the API for the Capstone Project
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default_secret"
	}

	requiredSkillRepo := repository.NewRequiredSkillRepository(db)
	requiredSkillService := service.NewRequiredSkillService(requiredSkillRepo)
	requiredSkillHandler := handlers.NewRequiredSkillHandler(requiredSkillService)

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, jwtSecret)
	authHandlder := handlers.NewAuthHandler(authService)

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", authHandlder.LoginUser)
		authRoutes.POST("/register", authHandlder.CreateUser)
	}

	gameRepo := repository.NewGameRepository(db)
	gameService := service.NewGameService(gameRepo)
	gameHandler := handlers.NewGameHandler(gameService)

	authMiddleware := middleware.AuthMiddleware(authService)

	gameRoutes := router.Group("/games")
	{
		gameRoutes.GET("/", gameHandler.ListGames)
		gameRoutes.GET("/:id", gameHandler.GetGame)
		gameRoutes.POST("/", authMiddleware, gameHandler.CreateGame)
		gameRoutes.PUT("/:id", authMiddleware, gameHandler.UpdateGame)
		gameRoutes.DELETE("/:id", authMiddleware, gameHandler.DeleteGame)
	}

	gameRequestRepo := repository.NewGameRequestRepository(db)
	gameRequestService := service.NewGameRequestService(gameRequestRepo)
	gameRequestHandler := handlers.NewGameRequestHandler(gameRequestService, gameService)

	gameCategoryRepo := repository.NewGameCategoryRepository(db)
	gameCategoryService := service.NewGameCategoryService(gameCategoryRepo)
	gameCategoryHandler := handlers.NewGameCategoryHandler(gameCategoryService)

	gameCategoryRoutes := router.Group("/game-categories")
	{
		gameCategoryRoutes.GET("/", gameCategoryHandler.ListGameCategories)
		gameCategoryRoutes.POST("/", gameCategoryHandler.CreateGameCategory)
	}

	gameRequestRoutes := router.Group("/game-requests")
	{
		gameRequestRoutes.POST("/", authMiddleware, gameRequestHandler.CreateGameRequest)
		gameRequestRoutes.PUT("/:id", authMiddleware, gameRequestHandler.UpdateGameRequest)
	}
	requiredSkillRoutes := router.Group("/required-skills")
	{
		requiredSkillRoutes.POST("/", authMiddleware, requiredSkillHandler.CreateRequiredSkills)
		requiredSkillRoutes.GET("/", requiredSkillHandler.ListRequiredSkills)
	}
	userService := service.NewUserService(*userRepo)
	userHandler := handlers.NewUserHandler(userService)

	userRoutes := router.Group("/players")
	{
		userRoutes.GET("/", userHandler.ListUsers)
	}
}
