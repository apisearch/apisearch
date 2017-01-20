package settings

import "github.com/apisearch/apisearch/model/elasticsearch"

const (
	indexName     = "settings"
	typeName      = "setting"
	indexSettings = `{
		"settings": {
			"number_of_shards": 1,
			"number_of_replicas": 0
		}
	}`
	typeSettings = `{
		"setting": {
			"properties": {
				"email": {
					"type": "keyword"
				},
				"token": {
					"type":"keyword"
				},
				"password": {
					"type":"keyword"
				},
				"feedUrl": {
					"type":"keyword"
				},
				"feedFormat": {
					"type":"keyword"
				}
			}
		}
	}`
)

func CreateIndex(force bool) error {
	var err error

	if force {
		err = elasticsearch.DeleteIndex(indexName)

		if err != nil {
			return err
		}
	}

	err = elasticsearch.CreateIndex(indexSettings, indexName)

	if err != nil {
		return err
	}

	return elasticsearch.PutMapping(typeSettings, indexName, typeName)
}
