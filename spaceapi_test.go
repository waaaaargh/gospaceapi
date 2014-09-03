package spaceapi

import "testing"
import "encoding/json"

var minimal_spaceapi string = `
{
    "api": "0.13",
    "space": "Slopspace",
    "logo": "http://your-space.org/img/logo.png",
    "url": "http://your-space.org",
    "location": {
        "address": "Ulmer Strasse 255, 70327 Stuttgart, Germany",
        "lon": 9.236,
        "lat": 48.777
    },
    "icon": {
        "open": "http://test.org/a.png",
        "closed": "http://test.org/b.png"
    },
    "contact": {
        "twitter": "@spaceapi"
    },
    "issue_report_channels": [
        "twitter"
    ],
    "state": {
        "open": true
    },
    "sensors": {
        "temperature": {
            "value": "23",
            "unit": "°C",
            "location": "inside"
        }
    },
    "feeds": {
        "blog": {
            "type": "rss",
            "url": "http://www.goatse.cx"
         }
    },
    "projects": [
        "http://www.weltraumpflege.org"
    ]
    

}`

func TestSpaceAPIUnmarshal(t *testing.T) {
	txt := []byte(minimal_spaceapi)
	var e SpaceAPI
	err := json.Unmarshal(txt, &e)

	if err != nil {
		t.Error(err)
	}

	if e.Space != "Slopspace" {
		t.Error("Error parsing main Object")
	}

	if e.Location.Address != "Ulmer Strasse 255, 70327 Stuttgart, Germany" {
		t.Error("Error parsing Location Object")
	}

	if e.State.Open != true {
		t.Error("Error parsing State Object")
	}

	if e.Contact.Twitter != "@spaceapi" {
		t.Error("Error parsing Contact Object")
	}

	if len(e.IssueReportChannels) < 1 {
		t.Error("Error parsing IssueReportChannels")
	}

	if len(e.Feeds) < 1 {
		t.Error("Error parsing Feeds")
	}
}
