package main

import (
	"log"
	v1 "veo/internal/api/v1"
	configs "veo/internal/config"
	"veo/internal/database"
	"veo/internal/repository"
	"veo/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := configs.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化数据库
	if err := database.Init(cfg.Database); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	// 初始化数据访问层
	userRepo := repository.NewUserRepository(database.GetDB())

	// 初始化业务逻辑层
	userService := service.NewUserService(userRepo)

	// 初始化 API 层
	accountAPI := v1.NewAccountAPI(userService)
	userAPI := v1.NewUserAPI(userService)

	// 启动 HTTP 服务器
	router := gin.Default()
	v1.SetupAccountRouter(router, accountAPI)
	v1.SetupUserRouter(router, userAPI)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
