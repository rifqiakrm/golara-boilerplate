package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rifqiakrm/golara-boilerplate/resources"
	"github.com/rifqiakrm/golara-boilerplate/utils/cache"
	"github.com/rifqiakrm/golara-boilerplate/utils/responses"
	"github.com/spf13/viper"
	"net/http"
)

func GetTokenById(ctx *gin.Context, db *sql.DB, id string, uid int64) (resources.OauthAccessToken, responses.ApiResponseList) {
	var oauth resources.OauthAccessToken

	var query = `
		SELECT
			id,
			user_id,
			client_id,
			name,
			scopes,
			revoked,
			multiple_login,
			expires_at,
			created_at,
			updated_at
		FROM
			oauth_access_tokens
		WHERE
			id = $1
		AND user_id = $2`

	bytes, _ := cache.CheckAndRetrieve("token", uid, "token-by-id", id)

	if bytes != nil {
		if err := json.Unmarshal(bytes, &oauth); err != nil {
			return resources.OauthAccessToken{}, responses.ErrorApiResponse(http.StatusInternalServerError, fmt.Sprintf("error while umarshal data from redis : %v", err.Error()))
		}

		return oauth, nil
	}

	err := db.QueryRowContext(ctx, query, id, uid).Scan(
		&oauth.ID,
		&oauth.UserID,
		&oauth.ClientID,
		&oauth.Name,
		&oauth.Scopes,
		&oauth.Revoked,
		&oauth.MultipleLogin,
		&oauth.ExpiresAt,
		&oauth.CreatedAt,
		&oauth.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return resources.OauthAccessToken{}, responses.ErrorApiResponse(http.StatusNotFound, "user not found")
		} else {
			if viper.GetString("app.env") == "local" || viper.GetString("app.env") == "development" {
				return resources.OauthAccessToken{}, responses.ErrorApiResponse(http.StatusInternalServerError, err.Error())
			} else {
				return resources.OauthAccessToken{}, responses.ErrorApiResponse(http.StatusInternalServerError, "server error")
			}
		}
	}

	if err := cache.Query("token", uid, "token-by-id", id, oauth, 3600); err != nil {
		return resources.OauthAccessToken{}, responses.ErrorApiResponse(http.StatusInternalServerError, fmt.Sprintf("Error while caching query : %v", err.Error()))
	}

	return oauth, nil
}
