package repositoryinterfaces

import entities "Elasticsearch/Elasticsearch.Core/Entities"

type IAuthorRepository interface {
	GetAllAsync(indexName string) ([]entities.Author, error)
	GetByIdAsync(indexName string, id string) (entities.Author, error)
	IndexAsync(indexName string, entity entities.Author) (entities.Author, error)
	UpdateAsync(indexName string, entity entities.Author) (entities.Author, error)
	DeleteAsync(indexName string, id string) (bool, error)
}
