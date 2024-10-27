package service

import "github.com/golang-class/api/model"

type CatService interface {
	FetchImage() ([]model.CatImage, error)
}
