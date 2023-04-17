package v1

type Server struct {
	Name        string `json:name`
	IPAddress   string `json:ip_address`
	Environment string `json:environment`
	Responder   string `json:responder`
}

var servers []Server
