package entities

import (
	"gorm.io/gorm"

	"github.com/bxcodec/faker/v3"
)

type BaseEntity interface {
	CreateFakeData(amount uint) []BaseEntity
}

type Song struct {
	gorm.Model
	Title string `gorm:"type:varchar(200);uniqueIndex:idx_title_author"`
	Author string `gorm:"type:varchar(200);uniqueIndex:idx_title_author"`
}

func (s *Song) CreateFakeData(amount uint) []BaseEntity {
	songs := []BaseEntity{}
	for i := uint(0); i < amount; i++ {
		songs = append(songs, &Song{
			Title: faker.Sentence(),
			Author: faker.Name(),
		})
	}

	return songs
}

func NewSong() BaseEntity {
	return &Song{}
}