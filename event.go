package sherdog

import "time"

type EventID string

// Event represents a MMA event on Sherdog.
type Event struct {
	Date      time.Time
	FightCard []*Fight
	ID        EventID
	Location  string
	Name      string
}

// TODO
func FindAllUpcomingEvents() []*Event {
	return nil
}

// TODO
func FindEventByName(name string) []*Event {
	return nil
}

// Fetch all upcoming events of a certain organization based on their unique OrganizationID
// (e.g. Bellator-MMA-1960). Returns a slice of Events.
func FindUpcomingEventsByOrganization(organizationID OrganizationID) ([]*Event, error) {
	organization, err := fetchOrganization(organizationID)
	if err != nil {
		return nil, err
	}

	return organization.UpcomingEvents, nil
}
