package sherdog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_FindOrganizationByID(t *testing.T) {
	tests := []struct {
		name             string
		organizationID   OrganizationID
		wantErr          error
		wantOrganization *Organization
	}{
		{
			"successUFC",
			UFCID,
			nil,
			&Organization{
				ID:   UFCID,
				Name: "Ultimate Fighting Championship (UFC)",
				UpcomingEvents: []*Event{
					{
						Date:     time.Date(2021, time.July, 24, 7, 0, 0, 0, time.UTC),
						ID:       "UFC-on-ESPN-27-Sandhagen-vs-Dillashaw-89119",
						Location: "UFC Apex, Las Vegas, Nevada, United States",
						Name:     "UFC on ESPN 27 - Sandhagen vs. Dillashaw",
					},
				},
			},
		},
		{
			"notFound",
			"non-existent-organization",
			ErrOrganizationNotFound,
			nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := FindOrganizationByID(tt.organizationID)
			assert.ErrorIs(t, tt.wantErr, err)

			if tt.wantErr == nil {
				assert.Equal(t, got.Name, tt.wantOrganization.Name)
				assert.Equal(t, got.ID, tt.wantOrganization.ID)
				assert.Equal(t, got.UpcomingEvents[0], tt.wantOrganization.UpcomingEvents[0])
			}
		})
	}
}

func Test_FindOrganizationByName(t *testing.T) {
	tests := []struct {
		name             string
		organizationName string
		wantErr          error
		wantOrganization *Organization
	}{
		{
			"successUFC",
			"Ultimate Fighting Championship (UFC)",
			nil,
			&Organization{
				ID:   UFCID,
				Name: "Ultimate Fighting Championship (UFC)",
				UpcomingEvents: []*Event{
					{
						Date:     time.Date(2021, time.July, 24, 7, 0, 0, 0, time.UTC),
						ID:       "UFC-on-ESPN-27-Sandhagen-vs-Dillashaw-89119",
						Location: "UFC Apex, Las Vegas, Nevada, United States",
						Name:     "UFC on ESPN 27 - Sandhagen vs. Dillashaw",
					},
				},
			},
		},
		{
			"notFound",
			"asdf123-non-existent-organization",
			ErrOrganizationNotFound,
			nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := FindOrganizationByName(tt.organizationName)
			assert.ErrorIs(t, tt.wantErr, err)

			if tt.wantErr == nil {
				for _, organization := range got {
					assert.Equal(t, organization.Name, tt.wantOrganization.Name)
					assert.Equal(t, organization.ID, tt.wantOrganization.ID)
					assert.Equal(t, organization.UpcomingEvents[0], tt.wantOrganization.UpcomingEvents[0])
				}
			}
		})
	}
}
