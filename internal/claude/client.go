package claude

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"chat-backend/pkg/config"
)

type ClaudeResponse struct {
	Content []struct {
		Text string `json:"text"`
	} `json:"content"`
}

func SendMessage(content string) (string, error) {
	cfg, err := config.Load()
	if err != nil {
		log.Printf("Error loading config: %v", err)
		return "", fmt.Errorf("error loading config: %v", err)
	}

	url := "https://api.anthropic.com/v1/messages"
	payload := map[string]interface{}{
		"model": "claude-3-opus-20240229",
		"messages": []map[string]string{
			{"role": "user", "content": content},
		},
		"max_tokens": 1000,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling payload: %v", err)
		return "", fmt.Errorf("error marshalling payload: %v", err)
	}

	log.Printf("Sending request to Claude API: %s", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", cfg.ClaudeAPIKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	log.Printf("Response status: %s", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	log.Printf("Response body: %s", string(body))

	var claudeResp ClaudeResponse
	if err := json.Unmarshal(body, &claudeResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return "", fmt.Errorf("error unmarshalling response: %v", err)
	}

	if len(claudeResp.Content) == 0 {
		log.Printf("No content in response")
		return "", fmt.Errorf("no content in response")
	}

	return claudeResp.Content[0].Text, nil
}
