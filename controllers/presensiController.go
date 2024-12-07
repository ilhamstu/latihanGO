package controllers

import (
	"net/http"
	"restapi/be/models"
	"restapi/be/validators"
	"time"

	"github.com/gin-gonic/gin"
)

func PresensiHandler(c *gin.Context) {
	var presensiRequest validators.PresensiRequest

	if err := c.ShouldBindJSON(&presensiRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	if err := validators.PresensiValidator(presensiRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var mahasiswa models.Mahasiswa

	// Mencari mahasiswa berdasarkan nim di tabel Mahasiswa
	if err := models.DB.Where("nim = ?", presensiRequest.Nim).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa not found"})
		return
	}

	// jika mahasiswa ketemu maka insert record presensinya kedalam table presensi
	presensi := models.Presensi{
		MahasiswaID: mahasiswa.Id,
		StartTime:   time.Now(),
	}
	models.DB.Create(&presensi)

	// Return response dengan data mahasiswa
	c.JSON(http.StatusOK, gin.H{
		"nim":      mahasiswa.Nim,
		"nama":     mahasiswa.Name,
		"check_in": presensi.StartTime,
	})
}
