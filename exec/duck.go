package main

import (
	"bytes"
	"io"
	"log"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type Duck struct {
	bin string
}

func (d *Duck) ExecCtlr(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[ERROR] Reading request %+v", err.Error())
		/*c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error(),
			"data": nil,
		})*/
		c.Data(500, "application/json", d.buildJSONOutput("error", "", []byte(""), []byte("")))
		return
	}
	query := string(body)
	resp, err := d.runQuery(query)
	if err != nil {
		log.Printf("[ERROR] Processing request %+v", err.Error())
		log.Printf("[ERROR] Output : %s", string(resp))
		//c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		c.Data(500, "application/json", d.buildJSONOutput("error", query, []byte(""), []byte(err.Error())))
		return
	}

	c.Data(200, "application/json", d.buildJSONOutput("ok", query, resp, []byte("")))
}

func (d *Duck) runQuery(query string) ([]byte, error) {
	log.Printf("[INFO] Processing with duckdb query : %s", query)

	duckCmd := exec.Command("duckdb", "-json", ":memory:", query)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	duckCmd.Stdout = &stdout
	duckCmd.Stderr = &stderr

	err := duckCmd.Run()
	if err != nil {
		return stderr.Bytes(), err
	}

	return stdout.Bytes(), nil
}

func (d *Duck) buildJSONOutput(status string, query string, resp []byte, err []byte) []byte {
	returned := []byte("{ \"status\":\"")
	returned = append(returned, []byte(status)...)
	returned = append(returned, []byte("\", \"query\":\"")...)
	returned = append(returned, []byte(query)...)
	returned = append(returned, "\" ,\"data\": "...)
	if len(resp) == 0 {
		returned = append(returned, []byte("[]")...)
	} else {
		returned = append(returned, resp...)
	}
	if len(err) > 0 {
		returned = append(returned, ",\"err\": \""...)
		returned = append(returned, err...)
		returned = append(returned, "\""...)
	}
	returned = append(returned, []byte("}")...)

	return returned
}
