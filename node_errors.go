package tree

import "github.com/pkg/errors"

var (
	ErrorNotFoundNode = errors.New("not found node")
	ErrorAlreadyExist = errors.New("key already exist in this tree")
)
