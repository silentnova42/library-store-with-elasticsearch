package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/silentnova42/library-store-with-elasticsearch/internal/model"
)

type Auth struct{}

func (a *Auth) NewToken(userId int, ext time.Duration, key []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   fmt.Sprintf("%d", userId),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ext)),
	})

	return token.SignedString(key)
}

func (a *Auth) Parse(accessToken string, accessKey []byte) (int, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return accessKey, nil
	})

	if err != nil {
		return -1, err
	}

	if !token.Valid {
		return -1, fmt.Errorf("invalid token")
	}

	climes, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return -1, fmt.Errorf("failed to cast token claims to jwt.MapClaims")
	}

	id, ok := climes["id"].(int)
	if !ok {
		return -1, fmt.Errorf("failed to parse 'id' from token claims: expected int, got %T", climes["id"])
	}

	return id, nil
}

func (a *Auth) Refresh(refreshToken string, data model.DataForRefresh) (*model.ResponsToken, error) {
	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return data.RefreshKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	climes, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	id, ok := climes["id"].(int)
	if !ok {
		return nil, fmt.Errorf("failed to parse 'id' from token claims: expected int, got %T", climes["id"])
	}

	newAccessToken, err := a.NewToken(id, data.AccessExt, data.AccessKey)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := a.NewToken(id, data.AccessExt, data.AccessKey)
	if err != nil {
		return nil, err
	}

	resp := model.ResponsToken{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}

	return &resp, nil
}
