package main

import (
    "net/http"
	"bytes"
    "fmt"
    "io"
    "os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
func generateUserIDtest() {
	id := uuid.New()
	fmt.Println(id.String())
}

func getUserIDs(c *gin.Context) {
	old := os.Stdout // keep backup of the real stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    generateUserIDtest()

    outC := make(chan string)
    // copy the output in a separate goroutine so printing can't block indefinitely
    go func() {
        var buf bytes.Buffer
        io.Copy(&buf, r)
        outC <- buf.String()
    }()

    // back to normal state
    w.Close()
    os.Stdout = old // restoring the real stdout
    out := <-outC
    c.IndentedJSON(http.StatusOK, out)
}


func deleteUserID(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, "deleting")
}

func main() {
	router := gin.Default()
	router.GET("/generate-user-ids", getUserIDs)
	router.DELETE("/delete-user-ids", deleteUserID)

	router.Run("localhost:8083")
}