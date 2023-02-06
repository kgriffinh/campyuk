package data

import (
	"campyuk-api/features/booking"
	"errors"
	"log"

	"gorm.io/gorm"
)

type bookingData struct {
	db *gorm.DB
}

func New(db *gorm.DB) booking.BookingData {
	return &bookingData{
		db: db,
	}
}

func (bd *bookingData) Create(userID uint, newBooking booking.Core) (booking.Core, error) {
	model := ToData(userID, newBooking)
	tx := bd.db.Create(&model)
	if tx.Error != nil {
		return booking.Core{}, tx.Error
	}

	return booking.Core{ID: model.ID}, nil
}

func (bd *bookingData) Update(userID uint, role string, updateBooking booking.Core) error {
	return nil
}

func (bd *bookingData) List(userID uint) ([]booking.Core, error) {
	return []booking.Core{}, nil
}

func (bd *bookingData) GetByID(userID uint, bookingID uint) (booking.Core, error) {
	return booking.Core{}, nil
}

func (bd *bookingData) Callback(ticket string, status string) error {
	err := bd.db.Model(&Booking{}).Where("ticket = ?", ticket).Update("status", status).Error
	if err != nil {
		return err
	}

	return nil
}

func (bd *bookingData) RequestHost(bookingID uint, status string) error {
	err := bd.db.Model(&Booking{}).Where("id = ?", bookingID).Update("status", status)
	if err != nil {
		return errors.New("failed to update status")
	}

	affrows := err.RowsAffected
	if affrows <= 0 {
		log.Println("no rows affected")
		return errors.New("no data updated")
	}

	return nil
}
