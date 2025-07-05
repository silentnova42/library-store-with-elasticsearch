package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silentnova42/library-store-with-elasticsearch/internal/model"
)

func (h *handler) SignIn(ctx *gin.Context) {

}

func (h *handler) SignUp(ctx *gin.Context) {

}

func (h *handler) CreateProfile(ctx *gin.Context) {
	var profile model.UserProfile

	if err := ctx.ShouldBind(&profile); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	password, err := h.hash.HashPassword(profile.Password)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if err := h.db.AddProfile(ctx.Request.Context(), profile, password); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, profile)
}
