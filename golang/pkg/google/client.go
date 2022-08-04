package google

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"iot/pkg/google/model"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

type ClientGoogle struct {
	client  *http.Client
	config  *oauth2.Config
	tokFile string
}

func NewClientGoogle() *ClientGoogle {
	return &ClientGoogle{
		tokFile: "token.json",
	}
}

func (c *ClientGoogle) Setup(ctx context.Context, credentials model.Credentials) (string, error) {
	credentialsFileName := "credentials.json"

	file, _ := json.MarshalIndent(credentials, "", " ")
	_ = ioutil.WriteFile(credentialsFileName, file, 0644)

	b, err := ioutil.ReadFile(credentialsFileName)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Unable to read client secret file: %v", err))
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Unable to parse client secret file to config: %v", err))
	}
	c.config = config

	return c.generateURLAuth(), nil
}

func (c *ClientGoogle) generateURLAuth() string {
	return c.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func (c *ClientGoogle) GenerateToken(ctx context.Context, authCode string) error {
	tok, err := c.config.Exchange(ctx, authCode)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(c.tokFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(tok)
	if err != nil {
		return err
	}

	return c.GetClient()
}

func (c *ClientGoogle) GetClient() error {
	tok, err := c.tokenFromFile()
	if err != nil {
		return err
	}
	c.client = c.config.Client(context.Background(), tok)
	return nil
}

// Retrieves a token from a local file.
func (c *ClientGoogle) tokenFromFile() (*oauth2.Token, error) {
	f, err := os.Open(c.tokFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}
