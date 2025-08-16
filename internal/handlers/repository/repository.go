package repository

import (
	"time"

	"github.com/jrovieri/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(*models.Reservation) (int, error)
	InsertRoomRestriction(*models.RoomRestriction) error
	SearchAvailabilityByDatesAndRoomId(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomByID(int) (models.Room, error)
	GetUserByID(int) (models.User, error)
	UpdateUser(models.User) error
	Authenticate(string, string) (int, string, error)
	AllReservations() ([]models.Reservation, error)
	GetReservationByID(int) (models.Reservation, error)
	UpdateReservation(models.Reservation) error
	DeleteReservation(int) error
	UpdateProcessedForReservations(int, int) error
}
