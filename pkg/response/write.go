package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Write(c *gin.Context, res *Response) {
	if res.HTTP != http.StatusOK {
		c.JSON(res.HTTP, res.Result)
		return
	}
	c.JSON(http.StatusOK, res.Result)
}
