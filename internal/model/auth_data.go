package model

import "time"

type ResponsToken struct {
	AccessToken  string
	RefreshToken string
}

type DataForRefresh struct {
	AccessKey  []byte
	RefreshKey []byte
	AccessExt  time.Duration
	RefreshExt time.Duration
}
