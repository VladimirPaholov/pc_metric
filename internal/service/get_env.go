package service

import (
	"fmt"

	"github.com/joho/godotenv"
)

// Загрузка файла переменных окружения
func GetENV() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("File .env not found")
	}

}
