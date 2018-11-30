package mongo

import (
	"log"
	"net/http"

	"github.com/onkiit/db-monitor/lib/helper"

	"github.com/onkiit/db-monitor/api"
)

type Controller struct {
	store api.Store
}

func (c *Controller) GetVersion(w http.ResponseWriter, r *http.Request) {
	v, err := c.store.GetVersion()
	if err != nil {
		log.Println(err)
		return
	}

	helper.RespondWithJSON(w, http.StatusOK, v)
}

func (c *Controller) GetActiveClient(w http.ResponseWriter, r *http.Request) {
	v, err := c.store.GetActiveClient()
	if err != nil {
		log.Println(err)
		return
	}

	helper.RespondWithJSON(w, http.StatusOK, v)
}
