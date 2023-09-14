package helpers

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ParseJson(body io.ReadCloser) (map[string]interface{}, error) {
    resBody, err := ioutil.ReadAll(body)
    if err != nil {
        return nil, err
    }
    resStr := string(resBody)
    resBytes := []byte(resStr)
    var jsonRes map[string]interface{}
    err = json.Unmarshal(resBytes, &jsonRes)
    return jsonRes, err
}
