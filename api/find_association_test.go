package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindAssociationByName(t *testing.T) {
	tests := []struct {
		name            string
		associationName string
		wantErr         error
		wantResponse    *FindAssocaitionResponse
	}{
		{
			"successChuteBoxeMonstro",
			"Chute Boxe Monstro",
			nil,
			&FindAssocaitionResponse{
				Error:      nil,
				TotalFound: 1,
				Limit:      8,
				Time:       "0.000",
				Collection: []Association{
					{
						ID:   "178571",
						Name: []string{"Chute Boxe Monstro"},
					},
				},
			},
		},
		{
			"successNoResult",
			"asdf123-non-existent-association",
			nil,
			&FindAssocaitionResponse{
				Error:      nil,
				TotalFound: 0,
				Limit:      8,
				Time:       "0.000",
				Collection: []Association{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindAssociationByName(tt.associationName)
			assert.ErrorIs(t, tt.wantErr, err)

			if tt.wantErr == nil {
				assert.Equal(t, tt.wantResponse, got)
			}
		})
	}
}
