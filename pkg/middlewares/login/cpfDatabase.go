package login

import (
	"fmt"

	"gorm.io/gorm"
	"microservicos.com/cmd/config"
	entites "microservicos.com/internal/entities"
)

// colocar senha aqui
func CpfDatabase(email string) error {
	db := config.GetDatabase()
	var p entites.Category
	fmt.Println("procurando usuario")

	fmt.Println(email)
	if err := db.Where("email = ?", email).First(&p).Error; err != nil {
		fmt.Printf("nome do email %s", email)
		if err == gorm.ErrRecordNotFound {
			err := err
			return err
		} else {
			err := err
			return err
		}

	}
	return nil
}
