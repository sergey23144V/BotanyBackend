package errors

import "errors"

var EmptyIdError = errors.New("id is empty")

var NilArgumentError = errors.New("argument is nil")

var NoTransactionError = errors.New("transaction is not opened")

var NotDelete = errors.New("not delete")

var NotAuthorization = errors.New("Not Authorization")

var BadRepeatedFieldMaskTpl = "unexpected fieldmask count %d for objects count %d"
