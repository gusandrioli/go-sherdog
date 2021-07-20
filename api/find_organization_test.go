package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindOrganizationByName(t *testing.T) {
	tests := []struct {
		name            string
		associationName string
		wantErr         error
		wantResponse    *FindOrganizationResponse
	}{
		{
			"successChuteBoxeMonstro",
			"Ultimate Fighting Championship (UFC)",
			nil,
			&FindOrganizationResponse{
				Error:      nil,
				TotalFound: 1,
				Limit:      8,
				Time:       "0.000",
				Collection: []Organization{
					{
						ID:   "2",
						Name: "Ultimate Fighting Championship (UFC)",
						URL:  "/organizations/Ultimate-Fighting-Championship-UFC-2",
					},
				},
			},
		},
		{
			"successNoResult",
			"asdf123-non-existent-association",
			nil,
			&FindOrganizationResponse{
				Error:      nil,
				TotalFound: 0,
				Limit:      8,
				Time:       "0.000",
				Collection: []Organization{},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := FindOrganizationByName(tt.associationName)
			assert.ErrorIs(t, tt.wantErr, err)

			if tt.wantErr == nil {
				assert.Equal(t, tt.wantResponse, got)
			}
		})
	}
}
