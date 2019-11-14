// Code generated by goa v3.0.7, DO NOT EDIT.
//
// collection HTTP server
//
// Command:
// $ goa gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package server

import (
	"context"
	"net/http"

	collection "github.com/artefactual-labs/enduro/internal/api/gen/collection"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the collection service endpoint HTTP handlers.
type Server struct {
	Mounts   []*MountPoint
	List     http.Handler
	Show     http.Handler
	Delete   http.Handler
	Cancel   http.Handler
	Retry    http.Handler
	Workflow http.Handler
	CORS     http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the collection service endpoints.
func New(
	e *collection.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"List", "GET", "/collection"},
			{"Show", "GET", "/collection/{id}"},
			{"Delete", "DELETE", "/collection/{id}"},
			{"Cancel", "POST", "/collection/{id}/cancel"},
			{"Retry", "POST", "/collection/{id}/retry"},
			{"Workflow", "GET", "/collection/{id}/workflow"},
			{"CORS", "OPTIONS", "/collection"},
			{"CORS", "OPTIONS", "/collection/{id}"},
			{"CORS", "OPTIONS", "/collection/{id}/cancel"},
			{"CORS", "OPTIONS", "/collection/{id}/retry"},
			{"CORS", "OPTIONS", "/collection/{id}/workflow"},
		},
		List:     NewListHandler(e.List, mux, dec, enc, eh),
		Show:     NewShowHandler(e.Show, mux, dec, enc, eh),
		Delete:   NewDeleteHandler(e.Delete, mux, dec, enc, eh),
		Cancel:   NewCancelHandler(e.Cancel, mux, dec, enc, eh),
		Retry:    NewRetryHandler(e.Retry, mux, dec, enc, eh),
		Workflow: NewWorkflowHandler(e.Workflow, mux, dec, enc, eh),
		CORS:     NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "collection" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.List = m(s.List)
	s.Show = m(s.Show)
	s.Delete = m(s.Delete)
	s.Cancel = m(s.Cancel)
	s.Retry = m(s.Retry)
	s.Workflow = m(s.Workflow)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the collection endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountListHandler(mux, h.List)
	MountShowHandler(mux, h.Show)
	MountDeleteHandler(mux, h.Delete)
	MountCancelHandler(mux, h.Cancel)
	MountRetryHandler(mux, h.Retry)
	MountWorkflowHandler(mux, h.Workflow)
	MountCORSHandler(mux, h.CORS)
}

// MountListHandler configures the mux to serve the "collection" service "list"
// endpoint.
func MountListHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleCollectionOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/collection", f)
}

// NewListHandler creates a HTTP handler which loads the HTTP request and calls
// the "collection" service "list" endpoint.
func NewListHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeListRequest(mux, dec)
		encodeResponse = EncodeListResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "list")
		ctx = context.WithValue(ctx, goa.ServiceKey, "collection")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountShowHandler configures the mux to serve the "collection" service "show"
// endpoint.
func MountShowHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleCollectionOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/collection/{id}", f)
}

// NewShowHandler creates a HTTP handler which loads the HTTP request and calls
// the "collection" service "show" endpoint.
func NewShowHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeShowRequest(mux, dec)
		encodeResponse = EncodeShowResponse(enc)
		encodeError    = EncodeShowError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "show")
		ctx = context.WithValue(ctx, goa.ServiceKey, "collection")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountDeleteHandler configures the mux to serve the "collection" service
// "delete" endpoint.
func MountDeleteHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleCollectionOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/collection/{id}", f)
}

// NewDeleteHandler creates a HTTP handler which loads the HTTP request and
// calls the "collection" service "delete" endpoint.
func NewDeleteHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteRequest(mux, dec)
		encodeResponse = EncodeDeleteResponse(enc)
		encodeError    = EncodeDeleteError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "delete")
		ctx = context.WithValue(ctx, goa.ServiceKey, "collection")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountCancelHandler configures the mux to serve the "collection" service
// "cancel" endpoint.
func MountCancelHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleCollectionOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/collection/{id}/cancel", f)
}

// NewCancelHandler creates a HTTP handler which loads the HTTP request and
// calls the "collection" service "cancel" endpoint.
func NewCancelHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeCancelRequest(mux, dec)
		encodeResponse = EncodeCancelResponse(enc)
		encodeError    = EncodeCancelError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "cancel")
		ctx = context.WithValue(ctx, goa.ServiceKey, "collection")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountRetryHandler configures the mux to serve the "collection" service
// "retry" endpoint.
func MountRetryHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleCollectionOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/collection/{id}/retry", f)
}

// NewRetryHandler creates a HTTP handler which loads the HTTP request and
// calls the "collection" service "retry" endpoint.
func NewRetryHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeRetryRequest(mux, dec)
		encodeResponse = EncodeRetryResponse(enc)
		encodeError    = EncodeRetryError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "retry")
		ctx = context.WithValue(ctx, goa.ServiceKey, "collection")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountWorkflowHandler configures the mux to serve the "collection" service
// "workflow" endpoint.
func MountWorkflowHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleCollectionOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/collection/{id}/workflow", f)
}

// NewWorkflowHandler creates a HTTP handler which loads the HTTP request and
// calls the "collection" service "workflow" endpoint.
func NewWorkflowHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeWorkflowRequest(mux, dec)
		encodeResponse = EncodeWorkflowResponse(enc)
		encodeError    = EncodeWorkflowError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "workflow")
		ctx = context.WithValue(ctx, goa.ServiceKey, "collection")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service collection.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = handleCollectionOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/collection", f)
	mux.Handle("OPTIONS", "/collection/{id}", f)
	mux.Handle("OPTIONS", "/collection/{id}/cancel", f)
	mux.Handle("OPTIONS", "/collection/{id}/retry", f)
	mux.Handle("OPTIONS", "/collection/{id}/workflow", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// handleCollectionOrigin applies the CORS response headers corresponding to
// the origin for the service collection.
func handleCollectionOrigin(h http.Handler) http.Handler {
	origHndlr := h.(http.HandlerFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "*") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS")
			}
			origHndlr(w, r)
			return
		}
		origHndlr(w, r)
		return
	})
}
