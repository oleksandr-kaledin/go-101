package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	fmt.Printf("c: %v\n", c)
}

func (h *Handler) signIn(c *gin.Context) {

}
