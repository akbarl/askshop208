package lazada_utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/akbarl/askshop208/pkg/utils/lazada_utils/lazada_constants"
	"github.com/imroc/req"
)

func GenerateSignatureRequest(restUrl string, param map[string]interface{}) string {
	keys := GetKeysFromParamRequest(param)
	SortKeys(keys)
	appendedKeys := AppendKeyAndValueParams(restUrl, keys, param)
	log.Println(appendedKeys)
	signedParam := EncryptKeyParams(appendedKeys, GetAppSecret(), GetDefaultSignMethod())
	log.Println(signedParam)
	return signedParam
}

func GetAppSecret() string {
	return os.Getenv("APP_SECRET")
}

func GetAppKey() string {
	return os.Getenv("APP_KEY")
}

func GetDefaultSignMethod() func() hash.Hash {
	return sha256.New
}

func GenerateCommonRequestParameters() map[string]interface{} {
	return req.Param{
		"app_key":     GetAppKey(),
		"timestamp":   GenerateTimeStampParam(),
		"sign_method": lazada_constants.DEFAULT_SIGN_METHOD_NAME,
	}
}

func GenerateTimeStampParam() string {
	return strconv.Itoa(int(time.Now().UnixNano() / 1000000))
}

func GetKeysFromParamRequest(param map[string]interface{}) []string {
	keys := []string{}
	for k := range param {
		keys = append(keys, k)
	}
	return keys
}

func SortKeys(keys []string) {
	sort.Strings(keys)
}

func AppendKeyAndValueParams(apiUrl string, sortedKeys []string, param map[string]interface{}) string {
	var result strings.Builder
	result.WriteString(apiUrl)
	for _, v := range sortedKeys {
		result.WriteString(fmt.Sprint(v))
		result.WriteString(fmt.Sprint(param[v]))
	}
	return result.String()
}

func EncryptKeyParams(appendedKeys string, secretKey string, hash func() hash.Hash) string {
	h := hmac.New(hash, []byte(secretKey))
	h.Write([]byte(appendedKeys))
	sha := hex.EncodeToString([]byte(h.Sum(nil)))
	return strings.ToUpper(sha)
}
