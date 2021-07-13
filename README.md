# Sherdog Go Wrapper

### Fighter Response Example
```
{
  name: 'Darren Till',
  nickname: 'The Gorilla',
  age: '28',
  birthday: '1992-12-24',
  locality: 'Liverpool, Merseyside',
  nationality: 'England',
  association: 'Astra Fight TeamTeam Kaobon',
  height: `6'0"`,
  weight: '185 lbs',
  weight_class: 'Middleweight',
  image_url: '/image_crop/200/300/_images/fighter/1601316708Darren_Till.jpg',
  wins: { total: 18, knockouts: 10, submissions: 2, decisions: 6, others: 0 },
  losses: { total: 3, knockouts: 1, submissions: 1, decisions: 1, others: 0 },
  no_contests: 1,
  fights: [
    {
      name: 'UFC on ESPN 14 - Whittaker vs. Till',
      date: 'Jul / 25 / 2020',
      url: '/events/UFC-on-ESPN-14-Whittaker-vs-Till-86682',
      result: 'loss',
      method: 'Decision (Unanimous)',
      referee: 'Herb Dean',
      round: '5',
      time: '5:00',
      opponent: 'Robert Whittaker'
    },
    .
    .
    .
  ],
}
```

https://github.com/gocolly/colly/blob/master/_examples/reddit/reddit.go