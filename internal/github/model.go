package github

import "time"

type Owner struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Repository struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	Owner           `json:"owner"`
	HTMLURL         string    `json:"html_url"`
	Description     string    `json:"description"`
	IssuesURL       string    `json:"issues_url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	PushedAt        time.Time `json:"pushed_at"`
	GitURL          string    `json:"git_url"`
	SSHURL          string    `json:"ssh_url"`
	CloneURL        string    `json:"clone_url"`
	Homepage        string    `json:"homepage"`
	StargazersCount int       `json:"stargazers_count"`
	WatchersCount   int       `json:"watchers_count"`
	Language        string    `json:"language"`
	Archived        bool      `json:"archived"`
	Disabled        bool      `json:"disabled"`
	Forks           int       `json:"forks"`
	OpenIssues      int       `json:"open_issues"`
	Watchers        int       `json:"watchers"`
}
