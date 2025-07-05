package api

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/silentnova42/library-store-with-elasticsearch/internal/model"
)

type BookApi interface{}

type ProfileApi interface {
	AddProfile(ctx context.Context, userProfile model.UserProfile, passwordHash string) error
}

type Db interface {
	BookApi
	ProfileApi
}

type Hash interface {
	HashPassword(password string) (string, error)
	CompareHashAndPassword(password, hash string) error
}

type Auth interface {
	NewToken(userId int, ext time.Duration, key []byte) (string, error)
	Parse(accessToken string, accessKey []byte) (int, error)
	Refresh(refreshToken string, data model.DataForRefresh) (*model.ResponsToken, error)
}

type handler struct {
	db   Db
	hash Hash
	auth Auth
}

func NewHandler(db Db, hash Hash, auch Auth) *handler {
	return &handler{
		db:   db,
		hash: hash,
		auth: auch,
	}
}

func (h *handler) InitRouter() *gin.Engine {
	router := gin.New()
	router.POST("/", h.CreateProfile)
	return router
}
