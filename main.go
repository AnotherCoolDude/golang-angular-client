package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	token, err := receiveToken()
	if err != nil {
		fmt.Println("token could not be received")
		fmt.Println(err)
		return
	}
	todos, err := getTodos(token)
	if err != nil {
		fmt.Println("todos could not be received")
		fmt.Println(err)
		return
	}
	fmt.Println(todos)
}

func receiveToken() (string, error) {
	url := "https://dev-3tt1ae45.eu.auth0.com/oauth/token"

	payload := strings.NewReader("{\"client_id\":\"Vi72HiE8p3clPpdji4L64i0axSba3q1u\",\"client_secret\":\"ONmN5BSndMCyPCHqnswN9d5ycoenQ72-XHmmgqr4WfqzqVAPbwCpaAU-ilTKNWq5\",\"audience\":\"https://my-golang-api\",\"grant_type\":\"client_credentials\"}")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Println("error creating request")
		return "", err
	}

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error sending request")
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("error reading response")
		return "", err
	}

	fmt.Println(res)
	fmt.Println()
	fmt.Println(string(body))
	return string(body), nil
}

func getTodos(bearerToken string) (string, error) {
	url := "http://localhost:3000/todo"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("error creating request")
		return "", err
	}

	req.Header.Add("authorization", "Bearer "+bearerToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error sending request")
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("error reading response")
		return "", err
	}
	fmt.Println(req)
	fmt.Println(res)
	fmt.Println(string(body))
	return string(body), nil
}
