package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/softclub-go-0-0/post-service/pkg/models"
	"gorm.io/gorm"
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
	var post models.Post
	if err := h.DB.Preload("ProcessingResult").First(&post, c.Param("postID")).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("post nor found:", err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "There is no such post",
			})
			return
		}
		log.Println("getting a post from DB:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *handler) SearchPosts(c *gin.Context) {
	var posts []models.Post
	if err := h.DB.Where("title ilike ?", "%"+c.Param("title")+"%").Find(&posts).Error; err != nil {
		log.Println("searching posts in DB:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, posts)
}
