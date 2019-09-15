package lazada_utils

import (
	"os"
)

func GetAppSecret() string {
	return os.Getenv("APP_SECRET")
}

func GetAppKey() string {
	return os.Getenv("APP_KEY")
}
