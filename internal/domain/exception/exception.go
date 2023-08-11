package exception

import "errors"

var ErrNoEntity = errors.New("entity not found")
var ErrUniqueField = errors.New("unique violation")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrInvalidForeinKey = errors.New("foreign key violation")
var ErrNoChange = errors.New("no change")
var ErrInvalidEmail = errors.New("email address invalid")
var ErrInvalidPhone = errors.New("phone number invalid")
var ErrAccountSuspended = errors.New("account suspended")
