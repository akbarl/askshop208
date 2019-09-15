package lazada_handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akbarl/askshop208/pkg/utils/lazada_client"
	"github.com/akbarl/askshop208/pkg/utils/lazada_utils/lazada_utils"
)

type Products struct {
	Code string `json:"code"`
	Data struct {
		TotalProducts string `json:"total_products"`
		Products      []struct {
			Skus []struct {
				Status          string   `json:"Status"`
				Quantity        int      `json:"quantity"`
				ProductWeight   string   `json:"product_weight"`
				Images          []string `json:"Images"`
				SellerSku       string   `json:"SellerSku"`
				ShopSku         string   `json:"ShopSku"`
				URL             string   `json:"Url"`
				PackageWidth    string   `json:"package_width"`
				SpecialToTime   string   `json:"special_to_time"`
				SpecialFromTime string   `json:"special_from_time"`
				PackageHeight   string   `json:"package_height"`
				SpecialPrice    int      `json:"special_price"`
				Price           int      `json:"price"`
				PackageLength   string   `json:"package_length"`
				PackageWeight   string   `json:"package_weight"`
				Available       int      `json:"Available"`
				SkuID           int      `json:"SkuId"`
				SpecialToDate   string   `json:"special_to_date"`
			} `json:"skus"`
			ItemID          string `json:"item_id"`
			PrimaryCategory string `json:"primary_category"`
			Attributes      struct {
				ShortDescription string `json:"short_description"`
				Name             string `json:"name"`
				Description      string `json:"description"`
				WarrantyType     string `json:"warranty_type"`
				Brand            string `json:"brand"`
			} `json:"attributes"`
		} `json:"products"`
	} `json:"data"`
	RequestID string `json:"request_id"`
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	cacheData, found := c.Get("AuthInfo")
	if found {
		appKey := lazada_utils.GetAppKey()
		appSecret := lazada_utils.GetAppSecret()
		authInfo := cacheData.(AuthorizationResponse)
		lazadaClient := lazada_client.NewLazadaClient(appKey, appSecret)
		data := []byte(lazadaClient.GetProducts(authInfo.AccessToken))
		var productList Products
		error := json.Unmarshal(data, &productList)
		if error != nil {

		}
		jsonResponse, error := json.Marshal(productList.Data.Products)
		if error != nil {
			fmt.Println(error)
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(jsonResponse))
	}
}
