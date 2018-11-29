package redis

import (
	"log"
	"net/http"

	"github.com/onkiit/db-monitor/api"
	"github.com/onkiit/db-monitor/lib/helper"
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
