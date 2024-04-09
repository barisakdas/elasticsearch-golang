package helperserviceinterfaces

import "net/http"

type IHelperService interface {
	IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler
	GetAccessToken(w http.ResponseWriter, r *http.Request)
}
