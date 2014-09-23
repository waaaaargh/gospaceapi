package spaceapi

import "fmt"
import "strings"
import "testing"
import "encoding/json"
import "net/http"
import "net/http/httptest"

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
        "open": false
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

	if e.State.Open != false {
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

	if len(e.Projects) < 1 {
		t.Error("Error parsing Projects")
	}

}

func TestSpaceApiGetEndpoint(t *testing.T) {
	test_server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, minimal_spaceapi)
	}))
	defer test_server.Close()

	test_endpoint := Endpoint{endpoint_url: test_server.URL}
	test_struct, err := test_endpoint.GetSpaceAPIData()
	if err != nil {
		t.Error(err)
	}
	if test_struct.Space != "Slopspace" {
		t.Error("Error parsing JSON over HTTP")
	}
	if len(test_struct.Projects) < 1 {
		t.Error("Error parsing JSON over HTTP")
	}

}

func TestSpaceAPIMarshaling(t *testing.T) {
	location := Location{Lat: 10.8867969, Lon: 48.3577121}
	state := State{Open: true}
	contact := Contact{Email: "kontakt@openlab-augsburg.de"}
	api_data := NewSpaceAPI("0.14", "OpenLab", "http://www.goatse.info", "http://www.openlab-augsburg.de", &location, &state, &contact, []string{"email"})
	json_str, err := api_data.ToJSON()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(json_str, "OpenLab") {
		t.Error(err)
	}
	if !strings.Contains(json_str, "10.8867") {
		t.Error(err)
	}
	if !strings.Contains(json_str, "kontakt@") {
		t.Error(err)
	}
	if strings.Contains(json_str, "projects") {
		t.Error(err)
	}
	if strings.Contains(json_str, "spacefed") {
		t.Error(api_data.SpaceFed)
	}

}
