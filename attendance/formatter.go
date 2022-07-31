package attendance

type AttendanceFormat struct {
	ID             int    `json:"id"`
	UserId         int    `json:"user_id"`
	TimeIn         string `json:"time_in"`
	TimeOut        string `json:"time_out"`
	TimeDuration   string `json:"time_duration"`
	DateAttendance string `json:"date_attendance"`
	File           string `json:"file"`
}

type JoinReportAttendance struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Fullname       string `json:"fullname"`
	Role           string `json:"role"`
	ID             int    `json:"id"`
	TimeIn         string `json:"time_in"`
	TimeOut        string `json:"time_out"`
	TimeDuration   string `json:"time_duration"`
	DateAttendance string `json:"date_attendance"`
	File           string `json:"file"`
}

func FormatAttendace(data Attendance, baseUrl string) AttendanceFormat {
	diff := data.TimeOut.Sub(data.TimeIn)

	var result AttendanceFormat
	result.ID = data.ID
	result.UserId = data.UserId
	result.TimeIn = data.TimeIn.Format("15:04:05")
	result.TimeOut = data.TimeOut.Format("15:04:05")
	result.TimeDuration = diff.String()
	result.DateAttendance = data.DateAttendance.Format("02 January 2006")
	result.File = baseUrl + "/" + data.File

	return result
}

func FormatReportAttendance(data JoinAttendanceUser, baseUrl string) JoinReportAttendance {
	var result JoinReportAttendance
	diff := data.TimeOut.Sub(data.TimeIn)
	result.Username = data.Username
	result.Email = data.Email
	result.Fullname = data.Fullname
	result.Role = data.Role
	result.ID = data.ID
	result.TimeIn = data.TimeIn.Format("15:04:05")
	result.TimeOut = data.TimeOut.Format("15:04:05")
	result.TimeDuration = diff.String()
	result.DateAttendance = data.DateAttendance.Format("02 January 2006")
	result.File = baseUrl + "/" + data.File

	return result
}

func FormatReportAttendances(data []JoinAttendanceUser, baseUrl string) []JoinReportAttendance {
	var result []JoinReportAttendance

	for _, attendance := range data {
		resultData := FormatReportAttendance(attendance, baseUrl)
		result = append(result, resultData)
	}

	return result
}
