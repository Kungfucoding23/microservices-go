package polo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	polo = "polo"
)

//MarcoController ..
func MarcoController(c *gin.Context) {
	c.String(http.StatusOK, polo)
}
