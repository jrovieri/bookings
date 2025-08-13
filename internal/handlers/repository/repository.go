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
	GetUserByID(id int) (models.User, error)
	UpdateUser(u models.User) error
	Authenticate(email, testPassword string) (int, string, error)
}
