package controllers

import (
	services "Elasticsearch/Elasticsearch.Application/Services"
	converters "Elasticsearch/Elasticsearch.Core/Converters"
	"encoding/json"
	"net/http"
)

type AuthorController struct {
}

var _authorService = services.NewAuthorService()

func GetAllAuthorsAsync(w http.ResponseWriter, r *http.Request) {
	indexName := r.URL.Query().Get("indexName")
	var responseData = _authorService.GetAllAuthorsAsync(indexName)
	if !responseData.IsSucceded {
		w.Header().Add("Api Status", "Bad Request")
	}

	response, _ := json.Marshal(&responseData)
	w.Write(response)
}

func GetAuthorByIdAsync(w http.ResponseWriter, r *http.Request) {
	indexName := r.URL.Query().Get("indexName")
	id := r.URL.Query().Get("id")
	var responseData = _authorService.GetAuthorByIdAsync(indexName, id)
	if !responseData.IsSucceded {
		w.Header().Add("Api Status", "Bad Request")
	}

	response, _ := json.Marshal(&responseData)
	w.Write(response)
}

func IndexAuthorAsync(w http.ResponseWriter, r *http.Request) {
	indexName := r.URL.Query().Get("indexName")
	model, err := converters.Convert_To_CreateAuthorModel(r)
	if err != nil {
		w.Header().Add("Api Status", "Bad Request")
	}

	var responseData = _authorService.IndexAuthorAsync(indexName, model)
	if !responseData.IsSucceded {
		w.Header().Add("Api Status", "Bad Request")
	}

	response, _ := json.Marshal(&responseData)
	w.Write(response)
}

func UpdateAuthorAsync(w http.ResponseWriter, r *http.Request) {
	indexName := r.URL.Query().Get("indexName")
	model, err := converters.Convert_To_UpdateAuthorModel(r)
	if err != nil {
		w.Header().Add("Api Status", "Bad Request")
	}

	var responseData = _authorService.UpdateAuthorAsync(indexName, model)
	if !responseData.IsSucceded {
		w.Header().Add("Api Status", "Bad Request")
	}

	response, _ := json.Marshal(&responseData)
	w.Write(response)
}

func DeleteAuthorAsync(w http.ResponseWriter, r *http.Request) {
	indexName := r.URL.Query().Get("indexName")
	id := r.URL.Query().Get("id")
	var responseData = _authorService.DeleteAuthorAsync(indexName, id)
	if !responseData.IsSucceded {
		w.Header().Add("Api Status", "Bad Request")
	}

	response, _ := json.Marshal(&responseData)
	w.Write(response)
}
