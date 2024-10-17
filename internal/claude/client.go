package claude

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		return "", fmt.Errorf("error marshalling payload: %v", err)
	}

	fmt.Printf("Sending request to Claude API: %s\n", url)
	fmt.Printf("Payload: %s\n", string(jsonPayload))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", cfg.ClaudeAPIKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("Response status: %s\n", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Printf("Raw response body: %s\n", string(body))

	var claudeResp ClaudeResponse
	if err := json.Unmarshal(body, &claudeResp); err != nil {
		return "", fmt.Errorf("error unmarshalling response: %v", err)
	}

	fmt.Printf("Unmarshalled response: %+v\n", claudeResp)

	if len(claudeResp.Content) == 0 || claudeResp.Content[0].Text == "" {
		return "", fmt.Errorf("unexpected response format: content not found or empty")
	}

	return claudeResp.Content[0].Text, nil
}
