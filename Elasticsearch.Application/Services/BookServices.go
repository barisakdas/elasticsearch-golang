package services

import (
	serviceinterfaces "Elasticsearch/Elasticsearch.Application/ServiceInterfaces"
	mappers "Elasticsearch/Elasticsearch.Core/Mappers"
	models "Elasticsearch/Elasticsearch.Core/Models"
	responsemodels "Elasticsearch/Elasticsearch.Core/ResponseModels"
	repositories "Elasticsearch/Elasticsearch.Infrastructure/Repositories"
	"reflect"
)

var _bookRepository = repositories.NewBookRepository()

type BookService struct{}

func NewBookService() serviceinterfaces.IBookService {
	return &BookService{}
}

func (b BookService) GetAllBooksAsync(indexName string) responsemodels.ResponseModel {
	if indexName == "" {
		return responsemodels.ResponseModel{}.Fail([]string{"Index ismi boş geçilemez"}, 400)
	}

	data, err := _bookRepository.GetAllAsync(indexName)
	if err != nil {
		return responsemodels.ResponseModel{}.Fail([]string{err.Error()}, 404)
	}

	return responsemodels.ResponseModel{}.Success(mappers.BooksToBookDtos(data))
}

func (b BookService) GetBookByIdAsync(indexName string, id string) responsemodels.ResponseModel {
	if indexName == "" {
		return responsemodels.ResponseModel{}.Fail([]string{"Index ismi boş geçilemez"}, 400)
	}

	if id == "" {
		return responsemodels.ResponseModel{}.Fail([]string{"Id boş geçilemez"}, 400)
	}

	data, err := _bookRepository.GetByIdAsync(indexName, id)
	if err != nil {
		return responsemodels.ResponseModel{}.Fail([]string{err.Error()}, 404)
	}

	return responsemodels.ResponseModel{}.Success(mappers.BookToBookDto(data))
}

func (b BookService) SearchBooksAsync(indexName string, searchText string) responsemodels.ResponseModel {
	if indexName == "" {
		return responsemodels.ResponseModel{}.Fail([]string{"Index ismi boş geçilemez"}, 400)
	}

	if searchText == "" {
		return responsemodels.ResponseModel{}.Fail([]string{"Id boş geçilemez"}, 400)
	}

	data, err := _bookRepository.SearchAsync(indexName, searchText)
	if err != nil {
		return responsemodels.ResponseModel{}.Fail([]string{err.Error()}, 404)
	}

	return responsemodels.ResponseModel{}.Success(mappers.BooksToBookDtos(data))
}

func (b BookService) SearchBooksByModelAsync(indexName string, model models.SearchBookModel) responsemodels.ResponseModel {
	if indexName == "" {
		return responsemodels.ResponseModel{}.Fail([]string{"Index ismi boş geçilemez"}, 400)
	}

	if reflect.DeepEqual(model, models.SearchBookModel{}) {
		return responsemodels.ResponseModel{}.Fail([]string{"Model boş geçilemez"}, 400)
	}

	datas, err := _bookRepository.SearchByModelAsync(indexName, model)
	if err != nil {
		return responsemodels.ResponseModel{}.Fail([]string{err.Error()}, 404)
	}

	return responsemodels.ResponseModel{}.Success(mappers.BooksToBookDtos(datas))
}

func (b BookService) IndexBookAsync(indexName string, model models.CreateBookModel) responsemodels.ResponseModel {
	if indexName == "" {
		return responsemodels.ResponseModel{}.Fail([]string{"Index ismi boş geçilemez"}, 400)
	}

	if reflect.DeepEqual(model, models.CreateBookModel{}) {
		return responsemodels.ResponseModel{}.Fail([]string{"Model boş geçilemez"}, 400)
	}

	data, err := _bookRepository.IndexAsync(indexName, mappers.CreateBookModelToBook(model))
	if err != nil {
		return responsemodels.ResponseModel{}.Fail([]string{err.Error()}, 404)
	}

	return responsemodels.ResponseModel{}.Success(mappers.BookToBookDto(data))
}

func (b BookService) UpdateBookAsync(indexName string, model models.UpdateBookModel) responsemodels.ResponseModel {
	if indexName == "" {
		return responsemodels.ResponseModel{}.Fail([]string{"Index ismi boş geçilemez"}, 400)
	}

	if reflect.DeepEqual(model, models.UpdateBookModel{}) {
		return responsemodels.ResponseModel{}.Fail([]string{"Model boş geçilemez"}, 400)
	}

	data, err := _bookRepository.IndexAsync(indexName, mappers.UpdateBookModelToBook(model))
	if err != nil {
		return responsemodels.ResponseModel{}.Fail([]string{err.Error()}, 404)
	}

	return responsemodels.ResponseModel{}.Success(mappers.BookToBookDto(data))
}

func (b BookService) DeleteBookAsync(indexName string, id string) responsemodels.ResponseModel {
	if indexName == "" {
		return responsemodels.ResponseModel{}.Fail([]string{"Index ismi boş geçilemez"}, 400)
	}

	if id == "" {
		return responsemodels.ResponseModel{}.Fail([]string{"Id boş geçilemez"}, 400)
	}

	data, err := _bookRepository.DeleteAsync(indexName, id)
	if err != nil {
		return responsemodels.ResponseModel{}.Fail([]string{err.Error()}, 404)
	}

	return responsemodels.ResponseModel{}.Success(data)
}
