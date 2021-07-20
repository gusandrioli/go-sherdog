package sherdog

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

// FighterID represents each Fighter's unique identifier (e.g. Robert-Whittaker-45132)
type FighterID string

// Fighter represents a fighter on Sherdog. It includes personal information
// fight history, record, etc.
type Fighter struct {
	Age             uint
	Assocaition     string
	Birthday        time.Time
	ProFightHisotry []Fight
	Height          string
	ImageURL        string
	ID              FighterID
	Locality        string
	Name            string
	Nationality     string
	Nickname        string
	Record          Record
	Weight          string
	WeightClass     string
}

// Record represents a record on Sherdog. Each Fighter has one Record.
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

// Fight represents a fight on Sherdog. Each Fighter has one or more fights
// on his FightHistory.
type Fight struct {
	Date     string
	Event    string
	Method   string
	Opponent string
	Referee  string
	Result   string
	Round    uint
	Time     string
}

// Find a fighter by their unique FighterID (e.g. Robert-Whittaker-45132). This ID
// can be fetched from FindFighterByName response.
func FindFighterByID(fighterID FighterID) (*Fighter, error) {
	return fetchFighter(fighterID)
}

// Find fighters by their name. Returns a slice of fighters.
func FindFighterByName(name string) ([]*Fighter, error) {
	return searchFighter(name)
}

func fetchFighter(fighterID FighterID) (*Fighter, error) {
	f := &Fighter{}

	c := colly.NewCollector()

	c.OnError(func(r *colly.Response, e error) {
		log.Printf("Error on %v: %v", r.Request.URL, e)
		return
	})

	// Age/Birthday
	c.OnHTML("span.birthday", func(h *colly.HTMLElement) {
		birthday, _ := time.Parse("2006-01-02", h.ChildText("span[itemprop=birthDate]"))
		f.Birthday = birthday

		ageWithLabel := strings.Split(h.ChildText("strong"), " ")
		ageInt, _ := strconv.Atoi(ageWithLabel[1])
		f.Age = uint(ageInt)
	})

	// Assocaition
	c.OnHTML("a.association", func(h *colly.HTMLElement) {
		f.Assocaition = h.Text
	})

	// Fight History
	var fights []Fight
	c.OnHTML("div.fight_history", func(g *colly.HTMLElement) {
		if g.ChildText("div.module_header") == "Fight History - Pro" {
			g.ForEach("table", func(_ int, table *colly.HTMLElement) {
				table.ForEach("tbody", func(_ int, tbody *colly.HTMLElement) {
					tbody.ForEach("tr", func(row int, tr *colly.HTMLElement) {
						if row != 0 {
							fight := Fight{}

							tr.ForEach("td", func(cell int, k *colly.HTMLElement) {
								switch cell {
								case 0:
									fight.Result = k.Text
								case 1:
									opponentURLSplitted := strings.Split(k.ChildAttr("a", "href"), "/")
									opponentID := opponentURLSplitted[len(opponentURLSplitted)-1]
									fight.Opponent = k.Text + fmt.Sprintf(" (%s)", opponentID)
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

					f.ProFightHisotry = fights
				})

			})
		}
	})

	// Height
	c.OnHTML("strong[itemprop=height]", func(h *colly.HTMLElement) {
		f.Height = strings.Replace(h.Text, "\\", "", -1)
	})

	// ImageURL
	c.OnHTML("img.profile_image", func(h *colly.HTMLElement) {
		f.ImageURL = baseURL + h.Attr("src")
	})

	// ID
	f.ID = fighterID

	// Locality
	c.OnHTML("span.locality", func(h *colly.HTMLElement) {
		f.Locality = h.Text
	})

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

	// Nationality
	c.OnHTML("strong[itemprop=nationality]", func(h *colly.HTMLElement) {
		f.Nationality = h.Text
	})

	// Record Wins/Losses/NoContests
	c.OnHTML("div.left_side", func(h *colly.HTMLElement) {
		h.ForEach("div.bio_graph", func(i int, j *colly.HTMLElement) {
			if i == 0 {
				wins, _ := strconv.Atoi(j.ChildText(".counter"))
				f.Record.WinsTotal = uint(wins)

				// Wins Stats
				j.ForEach("span.graph_tag", func(methodCounter int, k *colly.HTMLElement) {
					switch methodCounter {
					case 0:
						ko, _ := strconv.Atoi(string([]byte(k.Text)[0]))
						f.Record.WinsKnockouts = uint(ko)
					case 1:
						sub, _ := strconv.Atoi(string([]byte(k.Text)[0]))
						f.Record.WinsSubmissions = uint(sub)
					case 2:
						dec, _ := strconv.Atoi(string([]byte(k.Text)[0]))
						f.Record.WinsDecisions = uint(dec)
					default:
						break
					}
				})
			} else if i == 1 {
				losses, _ := strconv.Atoi(j.ChildText(".counter"))
				f.Record.LossesTotal = uint(losses)

				// Losses Stats
				j.ForEach("span.graph_tag", func(methodCounter int, k *colly.HTMLElement) {
					switch methodCounter {
					case 0:
						ko, _ := strconv.Atoi(string([]byte(k.Text)[0]))
						f.Record.LossesKnockouts = uint(ko)
					case 1:
						sub, _ := strconv.Atoi(string([]byte(k.Text)[0]))
						f.Record.LossesSubmissions = uint(sub)
					case 2:
						dec, _ := strconv.Atoi(string([]byte(k.Text)[0]))
						f.Record.LossesDecisions = uint(dec)
					default:
						break
					}
				})
			}
		})
	})

	c.OnHTML("div.count_history", func(h *colly.HTMLElement) {
		NC, _ := strconv.Atoi(h.ChildText("div.right_side div.bio_graph span.card span.counter"))
		f.Record.NoContests = uint(NC)
	})

	// Weight
	c.OnHTML("strong[itemprop=weight]", func(h *colly.HTMLElement) {
		f.Weight = h.Text
	})

	// WightClass
	c.OnHTML("strong.title", func(h *colly.HTMLElement) {
		f.WeightClass = h.Text
	})

	if err := c.Visit(URLFighter + string(fighterID)); err != nil {
		if err.Error() == "Not Found" {
			return nil, ErrFighterNotFound
		}

		return nil, err
	}

	return f, nil
}

func searchFighter(name string) ([]*Fighter, error) {
	c := colly.NewCollector()

	var fighterIDs []string

	// Search through result table
	c.OnHTML("table.fightfinder_result", func(h *colly.HTMLElement) {
		h.ForEach("tbody", func(_ int, i *colly.HTMLElement) {
			i.ForEach("tr", func(row int, j *colly.HTMLElement) {
				if row != 0 && !strings.Contains(j.ChildAttr("a", "href"), "/events") {
					fighterURLSplitted := strings.Split(j.ChildAttr("a", "href"), "/")
					fighterIDs = append(fighterIDs, fighterURLSplitted[len(fighterURLSplitted)-1])
				}
			})
		})
	})

	nameForURL := strings.Replace(name, " ", "+", -1)

	if err := c.Visit(URLFindFighter + "?SearchTxt=" + nameForURL); err != nil {
		return nil, err
	}

	if len(fighterIDs) == 0 {
		return nil, ErrFighterNotFound
	}

	var fighters []*Fighter

	for _, id := range fighterIDs {
		fighter, err := fetchFighter(FighterID(id))
		if err != nil {
			return nil, err
		}

		fighters = append(fighters, fighter)
	}

	return fighters, nil
}
