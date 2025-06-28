package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Get(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, h.db.GetBook())
}
