package storage

import (
	"context"

	"github.com/silentnova42/library-store-with-elasticsearch/internal/model"
)

func (d *db) AddProfile(ctx context.Context, userProfile model.UserProfile, passwordHash string) error {
	_, err := d.client.Exec(
		ctx,
		`
			INSERT INTO user_profile 
			(email, username, firstname, lastname, password) 
			VALUES ($1, $2, $3, $4, $5)
		`,
		userProfile.Email,
		userProfile.Username,
		userProfile.Firstname,
		userProfile.Lastname,
		passwordHash,
	)
	return err
}
