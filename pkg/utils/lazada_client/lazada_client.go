package lazada_client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/akbarl/askshop208/pkg/utils/http_utils"
	"github.com/imroc/req"
)

type LazadaClient struct {
	appKey    string
	appSecret string
}

func (client *LazadaClient) NewLazadaClient(appKey string, appSecret string) {
	client.appKey = appKey
	client.appSecret = appSecret
}

func NewLazadaClient(appKey string, appSecret string) *LazadaClient {
	return &LazadaClient{appKey, appSecret}
}

type LazadaRequest struct {
	httpMethod string
	endpoint   string
	apiUrl     string
	signMethod string
	params     map[string]interface{}
}

func NewLazadaRequest() (request *LazadaRequest) {
	lazadaRequest := new(LazadaRequest)
	lazadaRequest.params = make(map[string]interface{})
	lazadaRequest.signMethod = "sha256"
	return lazadaRequest
}

func (request *LazadaRequest) SetHttpMethod(methodName string) {
	request.httpMethod = methodName
}

func (request *LazadaRequest) SetEndPoint(endpoint string) {
	request.endpoint = endpoint
}

func (request *LazadaRequest) SetUrl(apiUrl string) {
	request.apiUrl = apiUrl
}

func (request *LazadaRequest) AddParam(name string, value interface{}) {
	request.params[name] = value
}

func (request *LazadaRequest) AddParams(params map[string]interface{}) {
	for k, v := range params {
		request.AddParam(k, v)
	}
}

func (request *LazadaRequest) GenerateSignature(client *LazadaClient) string {
	keys := GetKeysFromParamRequest(request.params)
	log.Println(keys)
	appendedKeys := AppendKeyAndValueParams(request.apiUrl, keys, request.params)
	log.Println(appendedKeys)
	signedParam := EncryptKeyParams(appendedKeys, client.appSecret, GetSignMethod(request.signMethod))
	log.Println(signedParam)
	return signedParam
}

func (request *LazadaRequest) GetFullPathUrl() string {
	return request.endpoint + request.apiUrl + BuildQueryString(request.params)
}

func (client *LazadaClient) Execute(request *LazadaRequest, accessToken string) string {
	if accessToken != "" {
		request.AddParam("access_token", accessToken)
	}
	commonParams := GenerateCommonRequestParameters(client, request.signMethod)
	request.AddParams(commonParams)
	request.AddParam("sign", request.GenerateSignature(client))
	fullPath := request.GetFullPathUrl()
	return http_utils.ExecuteGetMethod(fullPath)
}

func (client *LazadaClient) ExecuteWithoutAccessToken(request *LazadaRequest) string {
	return client.Execute(request, "")
}

func (client *LazadaClient) GetProducts(accessToken string) string {
	lazadaRequest := NewLazadaRequest()
	lazadaRequest.SetHttpMethod("GET")
	lazadaRequest.SetEndPoint("https://api.lazada.vn/rest")
	lazadaRequest.SetUrl("/products/get")
	lazadaRequest.AddParam("filter", "live")
	lazadaRequest.AddParam("limit", "10")
	return client.Execute(lazadaRequest, accessToken)
}

func (client *LazadaClient) GetAccessToken(code string) string {
	lazadaRequest := NewLazadaRequest()
	lazadaRequest.SetHttpMethod("GET")
	lazadaRequest.SetEndPoint("https://auth.lazada.com/rest")
	lazadaRequest.SetUrl("/auth/token/create")
	lazadaRequest.AddParam("code", code)
	return client.ExecuteWithoutAccessToken(lazadaRequest)
}

func (client *LazadaClient) RefreshAccessToken(refreshToken string) string {
	lazadaRequest := NewLazadaRequest()
	lazadaRequest.SetHttpMethod("GET")
	lazadaRequest.SetEndPoint("https://auth.lazada.com/rest")
	lazadaRequest.SetUrl("/auth/token/refresh")
	lazadaRequest.AddParam("refresh_token", refreshToken)
	return client.ExecuteWithoutAccessToken(lazadaRequest)
}
func GenerateCommonRequestParameters(client *LazadaClient, signMethodName string) map[string]interface{} {
	return req.Param{
		"app_key":     client.appKey,
		"timestamp":   GenerateTimeStampParam(),
		"sign_method": signMethodName,
	}
}

func GetKeysFromParamRequest(param map[string]interface{}) []string {
	keys := []string{}
	for k := range param {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func AppendKeyAndValueParams(apiUrl string, sortedKeys []string, param map[string]interface{}) string {
	var result strings.Builder
	result.WriteString(apiUrl)
	for _, key := range sortedKeys {
		result.WriteString(fmt.Sprint(key))
		result.WriteString(fmt.Sprint(param[key]))
	}
	return result.String()
}

func EncryptKeyParams(appendedKeys string, secretKey string, hash func() hash.Hash) string {
	h := hmac.New(hash, []byte(secretKey))
	h.Write([]byte(appendedKeys))
	sha := hex.EncodeToString([]byte(h.Sum(nil)))
	return strings.ToUpper(sha)
}

func GenerateTimeStampParam() string {
	return strconv.Itoa(int(time.Now().UnixNano() / 1000000))
}

func GetSignMethod(methodName string) func() hash.Hash {
	if methodName == "sha256" {
		return sha256.New
	}
	return nil
}

func BuildQueryString(param map[string]interface{}) string {
	result := ""
	if len(param) > 0 {
		var queryString strings.Builder
		queryString.WriteString("?")
		for k, v := range param {
			queryString.WriteString(k)
			queryString.WriteString("=")
			queryString.WriteString(fmt.Sprint(v))
			queryString.WriteString("&")
		}
		result = queryString.String()[:len(queryString.String())-1]
	}
	return result
}
