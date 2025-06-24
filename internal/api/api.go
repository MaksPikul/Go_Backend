package api

import "github.com/gin-gonic/gin"

func CreateServer() {

	router := gin.Default()

	api := router.Group("/api/v1") // edit this to change version easier

	api.Use()

}

func RunServer() {

}

func UploadImage() {

}

// azure functions for scheduled posts ?
