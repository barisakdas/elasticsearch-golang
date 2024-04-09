package converters

import (
	models "Elasticsearch/Elasticsearch.Core/Models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Convert_To_CreateAuthorModel(r *http.Request) (models.CreateAuthorModel, error) {
	var requestModel models.CreateAuthorModel
	bodyText, err := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &requestModel)

	if err != nil {
		return models.CreateAuthorModel{}, err
	}
	return requestModel, nil
}

func Convert_To_UpdateAuthorModel(r *http.Request) (models.UpdateAuthorModel, error) {
	var requestModel models.UpdateAuthorModel
	bodyText, err := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &requestModel)

	if err != nil {
		return models.UpdateAuthorModel{}, err
	}
	return requestModel, nil
}
