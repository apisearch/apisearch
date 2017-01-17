package settings

import (
	"errors"
	"github.com/apisearch/apisearch/model/elasticsearch"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"math/rand"
	"time"
)

type NewUser struct {
	UserId string `json:"userId"`
	Token  string `json:"token"`
}

type SignInData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const tokenLen int = 48

func SignIn(input SignInData) (NewUser, error) {
	var response NewUser

	found, _, err := findByEmail(input.Email)

	if err != nil {
		return response, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(input.Password))

	if err != nil {
		return response, errors.New("Invalid e-mail or password")
	}

	response.UserId = found.UserId
	response.Token = found.Token

	return response, nil
}

func (s *Settings) SignOut() error {
	client := elasticsearch.CreateClient()

	s.Token = generateToken()

	response, err := client.Index().
		Index(indexName).
		Type(typeName).
		Id(s.UserId).
		BodyJson(s).
		Do(context.TODO())

	if err != nil {
		return err
	}

	if response == nil {
		return errors.New("Empty response from DB")
	}

	return nil
}

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
