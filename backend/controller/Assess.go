package controller

import (
	"github.com/pokpak101/sa-64-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /users

func CreateAssess(c *gin.Context) {

	var assess entity.Assess

	if err := c.ShouldBindJSON(&assess); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&assess).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": assess})

}

// GET /user/:id

func GetAssess(c *gin.Context) {

	var assess entity.Assess

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM assesses WHERE id = ?", id).Scan(&assess).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": assess})

}

// GET /users

func ListAssess(c *gin.Context) {

	var assess []entity.Assess

	if err := entity.DB().Raw("SELECT * FROM assesses").Scan(&assess).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": assess})

}

// DELETE /users/:id

func DeleteAssess(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM assesses WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "assesses not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /users

func UpdateAssess(c *gin.Context) {

	var assess entity.Assess

	if err := c.ShouldBindJSON(&assess); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", assess.ID).First(&assess); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "assesses not found"})

		return

	}

	if err := entity.DB().Save(&assess).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": assess})

}
