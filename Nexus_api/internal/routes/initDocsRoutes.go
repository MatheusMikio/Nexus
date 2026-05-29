package routes

import (
	"net/http"
	"os"

	scalar "github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"
)

func initDocsRoutes(rg *gin.RouterGroup) {
	rg.StaticFile("/openapi.json", "./docs/swagger.json")

	rg.GET("/docs", func(ctx *gin.Context) {
		specContent, err := os.ReadFile("./docs/swagger.json")
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecContent: string(specContent),
			DarkMode:    true,
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Nexus API Docs",
			},
		})
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	})
}
