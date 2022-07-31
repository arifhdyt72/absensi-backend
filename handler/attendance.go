package handler

import (
	"absensi-backend/attendance"
	"absensi-backend/auth"
	"absensi-backend/config"
	"absensi-backend/helper"
	"absensi-backend/user"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type attendanceHandler struct {
	attendanceService attendance.Service
	authService       auth.Service
	userService       user.Service
	configENV         config.Config
}

func NewAttendanceHandler(attendanceService attendance.Service, authService auth.Service, userService user.Service, configENV config.Config) *attendanceHandler {
	return &attendanceHandler{attendanceService, authService, userService, configENV}
}

func (h *attendanceHandler) GetDetailAttendanceUserID(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)
	result, err := h.attendanceService.GetAttendanceNowByUserID(currentUser)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("get attencance failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("successfully get attendance", http.StatusOK, "success", attendance.FormatAttendace(result, h.configENV.BASE_URL))
	c.JSON(http.StatusOK, response)
}

func (h *attendanceHandler) AttendanceInHandler(c *gin.Context) {
	var input attendance.AttendanceIn
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("attendance in failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Failed to upload image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//simpan gambar di folder images/
	currentUser := c.MustGet("currentUser").(int)
	nameFile := fmt.Sprintf("%d-%d-%s", time.Now().Unix(), currentUser, file.Filename)
	path := fmt.Sprintf("images/%s", nameFile)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Failed to upload image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	result, err := h.attendanceService.AttendanceInService(currentUser, input, nameFile)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("attendance in failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully attendance in", http.StatusOK, "success", attendance.FormatAttendace(result, h.configENV.BASE_URL))
	c.JSON(http.StatusOK, response)

}

func (h *attendanceHandler) AttendanceOutHandler(c *gin.Context) {
	var input attendance.AttendanceOut
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("attendance out failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	result, err := h.attendanceService.AttendanceOutService(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("attendance out failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully attendance out", http.StatusOK, "success", attendance.FormatAttendace(result, h.configENV.BASE_URL))
	c.JSON(http.StatusOK, response)
}

func (h *attendanceHandler) AllAttendanceHandler(c *gin.Context) {
	input := c.Request.URL.Query()
	var data attendance.DatatableInput

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

	currentUser := c.MustGet("currentUser").(int)

	dataUser, err := h.userService.GetUserServiceByID(currentUser)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("failed get all attendance", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	result, err := h.attendanceService.GetAllAttendance(data, dataUser)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("failed get all attendance", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	count, err := h.attendanceService.GetCountDataAttendance(data, dataUser)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("failed get all attendance", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := attendance.FormatReportAttendances(result, h.configENV.BASE_URL)
	response := helper.ApiResponse("successfully get all attendance", http.StatusOK, "error", gin.H{"data": formatter, "countData": count})
	c.JSON(http.StatusOK, response)
}
