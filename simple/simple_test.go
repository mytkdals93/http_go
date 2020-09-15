package simple

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mytkdals93/http_go/app"
	"github.com/stretchr/testify/assert"
)

func TestSendSimpleGet(t *testing.T) {
	assert := assert.New(t)
	mux := app.NewHandler()
	ht := httptest.NewServer(mux)
	defer ht.Close()
	s := Sender{URL: ht.URL + "/"}
	resp, err := s.SimpleGet()
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	assert.NoError(err)
	assert.Equal("Hello World", string(body))

}
