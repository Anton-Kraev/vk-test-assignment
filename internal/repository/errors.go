package repository

import "fmt"

type ErrContainerNotFound struct {
	op string
	id int
}

func NewErrContainerNotFound(op string, id int) ErrContainerNotFound {
	return ErrContainerNotFound{op: op, id: id}
}

func (e ErrContainerNotFound) Error() string {
	return fmt.Sprintf("%s: container with id %d not found", e.op, e.id)
}
