package sherdog

import (
	"strconv"

	"github.com/gusandrioli/go-sherdog/api"
	"github.com/pkg/errors"
)

// Association represents an Association/Gym/Team on Sherdog.
// Each fighter belongs to an Association
type Association struct {
	ID   uint
	Name string
}

// Find an Association by their Name. Returns an slice of Associations
// and an error.
func FindAssociationByName(name string) ([]*Association, error) {
	res, err := api.FindAssociationByName(name)
	if err != nil {
		return nil, errors.Wrapf(err, "api.FindAssociationByName")
	}

	if res.TotalFound == 0 {
		return nil, ErrAssociationNotFound
	}

	var associations []*Association
	for _, association := range res.Collection {
		id, _ := strconv.Atoi(association.ID)

		a := &Association{
			ID:   uint(id),
			Name: association.Name[0],
		}

		associations = append(associations, a)
	}

	return associations, nil
}
