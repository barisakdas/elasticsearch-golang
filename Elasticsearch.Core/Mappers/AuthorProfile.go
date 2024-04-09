package mappers

import (
	dtos "Elasticsearch/Elasticsearch.Core/Dtos"
	entities "Elasticsearch/Elasticsearch.Core/Entities"
	models "Elasticsearch/Elasticsearch.Core/Models"
)

type AuthorProfile struct {
}

func AuthorToAuthorDto(author entities.Author) dtos.AuthorDto {
	return dtos.AuthorDto{
		Id:          author.Id,
		FirstName:   author.FirstName,
		LastName:    author.LastName,
		BirthDate:   author.BirthDate,
		CreatedDate: author.CreatedDate,
		UpdatedDate: author.UpdatedDate,
		CreatedBy:   author.CreatedBy,
		UpdatedBy:   author.UpdatedBy,
		IsActive:    author.IsActive,
	}
}

func AuthorDtoToAuthor(author dtos.AuthorDto) entities.Author {
	return entities.Author{
		Id:          author.Id,
		FirstName:   author.FirstName,
		LastName:    author.LastName,
		BirthDate:   author.BirthDate,
		CreatedDate: author.CreatedDate,
		UpdatedDate: author.UpdatedDate,
		CreatedBy:   author.CreatedBy,
		UpdatedBy:   author.UpdatedBy,
		IsActive:    author.IsActive,
	}
}

func AuthorsToAuthorDtos(authors []entities.Author) []dtos.AuthorDto {
	var authorDtos []dtos.AuthorDto

	for _, author := range authors {
		var authorDto = dtos.AuthorDto{
			Id:          author.Id,
			FirstName:   author.FirstName,
			LastName:    author.LastName,
			BirthDate:   author.BirthDate,
			CreatedDate: author.CreatedDate,
			UpdatedDate: author.UpdatedDate,
			CreatedBy:   author.CreatedBy,
			UpdatedBy:   author.UpdatedBy,
			IsActive:    author.IsActive,
		}
		authorDtos = append(authorDtos, authorDto)
	}

	return authorDtos
}

func AuthorDtosToAuthors(authorDtos []dtos.AuthorDto) []entities.Author {
	var authors []entities.Author

	for _, authorDto := range authorDtos {
		var author = entities.Author{
			Id:          authorDto.Id,
			FirstName:   authorDto.FirstName,
			LastName:    authorDto.LastName,
			BirthDate:   authorDto.BirthDate,
			CreatedDate: authorDto.CreatedDate,
			UpdatedDate: authorDto.UpdatedDate,
			CreatedBy:   authorDto.CreatedBy,
			UpdatedBy:   authorDto.UpdatedBy,
			IsActive:    authorDto.IsActive,
		}
		authors = append(authors, author)
	}

	return authors
}

func CreateAuthorModelToAuthor(model models.CreateAuthorModel) entities.Author {
	return entities.Author{
		FirstName: model.FirstName,
		LastName:  model.LastName,
		BirthDate: model.BirthDate,
	}
}

func UpdateAuthorModelToAuthor(model models.UpdateAuthorModel) entities.Author {
	return entities.Author{
		Id:        model.Id,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		BirthDate: model.BirthDate,
	}
}

func CreateAuthorModelToAuthorDto(model models.CreateAuthorModel) dtos.AuthorDto {
	return dtos.AuthorDto{
		FirstName: model.FirstName,
		LastName:  model.LastName,
		BirthDate: model.BirthDate,
	}
}

func UpdateAuthorModelToAuthorDto(model models.UpdateAuthorModel) dtos.AuthorDto {
	return dtos.AuthorDto{
		Id:        model.Id,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		BirthDate: model.BirthDate,
	}
}
