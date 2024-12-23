package config

import "os"

var ApiConfig = map[string]string{
	"PORT":       os.Getenv("PORT"),
	"SECRET_KEY": os.Getenv("SECRET_KEY"),
}
