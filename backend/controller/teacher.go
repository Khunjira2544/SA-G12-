package controller

import (
	"net/http"

	"github.com/Khunjira2544/sa-65-project/entity"
	"github.com/gin-gonic/gin"
)

// POST /videos      //ของเรา teacher  Teacher
func CreateTeacher(c *gin.Context) {
	var teacher entity.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&teacher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": teacher})
}

// GET /video/:id      //ของเรา teacher  Teacher
func GetTeacher(c *gin.Context) {
	var teacher entity.Teacher

	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&teacher); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "teacher not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": teacher})
}

// GET /videos         //ของเรา teacher  Teacher
func ListTeachers(c *gin.Context) {
	var teachers []entity.Teacher
	if err := entity.DB().Raw("SELECT * FROM teachers").Find(&teachers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": teachers})
}

// //// อันนี้ลองดูดีๆ  สาขาวิชา  ข้างล่างยังไม่ได้ทำต่อ  ตรง owner_id ต้องได้แก้อีกทีว่ามันคืออะไร
/*func ListT_Facultys(c *gin.Context) {
	owner_id := c.Param("owner_id") //
	var teachers []entity.Teacher
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM teacher WHERE owner_id=?", owner_id).Find(&teachers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": teachers})
}*/

// DELETE /videos/:id     //ของเรา teacher  Teacher
func DeleteTeacher(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM teachers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "teacher not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /videos         //ของเรา teacher  Teacher
func UpdateTeacher(c *gin.Context) {
	var teacher entity.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", teacher.ID).First(&teacher); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "teacher not found"})
		return
	}

	if err := entity.DB().Save(&teacher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": teacher})
}
