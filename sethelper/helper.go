package sethelper

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GetIndexOF
func GetIndexOF(patt, items interface{}) (int, error) {
	vpatt := reflect.ValueOf(patt)
	vitems := reflect.ValueOf(items)

	if vpatt.Type().Kind() != vitems.Type().Elem().Kind() {
		return -1, fmt.Errorf("wrong type! expect type array: %v, got type pattern: %v", vitems.Type().String(), vpatt.Type().String())
	}
	if vitems.Type().Kind() != reflect.Slice && vitems.Type().Kind() != reflect.Array {
		return -1, fmt.Errorf("arg not slice || array")
	}
	if vitems.Len() == 0 {
		return -1, fmt.Errorf("array is empty")
	}

	for i := 0; i < vitems.Len(); i++ {
		if router(vpatt, vitems.Index(i)) {
			return i, nil
		}
	}
	return -1, fmt.Errorf("no element in array")
}

func router(patt, elem reflect.Value) bool {
	switch elem.Kind() {
	case reflect.Int, reflect.Uint, reflect.Int16, reflect.Int32, reflect.Int64:
		return patt.Int() == elem.Int()
	case reflect.Float32, reflect.Float64:
		return patt.Float() == elem.Float()
	case reflect.String:
		return patt.String() == elem.String()
	}
	return false
}

// func float64Compare(patt, elem reflect.Value) bool { return patt.Float() == elem.Float() }

// func int64Compare(patt, elem reflect.Value) bool { return patt.Int() == elem.Int() }

// func stringCompare(patt, elem reflect.Value) bool { return patt.String() == elem.String() }

// func structCompare(patt, elem reflect.Value) bool { return string(patt.Bytes()) == string(elem.Bytes()) }

// GenerateSecret
func GenerateSecret(data string) (string, error) {
	hash1 := sha256.New()
	_, err := hash1.Write([]byte(data))
	if err != nil {
		return "", err
	}
	hash2 := md5.New()
	_, err = hash2.Write([]byte(data))
	if err != nil {
		return "", err
	}
	hash3 := sha256.New()
	_, err = hash3.Write([]byte(string(hash1.Sum(nil)) + string(hash2.Sum(nil))))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash3.Sum(nil)), nil
}

// GenerateJWTtoken
func GenerateJWTtoken(id, data, secret string, expires int64) (string, error) {
	claims := &JWTClaimsModel{
		ID:   id,
		Data: data,
		Time: time.Now(),
	}
	claims.ExpiresAt = time.Now().Add(time.Hour * time.Duration(expires)).Unix()

	JWTtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := JWTtoken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

// GetDataFromBodyRequest
func GetDataFromBodyRequest(r *http.Request, outModel interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("request body is empty")
	}
	buf, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, outModel)
}

// GetFilePathForSave
func GetFilePathForSave(nameFile, uploadDir string) (path string, err error) {

	return "", nil
}
