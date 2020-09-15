package simple

import (
	"net/http"
)

type Sender struct {
	URL string
}

//SimpleGet sends a GET request
func (s *Sender) SimpleGet() (*http.Response, error) {
	return http.Get(s.URL)
}
