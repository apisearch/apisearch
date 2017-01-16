package settings

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

type NewUser struct {
	UserId string `json:"userId"`
	Token  string `json:"token"`
}

const tokenLen int = 48

func generateToken() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, tokenLen)
	rand.Seed(time.Now().UnixNano())

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

//func verifyPassword() {
// err = bcrypt.CompareHashAndPassword(hashedPassword, password)
// fmt.Println(err) // nil means it is a match
//}
