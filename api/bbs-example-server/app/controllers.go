// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "bbs-example-server": Application Controllers
//
// Command:
// $ goagen
// --design=BBS-Example/api/bbs-example-server/design
// --out=$(GOPATH)/src/BBS-Example/api/bbs-example-server
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// PingController is the controller interface for the Ping actions.
type PingController interface {
	goa.Muxer
	Ping(*PingPingContext) error
}

// MountPingController "mounts" a Ping resource controller on the given service.
func MountPingController(service *goa.Service, ctrl PingController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewPingPingContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Ping(rctx)
	}
	service.Mux.Handle("GET", "/ping", ctrl.MuxHandler("ping", h, nil))
	service.LogInfo("mount", "ctrl", "Ping", "action", "Ping", "route", "GET /ping")
}

// UserPostController is the controller interface for the UserPost actions.
type UserPostController interface {
	goa.Muxer
	Create(*CreateUserPostContext) error
	Index(*IndexUserPostContext) error
}

// MountUserPostController "mounts" a UserPost resource controller on the given service.
func MountUserPostController(service *goa.Service, ctrl UserPostController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/user_posts/", ctrl.MuxHandler("preflight", handleUserPostOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateUserPostContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UserPostPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleUserPostOrigin(h)
	service.Mux.Handle("POST", "/user_posts/", ctrl.MuxHandler("create", h, unmarshalCreateUserPostPayload))
	service.LogInfo("mount", "ctrl", "UserPost", "action", "Create", "route", "POST /user_posts/")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewIndexUserPostContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Index(rctx)
	}
	h = handleUserPostOrigin(h)
	service.Mux.Handle("GET", "/user_posts/", ctrl.MuxHandler("index", h, nil))
	service.LogInfo("mount", "ctrl", "UserPost", "action", "Index", "route", "GET /user_posts/")
}

// handleUserPostOrigin applies the CORS response headers corresponding to the origin.
func handleUserPostOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateUserPostPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateUserPostPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &userPostPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
