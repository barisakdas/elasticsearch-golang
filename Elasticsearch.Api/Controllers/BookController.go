package controllers

import (
	services "Elasticsearch/Elasticsearch.Application/Services"
	converters "Elasticsearch/Elasticsearch.Core/Converters"
	"encoding/json"
	"net/http"
)

type BookController struct {
}

var _BookService = services.NewBookService()

func GetAllBooksAsync(w http.ResponseWriter, r *http.Request) {
	indexName := r.URL.Query().Get("indexName")
	var responseData = _BookService.GetAllBooksAsync(indexName)
	if !responseData.IsSucceded {
		w.Header().Add("Api Status", "Bad Request")
	}

	response, _ := json.Marshal(&responseData)
	w.Write(response)
}

func GetBookByIdAsync(w http.ResponseWriter, r *http.Request) {
	indexName := r.URL.Query().Get("indexName")
	id := r.URL.Query().Get("id")
	var responseData = _BookService.GetBookByIdAsync(indexName, id)
	if !responseData.IsSucceded {
		w.Header().Add("Api Status", "Bad Request")
	}

	response, _ := json.Marshal(&responseData)
	w.Write(response)
}

func SearchBookAsync(w http.ResponseWriter, r *http.Request) {
	indexName := r.URL.Query().Get("indexName")
	searchText := r.URL.Query().Get("searchText")
	var responseData = _BookService.SearchBooksAsync(indexName, searchText)
	if !responseData.IsSucceded {
		w.Header().Add("Api Status", "Bad Request")
	}

	response, _ := json.Marshal(&responseData)
	w.Write(response)
}

func SearchBookByModelAsync(w http.ResponseWriter, r *http.Request) {
	indexName := r.URL.Query().Get("indexName")
	model, err := converters.Convert_To_SearchBookModel(r)
	if err != nil {
		w.Header().Add("Api Status", "Bad Request")
	}
	var responseData = _BookService.SearchBooksByModelAsync(indexName, model)
	if !responseData.IsSucceded {
		w.Header().Add("Api Status", "Bad Request")
	}

	response, _ := json.Marshal(&responseData)
	w.Write(response)
}

func IndexBookAsync(w http.ResponseWriter, r *http.Request) {
	indexName := r.URL.Query().Get("indexName")
	model, err := converters.Convert_To_CreateBookModel(r)
	if err != nil {
		w.Header().Add("Api Status", "Bad Request")
	}

	var responseData = _BookService.IndexBookAsync(indexName, model)
	if !responseData.IsSucceded {
		w.Header().Add("Api Status", "Bad Request")
	}

	response, _ := json.Marshal(&responseData)
	w.Write(response)
}

func UpdateBookAsync(w http.ResponseWriter, r *http.Request) {
	indexName := r.URL.Query().Get("indexName")
	model, err := converters.Convert_To_UpdateBookModel(r)
	if err != nil {
		w.Header().Add("Api Status", "Bad Request")
	}

	var responseData = _BookService.UpdateBookAsync(indexName, model)
	if !responseData.IsSucceded {
		w.Header().Add("Api Status", "Bad Request")
	}

	response, _ := json.Marshal(&responseData)
	w.Write(response)
}

func DeleteBookAsync(w http.ResponseWriter, r *http.Request) {
	indexName := r.URL.Query().Get("indexName")
	id := r.URL.Query().Get("id")
	var responseData = _BookService.DeleteBookAsync(indexName, id)
	if !responseData.IsSucceded {
		w.Header().Add("Api Status", "Bad Request")
	}

	response, _ := json.Marshal(&responseData)
	w.Write(response)
}
