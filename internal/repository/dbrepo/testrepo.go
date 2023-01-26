package dbrepo

import (
	"errors"
	"time"

	"github.com/cxt314/learning-bnb-webapp/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservations inserts a reservation into the database
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	// if the room id is 2, then fail; otherwise, pass
	if res.RoomID == 2 {
		return -1, errors.New("some error")
	}
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	if r.RoomID == 1000 {
		return errors.New("some error")
	}
	return nil
}

// SearchAvailabilityByDatesByRoomID returns true if availability exists for roomID, and false if no availability exists
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	if roomID == 1000 {
		return false, errors.New("some error")
	}

	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of all available rooms, if any, for a given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room

	var testDate time.Time
	testDate, _ = time.Parse("2006-01-02", "2050-02-02")

	if start == testDate {
		room := models.Room{
			ID: 1,
		}
		rooms = append(rooms, room)
	}

	var testDBErrDate time.Time
	testDBErrDate, _ = time.Parse("2006-01-02", "2050-10-10")
	if start == testDBErrDate {
		return rooms, errors.New("some error")
	}

	return rooms, nil
}

// GetRoomByID gets a Room by room id
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room

	if id > 2 {
		return room, errors.New("Some error for testing")
	}

	return room, nil
}
