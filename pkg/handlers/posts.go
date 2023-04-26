package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/softclub-go-0-0/post-service/pkg/models"
	"log"
	"net/http"
)

func (h *handler) GetAllPosts(c *gin.Context) {
	var posts []models.Post
	if err := h.DB.Find(&posts).Error; err != nil {
		log.Println("getting posts from DB:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (h *handler) GetOnePost(c *gin.Context) {
}

func (h *handler) SearchPosts(c *gin.Context) {

}
