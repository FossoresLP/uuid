package uuid

import (
	"fmt"
)

func NewV2() (uuid UUID, err error) {
	err = fmt.Errorf("UUID version 2 is not part of RFC4122bis and is not supported by this package")
	return
}
