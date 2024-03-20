package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func ExecPythonCtlr(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[ERROR] Reading request %+v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error(),
			"data": nil,
		})
		return
	}
	code := string(body)
	resp, err := execPythonCode(code)
	if err != nil {
		log.Printf("[ERROR] Processing request %+v", err.Error())
		//c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		c.Data(500, "application/text", []byte(resp))
		//return
	}

	c.Data(200, "application/json", []byte(resp))
}

func execPythonCode(code string) ([]byte, error) {
	err := os.WriteFile("py_stub.py", []byte(code), 0644)
	if err != nil {
		return []byte{}, err
	}

	//pythonCmd := exec.Command("python", "-V") //, ":memory:", code)
	//conda run -n my-python-env python --version
	pythonCmd := exec.Command("conda", "run", "-n", "py312", "python", "py_stub.py")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	pythonCmd.Stdout = &stdout
	pythonCmd.Stderr = &stderr
	err = pythonCmd.Run()
	if err != nil {
		return stderr.Bytes(), err
	}

	return stdout.Bytes(), nil
}
