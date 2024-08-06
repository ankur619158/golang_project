package server

import (
	"golang_project/api/definition"
	"golang_project/serverconst"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userApi interface {
	CreatePerson(definition.CreatePersonRequest) definition.CreatePersonResponse
	GetPersonInfo(int64) definition.GetPersonInfoResponse
}

type userHandler struct {
	userApi userApi
}

func NewUserHandler(uapi userApi) *userHandler {
	return &userHandler{
		userApi: uapi,
	}
}

func (u *userHandler) CreatePerson() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var userDetails definition.CreatePersonRequest
		err := c.ShouldBindJSON(&userDetails)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		apiRes := u.userApi.CreatePerson(userDetails)

		if !apiRes.Error.Message.IsZero() {
			c.JSON(getErrorStatusCode(apiRes.Error.Message.String), apiRes.Error)
			return
		}
		c.JSON(http.StatusOK, apiRes)
	}
}

func (u *userHandler) GetPersonInfoByPersonID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")

		personId, err := strconv.ParseInt(ctx.Param("person_id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, serverconst.VErrIdNotProvided)
			return
		}

		apiRes := u.userApi.GetPersonInfo(personId)

		if !apiRes.Error.Message.IsZero() {
			ctx.JSON(getErrorStatusCode(apiRes.Error.Message.String), apiRes.Error)
			return
		}

		ctx.JSON(http.StatusOK, apiRes)
	}
}
