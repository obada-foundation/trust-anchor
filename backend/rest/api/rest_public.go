package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/obada-foundation/trust-anchor/rest"
	"github.com/obada-foundation/trust-anchor/ta/compliance"
)

type public struct {
	logger        *log.Logger
	complianceSvc *compliance.Service
	authTokenSvc  *jwtauth.JWTAuth
}

type VerifyTokenRequest struct {
	Token string `json:"token"`
}

type VerifyTokenResponse struct {
	IsCompliant bool   `json:"is_compliant"`
	Context     string `json:"context,omitempty"`
}

func (rpub *public) verifyToken(w http.ResponseWriter, r *http.Request) {
	var request VerifyTokenRequest
	var response VerifyTokenResponse

	if err := render.DecodeJSON(http.MaxBytesReader(w, r.Body, hardBodyLimit), &request); err != nil {
		response.Context = "can't decode request data"

		render.Status(r, http.StatusOK)
		render.JSON(w, r, &response)
		return
	}

	comp, err := rpub.complianceSvc.CheckByToken(request.Token)
	if err != nil {
		response.Context = err.Error()
	}

	response.IsCompliant = comp

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &response)
}

type UserRegistration struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (rpub *public) registerUser(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
}

func (rpub *public) token(w http.ResponseWriter, r *http.Request) {
	userID := uuid.New()

	_, tokenStr, err := rpub.authTokenSvc.Encode(map[string]interface{}{"user_id": fmt.Sprintf("%s", userID)})
	if err != nil {
		rest.SendErrorJSON(w, r, http.StatusBadRequest, err, "cannot create a wallet", rest.ErrDecode)
		return

	}

	resp := struct {
		Token string `json:"token"`
	}{
		Token: tokenStr,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &resp)
}
