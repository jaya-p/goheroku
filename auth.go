package goheroku

import (
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

	resp, err := http.Get("https://goherokuauth.herokuapp.com/api/v1/auth?token=" + token)
	if err != nil {
		log.Println("error accessing auth web service")

		return false, errors.New("error accessing auth web service")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("token is invalid")

		return false, errors.New("token is invalid")
	}

	return true, nil
}
