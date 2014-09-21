package spaceapi

import "encoding/json"
import "net/http"
import "io/ioutil"

type Location struct {
	Address string  `json:"address,omitempty"`
	Lat     float32 `json:"lat,omitempty"`
	Lon     float32 `json:"lon,omitempty"`
}

type State struct {
	Open          bool   `json:"open,omitempty"`
	Lastchange    int32  `json:"lastchange,omitempty"`
	TriggerPerson string `json:"trigger_person,omitempty"`
	Message       string `json:"message,omitempty"`
	Icon          struct {
		Open   string `json:"open,omitempty"`
		Closed string `json:"string,omitempty"`
	} `json:"icon,omitempty"`
}

type GoogleContact struct {
	Plus string `json:"plus,omitempty"`
}

type Contact struct {
	Phone      string        `json:"phone,omitempty"`
	Sip        string        `json:"sip,omitempty"`
	Irc        string        `json:"irc,omitempty"`
	Twitter    string        `json:"twitter,omitempty"`
	Facebook   string        `json:"facebook,omitempty"`
	Google     GoogleContact `json:"google,omitempty"`
	Identica   string        `json:"identica,omitempty"`
	Foursquare string        `json:"foursquare,omitempty"`
	Email      string        `json:"email,omitempty"`
	Ml         string        `json:"ml,omitempty"`
	Jabber     string        `json:"jabber,omitempty"`
	IssueMail  string        `json:"issue_mail,omitempty"`
}

type Feed struct {
	Type string `json:"type,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Cache struct {
	Schedule string `json:"schedule,omitempty"`
}

type SpaceFed struct {
	Spacenet   bool `json:"spacenet,omitempty"`
	Spacesaml  bool `json:"spacesaml,omitempty"`
	Spacephone bool `json:"spacephone,omitempty"`
}

type RadioShow struct {
	Name  string `json:"name,omitempty"`
	Url   string `json:"url,omitempty"`
	Type  string `json:"type,omitempty"`
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type SpaceAPI struct {
	Api                 string                 `json:"api,omitempty"`
	Space               string                 `json:"space,omitempty"`
	Logo                string                 `json:"logo,omitempty"`
	Url                 string                 `json:"url,omitempty"`
	Location            *Location              `json:"location,omitempty"`
	State               *State                 `json:"state,omitempty"`
	Contact             *Contact               `json:"contact,omitempty"`
	IssueReportChannels []string               `json:"issue_report_channels,omitempty"`
	Sensors             map[string]interface{} `json:"sensors,omitempty"`
	Feeds               map[string]Feed        `json:"feeds,omitempty"`
	Cache               *Cache                 `json:"cache,omitempty"`
	SpaceFed            *SpaceFed              `json:"spacefed,omitempty"`
	Projects            []string               `json:"projects,omitempty"`
	RadioShow           []RadioShow            `json:"radioshow,omitempty"`
}

func (s *SpaceAPI) ToJSON() (string, error) {
	bytes, err := json.Marshal(s)
	return string(bytes), err
}

func NewSpaceAPI(api_version string, space_name string, logo_url string, website_url string, location *Location, state *State, contact *Contact, issue_report_channels []string) SpaceAPI {
	return SpaceAPI{
		Api:                 api_version,
		Space:               space_name,
		Logo:                logo_url,
		Location:            location,
		Contact:             contact,
		State:               state,
		IssueReportChannels: issue_report_channels,
	}
}

type Endpoint struct {
	endpoint_url string
}

func (e *Endpoint) GetSpaceAPIData() (*SpaceAPI, error) {
	resp, err := http.Get(e.endpoint_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := new(SpaceAPI)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
