package webserviceTest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

func RunWebServiceTest(req *http.Request, object interface{}, address string, function func(http.ResponseWriter, *http.Request)) string {
	if address != "" {
		req.RemoteAddr = address
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(object)

	if err == nil {
		req.Body = ioutil.NopCloser(b)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(function)
		handler.ServeHTTP(rr, req)
		return strings.TrimSpace(rr.Body.String())
	}
	return ""
}
