package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestEchoGetRequest(t *testing.T) {
	r := gofight.New()
	r.GET("/").
		SetDebug(true).
		Run(EchoEngine(60, nil), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			data := []byte(r.Body.String())
			requestNum := gjson.GetBytes(data, "requestNum")
			assert.Equal(t, "1", requestNum.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestEchoGetRequestExceed(t *testing.T) {
	limit = 0
	r := gofight.New()
	r.GET("/").
		SetDebug(true).
		Run(EchoEngine(60, nil), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			data := []byte(r.Body.String())
			log.Println(string(data))
			message := gjson.GetBytes(data, "message")
			assert.Equal(t, "Error, request limit exceeded", message.String())
			assert.Equal(t, http.StatusTooManyRequests, r.Code)
		})
}
