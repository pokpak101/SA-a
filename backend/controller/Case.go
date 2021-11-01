package controller

import (
	"github.com/pokpak101/sa-64-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /users

func CreateCase(c *gin.Context) {

	var cases entity.Case

	if err := c.ShouldBindJSON(&cases); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&cases).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": cases})

}

// GET /user/:id

func GetCase(c *gin.Context) {

	var cases entity.Case

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM cases WHERE id = ?", id).Scan(&cases).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": cases})

}

// GET /users

func ListCase(c *gin.Context) {

	var cases []entity.Case

	if err := entity.DB().Raw("SELECT * FROM cases").Scan(&cases).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": cases})

}

// DELETE /users/:id

func DeleteCase(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM cases WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "cases not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /users

func UpdateCase(c *gin.Context) {

	var cases entity.Case

	if err := c.ShouldBindJSON(&cases); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", cases.ID).First(&cases); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "cases not found"})

		return

	}

	if err := entity.DB().Save(&cases).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": cases})

}
