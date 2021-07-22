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
func GetIndexOF(patt, array interface{}) (int, error) {
	refPatt := reflect.ValueOf(patt)
	refArray := reflect.ValueOf(array)

	if refPatt.Type() != refArray.Type() {
		return -1, fmt.Errorf("[error] wrong type! expect type array: %v, have type pattern: %v", refArray.String(), refPatt.String())
	}

	return -1, nil
}

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
