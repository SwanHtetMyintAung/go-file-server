package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Reply struct {
	code    int
	message string
}

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func getAllUsers(c *gin.Context) {
	data, err := os.ReadFile("../data.json")
	if err != nil {
		//i wanna console .log soo bad
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "opps",
		})
		return
	}
	var Users []User
	jsonErr := json.Unmarshal(data, &Users)
	if jsonErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "opps in array",
		})
		panic(jsonErr)
	}
	c.IndentedJSON(http.StatusOK, Users)

}
func fileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"message": "Error in Parsing the file",
		})
		return
	}
	saveErr := c.SaveUploadedFile(file, "./saved/"+file.Filename)
	if saveErr != nil {
		c.IndentedJSON(500, gin.H{
			"message": "Error in Saving the file",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "YESSSSS",
	})

}
func main() {

	fmt.Println("Hello World")

	server := gin.Default()
	server.Use(cors.Default()) // for dev only
	server.GET("/all-users", getAllUsers)
	server.POST("/file", fileUpload)

	server.Run("localhost:3000")
}
