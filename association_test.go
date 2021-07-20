package sherdog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindAssociationByName(t *testing.T) {
	tests := []struct {
		name            string
		associationName string
		wantErr         error
		wantAssociation *Association
	}{
		{
			"successChuteBoxeMonstro",
			"Chute Boxe Monstro",
			nil,
			&Association{
				ID:   178571,
				Name: "Chute Boxe Monstro",
			},
		},
		{
			"notFound",
			"asdf123-non-existent-association",
			ErrAssociationNotFound,
			nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := FindAssociationByName(tt.associationName)
			assert.ErrorIs(t, tt.wantErr, err)

			if tt.wantErr == nil {
				for _, association := range got {
					assert.Equal(t, association.Name, tt.wantAssociation.Name)
					assert.Equal(t, association.ID, tt.wantAssociation.ID)
				}
			}
		})
	}
}
