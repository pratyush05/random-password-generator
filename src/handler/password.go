package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"random-password-generator/src/param"
)

// Password parses incoming request on path /password and returns response
func Password(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Request received: %+v\n\n", request)

	switch request.Method {
	case http.MethodGet:
		response, err := Get(request)
		if err != nil {
			err := Error(http.StatusInternalServerError, err.Error())
			errJSON, _ := json.Marshal(err)
			writer.Header().Set(param.ContentType, "application/json")
			writer.Write(errJSON)
			return
		}
		responseJSON, err := json.Marshal(response)
		if err != nil {
			err := Error(http.StatusInternalServerError, err.Error())
			errJSON, _ := json.Marshal(err)
			writer.Header().Set(param.ContentType, "application/json")
			writer.Write(errJSON)
			return
		}
		writer.Header().Set(param.ContentType, "application/json")
		writer.Write(responseJSON)
	default:
		err := Error(http.StatusNotImplemented, "Method not implemented")
		errJSON, _ := json.Marshal(err)
		writer.Header().Set(param.ContentType, "application/json")
		writer.Write(errJSON)
	}
	return
}

// ErrorResponse is schema for error response
type ErrorResponse struct {
	ErrorCode    int32  `json:"error_code"`
	ErrorMessage string `json:"error_message,omitempty"`
}

// Error returns error response
func Error(errorCode int, errorMessage string) *ErrorResponse {
	return &ErrorResponse{
		ErrorCode:    int32(errorCode),
		ErrorMessage: errorMessage,
	}
}

// GetResponse is schema for get response
type GetResponse struct {
	Password       string `json:"random_password"`
	PasswordLength int32  `json:"length"`
}

// Get returns a random generated password and its length
func Get(request *http.Request) (*GetResponse, error) {
	var passwordLength int32 = 8
	lengthQueryParam, ok := request.URL.Query()["length"]
	if ok && len(lengthQueryParam) >= 1 {
		parsedPasswordLength, err := strconv.ParseInt(lengthQueryParam[0], 10, 32)
		if err != nil || parsedPasswordLength <= 0 {
			passwordLength = 8
		} else if parsedPasswordLength > 64 {
			passwordLength = 64
		} else {
			passwordLength = int32(parsedPasswordLength)
		}
	}

	var alphaNum bool = false
	alphaNumQueryParam, ok := request.URL.Query()["alphaNum"]
	if ok && len(alphaNumQueryParam) >= 1 {
		parsedAlphaNum, err := strconv.ParseBool(alphaNumQueryParam[0])
		if err != nil {
			alphaNum = false
		} else {
			alphaNum = parsedAlphaNum
		}
	}

	acceptableLetters := []byte(param.AcceptableLetters)
	if alphaNum {
		acceptableLetters = []byte(param.AcceptableLettersAlphaNum)
	}
	length := len(acceptableLetters)

	var buffer bytes.Buffer
	for i := int32(0); i < passwordLength; i++ {
		randomNumber := rand.Intn(length)
		randomLetter := acceptableLetters[randomNumber]
		buffer.WriteByte(randomLetter)
	}

	password := buffer.String()
	if password == "" {
		return nil, errors.New("Could not generate password")
	}

	return &GetResponse{
		Password:       password,
		PasswordLength: passwordLength,
	}, nil
}
