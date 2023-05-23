package domain

import (
	"errors"
	"fmt"
	"unicode"
)

type Contact struct {
	id          int64
	full_name   string
	phoneNumber string
}

func (c *Contact) GetId() int64 {
	return c.id
}

func (c *Contact) GetFullName() string {
	return c.full_name
}

func (c *Contact) GetPhoneNumber() string {
	return c.phoneNumber
}

func (c *Contact) SetPhoneNumber(phoneNumber string) error {
	for _, number := range phoneNumber {
		if !unicode.IsNumber(number) {
			return errors.New("phone number should have only numbers")
		}
	}
	c.phoneNumber = phoneNumber
	return nil
}

func NewContact(id int64, surname, name, patronymic, phoneNumber string) (*Contact, error) {
	contact := &Contact{}
	contact.id = id
	full_name := fmt.Sprintf("%s %s %s", surname, name, patronymic)
	contact.full_name = full_name
	err := contact.SetPhoneNumber(phoneNumber)

	if err != nil {
		return nil, err
	}

	return contact, nil
}
