package sherdog

import "errors"

var (
	ErrEventNotFound        = errors.New("Event was not found")
	ErrFighterNotFound      = errors.New("Fighter was not found")
	ErrOrganizationNotFound = errors.New("Organization was not found")
)
