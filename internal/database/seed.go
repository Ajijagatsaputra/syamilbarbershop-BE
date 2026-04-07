package database

import (
	"fmt"
	"syamilbarbershop-be/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func SeedData() {
	// Seed Admin
	var count int64
	DB.Model(&models.User{}).Count(&count)
	if count == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		admin := models.User{
			Name:     "Admin Syamil",
			Email:    "admin@syamil.com",
			Password: string(hashedPassword),
			Role:     "admin",
		}
		DB.Create(&admin)
		fmt.Println("Admin user seeded")
	}

	// Seed Services
	DB.Model(&models.Service{}).Count(&count)
	if count == 0 {
		services := []models.Service{
			{Name: "Gentleman's Haircut", Price: 50000, Duration: 45, Category: models.CategoryHaircut, Status: models.StatusActive},
			{Name: "Kids Haircut", Price: 35000, Duration: 30, Category: models.CategoryHaircut, Status: models.StatusActive},
			{Name: "Beard Trim & Shave", Price: 30000, Duration: 20, Category: models.CategoryGrooming, Status: models.StatusActive},
			{Name: "Hair Styling", Price: 25000, Duration: 15, Category: models.CategoryStyling, Status: models.StatusActive},
		}
		DB.Create(&services)
		fmt.Println("Services seeded")
	}

	// Seed Sample Customer
	DB.Model(&models.Customer{}).Count(&count)
	if count == 0 {
		customer := models.Customer{
			Name:  "John Doe",
			Phone: "08123456789",
			Email: "john@example.com",
		}
		DB.Create(&customer)
		fmt.Println("Sample customer seeded")
	}
}
