package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

type FindOrganizationResponse struct {
	Error      error          `json:"error"`
	TotalFound int            `json:"total_found"`
	Limit      int            `json:"limit"`
	Time       string         `json:"time"`
	Collection []Organization `json:"collection"`
}

type Organization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func FindOrganizationByName(name string) (*FindOrganizationResponse, error) {
	nameForURL := strings.Replace(name, " ", "+", -1)

	res, err := http.Get(SearchOrganizationsURL + "?q=" + nameForURL)
	if err != nil {
		return nil, errors.Wrapf(err, "http.Get")
	}
	defer res.Body.Close()

	organizationResp := &FindOrganizationResponse{}

	json.NewDecoder(res.Body).Decode(organizationResp)

	return organizationResp, nil
}
