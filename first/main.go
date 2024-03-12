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
	log.Printf("Duck db backend : %+v", duck)

	router := gin.Default()

	if true {
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{"*"} //app.AllowOrigins
		config.AllowHeaders = []string{"Origin"}
		//config.AllowCredentials = app.AllowCredentials
		router.Use(cors.New(config))
	}

	//Setup route group for the API v1
	v1 := router.Group("/v1")
	{
		v1.POST("/exec", duck.ExecCtlr)

		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Hello from datatask.io api proxy"})
		})

	}

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
	log.Printf("[INFO] Request body received : %+v", string(body))

	resp, err := d.runQuery(string(body))
	//log.Println(string(resp))
	if err != nil {
		log.Printf("[ERROR] Reading request %+v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	//c.JSON(200, gin.H{"message": fmt.Sprintf("Request body received : %+v", string(body))})
	//c.JSON(200, gin.H{"query": string(body), "response": resp})
	//c.HTML(http.StatusInternalServerError, resp, nil)
	c.Data(200, "txt/json", resp)
}

func (d *Duck) runQuery(query string) ([]byte, error) {
	log.Printf("[INFO] Processing : %s", query)
	duckCmd := exec.Command("duckdb", "-json", ":memory:", query)
	duckOut, err := duckCmd.Output()
	if err != nil {
		return []byte{}, err
	}
	/*var resp []interface{}
	err = json.Unmarshal(duckOut, &resp)
	if err != nil {
		panic(err)
	}*/

	return duckOut, nil

}
