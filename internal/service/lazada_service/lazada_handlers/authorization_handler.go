package lazada_handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/akbarl/askshop208/internal/service/lazada_service/lazada_models"
	"github.com/akbarl/askshop208/pkg/utils/http_utils"
	"github.com/akbarl/askshop208/pkg/utils/lazada_utils/lazada_constants"
	"github.com/akbarl/askshop208/pkg/utils/lazada_utils/lazada_utils"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
)

var c = cache.New(5*time.Minute, 10*time.Minute)

func NewAuthorizationHandler() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/auth", AuthenticationHandler)
	http.ListenAndServe(":8181", r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func AuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	var authInfo lazada_models.AuthorizationResponse
	item, found := c.Get("authorization_info")
	if found {
		authInfo = item.(lazada_models.AuthorizationResponse)
		log.Println(authInfo.AccessToken)
		return
	}

	code := r.URL.Query().Get("code")
	param := lazada_utils.GenerateCommonRequestParameters()
	param[lazada_constants.CODE] = code
	param[lazada_constants.SIGN] = lazada_utils.GenerateSignatureRequest(lazada_constants.GENERATE_ACCESS_TOKEN_REST_URL, param)
	queryString := http_utils.BuildQueryString(param)
	fullPath := lazada_constants.BASE_AUTH_REST_URL + lazada_constants.GENERATE_ACCESS_TOKEN_REST_URL + queryString
	log.Println(fullPath)
	bodyText := []byte(http_utils.ExecuteGetMethod(fullPath))
	var authResp lazada_models.AuthorizationResponse
	err := json.Unmarshal(bodyText, &authResp)
	var emptyAuthResp lazada_models.AuthorizationResponse = lazada_models.AuthorizationResponse{}
	if err != nil || emptyAuthResp == authResp {
		log.Println("Invalid authorization code: ", err)
		return
	}

	c.Add("authorization_info", authResp, cache.NoExpiration)

}
