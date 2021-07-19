package sherdog

import "errors"

var (
	ErrAssociationNotFound  = errors.New("Association was not found")
	ErrEventNotFound        = errors.New("Event was not found")
	ErrFighterNotFound      = errors.New("Fighter was not found")
	ErrOrganizationNotFound = errors.New("Organization was not found")
)
