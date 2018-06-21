package main

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	Serve := gin.Default()
	//gin.SetMode(gin.ReleaseMode)

	Serve.GET("/:package", func(c *gin.Context) {
		pkg := c.Param("package")
		pkgCmd := exec.Command("/usr/local/go/bin/godoc", pkg)
		pkgOut, pkgErr := pkgCmd.CombinedOutput()
		if pkgErr != nil {
			log.Fatal(pkgErr)
		}
		c.String(http.StatusOK, "%s", pkgOut)
	})

	Serve.Run(":8080")
}
