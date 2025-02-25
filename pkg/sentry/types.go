package sentry

type LogEntry struct {
	Formatted string      `json:"formatted"`
	Message   interface{} `json:"message"`
	Params    interface{} `json:"params"`
}

type Metadata struct {
	Title string `json:"title"`
}

type Event struct {
	EventID     string   `json:"event_id"`
	Level       string   `json:"level"`
	Version     string   `json:"version"`
	LogEntry    LogEntry `json:"logentry"`
	Platform    string   `json:"platform"`
	Timestamp   float64  `json:"timestamp"`
	Received    float64  `json:"received"`
	Environment string   `json:"environment"`
	Metadata    Metadata `json:"metadata"`
	Culprit     string   `json:"culprit"`
	Title       string   `json:"title"`
}

type Payload struct {
	ID          string `json:"id"`
	Project     string `json:"project"`
	ProjectName string `json:"project_name"`
	ProjectSlug string `json:"project_slug"`
	Level       string `json:"level"`
	Culprit     string `json:"culprit"`
	Message     string `json:"message"`
	URL         string `json:"url"`
	Event       Event  `json:"event"`
}
