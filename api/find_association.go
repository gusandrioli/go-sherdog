package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

type FindAssociationResponse struct {
	Error      error         `json:"error"`
	TotalFound int           `json:"total_found"`
	Limit      int           `json:"limit"`
	Time       string        `json:"time"`
	Collection []Association `json:"collection"`
}

type Association struct {
	ID   string   `json:"id"`
	Name []string `json:"name"`
}

func FindAssociationByName(name string) (*FindAssociationResponse, error) {
	nameForURL := strings.Replace(name, " ", "+", -1)

	res, err := http.Get(SearchAssociationsURL + "?q=" + nameForURL)
	if err != nil {
		return nil, errors.Wrapf(err, "http.Get")
	}
	defer res.Body.Close()

	associationResp := &FindAssociationResponse{}

	json.NewDecoder(res.Body).Decode(associationResp)

	return associationResp, nil
}
