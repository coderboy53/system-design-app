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
	var body []Module
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body passed"})
	}
	for _, module := range body {
		err := AddModules(h.Db, module)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error occurred while adding module"})
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Added all modules successfully"})
}

func (h *Handler) AddTopics(c *gin.Context) {
	var body []Topic
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body passed"})
	}
	for _, topic := range body {
		err := AddTopics(h.Db, topic)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error occurred while adding topic"})
			break
		}
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Added all topics successfully"})
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
