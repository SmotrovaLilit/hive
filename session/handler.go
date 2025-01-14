package session

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"

	"github.com/ory/herodot"
	"github.com/ory/hive/driver/configuration"
	"github.com/ory/hive/x"
)

type HandlerProvider interface {
	SessionHandler() *Handler
}

func NewHandler(
	r Registry,
	h herodot.Writer,
) *Handler {
	return &Handler{
		r: r,
		h: h,
	}
}

const (
	CheckPath = "/sessions/me"
)

type Handler struct {
	r Registry
	h herodot.Writer
}

func (h *Handler) RegisterPublicRoutes(public *x.RouterPublic) {
	public.GET(CheckPath, h.fromCookie)
}

func (h *Handler) RegisterAdminRoutes(admin *x.RouterAdmin) {
	admin.GET("/sessions/:sid", h.fromPath)
}

func (h *Handler) fromCookie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	s, err := h.r.SessionManager().FetchFromRequest(r.Context(), r)
	if err != nil {
		h.h.WriteError(w, r, err)
		return
	}

	// s.Devices = nil
	s.Identity = s.Identity.CopyWithoutCredentials()

	h.h.Write(w, r, s)
}

func (h *Handler) fromPath(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(505)
}

func (h *Handler) IsNotAuthenticated(wrap httprouter.Handle, onAuthenticated httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if _, err := h.r.SessionManager().FetchFromRequest(r.Context(), r); err != nil {
			if errors.Cause(err) == ErrNoActiveSessionFound {
				wrap(w, r, ps)
				return
			}
			h.h.WriteError(w, r, err)
			return
		}

		if onAuthenticated != nil {
			onAuthenticated(w, r, ps)
			return
		}

		h.h.WriteError(w, r, errors.WithStack(herodot.ErrForbidden.WithReason("This endpoint can only be accessed without a login session. Please log out and try again.")))
	}
}

func RedirectOnAuthenticated(c configuration.Provider) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		http.Redirect(w, r, c.DefaultReturnToURL().String(), http.StatusFound)
	}
}
