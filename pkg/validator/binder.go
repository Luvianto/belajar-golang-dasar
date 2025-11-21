package validator

import (
	"belajar-golang-dasar/pkg/handler"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func BindUri(c *gin.Context, req interface{}) (ok bool) {
	if err := c.ShouldBindUri(req); err != nil {
		handler.Error(c, http.StatusBadRequest, err.Error())
		log.Error().Err(err).Msg("Invalid uri data")
		return false
	}

	if err := requestValidator(c, req, "uri"); err != nil {
		return false
	}

	return true
}
