# **go-sherdog**
Wrapper of the most famous MMA website - [Sherdog](https://sherdog.com) - written in Go.
## Installation
```sh
go get -u github.com/gusandrioli/go-sherdog
```

## Quick Start
```go
package main

import (
    "fmt"
    sherdog "github.com/gusandrioli/go-sherdog"
)

func main() {
    darrenTill, err := sherdog.FindFighterByID("Darren-Till-73436")
    if err != nil {
      log.Errorf(err)
    }

    fmt.Pritnln(darrenTill)
}
```

## Features/Methods
```go
// Present
func FindFighterByID(fighterID FighterID) (*Fighter, error)
func FindFighterByName(name string) ([]*Fighter, error)

// Upcoming... (subject to change)
func GetOrganizationEvents(organization string) ([]*Event, error)
```

## Fighter Structure Example
```go
{
  Age: "28",
  Association: "Astra Fight TeamTeam Kaobon",
  Birthday: "1992-12-24",
  FightHistory: []*Fight{
    {
      Date: "2020-07-25",
      Event: "UFC on ESPN 14 - Whittaker vs. Till",
      Method: "Decision (Unanimous)",
      Opponent: "Robert Whittaker"
      Referee: "Herb Dean",
      Result: "loss",
      Round: "5",
      Time: "5:00",
    },
    ...
  },
  Height: "6'0\"",
  ID: "Darren-Till-73436"
  ImageURL: "https://sherdog.com/image_crop/200/300/_images/fighter/1601316708Darren_Till.jpg",
  Locality: "Liverpool, Merseyside",
  Name: "Darren Till",
  Nickname: "The Gorilla",
  Record: Record{
      LossesDecisions: 1,
      LossesKnockouts: 1,
      LossesSubmissions: 1,
      LossesTotal: 3,
      NoContests: 1,
      WinsDecisions: 6,
      WinsKnockouts: 10,
      WinsSubmissions: 2,
      WinsTotal: 18,
  }
  Nationality: "England",
  Weight: "185 lbs",
  WeightClass: "Middleweight",
}
```

## Bugs
Bugs or suggestions? Open an issue [here](https://github.com/gusandrioli/go-sherdog/issues/new).

## References
This project uses [Colly](https://github.com/gocolly/colly/) for scraping.
