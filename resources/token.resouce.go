package resources

import (
	"github.com/lib/pq"
)

type OauthAccessToken struct {
	ID            string      `json:"id"`
	UserID        int         `json:"user_id"`
	ClientID      int         `json:"client_id"`
	Name          string      `json:"name"`
	Scopes        string      `json:"scopes"`
	Revoked       bool        `json:"revoked"`
	MultipleLogin bool        `json:"multiple_login"`
	ExpiresAt     pq.NullTime `json:"expires_at"`
	CreatedAt     pq.NullTime `json:"created_at"`
	UpdatedAt     pq.NullTime `json:"updated_at"`
}

