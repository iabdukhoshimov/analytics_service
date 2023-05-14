package handlers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

var (
	errorBadRequest       = errors.New("bad request")
	errorInternal         = errors.New("internal error")
	errorWrongCredentials = errors.New("wrong credentials")
)

func (h *handler) handleError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	switch err {
	case errorBadRequest:
		c.JSON(400, gin.H{"error": err.Error()})
	case errorWrongCredentials:
		c.JSON(401, gin.H{"error": err.Error()})
	case errorInternal:
		c.JSON(500, gin.H{"error": err.Error()})
	default:
		c.JSON(500, gin.H{"error": err.Error()})
	}

	return true
}

func (h *handler) makeContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), h.config.Project.Timeout)
	return ctx, cancel
}
