package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Event struct {
      Name, Date, Location, Description string
  }

func main() {

  var events []Event
  // handlers
  c := colly.NewCollector(
      colly.AllowedDomains("www.su.rhul.ac.uk"),
    )

  // handling sending HTTP request
  c.OnRequest(func(r *colly.Request) {
      fmt.Println("Visiting ", r.URL)
  })

  // handling any errors in HTTP request
  c.OnError(func(_ *colly.Response, err error) {
      fmt.Println("Something went wrong: ", err)
  })

  // handling actually visiting the page
  c.OnResponse(func(r *colly.Response) {
      fmt.Println("Page visited: ", r.Request.URL)
  })

  c.OnScraped(func(r *colly.Response) {
    fmt.Println(r.Request.URL, " scraped!")
  })

  c.OnHTML("div.msl_eventlist", func(e *colly.HTMLElement) {

    event := Event{}

    event.Name = e.ChildText(".msl_event_name")
    event.Date = e.ChildText(".msl_event_time")
    event.Location = e.ChildText(".msl_event_description")
    event.Description = e.ChildText(".msl_event_description")

    events = append(events, event)
  })
  
  // open target URL
  c.Visit("www.su.rhul.ac.uk/events/calendar/")



}
