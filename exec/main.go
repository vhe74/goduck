package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	duck := Duck{bin: "duckdb"}

	router := gin.Default()

	// configure cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} //app.AllowOrigins
	config.AllowHeaders = []string{"Origin"}
	//config.AllowCredentials = app.AllowCredentials
	router.Use(cors.New(config))

	// Define routes
	router.POST("/exec", duck.ExecCtlr)
	//default route handler
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from datatask.io api proxy"})
	})

	router.Run(":8080")
}

type Duck struct {
	bin string
}

func (d *Duck) ExecCtlr(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[ERROR] Reading request %+v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error(),
			"data": nil,
		})
		return
	}
	query := string(body)
	resp, err := d.runQuery(query)
	if err != nil {
		log.Printf("[ERROR] Processing request %+v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.Data(200, "application/json", d.buildJSONOutput(query, resp))
}

func (d *Duck) runQuery(query string) ([]byte, error) {
	log.Printf("[INFO] Processing with duckdb query : %s", query)
	duckCmd := exec.Command("duckdb", "-json", ":memory:", query)
	duckOut, err := duckCmd.Output()
	if err != nil {
		return []byte{}, err
	}

	return duckOut, nil
}

func (d *Duck) buildJSONOutput(query string, resp []byte) []byte {
	returned := []byte("{ \"status\":\"ok\", \"query\":\"")
	returned = append(returned, []byte(query)...)
	returned = append(returned, "\" ,\"data\": "...)
	if len(resp) == 0 {
		returned = append(returned, []byte("[]")...)
	} else {
		returned = append(returned, resp...)
	}
	returned = append(returned, []byte("}")...)

	return returned
}
