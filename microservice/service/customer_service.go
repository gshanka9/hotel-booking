package service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"hotel_booking/microservice/model"
	"hotel_booking/microservice/utility"
)

func CreateCustomer(logger *logrus.Logger, customer *model.Customer) error {
	utility.Store.Mu.Lock()
	defer utility.Store.Mu.Unlock()

	if _, exists := utility.Store.Customers[customer.ID]; exists {
		return errors.New("customer already exists")
	}

	utility.Store.Customers[customer.ID] = customer
	logger.Infof("Customer created in memory: %+v", customer)
	return nil
}

func GetCustomer(logger *logrus.Logger, id string) (*model.Customer, error) {
	utility.Store.Mu.Lock()
	defer utility.Store.Mu.Unlock()

	customer, exists := utility.Store.Customers[id]
	if !exists {
		return nil, errors.New("customer not found")
	}

	return customer, nil
}

func UpdateCustomer(logger *logrus.Logger, id string, customer *model.Customer) error {
	utility.Store.Mu.Lock()
	defer utility.Store.Mu.Unlock()

	if _, exists := utility.Store.Customers[id]; !exists {
		return errors.New("customer not found")
	}

	utility.Store.Customers[id] = customer
	logger.Infof("Customer updated in memory: %+v", customer)
	return nil
}

func DeleteCustomer(logger *logrus.Logger, id string) error {
	utility.Store.Mu.Lock()
	defer utility.Store.Mu.Unlock()

	if _, exists := utility.Store.Customers[id]; !exists {
		return errors.New("customer not found")
	}

	delete(utility.Store.Customers, id)
	logger.Infof("Customer deleted from memory: %s", id)
	return nil
}
