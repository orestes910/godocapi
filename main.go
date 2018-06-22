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

	Serve.GET("/:package", func(c *gin.Context) {
		pkg := c.Param("package")
		valid, _ := regexp.MatchString("^[a-zA-Z]+$", pkg)
		if valid {
			pkgCmd := exec.Command("/usr/local/go/bin/godoc", pkg)
			pkgOut, pkgErr := pkgCmd.CombinedOutput()
			if pkgErr != nil {
				log.Fatal(pkgErr)
			}
			c.String(http.StatusOK, "%s", pkgOut)
		} else {
			c.String(http.StatusBadRequest, "%s", "Package may only contain letters\n")
		}
	})

	Serve.Run(":8080")
}
