package types

import (
	"time"
)

// Grant is the result of the Access Request
type Grant struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Status  int       `json:"status"` // 0=Ok, 1 = expired
	Expires time.Time `json:"expires"`
	Create  bool      `json:"create"`
	Read    bool      `json:"read"`
	Update  bool      `json:"update"`
	Delete  bool      `json:"delete"`
}
