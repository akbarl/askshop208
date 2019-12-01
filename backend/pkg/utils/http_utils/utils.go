package http_utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/imroc/req"
)

func ExecuteGetMethod(url string) string {
	r := req.New()
	resp, err := r.Get(url)
	if err != nil {
		log.Println(err)
	}

	log.Println(resp.String())
	return resp.String()
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
