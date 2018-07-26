package common

// Maker web hooks docs: https://maker.ifttt.com/use/d_n7RHixcSYWEu7DflLtwp

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://maker.ifttt.com/trigger/%s/with/key/%s"
)

func getBaseURL(event, key string) string {
	return fmt.Sprintf(baseURL, event, key)
}

type IFTTTClient interface {
	Trigger(event string, values ...string) error
}

type iftttClient struct {
	cfg Config
}

func (c *iftttClient) Trigger(event string, values ...string) error {
	if len(values) > 3 {
		return fmt.Errorf("Trigger() accepts up to 3 values")
	}
	targetURL := getBaseURL(event, c.cfg.GetIFTTTKey())

	formValues := url.Values{}
	if len(values) >= 1 {
		formValues["value1"] = []string{values[0]}
	}
	if len(values) >= 2 {
		formValues["value2"] = []string{values[1]}
	}
	if len(values) >= 3 {
		formValues["value3"] = []string{values[3]}
	}

	resp, err := http.PostForm(targetURL, formValues)
	if resp.StatusCode == 404 {
		return fmt.Errorf("A 404 was issued; is IFTTT_KEY correct?")
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Unexpected status code: %s", resp.Status)
	}

	fmt.Println(resp)
	return err
}

func NewIFTTTClient(cfg Config) IFTTTClient {
	return &iftttClient{cfg: cfg}
}
