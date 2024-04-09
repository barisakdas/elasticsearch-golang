package serviceinterfaces

import (
	models "Elasticsearch/Elasticsearch.Core/Models"
	responsemodels "Elasticsearch/Elasticsearch.Core/ResponseModels"
)

type IAuthorService interface {
	GetAllAuthorsAsync(indexName string) responsemodels.ResponseModel
	GetAuthorByIdAsync(indexName string, id string) responsemodels.ResponseModel
	IndexAuthorAsync(indexName string, model models.CreateAuthorModel) responsemodels.ResponseModel
	UpdateAuthorAsync(indexName string, model models.UpdateAuthorModel) responsemodels.ResponseModel
	DeleteAuthorAsync(indexName string, id string) responsemodels.ResponseModel
}
