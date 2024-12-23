package constant

import "os"

var Config = map[string]string{
	"PORT":       os.Getenv("PORT"),
	"SECRET_KEY": os.Getenv("SECRET_KEY"),
}
