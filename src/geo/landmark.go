package geo

import "errors"

type Landmark struct {
	name string
	Coordinates
}

// GET
func (l *Landmark) Name() string {
	return l.name
}

// SET
func (l *Landmark) SetName(name string) error {
	if name == "" {
		return errors.New("Invalid name")
	}
	l.name = name
	return nil
}
