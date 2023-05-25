package handlers

import (
	"auth/api/authentication"
	"auth/api/authentication/generation"
	"auth/cipher"
	"auth/dal"
	"auth/jwt"
	"auth/utils/httpUtils"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Register
// @Description to_register user
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			request body dal.User true "query params"
// @Success		200	{object}	dal.OkRegisterResponse
// @Router			/register [post]
func Register(repository dal.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqModel := dal.User{}
		err := json.NewDecoder(r.Body).Decode(&reqModel)
		if err != nil {
			log.Println(err)
			httpUtils.BadRequest(w)
			return
		}

		err = repository.Create(&reqModel)

		err = json.NewEncoder(w).Encode(&authentication.Access{
			AccessToken: authentication.Token{
				Token: accessToken(&reqModel),
				TTL:   86400,
			},
			RefreshToken: authentication.Token{
				Token: refreshToken(&reqModel),
				TTL:   604800,
			},
		})
		if err != nil {
			log.Println(err)
			httpUtils.BadRequest(w)
			return
		}
	}
}

func accessToken(user *dal.User) string {
	accessTokenClaims := generation.AccessTokenClaims{
		UserId:            user.Id,
		Email:             user.Email,
		CreationTimestamp: time.Now().UTC().Unix(),
		TTL:               86400,
		US:                false,
	}
	return jwt.GetToken(accessTokenClaims, "DhY5bNeb5n7R9DAqizaDB7klaUyDqHlV")
}

func refreshToken(user *dal.User) string {
	content := getContent(user.Id, false)
	token, err := cipher.Encrypt([]byte("wHnAg4MHONLLQj66YFpQGDnRddbZPPc6"), content)
	if err != nil {
		panic(err)
	}
	return token
}
func getContent(userId int, isSuperUser bool) []byte {
	claims := generation.RefreshTokenClaims{
		UserId:            userId,
		TTL:               604800,
		CreationTimestamp: time.Now().UTC().Unix(),
		US:                isSuperUser,
	}
	contentJson, err := json.Marshal(claims)
	if err != nil {
		panic(err) // TODO don't fail
	}
	return contentJson
}
