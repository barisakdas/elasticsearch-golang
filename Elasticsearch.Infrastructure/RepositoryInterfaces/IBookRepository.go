package repositoryinterfaces

import (
	entities "Elasticsearch/Elasticsearch.Core/Entities"
	models "Elasticsearch/Elasticsearch.Core/Models"
)

type IBookRepository interface {
	GetAllAsync(indexName string) ([]entities.Book, error)
	GetByIdAsync(indexName string, id string) (entities.Book, error)
	SearchAsync(indexName string, query string) ([]entities.Book, error)
	SearchByModelAsync(indexName string, searchModel models.SearchBookModel) ([]entities.Book, error)
	IndexAsync(indexName string, entity entities.Book) (entities.Book, error)
	UpdateAsync(indexName string, entity entities.Book) (entities.Book, error)
	DeleteAsync(indexName string, id string) (bool, error)
}
