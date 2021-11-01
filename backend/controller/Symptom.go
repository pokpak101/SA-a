package controller

import (
	"github.com/pokpak101/sa-64-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /users

func CreateSymptom(c *gin.Context) {

	var symptom entity.Symptom

	if err := c.ShouldBindJSON(&symptom); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&symptom).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": symptom})

}

// GET /user/:id

func GetSymptom(c *gin.Context) {

	var symptom entity.Symptom

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM symptoms WHERE id = ?", id).Scan(&symptom).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": symptom})

}

// GET /users

func ListSymptom(c *gin.Context) {

	var symptom []entity.Symptom

	if err := entity.DB().Raw("SELECT * FROM symptoms").Scan(&symptom).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": symptom})

}

// DELETE /users/:id

func DeleteSymptom(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM symptoms WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "symptoms not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /users

func UpdateSymptom(c *gin.Context) {

	var symptom entity.Symptom

	if err := c.ShouldBindJSON(&symptom); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", symptom.ID).First(&symptom); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "symptoms not found"})

		return

	}

	if err := entity.DB().Save(&symptom).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": symptom})

}
