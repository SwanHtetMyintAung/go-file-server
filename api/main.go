package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const PATH_NAME = "../data"

// type FileReponse struct {
// 	Name string
// }

// check if the folder exist, if not create one
// true for successfully creating folder OR folder already exist
func OpInit(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Folder Doesn't exist and Creating One.")
			os.Mkdir(path, 0755)
			return true, nil
		} else {
			fmt.Println("Error :", err)
			return false, err
		}
	}
	result := info.IsDir()

	if result == true {
		fmt.Println("Folder already exists")
		return true, nil
	} else {
		//case : the result is not a dir but a file
		os.Mkdir(path, 0755)
		return true, nil
	}
}
func getTheDirList(c *gin.Context) {
	files, err := os.ReadDir("../data")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't get the file list",
		})
	}
	var names = []string{}

	for _, file := range files {
		names = append(names, file.Name())

	}
	c.IndentedJSON(http.StatusOK, names)
}
func fileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"message": err,
		})
		return
	}
	saveErr := c.SaveUploadedFile(file, PATH_NAME+"/"+file.Filename)
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
	result, err := OpInit(PATH_NAME)
	if err != nil || result == false {
		fmt.Println("Couldn't Initiate The Folder.")
		return
	}

	fmt.Println("Hello World")

	server := gin.Default()
	server.Use(cors.Default()) // for dev only
	server.GET("all-file", getTheDirList)
	server.POST("/file", fileUpload)

	server.Run("localhost:3000")
}
