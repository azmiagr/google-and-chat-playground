package rest

import (
	"google-login/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) GoogleLogin(c *gin.Context) {
	url, state, err := r.service.OAuthService.GetGoogleLoginURL()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to generate google login url", err)
		return
	}

	c.SetCookie("oauth_state", state, 300, "/", "", false, true)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (r *Rest) GoogleCallback(c *gin.Context) {
	savedState, err := c.Cookie("oauth_state")
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid oauth state", err)
		return
	}
	c.SetCookie("oauth_state", "", -1, "/", "", false, true)

	// Get code dan state dari query params
	code := c.Query("code")
	state := c.Query("state")

	if code == "" {
		response.Error(c, http.StatusBadRequest, "Missing authorization code", nil)
		return
	}

	// Handle callback
	result, err := r.service.OAuthService.HandleGoogleCallback(code, state, savedState)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "OAuth callback failed", err)
		return
	}

	response.Success(c, http.StatusOK, "Login successful", result)
}
