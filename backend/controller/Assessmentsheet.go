package controller

import (
	"github.com/pokpak101/sa-64-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /users

func CreateAssessmentSheet(c *gin.Context) {

	var assessmentsheet entity.AssessmentSheet
	var state entity.State
	var assess entity.Assess
	var symptom entity.Symptom
	var cases entity.Case

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร watchVideo
	if err := c.ShouldBindJSON(&assessmentsheet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", assessmentsheet.CaseID).First(&cases); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Case not found"})
		return
	}

	// 9: ค้นหา video ด้วย id
	if tx := entity.DB().Where("id = ?", assessmentsheet.StateID).First(&state); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "State not found"})
		return
	}

	// 10: ค้นหา resolution ด้วย id
	if tx := entity.DB().Where("id = ?", assessmentsheet.SymptomID).First(&symptom); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Symptom not found"})
		return
	}

	// 11: ค้นหา playlist ด้วย id
	if tx := entity.DB().Where("id = ?", assessmentsheet.AssessID).First(&assess); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Assess not found"})
		return
	}
	ams := entity.AssessmentSheet{
		Case:       cases,
		Symptom:    symptom,
		State:      state,
		Assess:     assess,
		AssessTime: assessmentsheet.AssessTime,
	}
	if err := entity.DB().Create(&ams).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ams})
}

// GET /user/:id

func GetAssessmentSheet(c *gin.Context) {

	var assessmentsheet entity.AssessmentSheet

	id := c.Param("id")

	if err :=
		entity.DB().Preload("Case").Preload("Symptom").Preload("State").Preload("Assess").Raw("SELECT * FROM assessment_sheets WHERE id = ?", id).Find(&assessmentsheet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": assessmentsheet})
}

// GET /users

func ListAssessmentSheet(c *gin.Context) {

	var assessmentsheet []entity.AssessmentSheet

	if err :=
		entity.DB().Preload("Case").Preload("Symptom").Preload("State").Preload("Assess").Raw("SELECT * FROM assessment_sheets").Find(&assessmentsheet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": assessmentsheet})
}

// DELETE /users/:id

func DeleteAssessmentSheet(c *gin.Context) {

	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM assessment_sheets WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "assessment_sheets not found"})
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /users

func UpdateAssessmentSheet(c *gin.Context) {

	var assessmentsheet entity.AssessmentSheet

	if err := c.ShouldBindJSON(&assessmentsheet); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", assessmentsheet.ID).First(&assessmentsheet); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "assessment_sheets not found"})

		return

	}

	if err := entity.DB().Save(&assessmentsheet).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": assessmentsheet})

}
