package handlers

import (
	"github.com/alishchenko/discountaria/internal/server/helpers"
	"github.com/alishchenko/discountaria/internal/server/requests"
	"github.com/alishchenko/discountaria/internal/server/responses/problems"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"net/http"
)

func GetOffer(w http.ResponseWriter, r *http.Request) {
	id, err := requests.NewByIdRequest(r)
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to parse request").Error())
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to parse request")))
		return
	}
	offer, err := helpers.DB(r).NewOffers().FilterById(id).Get()
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to get offer").Error())
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to get offer")))
		return
	}
	if offer == nil {
		helpers.Log(r).Error("offer with such id not found")
		render.JSON(w, r, problems.NotFound())
		return
	}

	render.JSON(w, r, offer)
	return
}