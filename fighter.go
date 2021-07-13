package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type FighterID string

var (
	baseURL = "https://www.sherdog.com/"
)

// TODO
type Fighter struct {
	Age         uint
	Assocaition string
	Birthday    time.Time
	Fights      []Fight
	Height      string
	ImageURL    string
	Locality    string
	Name        string
	Nationality string
	Nickname    string
	Record      Record
	Weight      string
	WeightClass string
}

// TODO
type Record struct {
	LossesDecisions   uint
	LossesKnockouts   uint
	LossesSubmissions uint
	LossesTotal       uint
	NoContests        uint
	WinsDecisions     uint
	WinsKnockouts     uint
	WinsSubmissions   uint
	WinsTotal         uint
}

// TODO
type Fight struct {
	Date     string
	Event    string
	Method   string
	Opponent string // TODO: could be another Fighter struct but could cause performance issues.
	Referee  string
	Result   string
	Round    uint
	Time     string
}

func FindFighterByID(fighterID FighterID) *Fighter {
	return fetchFighter(fighterID)
}

func FindFighterByName(name string) []*Fighter {
	return searchFighter(name)
}

func fetchFighter(fighterID FighterID) *Fighter {
	f := &Fighter{}

	c := colly.NewCollector()

	// Name and Nickname
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		e.ForEach("span", func(i int, h *colly.HTMLElement) {
			switch i {
			case 0:
				f.Name = h.Text
			case 1:
				f.Nickname = strings.Replace(h.Text, "\\", "", -1) // TODO sanitize
			}
		})
	})

	// Age/Birthday
	c.OnHTML("span.birthday", func(h *colly.HTMLElement) {
		h.ForEach("", func(i int, h *colly.HTMLElement) {
			switch i {
			case 0:
				birthday, _ := time.Parse("2006-01-02", h.Text)
				f.Birthday = birthday
			case 1:

			}
		})
	})

	// Nationality
	c.OnHTML("strong[itemprop=nationality]", func(h *colly.HTMLElement) {
		f.Nationality = h.Text
	})

	// Locality
	c.OnHTML("span.locality", func(h *colly.HTMLElement) {
		f.Locality = h.Text
	})

	// Height
	c.OnHTML("strong[itemprop=height]", func(h *colly.HTMLElement) {
		f.Height = strings.Replace(h.Text, "\\", "", -1)
	})

	// Weight
	c.OnHTML("strong[itemprop=weight]", func(h *colly.HTMLElement) {
		f.Weight = h.Text
	})

	// WightClass
	c.OnHTML("strong.title", func(h *colly.HTMLElement) {
		f.WeightClass = h.Text
	})

	// Gym
	c.OnHTML("a.association", func(h *colly.HTMLElement) {
		f.Assocaition = h.Text
	})

	// TODO Record Wins
	c.OnHTML("div.bio_graph", func(h *colly.HTMLElement) {
		h.ForEach(".card", func(i int, h *colly.HTMLElement) {
			if i == 0 {
				// spew.Dump("Wins", h.Text)
			}
		})
	})

	// Record Losses
	c.OnHTML("div.loser", func(h *colly.HTMLElement) {
		losses, _ := strconv.Atoi(h.ChildText(".counter"))
		f.Record.LossesTotal = uint(losses)

		// Losses Stats
		h.ForEach(".graph_tag", func(i int, j *colly.HTMLElement) {
			switch i {
			case 0:
				dec, _ := strconv.Atoi(string([]byte(j.Text)[0]))
				f.Record.LossesKnockouts = uint(dec)
			case 1:
				dec, _ := strconv.Atoi(string([]byte(j.Text)[0]))
				f.Record.LossesSubmissions = uint(dec)
			case 2:
				dec, _ := strconv.Atoi(string([]byte(j.Text)[0]))
				f.Record.LossesDecisions = uint(dec)
			default:
				break
			}
		})
	})

	// Fights
	var fights []Fight
	c.OnHTML("tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(row int, j *colly.HTMLElement) {
			if row != 0 {
				fight := Fight{}

				j.ForEach("td", func(cell int, k *colly.HTMLElement) {
					switch cell {
					case 0:
						fight.Result = k.Text
					case 1:
						fight.Opponent = k.Text
					case 2:
						fight.Event = k.ChildText("span[itemprop=award]")
						fight.Date = k.ChildText("span.sub_line")
					case 3:
						ref := k.ChildText("span.sub_line")
						methodWithRef := []byte(k.Text)
						fight.Method = string(methodWithRef[:(len(methodWithRef) - len(ref))]) // Hacky way to fetch method
						fight.Referee = ref
					case 4:
						r, _ := strconv.Atoi(k.Text)
						fight.Round = uint(r)
					case 5:
						fight.Time = k.Text
					}

				})

				fights = append(fights, fight)
			}
		})

		f.Fights = fights
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Println("error:", e, r.Request.URL, string(r.Body))
	})

	if err := c.Visit("https://www.sherdog.com/fighter/" + string(fighterID)); err != nil {
		log.Fatal(err)
	}

	return f
}

func searchFighter(name string) []*Fighter {
	return nil
}
