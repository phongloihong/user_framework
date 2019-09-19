package handlers

import (
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo"
	"github.com/phongloihong/user_framework/db/types"
	"github.com/phongloihong/user_framework/repository"
)

// AddStudent insert new studen to db
func AddStudent(c echo.Context) error {
	var req types.StudentAddReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "BadRequest", Message: err.Error()},
		)
	}

	v := validator.New()
	if errs := v.Struct(req); errs != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "Badrequest", Message: errs.Error()},
		)
	}

	newStudent, err := repository.Insert(req)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "Badrequest", Message: err.Error()},
		)
	}

	return c.JSON(http.StatusCreated, newStudent)
}

func UpdateStudent(c echo.Context) error {
	var req types.StudentUpdateReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "BadRequest", Message: err.Error()},
		)
	}

	objectId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "BadRequest", Message: "Invalid id"},
		)
	}

	v := validator.New()
	if errs := v.Struct(req); errs != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "Badrequest", Message: errs.Error()},
		)
	}

	result, err := repository.UpdateById(objectId, req)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "Badrequest", Message: err.Error()},
		)
	}

	return c.JSON(http.StatusCreated, result)
}

func DeleteStudent(c echo.Context) error {
	objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "BadRequest", Message: "Invalid id"},
		)
	}

	result, err := repository.DeleteById(objectId)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "Badrequest", Message: err.Error()},
		)
	}

	return c.JSON(http.StatusOK, result)
}
