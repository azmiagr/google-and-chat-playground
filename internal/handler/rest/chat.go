package rest

import (
	"google-login/entity"
	"google-login/model"
	"google-login/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) SendMessage(c *gin.Context) {
	idStr := c.Param("convoID")
	convoID, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid conversation id", err)
		return
	}

	var param model.SendMessageInput
	err = c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid input", err)
		return
	}

	user := c.MustGet("user").(*entity.User)
	param.UserID = user.UserID

	err = r.service.ChatService.SendMessage(param, convoID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to send message", err)
		return
	}

	response.Success(c, http.StatusOK, "success", nil)
}

func (r *Rest) CreateConversation(c *gin.Context) {
	var param model.CreateConversationInput
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid input", err)
		return
	}

	user := c.MustGet("user").(*entity.User)
	param.UserIDs = append(param.UserIDs, user.UserID)

	err = r.service.ChatService.CreateConversation(param)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to create conversation", err)
		return
	}

	response.Success(c, http.StatusCreated, "success to create conversation", nil)
}

func (r *Rest) GetMessagesByConversationID(c *gin.Context) {
	idStr := c.Param("convoID")
	convoID, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid conversation id", err)
		return
	}

	messages, err := r.service.ChatService.GetMessagesByConversationID(convoID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get conversation", err)
		return
	}

	response.Success(c, http.StatusOK, "success to get conversation", messages)
}

func (r *Rest) GetUserConversations(c *gin.Context) {
	user := c.MustGet("user").(*entity.User)

	convos, err := r.service.ChatService.GetConversationsByUser(user.UserID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get conversations", err)
		return
	}

	response.Success(c, http.StatusOK, "success to get conversations", convos)
}
