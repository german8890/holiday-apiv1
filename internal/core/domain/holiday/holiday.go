package holiday

import (
	"context"
	"encoding/json"
	"errors"
)

func UnmarshalHolidays(data []byte) (Holidays, error) {
	var r Holidays
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Holidays) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Holiday) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Holiday) UnmarshalToXML(result interface{}) error {
	b, err := r.Marshal()
	if err != nil {
		return err
	}
	return json.Unmarshal(b, result)
}

func (r *Holiday) ToXMLInterface() (HolidayXML, error) {
	var xml HolidayXML
	err := r.UnmarshalToXML(&xml)
	return xml, err
}

// Holidays is an array of Holiday
//
// This type is used for unmarshaling JSON.
//
// swagger:model Holiday
type Holidays struct {
	// The holidays
	// in: body
	Holidays []Holiday `json:"holidays,omitempty" xml:"holidays,omitempty"`
}

// Holiday is a holiday
//
// This type is used for unmarshaling JSON.
//
// swagger:model Holiday
type Holiday struct {
	Date        string `json:"date,omitempty" xml:"date,omitempty"`
	Title       string `json:"title,omitempty" xml:"title,omitempty"`
	Type        string `json:"type,omitempty" xml:"type,omitempty"`
	Inalienable bool   `json:"inalienable,omitempty" xml:"inalienable,omitempty"`
	Extra       string `json:"extra,omitempty" xml:"extra,omitempty"`
}

// HolidayXML is a holiday
//
// This type is used for unmarshaling JSON.
//
// swagger:model HolidayXML
type HolidayXML struct {
	// date
	Date string `xml:"date,omitempty"`
	// tittle
	Title string `xml:"title,omitempty"`
	//type
	Type string `xml:"type,omitempty"`
	//inalienable
	Inalienable bool `xml:"inalienable,omitempty"`
	// extra
	Extra string `xml:"extra,omitempty"`
}

//go:generate mockgen -package mocks -destination mocks/holidays_repository_mocks.go . Repository

func ValidateHolidayExtra(name string) error {
	if name == "" {
		return errors.New("extra is empty")
	}
	return nil
}

func ValidateHolidayDate(date string) error {
	if date == "" {
		return errors.New("Date is empty")
	}
	return nil
}

// Repository is the storage abstraction
type Repository interface {
	Find(ctx context.Context, extra string) ([]Holiday, error)
	FindOne(ctx context.Context, extra string) (Holiday, error)
	LoadAll(ctx context.Context, holidays []Holiday) error
}
