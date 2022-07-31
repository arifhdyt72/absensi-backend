package handler

import (
	"absensi-backend/auth"
	"absensi-backend/helper"
	"absensi-backend/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var input user.UserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("create user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	result, err := h.userService.Insert(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("create user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully created user", http.StatusOK, "success", user.FormatUser(result))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) GetUserByID(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)
	result, err := h.userService.GetUserServiceByID(currentUser)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Get detail user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully get detail user", http.StatusOK, "success", user.FormatUser(result))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginHandler(c *gin.Context) {
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("login user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.LoginService(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ApiResponse("login user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(user)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("login user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	result := gin.H{"token": token}
	response := helper.ApiResponse("successfully login user", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) AllUsersHandler(c *gin.Context) {
	input := c.Request.URL.Query()
	var data user.DatatableInput

	page, _ := strconv.Atoi(input.Get("page"))
	first, _ := strconv.Atoi(input.Get("first"))
	rows, _ := strconv.Atoi(input.Get("rows"))
	pageCount, _ := strconv.Atoi(input.Get("pageCount"))

	data.Filters = input.Get("filters")
	data.Page = page
	data.First = first
	data.Rows = rows
	data.PageCount = pageCount
	data.SortField = input.Get("sortField")
	data.SortOrder = input.Get("sortOrder")

	result, err := h.userService.GetAllUsers(data)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("failed get all users", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	count, err := h.userService.GetCountDataUser(data)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("failed get all users", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUsers(result)
	response := helper.ApiResponse("successfully get all users", http.StatusOK, "error", gin.H{"data": formatter, "countData": count})
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateUserHandler(c *gin.Context) {
	var input user.UserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("update user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.UpdateUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("update user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully updated user", http.StatusOK, "success", user.FormatUser(newUser))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) DeleteUserHandler(c *gin.Context) {
	var input user.InputId
	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("delete user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	result, err := h.userService.DeleteUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("delete user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("successfully delete user", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}
