package serviceinterfaces

import (
	models "Elasticsearch/Elasticsearch.Core/Models"
	responsemodels "Elasticsearch/Elasticsearch.Core/ResponseModels"
)

type IBookService interface {
	GetAllBooksAsync(indexName string) responsemodels.ResponseModel
	GetBookByIdAsync(indexName string, id string) responsemodels.ResponseModel
	SearchBooksAsync(indexName string, searchText string) responsemodels.ResponseModel
	SearchBooksByModelAsync(indexName string, model models.SearchBookModel) responsemodels.ResponseModel
	IndexBookAsync(indexName string, model models.CreateBookModel) responsemodels.ResponseModel
	UpdateBookAsync(indexName string, model models.UpdateBookModel) responsemodels.ResponseModel
	DeleteBookAsync(indexName string, id string) responsemodels.ResponseModel
}
