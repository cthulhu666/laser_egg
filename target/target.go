package target

import "github.com/cthulhu666/laser-egg/laseregg"

type Target interface {
	Send(measurement api.Measurement) error
}
