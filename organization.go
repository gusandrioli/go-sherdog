package sherdog

import (
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/gusandrioli/go-sherdog/api"
	"github.com/pkg/errors"
)

const (
	UFCID      OrganizationID = "Ultimate-Fighting-Championship-UFC-2"
	BellatorID OrganizationID = "Bellator-MMA-1960"
	PFLID      OrganizationID = "Professional-Fighters-League-12241"
	OneID      OrganizationID = "One-Championship-3877"
)

// OrganizationID represents each Organization's unique identifier (e.g. Bellator-MMA-1960).
type OrganizationID string

// Organization represents an Organization/Company on Sherdog.
// E.g. UFC, Bellator, PFL. Each event belongs to an Organization.
type Organization struct {
	ID             OrganizationID
	Name           string
	PastEvents     []*Event
	UpcomingEvents []*Event
}

// Find an Organization by their unique OrganizationID (e.g. Bellator-MMA-1960).
// Returns an Organization
func FindOrganizationByID(organizationID OrganizationID) (*Organization, error) {
	return fetchOrganization(organizationID)
}

// Find an Organization by their Name. Returns an slice of Organizations
// and an error.
func FindOrganizationByName(name string) ([]*Organization, error) {
	res, err := api.FindOrganizationByName(name)
	if err != nil {
		return nil, errors.Wrapf(err, "api.FindOrganizationByName")
	}

	if res.TotalFound == 0 {
		return nil, ErrOrganizationNotFound
	}

	var organizations []*Organization
	for _, organization := range res.Collection {
		urlSplitted := strings.Split(organization.URL, "/")

		id := urlSplitted[len(urlSplitted)-1]

		org, err := FindOrganizationByID(OrganizationID(id))
		if err != nil {
			return nil, errors.Wrapf(err, "FindOrganizationByID(%v)", id)
		}

		organizations = append(organizations, org)
	}

	return organizations, nil
}

func fetchOrganization(organizationID OrganizationID) (*Organization, error) {
	organization := &Organization{}

	c := colly.NewCollector()

	// ID
	organization.ID = organizationID

	// Name
	c.OnHTML("h2[itemprop=name]", func(h *colly.HTMLElement) {
		organization.Name = h.Text
	})

	// Past Events
	c.OnHTML("div#recent_tab", func(div *colly.HTMLElement) {
		div.ForEach("table", func(_ int, table *colly.HTMLElement) {
			table.ForEach("tbody", func(_ int, tbody *colly.HTMLElement) {
				tbody.ForEach("tr", func(trCounter int, tr *colly.HTMLElement) {
					if trCounter != 0 {
						event := &Event{}

						tr.ForEach("td", func(tdCounter int, td *colly.HTMLElement) {
							switch tdCounter {
							case 0:
								date := td.ChildAttr("meta[itemprop=startDate]", "content")
								dateParsed, _ := time.Parse(time.RFC3339, date)
								event.Date = dateParsed.UTC()
							case 1:
								eventURLSplitted := strings.Split(td.ChildAttr("a[itemprop=url]", "href"), "/")
								event.ID = EventID(eventURLSplitted[len(eventURLSplitted)-1])
								event.Name = td.ChildText("span[itemprop=name]")
							case 2:
								event.Location = strings.TrimSpace(td.Text)
							}
						})

						organization.PastEvents = append(organization.PastEvents, event)
					}
				})
			})
		})
	})

	// Upcoming Events
	c.OnHTML("div#upcoming_tab", func(div *colly.HTMLElement) {
		div.ForEach("table", func(_ int, table *colly.HTMLElement) {
			table.ForEach("tbody", func(_ int, tbody *colly.HTMLElement) {
				tbody.ForEach("tr", func(trCounter int, tr *colly.HTMLElement) {
					if trCounter != 0 {
						event := &Event{}

						tr.ForEach("td", func(tdCounter int, td *colly.HTMLElement) {
							switch tdCounter {
							case 0:
								date := td.ChildAttr("meta[itemprop=startDate]", "content")
								dateParsed, _ := time.Parse(time.RFC3339, date)
								event.Date = dateParsed.UTC()
							case 1:
								eventURLSplitted := strings.Split(td.ChildAttr("a[itemprop=url]", "href"), "/")
								event.ID = EventID(eventURLSplitted[len(eventURLSplitted)-1])
								event.Name = td.ChildText("span[itemprop=name]")
							case 2:
								event.Location = strings.TrimSpace(td.Text)
							}
						})

						organization.UpcomingEvents = append(organization.UpcomingEvents, event)
					}
				})
			})
		})
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Printf("Error on %v: %v", r.Request.URL, e)
	})

	if err := c.Visit(URLOrganizations + string(organizationID)); err != nil {
		if err.Error() == "Not Found" {
			return nil, ErrOrganizationNotFound
		}

		return nil, err
	}

	return organization, nil
}
