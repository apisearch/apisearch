package settings

import (
	"encoding/json"
	"errors"
	"github.com/apisearch/apisearch/model/elasticsearch"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"reflect"
)

type Settings struct {
	UserId     string `json:"-"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	Password   string `json:"password"`
	FeedUrl    string `json:"feedUrl"`
	FeedFormat string `json:"feedFormat"`
}

func (s *Settings) Create() (NewUser, error) {
	client := elasticsearch.CreateClient()
	var err error
	s.Password, err = hashPassword(s.Password)
	s.Token = generateToken()

	if err != nil {
		return NewUser{}, err
	}

	var found bool
	_, found, err = findByEmail(s.Email)

	if err != nil {
		return NewUser{}, err
	}

	if found == true {
		return NewUser{}, errors.New("Email already exists")
	}

	response, err := client.Index().
		Index(indexName).
		Type(typeName).
		BodyJson(s).
		Do(context.TODO())

	if response == nil || err != nil {
		return NewUser{}, err
	}

	return NewUser{response.Id, s.Token}, nil
}

func (s *Settings) Update(userId string) error {
	client := elasticsearch.CreateClient()

	response, err := client.Index().
		Index(indexName).
		Type(typeName).
		Id(userId).
		BodyJson(s).
		Do(context.TODO())

	if response == nil || err != nil {
		return err
	}

	return nil
}

func (s *Settings) Remove(userId string) (bool, error) {
	client := elasticsearch.CreateClient()

	res, err := client.Delete().Index(indexName).Type(typeName).Id(userId).Do(context.TODO())

	if err != nil {
		return false, err
	}

	if res.Found != true {
		return false, nil
	}

	return true, nil
}

func (s *Settings) Find(userId string) (bool, error) {
	client := elasticsearch.CreateClient()

	res, err := client.Get().Index(indexName).Type(typeName).Id(userId).Do(context.TODO())

	if err != nil {
		return false, err
	}

	if res.Found != true || res.Source == nil {
		return false, nil
	}

	if err := json.Unmarshal(*res.Source, &s); err != nil {
		return false, err
	}

	return true, nil
}

func (s *Settings) FindAll() ([]Settings, error) {
	client := elasticsearch.CreateClient()

	res, err := client.Search().Index(indexName).Query(elastic.NewMatchAllQuery()).Size(10000).Do(context.TODO())

	if err != nil {
		return nil, err
	}

	var ttyp Settings
	var result = []Settings{}

	for id, item := range res.Each(reflect.TypeOf(ttyp)) {
		if s, ok := item.(Settings); ok {
			s.UserId = string(id)
			result = append(result, s)
		}
	}

	return result, nil
}

func findByEmail(email string) (Settings, bool, error) {
	client := elasticsearch.CreateClient()
	var ttyp Settings
	res, err := client.Search().Index(indexName).Query(elastic.NewTermQuery("email", email)).Do(context.TODO())

	if err != nil {
		return ttyp, false, err
	}

	for id, item := range res.Each(reflect.TypeOf(ttyp)) {
		if s, ok := item.(Settings); ok {
			s.UserId = string(id)
			return s, true, nil
		}
	}

	return ttyp, false, nil
}
