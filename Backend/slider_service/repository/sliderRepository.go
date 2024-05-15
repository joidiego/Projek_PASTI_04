package repository

import (
	"log"

	"github.com/joidiego/PASTI_KEl04/entity"
	"gorm.io/gorm"
)

type SliderRepository interface {
	InsertSlider(slider entity.Slider) entity.Slider
	UpdateSlider(slider entity.Slider) entity.Slider
	All() []entity.Slider
	FindByID(sliderID uint64) entity.Slider
	DeleteSlider(slider entity.Slider)
	Count() int64
}

type sliderConnection struct {
	connection *gorm.DB
}

func NewSliderRepository(db *gorm.DB) SliderRepository {
	return &sliderConnection{
		connection: db,
	}
}

func (db *sliderConnection) InsertSlider(slider entity.Slider) entity.Slider {
	db.connection.Save(&slider)
	return slider
}

func (db *sliderConnection) UpdateSlider(slider entity.Slider) entity.Slider {
	db.connection.Save(&slider)
	return slider
}
func (db *sliderConnection) All() []entity.Slider {
	var sliders []entity.Slider
	result := db.connection.Find(&sliders)
	if result.Error != nil {
		log.Printf("Error fetching all sliders: %v", result.Error)
	} else {
		log.Printf("Fetched %d sliders", len(sliders))
	}
	return sliders
}

func (db *sliderConnection) FindByID(sliderID uint64) entity.Slider {
	var slider entity.Slider
	result := db.connection.First(&slider, sliderID)
	if result.Error != nil {
		log.Printf("Error finding slider by ID %d: %v", sliderID, result.Error)
	} else {
		log.Printf("Found slider: %v", slider)
	}
	return slider
}

func (db *sliderConnection) DeleteSlider(slider entity.Slider) {
	result := db.connection.Delete(&slider)
	if result.Error != nil {
		log.Printf("Error deleting slider: %v", result.Error)
	}
}

func (db *sliderConnection) Count() int64 {
	var count int64
	result := db.connection.Model(&entity.Slider{}).Count(&count)
	if result.Error != nil {
		log.Printf("Error counting sliders: %v", result.Error)
	}
	return count
}
