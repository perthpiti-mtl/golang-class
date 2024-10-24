package service

import "github.com/golang-class/di-api/model"

type CatService interface {
	FetchImage() ([]model.CatImage, error)
}
