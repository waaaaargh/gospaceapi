package spaceapi

import "encoding/json"
import "net/http"
import "io/ioutil"

type SpaceAPI struct {
	Api      string `json:"api"`
	Space    string `json:"space"`
	Logo     string `json:"logo"`
	Url      string `json:"url"`
	Location struct {
		Address string  `json:"address"`
		Lat     float32 `json:"lat"`
		Lon     float32 `json:"lon"`
	} `json:"location"`
	State struct {
		Open          bool   `json:"open"`
		Lastchange    int32  `json:"lastchange"`
		TriggerPerson string `json:"trigger_person"`
		Message       string `json:"message"`
		Icon          struct {
			Open   string `json:"open"`
			Closed string `json:"string"`
		} `json:"icon"`
	} `json:"state"`
	Contact struct {
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
	} `json:"contact"`
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



}
