package modules

import (
	"database/sql"
	"net/http"

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
	var body Module
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Invalid body passed"})
	}
	AddModules(h.Db, body)
}

func (h *Handler) AddTopics(c *gin.Context) {
	var body Topic
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Invalid body passed"})
	}
	AddTopics(h.Db, body)
}

// production endpoints start here
func (h *Handler) GetModules(c *gin.Context) {
	
}

func (h *Handler) GetModule(c *gin.Context) {

}

func GetTopics(c *gin.Context) {

}

func GetTopic(c *gin.Context) {

}
