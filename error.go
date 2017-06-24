package door

import (
	"errors"
	"fmt"
)

// ErrorNotFound 没找到错误.
var ErrorNotFound = errors.New("Not found")

// ErrorNotNum 编号错误.
type ErrorNotNum struct {
	Num uint32
}

func (errorNotNum ErrorNotNum) Error() string {
	return fmt.Sprintf("Num is error: %d", errorNotNum.Num)
}
