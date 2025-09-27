package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"quizgen-backend/internal/config"
	"quizgen-backend/internal/handlers"
	"quizgen-backend/internal/middleware"
	"quizgen-backend/internal/services"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize configuration
	cfg := config.New()

	// Initialize Firebase services
	firebaseClient, err := services.NewFirebaseClient(cfg.FirebaseCredentialsPath)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase client: %v", err)
	}

	// Initialize OpenAI service
	openaiService := services.NewOpenAIService(cfg.OpenAIAPIKey)

	// Initialize quiz service with dependencies
	quizService := services.NewQuizService(firebaseClient, openaiService)
	
	// Initialize handlers
	quizHandler := handlers.NewQuizHandler(quizService)
	authHandler := handlers.NewAuthHandler(firebaseClient)
	leaderboardHandler := handlers.NewLeaderboardHandler(firebaseClient)

	// Setup Gin router
	router := gin.New()
	
	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.RequestID())
	router.Use(middleware.Logger())
	
	// CORS configuration for Flutter apps
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{
		"http://localhost:3000",
		"https://*.vercel.app",
		"https://*.replit.app",
		"capacitor://localhost", // Capacitor iOS
		"ionic://localhost",     // Capacitor Android
	}
	corsConfig.AllowHeaders = []string{
		"Origin", "Content-Length", "Content-Type", "Authorization",
		"X-Requested-With", "X-Request-ID",
	}
	corsConfig.AllowMethods = []string{
		"GET", "POST", "PUT", "DELETE", "OPTIONS",
	}
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "healthy",
			"timestamp": time.Now().UTC(),
			"version":   "1.0.0",
		})
	})

	// API routes
	api := router.Group("/api/v1")
	{
		// Authentication routes
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
			auth.POST("/logout", middleware.AuthRequired(firebaseClient), authHandler.Logout)
			auth.GET("/profile", middleware.AuthRequired(firebaseClient), authHandler.GetProfile)
			auth.PUT("/profile", middleware.AuthRequired(firebaseClient), authHandler.UpdateProfile)
		}

		// Quiz routes
		quizzes := api.Group("/quizzes")
		quizzes.Use(middleware.AuthRequired(firebaseClient))
		{
			quizzes.POST("/generate", quizHandler.GenerateQuiz)
			quizzes.GET("/:id", quizHandler.GetQuiz)
			quizzes.POST("/:id/start", quizHandler.StartQuiz)
			quizzes.POST("/:id/answer", quizHandler.SubmitAnswer)
			quizzes.POST("/:id/complete", quizHandler.CompleteQuiz)
			quizzes.GET("/history", quizHandler.GetQuizHistory)
		}

		// Multiplayer routes
		multiplayer := api.Group("/multiplayer")
		multiplayer.Use(middleware.AuthRequired(firebaseClient))
		{
			multiplayer.POST("/create", quizHandler.CreateMultiplayerRoom)
			multiplayer.POST("/join/:roomId", quizHandler.JoinMultiplayerRoom)
			multiplayer.GET("/room/:roomId", quizHandler.GetMultiplayerRoom)
			multiplayer.POST("/room/:roomId/start", quizHandler.StartMultiplayerQuiz)
		}

		// Leaderboard routes
		leaderboard := api.Group("/leaderboard")
		leaderboard.Use(middleware.AuthRequired(firebaseClient))
		{
			leaderboard.GET("/global", leaderboardHandler.GetGlobalLeaderboard)
			leaderboard.GET("/friends", leaderboardHandler.GetFriendsLeaderboard)
			leaderboard.GET("/subject/:subjectId", leaderboardHandler.GetSubjectLeaderboard)
		}

		// Admin routes
		admin := api.Group("/admin")
		admin.Use(middleware.AuthRequired(firebaseClient))
		admin.Use(middleware.AdminRequired())
		{
			admin.POST("/subjects", quizHandler.CreateSubject)
			admin.PUT("/subjects/:id", quizHandler.UpdateSubject)
			admin.DELETE("/subjects/:id", quizHandler.DeleteSubject)
			admin.POST("/topics", quizHandler.CreateTopic)
			admin.PUT("/topics/:id", quizHandler.UpdateTopic)
			admin.DELETE("/topics/:id", quizHandler.DeleteTopic)
			admin.POST("/content/upload", quizHandler.UploadContent)
		}
	}

	// WebSocket endpoint for real-time multiplayer
	router.GET("/ws", middleware.WebSocketUpgrade(), quizHandler.HandleWebSocket)

	// Start server with graceful shutdown
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on port %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}