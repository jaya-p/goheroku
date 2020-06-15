package goheroku

import (
	"crypto/tls"
	"errors"
	"log"
	"net/http"
	"os"
)

// CheckToken checks whether token exists
func CheckToken(token string) (bool, error) {
	log.SetOutput(os.Stdout)

	if token == "" {
		return false, errors.New("token is empty")
	}

	//disable ssl verification, due to using SCRATCH docker image
	t := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: t}
	resp, err := client.Get("https://goherokuauth.herokuapp.com/api/v1/auth?token=" + token)
	if err != nil {
		log.Println(err)

		return false, errors.New("error accessing auth web service")
	}
	resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("token is invalid")

		return false, errors.New("token is invalid")
	}

	return true, nil
}
