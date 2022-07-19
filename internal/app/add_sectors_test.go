package app

import (
	"net/http"
	"strings"
	"testing"

	dns_storage "github.com/DavidBagaryan/drone-attack/internal/storage/dns"
	sector_storage "github.com/DavidBagaryan/drone-attack/internal/storage/sector"
	"github.com/stretchr/testify/assert"
)

func TestImplementation_AddSectors(t *testing.T) {
	badReq, err := http.NewRequest("POST", "", strings.NewReader(""))
	assert.NoError(t, err)

	okReq, err := http.NewRequest(
		"POST",
		"",
		strings.NewReader(`[{"count_dns": 12},{"count_dns": 10}]`),
	)
	assert.NoError(t, err)

	tt := map[string]struct {
		statusCode int
		response   string
		req        *http.Request
	}{
		"method not allowed": {
			statusCode: 405,
			response:   "method not allowed",
			req:        &http.Request{Method: "PUT"},
		},
		"bad request": {
			statusCode: 400,
			response:   "an error occurred",
			req:        badReq,
		},
		"all ok": {
			response: `[{"id":0,"available_dns":12,"deployed_dns":0,"drone_count":0},{"id":1,"available_dns":10,"deployed_dns":0,"drone_count":0}]`,
			req:      okReq,
		},
	}

	impl := New(sector_storage.New(), dns_storage.New())

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			rw := &trw{}
			impl.AddSectors(rw, tc.req)
			assert.Equal(t, tc.response, rw.response)
			assert.Equal(t, tc.statusCode, rw.statusCode)
		})
	}
}