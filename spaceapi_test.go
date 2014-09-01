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
    "contact": {
        "twitter": "@spaceapi"
    },
    "issue_report_channels": [
        "twitter"
    ],
    "state": {
        "open": true
    }
}`

func TestSpaceAPI(t *testing.T) {
	txt := []byte(minimal_spaceapi)
	var e Endpoint
	json.Unmarshal(txt, &e)

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
}
