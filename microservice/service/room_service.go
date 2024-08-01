package service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"hotel_booking/microservice/model"
	"hotel_booking/microservice/utility"
)

func CreateRoom(logger *logrus.Logger, room *model.Room) error {
	utility.Store.Mu.Lock()
	defer utility.Store.Mu.Unlock()

	if _, exists := utility.Store.Rooms[room.ID]; exists {
		return errors.New("room already exists")
	}

	utility.Store.Rooms[room.ID] = room
	logger.Infof("Room created in memory: %+v", room)
	return nil
}

func GetRoom(logger *logrus.Logger, id string) (*model.Room, error) {
	utility.Store.Mu.Lock()
	defer utility.Store.Mu.Unlock()

	room, exists := utility.Store.Rooms[id]
	if !exists {
		return nil, errors.New("room not found")
	}

	return room, nil
}

func UpdateRoom(logger *logrus.Logger, id string, room *model.Room) error {
	utility.Store.Mu.Lock()
	defer utility.Store.Mu.Unlock()

	if _, exists := utility.Store.Rooms[id]; !exists {
		return errors.New("room not found")
	}

	utility.Store.Rooms[id] = room
	logger.Infof("Room updated in memory: %+v", room)
	return nil
}

func DeleteRoom(logger *logrus.Logger, id string) error {
	utility.Store.Mu.Lock()
	defer utility.Store.Mu.Unlock()

	if _, exists := utility.Store.Rooms[id]; !exists {
		return errors.New("room not found")
	}

	delete(utility.Store.Rooms, id)
	logger.Infof("Room deleted from memory: %s", id)
	return nil
}
