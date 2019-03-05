package openapi

import "github.com/bmc-toolbox/redgopher/oem"

// Oem specific
type Oem struct {
	Hpe *oem.Hpe `json:"Hpe"`
}
