package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	url := "https://api.wazzpay.com/api/v1/user/auth/login"
	method := "POST"

	payload := strings.NewReader(`{
	  "email": "w@wazzpay.com",
	  "password": "abc123456"
  	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
