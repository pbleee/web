package controllers

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pbleee/backend/config"
	"github.com/pbleee/backend/models"
	"golang.org/x/crypto/bcrypt"
)

// regis
func Regis(c fiber.Ctx) error {
	user := new(models.User)

	if err := c.Bind().JSON(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "format tidak valid",
		})
	}

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashPass)

	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "gagal menyimpan data ke database",
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "berhasil registrasi",
		"data":    user,
	})
}

// login
func Login(c fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind().JSON(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "format tidak valid",
		})
	}
	var user models.User

	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "email tidak terdaftar",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "password salah",
		})
	}

	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("rahasia"))

	return c.JSON(fiber.Map{
		"message": "berhasil login",
		"token":   t,
	})
}

// course
func GetCourse(c fiber.Ctx) error {
	var course []models.Course
	config.DB.Find(&course)
	return c.JSON(fiber.Map{
		"data": course,
	})
}

func CreateCourse(c fiber.Ctx) error {
	var course models.Course

	if err := c.Bind().Body(&course); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "input tidak valid",
		})
	}
	config.DB.Create(&course)
	return c.Status(201).JSON(fiber.Map{
		"message": "berhasil membuat",
		"data":    course,
	})
}

func UpdateCourse(c fiber.Ctx) error {
	var course models.Course
	id := c.Params("id")

	if err := config.DB.First(&course, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "data tidak ditemukan",
		})
	}

	if err := c.Bind().Body(&course); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "input tidak valid",
		})
	}
	config.DB.Save(&course)
	return c.JSON(fiber.Map{
		"message": "berhasil diperbarui",
		"data":    course,
	})
}

func DelCourse(c fiber.Ctx) error {
	var course models.Course
	id := c.Params("id")

	result := config.DB.First(&course, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "data tidak ditemukan",
		})
	}
	if err := config.DB.Delete(&course).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "gagal menghapus",
		})
	}
	return c.JSON(fiber.Map{
		"message": "berhasil menghapus",
	})
}
