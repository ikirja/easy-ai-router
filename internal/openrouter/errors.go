package openrouter

type OpenrouterErrors struct {
	CreateRequest string
	DoRequest     string
	ReadBody      string
	UnmarshalJson string
	MarshalJson   string
}

var openrouterErrors = OpenrouterErrors{
	CreateRequest: "couldn't create request",
	DoRequest:     "request failed",
	ReadBody:      "couldn't read response body",
	UnmarshalJson: "couldn't unmarshal json body",
	MarshalJson:   "couldn't marshal post body",
}
