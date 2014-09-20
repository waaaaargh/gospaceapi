package spaceapi

import "encoding/json"
import "net/http"
import "io/ioutil"

type Location struct {
	Address string  `json:"address"`
	Lat     float32 `json:"lat"`
	Lon     float32 `json:"lon"`
}

type State struct {
	Open          bool   `json:"open"`
	Lastchange    int32  `json:"lastchange"`
	TriggerPerson string `json:"trigger_person"`
	Message       string `json:"message"`
	Icon          struct {
		Open   string `json:"open"`
		Closed string `json:"string"`
	} `json:"icon"`
}

type Contact struct {
	Phone    string `json:"phone"`
	Sip      string `json:"sip"`
	Irc      string `json:"irc"`
	Twitter  string `json:"twitter"`
	Facebook string `json:"facebook"`
	Google   struct {
		Plus string `json:"plus"`
	} `json:"google"`
	Identica   string `json:"identica"`
	Foursquare string `json:"foursquare"`
	Email      string `json:"email"`
	Ml         string `json:"ml"`
	Jabber     string `json:"jabber"`
	IssueMail  string `json:"issue_mail"`
}

type SpaceAPI struct {
	Api                 string                 `json:"api"`
	Space               string                 `json:"space"`
	Logo                string                 `json:"logo"`
	Url                 string                 `json:"url"`
	Location            Location               `json:"location"`
	State               State                  `json:"state"`
	Contact             Contact                `json:"contact"`
	IssueReportChannels []string               `json:"issue_report_channels"`
	Sensors             map[string]interface{} `json:"sensors"`
	Feeds               map[string]struct {
		Type string `json:"type"`
		Url  string `json:"url"`
	} `json:"feeds"`
	Cache struct {
		Schedule string `json:"schedule"`
	} `json:"cache"`
	SpaceFed struct {
		Spacenet   bool `json:"spacenet"`
		Spacesaml  bool `json:"spacesaml"`
		Spacephone bool `json:"spacephone"`
	} `json:"spacefed"`
	Projects  []string `json:"projects"`
	RadioShow []struct {
		Name  string `json:"name"`
		Url   string `json:"url"`
		Type  string `json:"type"`
		Start string `json:"start"`
		End   string `json:"end"`
	}
}

func (s *SpaceAPI) ToJSON() (string, error) {
	bytes, err := json.Marshal(s)
	return string(bytes), err
}

func NewSpaceAPI(api_version string, space_name string, logo_url string, website_url string, location Location, state State, contact Contact, issue_report_channels []string) SpaceAPI {
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
