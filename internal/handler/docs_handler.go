package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ScalarHTML is the Scalar API Reference UI page that loads from CDN.
const ScalarHTML = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Mintrix API Docs</title>
</head>
<body style="margin:0">
  <script
    id="api-reference"
    data-url="/docs/openapi.yaml"
  ></script>
  <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
</body>
</html>`

// DocsHandler serves the Scalar API documentation UI and the OpenAPI spec.
func DocsHandler(router *gin.Engine) {
	router.GET("/docs", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, ScalarHTML)
	})

	// Serve the OpenAPI YAML spec so Scalar can fetch it.
	router.StaticFile("/docs/openapi.yaml", "./docs/openapi.yaml")
}
