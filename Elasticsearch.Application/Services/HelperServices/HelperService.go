package helperservices

import (
	helperserviceinterfaces "Elasticsearch/Elasticsearch.Application/ServiceInterfaces/HelperServiceInterfaces"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type HelperService struct{}

var mySigningKey = []byte("a65ed4d2-e7ba-4a18-8891-91d76cb41cdd")

func NewHelperService() helperserviceinterfaces.IHelperService {
	return &HelperService{}
}

func (h HelperService) IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			if token.Valid {
				fmt.Println("Token authorization successful.")
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func (h HelperService) GetAccessToken(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "client"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
	}

	resp := make(map[string]string)
	resp["Token_Expire_Duration"] = "30 Minutes"
	resp["Token"] = tokenString
	resp["Token_Authorization"] = "Full Permission"
	respJson, _ := json.Marshal(resp)

	w.Write(respJson)
}
