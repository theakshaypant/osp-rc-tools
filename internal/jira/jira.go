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

type issueFields struct {
	Summary string      `json:"summary"`
	Status  issueStatus `json:"status"`
}

type issueStatus struct {
	Name string `json:"name"`
}

type searchIssue struct {
	Key    string      `json:"key"`
	Fields issueFields `json:"fields"`
}

type searchResponse struct {
	Issues        []searchIssue `json:"issues"`
	NextPageToken string        `json:"nextPageToken"`
}

type Ticket struct {
	Key     string `json:"key"`
	URL     string `json:"url"`
	Summary string `json:"summary"`
	Status  string `json:"status"`
}

func (c *Client) browseURL(key string) string {
	return fmt.Sprintf("%s/browse/%s", c.baseURL, key)
}

func KeyFromURL(url string) string {
	if idx := strings.LastIndex(url, "/"); idx >= 0 {
		return url[idx+1:]
	}
	return url
}

func (c *Client) searchJQL(jql string, fields []string, maxResults int, nextPageToken string) (*searchResponse, error) {
	body := map[string]any{
		"jql":        jql,
		"fields":     fields,
		"maxResults": maxResults,
	}
	if nextPageToken != "" {
		body["nextPageToken"] = nextPageToken
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest("POST", c.baseURL+"/rest/api/3/search/jql", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var result searchResponse
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("jira search failed (status %d): %s", resp.StatusCode, body)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	resp.Body.Close()
	return &result, err
}

func (c *Client) FindTicketForPR(prURL string) string {
	jql := fmt.Sprintf(`cf[%s] ~ "%s"`, strings.TrimPrefix(c.prField, "customfield_"), prURL)

	result, err := c.searchJQL(jql, []string{"key"}, 1, "")
	if err != nil || len(result.Issues) == 0 {
		return ""
	}

	return c.browseURL(result.Issues[0].Key)
}

func (c *Client) FindTicketsForFixVersion(version string) ([]Ticket, error) {
	fixVersion := fmt.Sprintf("Pipeline %s", version)
	jql := fmt.Sprintf(`project = SRVKP AND fixVersion = "%s"`, fixVersion)

	var tickets []Ticket
	nextPage := ""

	for {
		result, err := c.searchJQL(jql, []string{"summary", "status"}, 50, nextPage)
		if err != nil {
			return nil, err
		}

		for _, issue := range result.Issues {
			tickets = append(tickets, Ticket{
				Key:     issue.Key,
				URL:     c.browseURL(issue.Key),
				Summary: issue.Fields.Summary,
				Status:  issue.Fields.Status.Name,
			})
		}

		if result.NextPageToken == "" || len(result.Issues) == 0 {
			break
		}
		nextPage = result.NextPageToken
	}

	return tickets, nil
}
