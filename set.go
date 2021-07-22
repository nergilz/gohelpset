package goset

import (
	"goset/dispatch"
	"goset/sethelper"
	"goset/setvalidator"
	"net/http"
	"sync"
)

type GoSet interface {
	GetBodyRequest(r *http.Request, out interface{}) error
	InsexOF(p interface{}, a ...interface{}) (int, error)
	GetPath(nameFile, uploadDir string) (string, error)
	Validator(args ...interface{}) error
	Dispatch(wg *sync.WaitGroup, fu interface{}, args ...interface{}) error
	GenSecret(data string) (string, error)
	GenJWTtoken(id int64, data, secret string, expires int64) (string, error)
}

func GetBodyRequest(r *http.Request, out interface{}) error {
	return sethelper.GetDataFromBodyRequest(r, out)
}

func InsexOF(p interface{}, a ...interface{}) (int, error) {
	return sethelper.GetIndexOF(p, a)
}

func GetPath(nameFile, uploadDir string) (string, error) {
	return sethelper.GetFilePathForSave(nameFile, uploadDir)
}

func Validator(args ...interface{}) error {
	return setvalidator.Validator(args)
}

func Dispatch(wg *sync.WaitGroup, fu interface{}, args ...interface{}) error {
	return dispatch.Dispatch(wg, fu, args)
}

func GenSecret(data string) (string, error) {
	return sethelper.GenerateSecret(data)
}

func GenJWTtoken(id int64, data, secret string, expires int64) (string, error) {
	return sethelper.GenerateJWTtoken(id, data, secret, expires)
}
