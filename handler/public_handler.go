package handler

import (
	"net/http"

	"github.com/Triluong/nc-student/db"
	"github.com/labstack/echo/v4"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func TestPublic(c echo.Context) error {
	db.Test()
	return c.String(http.StatusOK, "TestPublic")
}

func GetAllStudents(c echo.Context) error {
	students, err := db.GetStudents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, students)
}

func AddStudent(c echo.Context) error {
	var student db.Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := db.AddStudent(student)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, result)
}

func SearchStudentSimple(c echo.Context) error {
	var req db.StudentSearchRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	students, err := db.SerchStudentSimple(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, students)
}
