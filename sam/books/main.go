package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type appParams struct {
	Password string `json:"Password"`
	Field2   string `json:"Field2"`
	Field3   string `json:"Field3"`
}
type book struct {
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var isbnRegexp = regexp.MustCompile(`[0-9]{3}\-[0-9]{10}`)
var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)
var awsregion, secretName = os.Getenv("AWSRegion"), os.Getenv("secretName")

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return show(req)
	case "POST":
		return create(req)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

func show(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	/* 	bk := &book{
		ISBN:   "978-1420931693",
		Title:  "The Republic",
		Author: "Plato",
	} */
	//var decodedBinarySecret, secretString string
	fmt.Println("Executing function v 0.0.1")
	fmt.Printf("Execuing Function with the Region variable set to [%s] \n", awsregion)
	fmt.Println("Getting Secret Trough Functions")

	secretString, err := getSecret()
	if err != nil {
		return serverError(err)
	}

	fmt.Println("Function Executed")

	if secretString == "" {
		fmt.Println("decodedBinarySecret is empty")
		return serverError(err)
	}
	secret := new(appParams)
	err = json.Unmarshal([]byte(secretString), &secret)
	if err != nil {
		return clientError(http.StatusUnprocessableEntity)
	}

	fmt.Printf("Returning my secret:[ %s ] \n", secret.Password)
	fmt.Println("Populating ISBN")
	isbn := request.QueryStringParameters["isbn"]
	if !isbnRegexp.MatchString(isbn) {
		return clientError(http.StatusBadRequest)
	}
	if isbn == "" {
		return clientError(http.StatusBadRequest)
	}
	fmt.Println("Preparign  to send a get request")
	bk, err := getItem(isbn)
	if err != nil {
		return serverError(err)
	}
	if bk == nil {
		return clientError(http.StatusNotFound)
	}

	// The APIGatewayProxyResponse.Body field needs to be a string, so
	// we marshal the book record into JSON.
	js, err := json.Marshal(bk)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

func create(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.Headers["content-type"] != "application/json" && req.Headers["Content-Type"] != "application/json" {
		return clientError(http.StatusNotAcceptable)
	}

	bk := new(book)
	err := json.Unmarshal([]byte(req.Body), bk)
	if err != nil {
		return clientError(http.StatusUnprocessableEntity)
	}

	if !isbnRegexp.MatchString(bk.ISBN) {
		return clientError(http.StatusBadRequest)
	}
	if bk.Title == "" || bk.Author == "" {
		return clientError(http.StatusBadRequest)
	}

	err = putItem(bk)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Location": fmt.Sprintf("/books?isbn=%s", bk.ISBN)},
	}, nil
}

// Add a helper for handling errors. This logs any error to os.Stderr
// and returns a 500 Internal Server Error response that the AWS API
// Gateway understands.
func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

// Similarly add a helper for send responses relating to client errors.
func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func main() {
	lambda.Start(router)
}
