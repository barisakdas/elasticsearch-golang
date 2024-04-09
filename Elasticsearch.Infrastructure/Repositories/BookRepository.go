package repositories

import (
	entities "Elasticsearch/Elasticsearch.Core/Entities"
	models "Elasticsearch/Elasticsearch.Core/Models"
	repositoryinterfaces "Elasticsearch/Elasticsearch.Infrastructure/RepositoryInterfaces"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
)

type BookRepository struct {
	client *elasticsearch.Client
}

func NewBookRepository() repositoryinterfaces.IBookRepository {
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

	return &BookRepository{
		client: client,
	}
}

func (repo BookRepository) GetAllAsync(indexName string) ([]entities.Book, error) {
	var r map[string]interface{}
	var books []entities.Book

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

	// Parse the hits and append to books slice
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var book entities.Book
		source := hit.(map[string]interface{})["_source"]
		marshalled, _ := json.Marshal(source)
		json.Unmarshal(marshalled, &book)
		books = append(books, book)
	}

	return books, nil
}

func (repo BookRepository) GetByIdAsync(indexName string, id string) (entities.Book, error) {
	var book entities.Book

	// GetRequest oluşturuluyor
	getReq := esapi.GetRequest{
		Index:      indexName,
		DocumentID: id,
	}

	// GetRequest Elasticsearch'e gönderiliyor
	res, err := getReq.Do(context.Background(), repo.client)
	if err != nil {
		return book, fmt.Errorf("error getting document: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return book, fmt.Errorf("error with the response: %s", res.String())
	}

	// Yanıtın içeriği JSON'dan book nesnesine dönüştürülüyor
	if err := json.NewDecoder(res.Body).Decode(&book); err != nil {
		return book, fmt.Errorf("error parsing the response body: %s", err)
	}

	return book, nil
}

func (repo *BookRepository) SearchAsync(indexName string, searchQuery string) ([]entities.Book, error) {
	var books []entities.Book

	// Fuzziness değerini hesapla
	fuzziness := "AUTO" // Varsayılan olarak AUTO kullan
	if len(searchQuery) < 5 {
		fuzziness = "1" // Eğer searchQuery 5 karakterden az ise
	} else {
		fuzzinessValue := len(searchQuery) / 5
		fuzziness = fmt.Sprintf("%d", fuzzinessValue) // Karakter sayısının 5'te biri
	}

	// Elasticsearch sorgusu
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":     searchQuery,
				"fields":    []string{"title", "abstract"},
				"type":      "best_fields",
				"fuzziness": fuzziness,
			},
		},
	}

	// Sorguyu JSON'a dönüştür
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	// Elasticsearch'e sorguyu gönder
	res, err := repo.client.Search(
		repo.client.Search.WithContext(context.Background()),
		repo.client.Search.WithIndex(indexName),
		repo.client.Search.WithBody(bytes.NewReader(queryJSON)),
		repo.client.Search.WithSize(10), // İstenilen sonuç sayısı
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error with the response: %s", res.String())
	}

	// Yanıtı ayrıştır ve books slice'ına ekle
	err = json.NewDecoder(res.Body).Decode(&books)
	if err != nil {
		return nil, fmt.Errorf("error parsing the response body: %s", err)
	}

	return books, nil
}

func (repo *BookRepository) SearchByModelAsync(indexName string, searchModel models.SearchBookModel) ([]entities.Book, error) {
	var books []entities.Book
	var boolQuery []map[string]interface{}

	// Title için match query ekle
	if searchModel.Title != "" {
		boolQuery = append(boolQuery, map[string]interface{}{
			"match": map[string]interface{}{
				"title": searchModel.Title,
			},
		})
	}

	// Abstract için match query ekle
	if searchModel.Abstract != "" {
		boolQuery = append(boolQuery, map[string]interface{}{
			"match": map[string]interface{}{
				"abstract": searchModel.Abstract,
			},
		})
	}

	// PublishDateStart için range query ekle
	if !searchModel.PublishDateStart.IsZero() {
		boolQuery = append(boolQuery, map[string]interface{}{
			"range": map[string]interface{}{
				"publish_date": map[string]interface{}{
					"gte": searchModel.PublishDateStart,
				},
			},
		})
	}

	// MinPrice için range query ekle
	if searchModel.MinPrice > 0 {
		boolQuery = append(boolQuery, map[string]interface{}{
			"range": map[string]interface{}{
				"price": map[string]interface{}{
					"gte": searchModel.MinPrice,
				},
			},
		})
	}

	// MinStock için range query ekle
	if searchModel.MinStock > 0 {
		boolQuery = append(boolQuery, map[string]interface{}{
			"range": map[string]interface{}{
				"stock": map[string]interface{}{
					"gte": searchModel.MinStock,
				},
			},
		})
	}

	// Dinamik bool query oluştur
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": boolQuery,
			},
		},
	}

	// Sorguyu JSON'a dönüştür
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	// Elasticsearch'e sorguyu gönder
	res, err := repo.client.Search(
		repo.client.Search.WithContext(context.Background()),
		repo.client.Search.WithIndex(indexName),
		repo.client.Search.WithBody(bytes.NewReader(queryJSON)),
		repo.client.Search.WithSize(10), // İstenilen sonuç sayısı
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error with the response: %s", res.String())
	}

	// Yanıtı ayrıştır ve books slice'ına ekle
	err = json.NewDecoder(res.Body).Decode(&books)
	if err != nil {
		return nil, fmt.Errorf("error parsing the response body: %s", err)
	}

	return books, nil
}

func (repo BookRepository) IndexAsync(indexName string, entity entities.Book) (entities.Book, error) {
	// book nesnesini JSON'a dönüştür
	data, err := json.Marshal(entity)
	if err != nil {
		return entities.Book{}, err
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
		return entities.Book{}, err
	}
	defer res.Body.Close()

	// Yanıtı kontrol et
	if res.IsError() {
		return entities.Book{}, fmt.Errorf("error indexing document ID=%s: %s", entity.Id, res.String())
	}

	// Başarılı indexleme sonrası entity'yi geri döndür
	return entity, nil
}

func (repo BookRepository) UpdateAsync(indexName string, entity entities.Book) (entities.Book, error) {
	// book nesnesini JSON'a dönüştür
	data, err := json.Marshal(entity)
	if err != nil {
		return entities.Book{}, err
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
		return entities.Book{}, err
	}
	defer res.Body.Close()

	// Yanıtı kontrol et
	if res.IsError() {
		return entities.Book{}, fmt.Errorf("error updating document ID=%s: %s", entity.Id, res.String())
	}

	// Başarılı güncelleme sonrası entity'yi geri döndür
	return entity, nil
}

func (repo BookRepository) DeleteAsync(indexName string, id string) (bool, error) {
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
