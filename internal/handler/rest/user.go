package rest

import (
	"google-login/model"
	"google-login/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) Register(c *gin.Context) {
	var param model.UserRegisterParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	err = r.service.UserService.Register(param)
	if err != nil {
		if err.Error() == "email already exists" {
			response.Error(c, http.StatusConflict, "email already exists", err)
			return
		} else if err.Error() == "password doesn't match" {
			response.Error(c, http.StatusConflict, "password doesn't match", err)
			return
		} else {
			response.Error(c, http.StatusInternalServerError, "failed to register user", err)
			return
		}
	}

	response.Success(c, http.StatusOK, "success to register user", nil)
}

func (r *Rest) Login(c *gin.Context) {
	var param model.UserLoginParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	token, err := r.service.UserService.Login(param)
	if err != nil {
		if err.Error() == "invalid email or password" {
			response.Error(c, http.StatusUnauthorized, "invalid email or password", err)
			return
		} else {
			response.Error(c, http.StatusInternalServerError, "failed to login user", err)
			return
		}
	}

	response.Success(c, http.StatusOK, "success to login user", token)

}
