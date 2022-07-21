package google

import (
	"context"
	"time"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type CalendarGoogle struct {
	ClientGoogle *ClientGoogle
}

func NewCalendarGoogle(cg *ClientGoogle) *CalendarGoogle {
	return &CalendarGoogle{
		ClientGoogle: cg,
	}
}

func (cal *CalendarGoogle) Get(ctx context.Context) (*calendar.Event, error) {
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(cal.ClientGoogle.client))
	if err != nil {
		return nil, err
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		return nil, err
	}
	if len(events.Items) == 0 {
		return nil, nil
	} else {
		return events.Items[0], nil
		// for _, item := range events.Items {
		// 	date := item.Start.DateTime
		// 	if date == "" {
		// 		date = item.Start.Date
		// 	}
		// 	aJSON, _ := json.Marshal(item)
		// 	fmt.Printf("JSON Print - \n%s\n", string(aJSON))

		// 	// fmt.Printf("%v (%v)\n%v", item.Summary, date, item)
		// }
	}
}
