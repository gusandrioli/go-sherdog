package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_FindFighterByID(t *testing.T) {
	tests := []struct {
		name        string
		fighterID   FighterID
		wantErr     error
		wantFighter *Fighter
	}{
		{
			name:      "successRobertWittaker",
			fighterID: "Robert-Whittaker-45132",
			wantErr:   nil,
			wantFighter: &Fighter{
				Age:         30,
				Assocaition: "PMA Super Martial Arts",
				Birthday:    time.Date(1990, time.December, 20, 0, 0, 0, 0, time.UTC),
				FightHisotry: []Fight{
					{
						Date:     "Apr / 17 / 2021",
						Event:    "UFC on ESPN 22 - Whittaker vs. Gastelum",
						Method:   "Decision (Unanimous)",
						Opponent: "Kelvin Gastelum (Kelvin-Gastelum-74700)",
						Referee:  "Herb Dean",
						Result:   "win",
						Round:    5,
						Time:     "5:00",
					},
					{
						Date:     "Oct / 24 / 2020",
						Event:    "UFC 254 - Nurmagomedov vs. Gaethje",
						Method:   "Decision (Unanimous)",
						Opponent: "Jared Cannonier (Jared-Cannonier-78628)",
						Referee:  "Anders Ohlsson",
						Result:   "win",
						Round:    3,
						Time:     "5:00",
					},
					{
						Date:     "Jul / 25 / 2020",
						Event:    "UFC on ESPN 14 - Whittaker vs. Till",
						Method:   "Decision (Unanimous)",
						Opponent: "Darren Till (Darren-Till-73436)",
						Referee:  "Herb Dean",
						Result:   "win",
						Round:    5,
						Time:     "5:00",
					},
					{
						Date:     "Oct / 05 / 2019",
						Event:    "",
						Method:   "KO (Punches)",
						Opponent: "Israel Adesanya (Israel-Adesanya-56374)",
						Referee:  "Marc Goddard",
						Result:   "loss",
						Round:    2,
						Time:     "3:33",
					},
					{
						Date:     "Jun / 09 / 2018",
						Event:    "UFC 225 - Whittaker vs. Romero 2",
						Method:   "Decision (Split)",
						Opponent: "Yoel Romero (Yoel-Romero-60762)",
						Referee:  "Dan Miragliotta",
						Result:   "win",
						Round:    5,
						Time:     "5:00",
					},
					{
						Date:     "Jul / 08 / 2017",
						Event:    "UFC 213 - Romero vs. Whittaker",
						Method:   "Decision (Unanimous)",
						Opponent: "Yoel Romero (Yoel-Romero-60762)",
						Referee:  "John McCarthy",
						Result:   "win",
						Round:    5,
						Time:     "5:00",
					},
					{
						Date:     "Apr / 15 / 2017",
						Event:    "UFC on Fox 24 - Johnson vs. Reis",
						Method:   "TKO (Head Kick and Punches)",
						Opponent: "Ronaldo Souza (Ronaldo-Souza-8394)",
						Referee:  "Mario Yamasaki",
						Result:   "win",
						Round:    2,
						Time:     "3:28",
					},
					{
						Date:     "Nov / 26 / 2016",
						Event:    "UFC Fight Night 101 - Whittaker vs. Brunson",
						Method:   "TKO (Head Kick and Punches)",
						Opponent: "Derek Brunson (Derek-Brunson-68494)",
						Referee:  "Herb Dean",
						Result:   "win",
						Round:    1,
						Time:     "4:07",
					},
					{
						Date:     "Apr / 23 / 2016",
						Event:    "UFC 197 - Jones vs. St. Preux",
						Method:   "Decision (Unanimous)",
						Opponent: "Rafael Natal (Rafael-Natal-13968)",
						Referee:  "John McCarthy",
						Result:   "win",
						Round:    3,
						Time:     "5:00",
					},
					{
						Date:     "Nov / 14 / 2015",
						Event:    "UFC 193 - Rousey vs. Holm",
						Method:   "Decision (Unanimous)",
						Opponent: "Uriah Hall (Uriah-Hall-14210)",
						Referee:  "Steve Perceval",
						Result:   "win",
						Round:    3,
						Time:     "5:00",
					},
					{
						Date:     "May / 10 / 2015",
						Event:    "UFC Fight Night 65 - Miocic vs. Hunt",
						Method:   "KO (Punches)",
						Opponent: "Brad Tavares (Brad-Tavares-33095)",
						Referee:  "Steve Perceval",
						Result:   "win",
						Round:    1,
						Time:     "0:44",
					},
					{
						Date:     "Nov / 07 / 2014",
						Event:    "UFC Fight Night 55 - Rockhold vs. Bisping",
						Method:   "TKO (Knee and Punches)",
						Opponent: "Clint Hester (Clint-Hester-43866)",
						Referee:  "John Sharp",
						Result:   "win",
						Round:    2,
						Time:     "2:43",
					},
					{
						Date:     "Jun / 28 / 2014",
						Event:    "UFC Fight Night 43 - Te Huna vs. Marquardt",
						Method:   "Decision (Unanimous)",
						Opponent: "Mike Rhodes (Mike-Rhodes-87296)",
						Referee:  "Steve Perceval",
						Result:   "win",
						Round:    3,
						Time:     "5:00",
					},
					{
						Date:     "Feb / 22 / 2014",
						Event:    "",
						Method:   "TKO (Punches)",
						Opponent: "Stephen Thompson (Stephen-Thompson-59608)",
						Referee:  "Mario Yamasaki",
						Result:   "loss",
						Round:    1,
						Time:     "3:43",
					},
					{
						Date:     "Aug / 28 / 2013",
						Event:    "",
						Method:   "Decision (Split)",
						Opponent: "Court McGee (Court-McGee-34124)",
						Referee:  "Herb Dean",
						Result:   "loss",
						Round:    3,
						Time:     "5:00",
					},
					{
						Date:     "May / 25 / 2013",
						Event:    "UFC 160 - Velasquez vs. Bigfoot 2",
						Method:   "TKO (Punches)",
						Opponent: "Colton Smith (Colton-Smith-63163)",
						Referee:  "Chris Tognoni",
						Result:   "win",
						Round:    3,
						Time:     "0:41",
					},
					{
						Date:     "Dec / 15 / 2012",
						Event:    "UFC on FX 6 - Sotiropoulos vs. Pearson",
						Method:   "Decision (Unanimous)",
						Opponent: "Bradley Scott (Bradley-Scott-44442)",
						Referee:  "Steve Perceval",
						Result:   "win",
						Round:    3,
						Time:     "5:00",
					},
					{
						Date:     "May / 18 / 2012",
						Event:    "",
						Method:   "Decision (Unanimous)",
						Opponent: "Jesse Juarez (Jesse-Juarez-22291)",
						Referee:  "Steve Perceval",
						Result:   "loss",
						Round:    5,
						Time:     "5:00",
					},
					{
						Date:     "Mar / 23 / 2012",
						Event:    "SA - Superfight Australia 13",
						Method:   "TKO (Punches)",
						Opponent: "Shaun Spooner (Shaun-Spooner-21796)",
						Referee:  "N/A",
						Result:   "win",
						Round:    1,
						Time:     "4:01",
					},
					{
						Date:     "Dec / 09 / 2011",
						Event:    "CFC 19 - Falciroli vs. Honstein",
						Method:   "TKO (Punches)",
						Opponent: "Ian Bone (Ian-Bone-23652)",
						Referee:  "John Sharp",
						Result:   "win",
						Round:    2,
						Time:     "3:15",
					},
					{
						Date:     "Oct / 30 / 2011",
						Event:    "",
						Method:   "Submission (Triangle Choke)",
						Opponent: "Hoon Kim (Hoon-Kim-36879)",
						Referee:  "Warren Wang",
						Result:   "loss",
						Round:    1,
						Time:     "3:01",
					},
					{
						Date:     "Aug / 26 / 2011",
						Event:    "CFC 18 - Juarez vs. Rodriguez",
						Method:   "Submission (Armbar)",
						Opponent: "Corey Nelson (Corey-Nelson-56759)",
						Referee:  "Steve Perceval",
						Result:   "win",
						Round:    2,
						Time:     "4:40",
					},
					{
						Date:     "Jun / 03 / 2011",
						Event:    "CFC - Cage Fighting Championship 17",
						Method:   "Submission (Rear-Naked Choke)",
						Opponent: "Ben Alloway (Ben-Alloway-60790)",
						Referee:  "N/A",
						Result:   "win",
						Round:    2,
						Time:     "4:07",
					},
					{
						Date:     "Oct / 08 / 2010",
						Event:    "CFC - Cage Fighting Championships 15",
						Method:   "Submission (Rear-Naked Choke)",
						Opponent: "Nate Thomson (Nate-Thomson-62718)",
						Referee:  "N/A",
						Result:   "win",
						Round:    1,
						Time:     "2:21",
					},
					{
						Date:     "Jun / 05 / 2010",
						Event:    "CFC  - Cage Fighting Championships 14",
						Method:   "Submission (Armbar)",
						Opponent: "Jay Cobain (Jay-Cobain-44247)",
						Referee:  "N/A",
						Result:   "win",
						Round:    2,
						Time:     "0:32",
					},
					{
						Date:     "Mar / 12 / 2010",
						Event:    "CFC 12 - Lombard vs. Santore",
						Method:   "KO (Punch)",
						Opponent: "Nick Ariel (Nick-Ariel-55020)",
						Referee:  "John Sharp",
						Result:   "win",
						Round:    2,
						Time:     "2:50",
					},
					{
						Date:     "Nov / 20 / 2009",
						Event:    "CFC - Cage Fighting Championships 11",
						Method:   "Submission (Rear-Naked Choke)",
						Opponent: "Richard Walsh (Richard-Walsh-55587)",
						Referee:  "N/A",
						Result:   "win",
						Round:    2,
						Time:     "2:40",
					},
					{
						Date:     "Mar / 14 / 2009",
						Event:    "XFC - Return of the Hulk",
						Method:   "TKO (Punches) ",
						Opponent: "Chris Tallowin (Chris-Tallowin-45133)",
						Referee:  "N/A",
						Result:   "win",
						Round:    1,
						Time:     "N/A",
					},
				},
				Height:      "5'11\"",
				ID:          "Robert-Whittaker-45132",
				ImageURL:    "https://www.sherdog.com/image_crop/200/300/_images/fighter/1601316224Robert_Whittaker.jpg",
				Locality:    "Sydney, New South Wales",
				Name:        "Robert Whittaker",
				Nationality: "Australia",
				Nickname:    "\"The Reaper / Bobby Knuckles\"",
				Record: Record{
					LossesDecisions:   2,
					LossesKnockouts:   2,
					LossesSubmissions: 1,
					LossesTotal:       5,
					NoContests:        0,
					WinsDecisions:     9,
					WinsKnockouts:     9,
					WinsSubmissions:   5,
					WinsTotal:         23,
				},
				Weight:      "185 lbs",
				WeightClass: "Middleweight",
			},
		},
		{
			name:        "notFound",
			fighterID:   "123-nonexistent",
			wantErr:     ErrFighterNotFound,
			wantFighter: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindFighterByID(tt.fighterID)
			assert.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				assert.Equal(t, tt.wantFighter, got)
			}
		})
	}
}

func Test_FindFighterByName(t *testing.T) {
	tests := []struct {
		name         string
		fighterName  string
		wantErr      error
		wantFighters []*Fighter
	}{
		{
			"successJohnAllan",
			"John Allan",
			nil,
			[]*Fighter{
				{
					Name:        "Allan Johns",
					Nationality: "United States",
				},
				{
					Name:        "John Allan",
					Nationality: "Brazil",
				},
			},
		},
		{
			"notFound",
			"non-existent-person 123",
			ErrFighterNotFound,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindFighterByName(tt.fighterName)
			assert.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				for i, f := range tt.wantFighters {
					assert.Equal(t, f.Name, got[i].Name)
					assert.Equal(t, f.Nationality, got[i].Nationality)
				}
			}
		})
	}
}
