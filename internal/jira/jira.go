package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Client struct {
	baseURL    string
	email      string
	token      string
	prField    string
	httpClient *http.Client
}

func NewClient() (*Client, error) {
	jiraURL := os.Getenv("JIRA_URL")
	jiraEmail := os.Getenv("JIRA_EMAIL")
	jiraToken := os.Getenv("JIRA_TOKEN")
	prField := os.Getenv("JIRA_PR_FIELD")
	if prField == "" {
		prField = "customfield_10875"
	}
	if jiraURL == "" || jiraEmail == "" || jiraToken == "" {
		return nil, nil
	}

	return &Client{
		baseURL: strings.TrimRight(jiraURL, "/"),
		email:   jiraEmail,
		token:   jiraToken,
		prField: prField,
		httpClient: &http.Client{
			Transport: &http.Transport{ForceAttemptHTTP2: true},
		},
	}, nil
}

func (c *Client) newRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.email, c.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	return req, nil
}

type searchIssue struct {
	Key string `json:"key"`
}

type searchResponse struct {
	Issues []searchIssue `json:"issues"`
}

func (c *Client) FindTicketForPR(prURL string) string {
	jql := fmt.Sprintf(`cf[%s] ~ "%s"`, strings.TrimPrefix(c.prField, "customfield_"), prURL)

	payload, err := json.Marshal(map[string]interface{}{
		"jql":        jql,
		"fields":     []string{"key"},
		"maxResults": 1,
	})
	if err != nil {
		return ""
	}

	req, err := c.newRequest("POST", c.baseURL+"/rest/api/3/search/jql", bytes.NewBuffer(payload))
	if err != nil {
		return ""
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return ""
	}

	var result searchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil || len(result.Issues) == 0 {
		return ""
	}

	return fmt.Sprintf("%s/browse/%s", c.baseURL, result.Issues[0].Key)
}
