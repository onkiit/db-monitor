package mongo

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/onkiit/db-monitor/lib/helper"

	"github.com/onkiit/db-monitor/api"
)

type Controller struct {
	store api.Store
}

func (c *Controller) GetVersion(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	v, err := c.store.GetVersion(ctx)
	if err != nil {
		helper.ErrorResponse(w, http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	helper.RespondWithJSON(w, http.StatusOK, v)
}

func (c *Controller) GetActiveClient(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	v, err := c.store.GetActiveClient(ctx)
	if err != nil {
		helper.ErrorResponse(w, http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	helper.RespondWithJSON(w, http.StatusOK, v)
}

func (c *Controller) GetHealth(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	h, err := c.store.GetHealth(ctx)
	if err != nil {
		helper.ErrorResponse(w, http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	helper.RespondWithJSON(w, http.StatusOK, h)
}
