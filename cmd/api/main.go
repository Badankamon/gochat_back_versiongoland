package main

import (
	"log"

	"github.com/Badankamon/gochat_backend/internal/config"
	"github.com/Badankamon/gochat_backend/internal/platform/cache"
	"github.com/Badankamon/gochat_backend/internal/platform/database"
	"github.com/Badankamon/gochat_backend/internal/platform/storage"
	"github.com/Badankamon/gochat_backend/internal/shared/logger"
	"github.com/Badankamon/gochat_backend/internal/shared/middleware"
	"go.uber.org/zap"

	// Auth Dependencies
	authUC "github.com/Badankamon/gochat_backend/internal/modules/auth/application/usecase"
	authEntity "github.com/Badankamon/gochat_backend/internal/modules/auth/domain/entity"
	authExt "github.com/Badankamon/gochat_backend/internal/modules/auth/infrastructure/external"
	authPostgres "github.com/Badankamon/gochat_backend/internal/modules/auth/infrastructure/persistence/postgres"
	authRedis "github.com/Badankamon/gochat_backend/internal/modules/auth/infrastructure/persistence/redis"
	authHttp "github.com/Badankamon/gochat_backend/internal/modules/auth/presentation/http"

	// QR Dependencies
	qrUC "github.com/Badankamon/gochat_backend/internal/modules/qr/application/usecase"
	qrEntity "github.com/Badankamon/gochat_backend/internal/modules/qr/domain/entity"
	qrPostgres "github.com/Badankamon/gochat_backend/internal/modules/qr/infrastructure/persistence/postgres"
	qrHttp "github.com/Badankamon/gochat_backend/internal/modules/qr/presentation/http"

	// User Dependencies
	userUC "github.com/Badankamon/gochat_backend/internal/modules/user/application/usecase"
	userHttp "github.com/Badankamon/gochat_backend/internal/modules/user/presentation/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. Logger
	logger.Init(cfg.App.Env)
	defer logger.Sync()

	// 3. Infrastructure
	database.Connect(cfg.Database)
	cache.Connect(cfg.Redis)

	// Local storage for uploads
	storageService := storage.NewLocalStorage("uploads", "http://localhost:"+cfg.App.Port+"/uploads")

	// 4. Migration (Auto-migrate for dev simplicity)
	if cfg.App.Env != "production" {
		// Enable UUID extension
		database.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

		if err := database.DB.AutoMigrate(&authEntity.User{}, &authEntity.Session{}, &qrEntity.QRTicket{}); err != nil {
			logger.Fatal("Failed to auto-migrate", zap.Error(err))
		}
	}

	// 5. Auth Module Setup
	userRepo := authPostgres.NewUserRepository(database.DB)
	sessionRepo := authPostgres.NewSessionRepository(database.DB)
	verificationRepo := authRedis.NewVerificationRepository(cache.RDB)
	smsService := authExt.NewMockSMSService()

	sendCodeUC := authUC.NewSendVerificationCodeUseCase(verificationRepo, smsService)
	registerUC := authUC.NewRegisterUseCase(userRepo, verificationRepo, sessionRepo, cfg)
	loginUC := authUC.NewLoginUseCase(userRepo, verificationRepo, sessionRepo, cfg)

	authHandler := authHttp.NewAuthHandler(sendCodeUC, registerUC, loginUC)

	// 6. User Module Setup
	getProfileUC := userUC.NewGetProfileUseCase(userRepo)
	updateProfileUC := userUC.NewUpdateProfileUseCase(userRepo)
	uploadAvatarUC := userUC.NewUploadAvatarUseCase(userRepo, storageService)

	userHandler := userHttp.NewUserHandler(getProfileUC, updateProfileUC, uploadAvatarUC)

	// 7. QR Module Setup
	ticketRepo := qrPostgres.NewTicketRepository(database.DB)

	genTicketUC := qrUC.NewGenerateTicketUseCase(ticketRepo, cfg)
	scanTicketUC := qrUC.NewScanTicketUseCase(ticketRepo, userRepo)

	qrHandler := qrHttp.NewQRHandler(genTicketUC, scanTicketUC)

	// 8. Router
	r := gin.Default()

	// Serve static files for uploads
	r.Static("/uploads", "./uploads")

	// Routes
	api := r.Group("/api/v1")
	authHttp.RegisterRoutes(api, authHandler)
	userHttp.RegisterRoutes(api, userHandler, middleware.AuthMiddleware(cfg.JWT))
	qrHttp.RegisterRoutes(api, qrHandler, middleware.AuthMiddleware(cfg.JWT))

	// 7. Start
	port := cfg.App.Port
	if port == "" {
		port = "8080"
	}
	logger.Info("Starting server on port " + port)
	if err := r.Run(":" + port); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
