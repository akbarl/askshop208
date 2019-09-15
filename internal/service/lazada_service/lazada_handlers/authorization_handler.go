package lazada_handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/akbarl/askshop208/pkg/utils/lazada_client"
	"github.com/akbarl/askshop208/pkg/utils/lazada_utils/lazada_utils"
	"github.com/patrickmn/go-cache"
)

var c = cache.New(5*time.Minute, 10*time.Minute)

type AuthorizationResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpriesIn    int64  `json:"expires_in"`
	Account      string `json:"account"`
}

func AuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	appKey := lazada_utils.GetAppKey()
	appSecret := lazada_utils.GetAppSecret()
	code := r.URL.Query().Get("code")
	lazadaClient := lazada_client.NewLazadaClient(appKey, appSecret)
	data := []byte(lazadaClient.GetAccessToken(code))
	var authInfo AuthorizationResponse
	error := json.Unmarshal(data, &authInfo)
	if error != nil {
		fmt.Println("Could not find access token")
		return
	}

	if authInfo.AccessToken != "" {
		c.Set("AuthInfo", authInfo, cache.NoExpiration)
	}
}
