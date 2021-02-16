package polo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	polo = "polo"
)

//PoloController
func PoloController(c *gin.Context) {
	c.String(http.StatusOK, polo)
}
