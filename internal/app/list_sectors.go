package app

import (
	"net/http"
)

// ListSectors fetches all existed sectors
func (i Implementation) ListSectors(writer http.ResponseWriter, request *http.Request) {
	response200(writer, i.sectors.List())
}
