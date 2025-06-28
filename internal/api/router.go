package api

import "github.com/gin-gonic/gin"

type bookApi interface {
	GetBook() string
}

type profileApi interface{}

type Db interface {
	bookApi
	profileApi
}

type handler struct {
	db Db
}

func NewHandler(db Db) *handler {
	return &handler{db: db}
}

func (h *handler) InitRouter() *gin.Engine {
	router := gin.New()
	router.GET("/", h.Get)
	return router
}
