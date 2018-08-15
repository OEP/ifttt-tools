package common

// Maker web hooks docs: https://maker.ifttt.com/use/d_n7RHixcSYWEu7DflLtwp

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL    = "https://maker.ifttt.com/trigger/%s/with/key/%s"
	retryCount = 3
)

var retryStatusCodes = map[int]bool{
	502: true,
	504: true,
}

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

	for i := 0; i < retryCount; i++ {
		log.Println("Sending maker event", event, "with values", values)
		resp, err := http.PostForm(targetURL, formValues)
		if err != nil {
			return err
		}
		if resp.StatusCode == 200 {
			return nil
		} else if resp.StatusCode == 404 {
			return fmt.Errorf("A 404 was issued; is IFTTT_KEY correct?")
		} else if retryStatusCodes[resp.StatusCode] {
			log.Printf("Got HTTP status %s, retrying...\n", resp.Status)
			time.Sleep(5 * time.Second)
		} else {
			return fmt.Errorf("Unexpected status code: %s", resp.Status)
		}
	}
	return fmt.Errorf("Maximum retry count (%d) exceeded.", retryCount)
}

func NewIFTTTClient(cfg Config) IFTTTClient {
	return &iftttClient{cfg: cfg}
}
