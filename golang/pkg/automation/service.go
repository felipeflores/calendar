package automation

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type AutomationService struct {
	url string
}

func NewAutomationService() *AutomationService {
	return &AutomationService{
		url: "http://192.168.50.50",
	}
}

func (as *AutomationService) Blink(ctx context.Context, on string) error {
	client := &http.Client{}

	url := fmt.Sprintf("%s/blink?on=%s", as.url, on)

	fmt.Println("Link", url)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("Errou")
	}
	return nil
}

func (as *AutomationService) Write(ctx context.Context, text string) error {
	client := &http.Client{}

	url := fmt.Sprintf("%s/write?text=%s", as.url, text)

	fmt.Println("Link", url)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("Errou")
	}
	return nil
}
