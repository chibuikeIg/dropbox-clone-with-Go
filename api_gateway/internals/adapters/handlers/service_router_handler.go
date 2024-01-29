package handlers

import (
	"api-gateway/internals/app"
	consul_api "api-gateway/internals/app/consul"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceRouterHandler struct{}

func NewServiceRouterHandler() *ServiceRouterHandler {
	return &ServiceRouterHandler{}
}

func (sr *ServiceRouterHandler) RouteRequests(c *gin.Context) {

	url, statusCode := consul_api.ServiceDiscovery(c.Param("service"), c.Param("req_path"))
	queryParams := c.Request.URL.Query()

	if len(queryParams) > 0 {
		url += "?"
		i := 0

		for k, vals := range queryParams {
			for _, val := range vals {
				if i > 0 {
					url += "&"
				}
				url += k + "=" + val
				i++
			}
		}
	}

	// check for error status code
	if statusCode == http.StatusNotFound {
		c.JSON(statusCode, gin.H{
			"message": "Not Found",
		})
		return
	} else if statusCode == http.StatusInternalServerError {
		c.JSON(statusCode, gin.H{
			"message": "Internal server error occured",
		})
		return
	}

	/// setup/send requests
	headers := c.Request.Header

	headers.Del("Authorization")

	resp, respStatusCode, err := app.SendRequest(c.Request.Method, url, c.Request.Body, headers)

	if err != nil {

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error occured",
		})
		return
	}

	c.JSON(respStatusCode, resp)
}
