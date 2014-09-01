package spaceapi

type Location struct {
	Address string `json:"address"`
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type State struct {
	Open bool `json:"open"`
	Lastchange int32 `json:"lastchange"`
}

type Contact struct {
	Twitter string `json:"twitter"`
	Irc string `json:"irc"`
	Facebook string `json:"facebook"`
	Email string `json:"email"`
	Ml string `json:"ml"`
}

type Endpoint struct {
	Api string `json:"api"`
	Space string `json:"space"`
	Logo string `json:"logo"`
	Url string `json:"url"`
	Location Location `json:"location"`
	State State `json:"state"`
	Contact Contact `json:"contact"`
}
