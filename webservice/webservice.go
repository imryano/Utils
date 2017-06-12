package webservice

import (
	"net/http"
)

type webService struct {
	Name    string
	Address string
}

var webServiceAddresses = []webService{
	webService{Name: "UserRead", Address: "localhost:8080"},
	webService{Name: "UserWrite", Address: "localhost:8081"},
}

func GetServiceAddress(serviceName string) string {
	for _, w := range webServiceAddresses {
		if w.Name == serviceName {
			return w.Address
		}
	}
	return ""
}

func SetCORSProps(w http.ResponseWriter, origin string) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	return w
}
