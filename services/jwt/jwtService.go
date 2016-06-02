package servicesJWT

import (
	"errors"
	"fmt"
	"time"

	r "github.com/dancannon/gorethink"
	"github.com/dchest/uniuri"
	"github.com/dgrijalva/jwt-go"
	"github.com/nfrush/Go-MarketPlace/database"
	"github.com/nfrush/Go-MarketPlace/models/jwt"
	"github.com/nfrush/Go-MarketPlace/models/user"
)

var session = db.GetSession()

//signingKey - Signing Key For Cookies
var signingKey = InitSigningKey()

//InitSigningKey - Initalize Our Key To Sign With
func InitSigningKey() string {
	return uniuri.NewLen(32)
}

//GetSigningKey - get the current signing key
func GetSigningKey() string {
	return signingKey
}

//IssueToken - Issue New JWT Token
func IssueToken(u *modelUser.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["iss"] = "Aurelia Development LTD"
	token.Claims["aud"] = u.Name
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token.Claims["jti"] = "http://example.com"

	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	issuedToken := modelJWT.JWT{Token: tokenString, Issuer: "Aurelia Development LTD", Audience: u.Name, IssuedAt: time.Now().Unix(), Expires: time.Now().Add(time.Hour * 72).Unix(), JTI: "http://example.com"}

	if err := r.Table("tokens").Insert(issuedToken).Exec(session); err != nil {
		return "", err
	}
	fmt.Println("Issued Token Successfully")
	return tokenString, nil
}

//RevokeToken - Revoke the JWT Token
func RevokeToken(u *modelUser.User) error {
	result, err := r.Table("tokens").Filter(map[string]interface{}{"Audience": u.Name}).Run(session)
	if err != nil {
		return err
	}
	var transformToken modelJWT.JWT
	result.One(&transformToken)
	result.Close()

	if err := r.Table("tokens").Filter("Audience: u.Name").Delete().Exec(session); err != nil {
		return err
	}

	if err := r.Table("blacklist").Insert(&transformToken).Exec(session); err != nil {
		return err
	}

	return nil
}

//RefreshToken - Reissue a new token
func RefreshToken(u *modelUser.User) (string, error) {
	if err := RevokeToken(u); err != nil {
		return "error", err
	}
	token, errB := IssueToken(u)
	if errB != nil {
		return "error", errB
	}
	return token, nil
}

//TokenExists - Check if Token Exists
func TokenExists(token string) (bool, error) {
	if err := r.Table("tokens").Filter(map[string]interface{}{"Token": token}).Exec(session); err != nil {
		return false, err
	}
	return true, nil
}

//TokenExistsUser - Checks if a user has an assigned token
func TokenExistsUser(u *modelUser.User) (bool, error) {
	if err := r.Table("tokens").Filter(map[string]interface{}{"Audience": u.Name}).Exec(session); err != nil {
		return false, err
	}
	return true, nil
}

//RequiresAuth - Authenicates user on API Route
func RequiresAuth(token string) (bool, error) {
	exists, err := TokenExists(token)
	if err != nil {
		return false, err
	}
	if exists {
		res, err := r.Table("tokens").Filter(map[string]interface{}{"Token": token}).Run(session)
		if err != nil {
			return false, err
		}
		var transformToken modelJWT.JWT
		res.One(&transformToken)
		res.Close()

		resu, err := r.Table("users").Filter(map[string]interface{}{"Name": transformToken.Audience}).Run(session)
		if err != nil {
			return false, err
		}
		var user modelUser.User
		resu.One(&user)
		resu.Close()

		if transformToken.Expires <= time.Now().Unix() {
			return true, nil
		}
		if transformToken.Expires > time.Now().Unix() {
			RevokeToken(&user)
			return false, errors.New("Token has expired and been revoked.")
		}
	}
	return false, errors.New("The Token Does Not Exist")
}
