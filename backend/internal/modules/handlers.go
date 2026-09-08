package modules

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Db *sql.DB
}

func NewHandlers(db *sql.DB) *Handler {
	return &Handler{Db: db}
}

// tmp endpoints to just push data
func (h *Handler) AddModules(c *gin.Context) {

}

func (h *Handler) AddTopics(c *gin.Context) {

}

func (h *Handler) GetModules(c *gin.Context) {

}

func (h *Handler) GetModule(c *gin.Context) {

}

func GetTopics(c *gin.Context) {

}

func GetTopic(c *gin.Context) {

}
