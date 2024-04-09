package api

import (
	controllers "Elasticsearch/Elasticsearch.Api/Controllers"
	helperservices "Elasticsearch/Elasticsearch.Application/Services/HelperServices"
	"net/http"
)

var _helperService = helperservices.NewHelperService()

func Run() {
	Endpoints()
	http.ListenAndServe(":8080", nil)
}

func Endpoints() {
	//Token Operations (Header içine Token : "{active_token}" şeklinde eklenir)
	http.HandleFunc("/get-access-token", _helperService.GetAccessToken) // For access token

	// Author Operations
	http.Handle("Author/getall", _helperService.IsAuthorized(controllers.GetAllAuthorsAsync))
	http.Handle("Author/getbyid", _helperService.IsAuthorized(controllers.GetAuthorByIdAsync))
	http.Handle("Author/index", _helperService.IsAuthorized(controllers.IndexAuthorAsync))
	http.Handle("Author/update", _helperService.IsAuthorized(controllers.UpdateAuthorAsync))
	http.Handle("Author/delete", _helperService.IsAuthorized(controllers.DeleteAuthorAsync))

	// Book Operations
	http.Handle("Book/getall", _helperService.IsAuthorized(controllers.GetAllBooksAsync))
	http.Handle("Book/getbyid", _helperService.IsAuthorized(controllers.GetBookByIdAsync))
	http.Handle("Book/search", _helperService.IsAuthorized(controllers.SearchBookAsync))
	http.Handle("Book/searchbymodel", _helperService.IsAuthorized(controllers.SearchBookByModelAsync))
	http.Handle("Book/index", _helperService.IsAuthorized(controllers.IndexBookAsync))
	http.Handle("Book/update", _helperService.IsAuthorized(controllers.UpdateBookAsync))
	http.Handle("Book/delete", _helperService.IsAuthorized(controllers.DeleteBookAsync))
}
