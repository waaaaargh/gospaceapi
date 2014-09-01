package spaceapi

type Location struct {
	Address string  `json:"address"`
	Lat     float32 `json:"lat"`
	Lon     float32 `json:"lon"`
}

type StateIcon struct {
	Open   string `json:"open"`
	Closed string `json:"string"`
}

type State struct {
	Open          bool      `json:"open"`
	Lastchange    int32     `json:"lastchange"`
	TriggerPerson string    `json:"trigger_person"`
	Message       string    `json:"message"`
	Icon          StateIcon `json:"icon"`
}

type SpaceFed struct {
	Spacenet   bool `json:"spacenet"`
	Spacesaml  bool `json:"spacesaml"`
	Spacephone bool `json:"spacephone"`
}

type GoogleContact struct {
	Plus string `json:"plus"`
}

type Contact struct {
	Phone      string        `json:"phone"`
	Sip        string        `json:"sip"`
	Irc        string        `json:"irc"`
	Twitter    string        `json:"twitter"`
	Facebook   string        `json:"facebook"`
	Google     GoogleContact `json:"google"`
	Identica   string        `json:"identica"`
	Foursquare string        `json:"foursquare"`
	Email      string        `json:"email"`
	Ml         string        `json:"ml"`
	Jabber     string        `json:"jabber"`
	IssueMail  string        `json:"issue_mail"`
}

type Feed struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type Cache struct {
	Schedule string `json:"schedule"`
}

type RadioShow struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Type  string `json:"type"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type Endpoint struct {
	Api                 string                 `json:"api"`
	Space               string                 `json:"space"`
	Logo                string                 `json:"logo"`
	Url                 string                 `json:"url"`
	Location            Location               `json:"location"`
	State               State                  `json:"state"`
	Contact             Contact                `json:"contact"`
	IssueReportChannels []string               `json:"issue_report_channels"`
	Sensors             map[string]interface{} `json:"sensors"`
	Feeds               map[string]Feed        `json:"feeds"`
	Cache               Cache                  `json:"cache"`
	Projects            []string               `json:"projects"`
}
