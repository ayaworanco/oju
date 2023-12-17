package entities

type Journey struct {
	Resources map[string]map[string][]JourneyData
}

/*
TRACE 3FAFCF87-BF66-4DC5-84C1-34E178FF55CC AWO\n{"name":"span-name","service":"","attributes":{"http.url":"http://products.api.svc.cluster.local","http.method":"POST","http.body.email":"test@email.com"}}
*/

type JourneyData struct {
	Action     string            `json:"action"`
	Service    string            `json:"service"`
	Attributes map[string]string `json:"attributes"`
}
