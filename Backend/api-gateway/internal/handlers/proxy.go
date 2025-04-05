package handlers

import (
	"net/http/httputil"
	"net/url"

	"api-gateway/internal/config"

	"github.com/gin-gonic/gin"
)

func ProxyToInventoryService(cfg *config.Config) gin.HandlerFunc {
	return createReverseProxy(cfg.InventoryServiceURL)
}

func ProxyToOrderService(cfg *config.Config) gin.HandlerFunc {
	return createReverseProxy(cfg.OrderServiceURL)
}

func createReverseProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetURL, _ := url.Parse(target)
		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		// Modify request
		c.Request.URL.Path = c.Param("path")
		c.Request.Host = targetURL.Host

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
