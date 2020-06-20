package types

import (
	"time"
)

// Grant is the result of the Access Request
type Grant struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
	Create  bool      `json:"create"`
	Read    bool      `json:"read"`
	Update  bool      `json:"update"`
	Delete  bool      `json:"delete"`
}
