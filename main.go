package main

import (
	"log"
	"net/http"
	"os/exec"
	"regexp"

	"github.com/gin-gonic/gin"
)

func main() {
	Serve := gin.Default()
	//gin.SetMode(gin.ReleaseMode)

	Serve.GET("/doc", func(c *gin.Context) {
		pkg := c.Query("package")
		sub := c.Query("subdir")
		fn := c.Query("function")

		queries := c.Request.URL.Query()
		valid := true
		for _, v := range queries {
			if regexpAZ(v[0]) != true {
				valid = false
				break
			}
		}

		// Figure out how to successfully run godoc with missing arguments without have a condition for ever possible combination.
		if valid == false {
			c.String(http.StatusBadRequest, "%s", "Queries may only contain letters\n")
		} else {
			if sub != "" {
				arg := 
				cmd := exec.Command("/usr/local/go/bin/godoc", pkg+"/"+sub, fn)
				cmdOut, cmdErr := cmd.CombinedOutput()
				if cmdErr != nil {
					log.Fatal(cmdErr)
				}
				c.String(http.StatusOK, "%s", cmdOut)
			} else {
				cmd := exec.Command("/usr/local/go/bin/godoc", pkg, fn)
				cmdOut, cmdErr := cmd.CombinedOutput()
				if cmdErr != nil {
					log.Fatal(cmdErr)
				}
				c.String(http.StatusOK, "%s", cmdOut)
			}
		}
	})

	Serve.Run(":8080")
}

func regexpAZ(str string) bool {
	if valid, _ := regexp.MatchString("^[a-zA-Z]+$", str); !valid {
		return false
	}
	return true
}
