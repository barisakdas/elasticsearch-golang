package converters

import (
	models "Elasticsearch/Elasticsearch.Core/Models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Convert_To_CreateBookModel(r *http.Request) (models.CreateBookModel, error) {
	var requestModel models.CreateBookModel
	bodyText, err := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &requestModel)

	if err != nil {
		return models.CreateBookModel{}, err
	}
	return requestModel, nil
}

func Convert_To_UpdateBookModel(r *http.Request) (models.UpdateBookModel, error) {
	var requestModel models.UpdateBookModel
	bodyText, err := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &requestModel)

	if err != nil {
		return models.UpdateBookModel{}, err
	}
	return requestModel, nil
}

func Convert_To_SearchBookModel(r *http.Request) (models.SearchBookModel, error) {
	var requestModel models.SearchBookModel
	bodyText, err := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &requestModel)

	if err != nil {
		return models.SearchBookModel{}, err
	}
	return requestModel, nil
}
