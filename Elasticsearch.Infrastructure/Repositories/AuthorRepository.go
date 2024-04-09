package repositories

import (
	entities "Elasticsearch/Elasticsearch.Core/Entities"
	repositoryinterfaces "Elasticsearch/Elasticsearch.Infrastructure/RepositoryInterfaces"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
)

type AuthorRepository struct {
	client *elasticsearch.Client
}

func NewAuthorRepository() repositoryinterfaces.IAuthorRepository {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200/",
		},
		Username: "elastic",
		Password: "DkIedPPSCb",
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	return &AuthorRepository{
		client: client,
	}
}

func (repo AuthorRepository) GetAllAsync(indexName string) ([]entities.Author, error) {
	var r map[string]interface{}
	var authors []entities.Author

	res, err := repo.client.Search(
		repo.client.Search.WithContext(context.Background()),
		repo.client.Search.WithIndex(indexName),
		repo.client.Search.WithSize(1000), // Adjust the size as needed
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error getting response: %s", res.String())
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("error parsing the response body: %s", err)
	}

	// Parse the hits and append to authors slice
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var author entities.Author
		source := hit.(map[string]interface{})["_source"]
		marshalled, _ := json.Marshal(source)
		json.Unmarshal(marshalled, &author)
		authors = append(authors, author)
	}

	return authors, nil
}

func (repo *AuthorRepository) GetByIdAsync(indexName string, id string) (entities.Author, error) {
	var author entities.Author

	// GetRequest oluşturuluyor
	getReq := esapi.GetRequest{
		Index:      indexName,
		DocumentID: id,
	}

	// GetRequest Elasticsearch'e gönderiliyor
	res, err := getReq.Do(context.Background(), repo.client)
	if err != nil {
		return author, fmt.Errorf("error getting document: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return author, fmt.Errorf("error with the response: %s", res.String())
	}

	// Yanıtın içeriği JSON'dan Author nesnesine dönüştürülüyor
	if err := json.NewDecoder(res.Body).Decode(&author); err != nil {
		return author, fmt.Errorf("error parsing the response body: %s", err)
	}

	return author, nil
}

func (repo *AuthorRepository) IndexAsync(indexName string, entity entities.Author) (entities.Author, error) {
	// Author nesnesini JSON'a dönüştür
	data, err := json.Marshal(entity)
	if err != nil {
		return entities.Author{}, err
	}

	// Elasticsearch'e indexleme isteği gönder
	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: entity.Id, // Varsayılan olarak ID'yi belirtin, yoksa otomatik olarak oluşturulur
		Body:       bytes.NewReader(data),
		Refresh:    "true", // İndexleme işleminden sonra arama yapılabilir olmasını sağlar
	}

	// İsteği yürüt
	res, err := req.Do(context.Background(), repo.client)
	if err != nil {
		return entities.Author{}, err
	}
	defer res.Body.Close()

	// Yanıtı kontrol et
	if res.IsError() {
		return entities.Author{}, fmt.Errorf("error indexing document ID=%s: %s", entity.Id, res.String())
	}

	// Başarılı indexleme sonrası entity'yi geri döndür
	return entity, nil
}

func (repo *AuthorRepository) UpdateAsync(indexName string, entity entities.Author) (entities.Author, error) {
	// Author nesnesini JSON'a dönüştür
	data, err := json.Marshal(entity)
	if err != nil {
		return entities.Author{}, err
	}

	// Elasticsearch'e güncelleme isteği gönder
	req := esapi.UpdateRequest{
		Index:      indexName,
		DocumentID: entity.Id,
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}

	// İsteği yürüt
	res, err := req.Do(context.Background(), repo.client)
	if err != nil {
		return entities.Author{}, err
	}
	defer res.Body.Close()

	// Yanıtı kontrol et
	if res.IsError() {
		return entities.Author{}, fmt.Errorf("error updating document ID=%s: %s", entity.Id, res.String())
	}

	// Başarılı güncelleme sonrası entity'yi geri döndür
	return entity, nil
}

func (repo *AuthorRepository) DeleteAsync(indexName string, id string) (bool, error) {
	// Elasticsearch'e silme isteği gönder
	req := esapi.DeleteRequest{
		Index:      indexName,
		DocumentID: id,
		Refresh:    "true",
	}

	// İsteği yürüt
	res, err := req.Do(context.Background(), repo.client)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	// Yanıtı kontrol et
	if res.IsError() {
		return false, fmt.Errorf("error deleting document ID=%s: %s", id, res.String())
	}

	// Başarılı silme işlemi sonrası true döndür
	return true, nil
}
