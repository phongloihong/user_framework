package handlers

import (
	"net/http"

	"github.com/phongloihong/user_framework/utils"

	"github.com/labstack/echo"
	"github.com/phongloihong/user_framework/db/types"
	"github.com/phongloihong/user_framework/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
)

// GetStudent return all studen in json
func GetStudents(c echo.Context) error {
	students, err := repository.Fetch()
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "BadRequest", Message: err.Error()},
		)
	}

	return c.JSON(http.StatusOK, students)
}

func GetStudent(c echo.Context) error {
	objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "BadRequest", Message: "Invalid id"},
		)
	}

	student, err := repository.GetById(objectId)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "BadRequest", Message: err.Error()},
		)
	}

	return c.JSON(http.StatusOK, student)
}

func SearchStudent(c echo.Context) error {
	var req types.StudentSearchRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "BadRequest", Message: err.Error()},
		)
	}

	students, err := repository.Find(req)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "BadRequest", Message: err.Error()},
		)
	}

	return c.JSON(http.StatusOK, students)
}

func GroupStudent(c echo.Context) error {
	result, err := repository.GroupStudent(c.Param("name"))
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "BadRequest", Message: err.Error()},
		)
	}

	return c.JSON(http.StatusOK, result)
}

func CreateUser(c echo.Context) error {
	var req types.UserAddReq
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

	newUser, err := repository.UserRepo.Create(req)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "Badrequest", Message: err.Error()},
		)
	}

	token, err := utils.GenerateToken(types.TokenClaims{
		newUser.InsertedID.(primitive.ObjectID).Hex(),
		req.Email,
	})

	return c.JSON(http.StatusCreated, bson.M{"token": token})
}

func Auth(c echo.Context) error {
	var req types.UserAddReq
	var user types.User
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

	err := repository.UserRepo.GetByEmail(req.Email).Decode(&user)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "Badrequest", Message: err.Error()},
		)
	}

	err = utils.CompareHashWithPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "Badrequest", Message: err.Error()},
		)
	}

	token, err := utils.GenerateToken(types.TokenClaims{
		user.ID.Hex(),
		req.Email,
	})

	return c.JSON(http.StatusOK, bson.M{"id": user.ID, "email": user.Email, "token": token})
}

// CheckHeath check server status
func CheckHeath(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
