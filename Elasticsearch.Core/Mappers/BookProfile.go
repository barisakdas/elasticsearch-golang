package mappers

import (
	dtos "Elasticsearch/Elasticsearch.Core/Dtos"
	entities "Elasticsearch/Elasticsearch.Core/Entities"
	models "Elasticsearch/Elasticsearch.Core/Models"
)

type BookProfile struct {
}

func BookToBookDto(book entities.Book) dtos.BookDto {
	return dtos.BookDto{

		Id:          book.Id,
		Title:       book.Title,
		Abstract:    book.Abstract,
		PublishDate: book.PublishDate,
		Price:       book.Price,
		Stock:       book.Stock,
		Categories:  book.Categories,
		CreatedDate: book.CreatedDate,
		UpdatedDate: book.UpdatedDate,
		CreatedBy:   book.CreatedBy,
		UpdatedBy:   book.UpdatedBy,
		IsActive:    book.IsActive,
		Author:      AuthorToAuthorDto(book.Author),
	}
}

func BookDtoToBook(book dtos.BookDto) entities.Book {
	return entities.Book{

		Id:          book.Id,
		Title:       book.Title,
		Abstract:    book.Abstract,
		PublishDate: book.PublishDate,
		Price:       book.Price,
		Stock:       book.Stock,
		Categories:  book.Categories,
		CreatedDate: book.CreatedDate,
		UpdatedDate: book.UpdatedDate,
		CreatedBy:   book.CreatedBy,
		UpdatedBy:   book.UpdatedBy,
		IsActive:    book.IsActive,
		Author:      AuthorDtoToAuthor(book.Author),
	}
}

func BooksToBookDtos(books []entities.Book) []dtos.BookDto {
	var bookDtos []dtos.BookDto

	for _, book := range books {

		var bookDto = dtos.BookDto{
			Id:          book.Id,
			Title:       book.Title,
			Abstract:    book.Abstract,
			PublishDate: book.PublishDate,
			Price:       book.Price,
			Stock:       book.Stock,
			Categories:  book.Categories,
			CreatedDate: book.CreatedDate,
			UpdatedDate: book.UpdatedDate,
			CreatedBy:   book.CreatedBy,
			UpdatedBy:   book.UpdatedBy,
			IsActive:    book.IsActive,
			Author:      AuthorToAuthorDto(book.Author),
		}

		bookDtos = append(bookDtos, bookDto)
	}
	return bookDtos
}

func BookDtosToBooks(bookDtos []dtos.BookDto) []entities.Book {
	var books []entities.Book

	for _, bookDto := range bookDtos {

		var book = entities.Book{
			Id:          bookDto.Id,
			Title:       bookDto.Title,
			Abstract:    bookDto.Abstract,
			PublishDate: bookDto.PublishDate,
			Price:       bookDto.Price,
			Stock:       bookDto.Stock,
			Categories:  bookDto.Categories,
			CreatedDate: bookDto.CreatedDate,
			UpdatedDate: bookDto.UpdatedDate,
			CreatedBy:   bookDto.CreatedBy,
			UpdatedBy:   bookDto.UpdatedBy,
			IsActive:    bookDto.IsActive,
			Author:      AuthorDtoToAuthor(bookDto.Author),
		}

		books = append(books, book)
	}
	return books
}

func CreateBookModelToBook(model models.CreateBookModel) entities.Book {
	return entities.Book{
		Title:       model.Title,
		Abstract:    model.Abstract,
		PublishDate: model.PublishDate,
		Price:       model.Price,
		Stock:       model.Stock,
		Categories:  model.Categories,
	}
}

func UpdateBookModelToBook(model models.UpdateBookModel) entities.Book {
	return entities.Book{
		Id:          model.Id,
		Title:       model.Title,
		Abstract:    model.Abstract,
		PublishDate: model.PublishDate,
		Price:       model.Price,
		Stock:       model.Stock,
		Categories:  model.Categories,
	}
}

func CreateBookModelToBookDto(model models.CreateBookModel) dtos.BookDto {
	return dtos.BookDto{
		Title:       model.Title,
		Abstract:    model.Abstract,
		PublishDate: model.PublishDate,
		Price:       model.Price,
		Stock:       model.Stock,
		Categories:  model.Categories,
	}
}

func UpdateBookModelToBookDto(model models.UpdateBookModel) dtos.BookDto {
	return dtos.BookDto{
		Id:          model.Id,
		Title:       model.Title,
		Abstract:    model.Abstract,
		PublishDate: model.PublishDate,
		Price:       model.Price,
		Stock:       model.Stock,
		Categories:  model.Categories,
	}
}
