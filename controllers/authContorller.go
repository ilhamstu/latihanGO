package controllers

import (
	"net/http"
	"restapi/be/models"
	"restapi/be/utils"
	"restapi/be/validators"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var loginRequest validators.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	if errorMsgs, err := validators.ValidateLoginRequest(loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMsgs,
		})
		return
	}

	var user models.User
	var mahasiswa models.Mahasiswa

	// Mencari user berdasarkan email di tabel User
	if err := models.DB.Where("email = ?", loginRequest.Email).First(&user).Error; err == nil {
		// Jika ditemukan, memeriksa password di tabel User
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// Generate access dan refresh token untuk User
		accessToken, err := utils.GenerateAccessToken(user.Id, user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate access token"})
			return
		}

		refreshToken, err := utils.GenerateRefreshToken(user.Id, user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate refresh token"})
			return
		}

		// Return response dengan access token dan refresh token
		c.JSON(http.StatusOK, gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
		return
	}

	// Jika tidak ditemukan di tabel User, mencari mahasiswa berdasarkan email
	if err := models.DB.Where("email = ?", loginRequest.Email).First(&mahasiswa).Error; err == nil {
		// Jika ditemukan, memeriksa password di tabel Mahasiswa
		if err := bcrypt.CompareHashAndPassword([]byte(mahasiswa.Password), []byte(loginRequest.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// Generate access dan refresh token untuk Mahasiswa
		accessToken, err := utils.GenerateAccessToken(mahasiswa.Id, mahasiswa.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate access token"})
			return
		}

		refreshToken, err := utils.GenerateRefreshToken(mahasiswa.Id, mahasiswa.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate refresh token"})
			return
		}

		// Return response dengan access token dan refresh token
		c.JSON(http.StatusOK, gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
}

func RefreshTokenHandler(c *gin.Context) {
	var refreshRequest struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&refreshRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid refresh token"})
		return
	}

	claims, err := utils.ValidateToken(refreshRequest.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	accessToken, err := utils.GenerateAccessToken(claims.UserID, claims.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate new access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}

func RegisterUser(c *gin.Context) {
	var userRequest validators.UserRegisterRequest

	// Bind JSON dari request body ke struct UserRegisterRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	// Validasi request menggunakan validasi yang sudah dibuat
	if errorMsgs, err := validators.ValidateUserRegisterRequest(userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMsgs,
		})
		return
	}

	// Hash password menggunakan bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	input := models.User{
		Nik:      userRequest.Nik,
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: string(hashedPassword), // Menyimpan password dalam bentuk hash
	}
	models.DB.Create(&input)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

// Fungsi untuk registrasi mahasiswa
func RegisterMahasiswa(c *gin.Context) {
	var mahasiswaRequest validators.MahasiswaRegisterRequest

	// Bind JSON dari request body ke struct MahasiswaRegisterRequest
	if err := c.ShouldBindJSON(&mahasiswaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	// Validasi request menggunakan validasi yang sudah dibuat
	if errorMsgs, err := validators.ValidateMahasiswaRegisterRequest(mahasiswaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMsgs,
		})
		return
	}

	// Hash password menggunakan bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(mahasiswaRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	input := models.Mahasiswa{
		Nim:      mahasiswaRequest.Nim,
		Name:     mahasiswaRequest.Name,
		Email:    mahasiswaRequest.Email,
		Password: string(hashedPassword), // Menyimpan password dalam bentuk hash
	}
	models.DB.Create(&input)

	// Jika validasi sukses, lanjutkan proses registrasi
	c.JSON(http.StatusCreated, gin.H{
		"message": "Mahasiswa registered successfully",
	})
}
