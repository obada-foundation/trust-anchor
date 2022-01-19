package api

import (
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/obada-foundation/trust-anchor/rest"
	"github.com/obada-foundation/trust-anchor/ta"
)

type private struct {
	taTokenSvc *ta.TokenService
	logger     *log.Logger
}

type RegisterNFTRequest struct {
	DID string `json:"did"`
}

type RegisterNFTResponse struct {
	Token string `json:"token"`
}

func (rpriv *private) registerNFT(w http.ResponseWriter, r *http.Request) {
	var request RegisterNFTRequest

	if err := render.DecodeJSON(http.MaxBytesReader(w, r.Body, hardBodyLimit), &request); err != nil {
		rest.SendErrorJSON(w, r, http.StatusBadRequest, err, "can't decode request data", rest.ErrDecode)
		return
	}

	if len(request.DID) == 0 {
		rest.SendErrorJSON(w, r, http.StatusBadRequest, nil, "OBADA DID is required for this request", rest.ErrDecode)
		return
	}

	var resp RegisterNFTResponse

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, &resp)
}
