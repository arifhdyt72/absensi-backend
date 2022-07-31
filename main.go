package main

import (
	"absensi-backend/attendance"
	"absensi-backend/auth"
	"absensi-backend/config"
	"absensi-backend/handler"
	"absensi-backend/helper"
	"absensi-backend/user"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config, err := config.LoadConfig(".", os.Args)
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", config.DBUser, config.DBPassword, config.ServerAddress, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	userRepository := user.NewRepository(db)
	attendanceRepository := attendance.NewRepository(db)

	userService := user.NewService(userRepository)
	attendanceService := attendance.NewService(attendanceRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	attendanceHandler := handler.NewAttendanceHandler(attendanceService, authService, userService, config)

	router := gin.Default()
	configCors := cors.DefaultConfig()
	configCors.AllowAllOrigins = true
	configCors.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"}
	configCors.ExposeHeaders = []string{"Content-Length"}
	configCors.AllowCredentials = true
	router.Use(cors.New(configCors))
	router.Static("/images", "./images")

	api := router.Group("api/v1")

	// user endpoint
	api.POST("/user", authMiddleware(authService), userHandler.CreateUser)
	api.POST("/login", userHandler.LoginHandler)
	api.GET("/users", authMiddleware(authService), userHandler.AllUsersHandler)
	api.GET("/user", authMiddleware(authService), userHandler.GetUserByID)
	api.PUT("/user", authMiddleware(authService), userHandler.UpdateUserHandler)
	api.DELETE("/user/:id", authMiddleware(authService), userHandler.DeleteUserHandler)

	// attendance endpoint
	api.GET("/attendance", authMiddleware(authService), attendanceHandler.GetDetailAttendanceUserID)
	api.POST("/attendance-in", authMiddleware(authService), attendanceHandler.AttendanceInHandler)
	api.POST("/attendance-out", authMiddleware(authService), attendanceHandler.AttendanceOutHandler)
	api.GET("/attendances", authMiddleware(authService), attendanceHandler.AllAttendanceHandler)

	router.Run(config.PORT)
}

func authMiddleware(authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			errorMessage := gin.H{"error": "Unauthorized"}
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", errorMessage)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			errorMessage := gin.H{"error": err.Error()}
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", errorMessage)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			errorMessage := gin.H{"error": "Unauthorized"}
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", errorMessage)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))
		c.Set("currentUser", userID)
	}
}
