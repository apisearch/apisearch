package settings

import (
	"encoding/json"
	"errors"
	"github.com/apisearch/apisearch/model/elasticsearch"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
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

	s.UserId = response.Id

	client.Refresh(indexName)

	return NewUser{response.Id, s.Token}, nil
}

func (s *Settings) Update(newSettings Settings) error {
	client := elasticsearch.CreateClient()
	var err error

	s.FeedUrl = newSettings.FeedUrl
	s.FeedFormat = newSettings.FeedFormat
	s.Email = newSettings.Email

	if s.Password != "" {
		s.Password, err = hashPassword(s.Password)

		if err != nil {
			return err
		}
	}

	response, err := client.Index().
		Index(indexName).
		Type(typeName).
		Id(s.UserId).
		BodyJson(s).
		Do(context.TODO())

	if response == nil || err != nil {
		return err
	}

	return nil
}

func (s *Settings) Remove() (bool, error) {
	client := elasticsearch.CreateClient()

	res, err := client.Delete().Index(indexName).Type(typeName).Id(s.UserId).Do(context.TODO())

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

	var result = []Settings{}

	for _, hit := range res.Hits.Hits {
		err = json.Unmarshal(*hit.Source, &s)

		if err == nil {
			s.UserId = hit.Id
			result = append(result, *s)
		}
	}

	return result, nil
}

func (s *Settings) FindByToken(token string) (bool, error) {
	client := elasticsearch.CreateClient()
	res, err := client.Search().Index(indexName).Query(elastic.NewTermQuery("token", token)).Do(context.TODO())

	if err != nil {
		return false, err
	}

	for _, hit := range res.Hits.Hits {
		err := json.Unmarshal(*hit.Source, &s)

		if err != nil {
			return false, err
		}

		s.UserId = hit.Id

		return true, nil
	}

	return false, nil
}

func findByEmail(email string) (Settings, bool, error) {
	client := elasticsearch.CreateClient()
	var ttyp Settings
	var s Settings
	res, err := client.Search().Index(indexName).Query(elastic.NewTermQuery("email", email)).Do(context.TODO())

	if err != nil {
		return ttyp, false, err
	}

	for _, hit := range res.Hits.Hits {
		err := json.Unmarshal(*hit.Source, &s)

		if err != nil {
			return ttyp, false, err
		}

		s.UserId = hit.Id

		return s, true, nil
	}

	return ttyp, false, nil
}
