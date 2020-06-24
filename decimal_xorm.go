package decimal

import (
	"database/sql/driver"
	"errors"
)

// FromDB implements Conversion interface for xorm.
// This is needed because Decimal is a struct not a newtype pattern.
func (d *Decimal) FromDB(payload []byte) (err error) {
	err = d.Scan(payload)
	return
}

var (
	// ErrEmptyDecimalString specifies the empty string of decimal value.
	ErrEmptyDecimalString = errors.New("Invalid empty decimal string")
)

// ToDB implements Conversion interface for xorm.
// This is needed because Decimal is a struct not a newtype pattern.
func (d Decimal) ToDB() (payload []byte, err error) {
	var result driver.Value
	result, err = d.Value()
	if err != nil {
		return
	}
	var (
		resultString string
		ok           bool
	)
	resultString, ok = result.(string)
	if !ok {
		err = ErrEmptyDecimalString
		return
	}
	payload = []byte(resultString)
	return
}
