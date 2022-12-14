package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// @title Hello World API
// @description This is a sample server for a hello world API.
// @version 1.0.0
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", helloWorldHandler)
	router.HandleFunc("/", userPostHandler).Methods("POST")
	router.HandleFunc("deneme", denemeHandler)
	http.ListenAndServe(":8080", router)
}

// @Summary Hello World
// @Description Hello World Description
// @Tags User Service
// @Router / [get]
// @Success 200 {string} string "Hoş Geldin Dünyaya"
// @Failure 400 {string} string "Hata"
// @Failure 500 {string} string "Sunucu Hatası"
// @Failure default {string} string "Bilinmeyen Hata"
func helloWorldHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hoş Geldin Dünyaya"))
}

// @Summary User Post
// @Description User Post Description
// @Tags User Service
// @Router / [post]
// @Accept json
// @Produce json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {string} string "Hoş Geldin"
// @Failure 400 {string} string "Hata"
// @Failure 500 {string} string "Sunucu Hatası"
func userPostHandler(writer http.ResponseWriter, request *http.Request) {
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Hata"))
		return
	}
	writer.Write([]byte("Hoş Geldin" + requestBody.Username))
}

// @Summary Zamanı gösteren endpoint
// @Description Zamanı gösteren endpoint
// @Tags User Service
// @Router /deneme [get]
// @Success 200 {string} string "2022-12-30 12:50:00"
// @Failure 400 {string} string "Hata"
// @Failure 500 {string} string "Sunucu Hatası"
// @Failure default {string} string "Bilinmeyen Hata"
func denemeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
}
