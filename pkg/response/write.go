package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Write(c *gin.Context, res InterfaceResponse) {
	if res.error() != nil {
		c.JSON(res.error().getHTTP(), res.error().getResult())
		return
	}
	c.JSON(http.StatusOK, res.data().getResult())
}
