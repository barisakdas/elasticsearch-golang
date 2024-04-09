package services

import (
	serviceinterfaces "Elasticsearch/Elasticsearch.Application/ServiceInterfaces"
	mappers "Elasticsearch/Elasticsearch.Core/Mappers"
	models "Elasticsearch/Elasticsearch.Core/Models"
	responsemodels "Elasticsearch/Elasticsearch.Core/ResponseModels"
	repositories "Elasticsearch/Elasticsearch.Infrastructure/Repositories"
	"reflect"
)

var _authorRepository = repositories.NewAuthorRepository()

type AuthorService struct{}

func NewAuthorService() serviceinterfaces.IAuthorService {
	return &AuthorService{}
}

func (a AuthorService) GetAllAuthorsAsync(indexName string) responsemodels.ResponseModel {
	if indexName == "" {
		return responsemodels.ResponseModel{}.Fail([]string{"Index ismi boş geçilemez"}, 400)
	}

	datas, err := _authorRepository.GetAllAsync(indexName)

	if err != nil {
		return responsemodels.ResponseModel{}.Fail([]string{err.Error()}, 404)
	}

	return responsemodels.ResponseModel{}.Success(mappers.AuthorsToAuthorDtos(datas))
}

func (a AuthorService) GetAuthorByIdAsync(indexName string, id string) responsemodels.ResponseModel {
	if indexName == "" {
		return responsemodels.ResponseModel{}.Fail([]string{"Index ismi boş geçilemez"}, 400)
	}

	if id == "" {
		return responsemodels.ResponseModel{}.Fail([]string{"Id boş geçilemez"}, 400)
	}

	data, err := _authorRepository.GetByIdAsync(indexName, id)

	if err != nil {
		return responsemodels.ResponseModel{}.Fail([]string{err.Error()}, 404)
	}

	return responsemodels.ResponseModel{}.Success(mappers.AuthorToAuthorDto(data))
}

func (a AuthorService) IndexAuthorAsync(indexName string, model models.CreateAuthorModel) responsemodels.ResponseModel {
	if indexName == "" {
		return responsemodels.ResponseModel{}.Fail([]string{"Index ismi boş geçilemez"}, 400)
	}

	if reflect.DeepEqual(model, models.CreateAuthorModel{}) {
		return responsemodels.ResponseModel{}.Fail([]string{"Model boş geçilemez"}, 400)
	}

	data, err := _authorRepository.IndexAsync(indexName, mappers.CreateAuthorModelToAuthor(model))

	if err != nil {
		return responsemodels.ResponseModel{}.Fail([]string{err.Error()}, 404)
	}

	return responsemodels.ResponseModel{}.Success(mappers.AuthorToAuthorDto(data))
}

func (a AuthorService) UpdateAuthorAsync(indexName string, model models.UpdateAuthorModel) responsemodels.ResponseModel {
	if indexName == "" {
		return responsemodels.ResponseModel{}.Fail([]string{"Index ismi boş geçilemez"}, 400)
	}

	if reflect.DeepEqual(model, models.UpdateAuthorModel{}) {
		return responsemodels.ResponseModel{}.Fail([]string{"Model boş geçilemez"}, 400)
	}

	data, err := _authorRepository.UpdateAsync(indexName, mappers.UpdateAuthorModelToAuthor(model))

	if err != nil {
		return responsemodels.ResponseModel{}.Fail([]string{err.Error()}, 404)
	}

	return responsemodels.ResponseModel{}.Success(mappers.AuthorToAuthorDto(data))
}

func (a AuthorService) DeleteAuthorAsync(indexName string, id string) responsemodels.ResponseModel {
	if indexName == "" {
		return responsemodels.ResponseModel{}.Fail([]string{"Index ismi boş geçilemez"}, 400)
	}

	if id == "" {
		return responsemodels.ResponseModel{}.Fail([]string{"Id boş geçilemez"}, 400)
	}

	data, err := _authorRepository.DeleteAsync(indexName, id)

	if err != nil {
		return responsemodels.ResponseModel{}.Fail([]string{err.Error()}, 404)
	}

	return responsemodels.ResponseModel{}.Success(data)
}
