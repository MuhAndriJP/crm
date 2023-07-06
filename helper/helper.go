package helper

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal memuat file .env")
	}
}

func StrToInt64(input string) int64 {
	if input == "" {
		return 0
	}
	result, _ := strconv.Atoi(input)
	return int64(result)
}

func CheckPassword(Existing, Incoming string) error {
	hash := []byte(Existing)
	Password := []byte(Incoming)
	return bcrypt.CompareHashAndPassword(hash, Password)
}

func CryptPassword(password string) string {
	Password := []byte(password)
	cryptedPassword, _ := bcrypt.GenerateFromPassword(Password, bcrypt.DefaultCost)
	return string(cryptedPassword)
}

func GenerateCode(input int64) string {
	inputStr := strconv.Itoa(int(input))
	now := time.Now()

	year := now.Year() % 100
	yearStr := fmt.Sprintf("%02d", year)

	month := int(now.Month())
	monthStr := fmt.Sprintf("%02d", month)

	id := ""
	if input > 0 {
		id = "000" + inputStr
	} else if input > 9 {
		id = "00" + inputStr
	} else if input > 99 {
		id = "0" + inputStr
	} else if input > 999 {
		id = inputStr
	} else {
		id = inputStr
	}

	return yearStr + monthStr + id
}
