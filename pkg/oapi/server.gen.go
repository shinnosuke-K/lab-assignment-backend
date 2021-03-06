// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.1 DO NOT EDIT.
package oapi

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 進学情報の修正
	// (PATCH /graduate/{userId}/fix)
	PostGraduateFix(w http.ResponseWriter, r *http.Request, userId UserId, params PostGraduateFixParams)
	// 進学情報の保存
	// (POST /graduate/{userId}/save)
	PostGraduateSave(w http.ResponseWriter, r *http.Request, userId UserId, params PostGraduateSaveParams)
	// Your GET endpoint
	// (GET /home)
	GetHome(w http.ResponseWriter, r *http.Request)
	// 研究室調査結果の修正
	// (PATCH /lab/{userId}/fix)
	PostLabFix(w http.ResponseWriter, r *http.Request, userId UserId, params PostLabFixParams)
	// 研究室調査結果の保存
	// (POST /lab/{userId}/save)
	PostLabSave(w http.ResponseWriter, r *http.Request, userId UserId, params PostLabSaveParams)
	// ログイン
	// (POST /login)
	Login(w http.ResponseWriter, r *http.Request)
	// ログアウト
	// (GET /logout)
	Logout(w http.ResponseWriter, r *http.Request, params LogoutParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// PostGraduateFix operation middleware
func (siw *ServerInterfaceWrapper) PostGraduateFix(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameter("simple", false, "userId", chi.URLParam(r, "userId"), &userId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter userId: %s", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params PostGraduateFixParams

	headers := r.Header

	// ------------- Optional header parameter "Set-Cookie" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Set-Cookie")]; found {
		var SetCookie SetCookie
		n := len(valueList)
		if n != 1 {
			http.Error(w, fmt.Sprintf("Expected one value for Set-Cookie, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "Set-Cookie", runtime.ParamLocationHeader, valueList[0], &SetCookie)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter Set-Cookie: %s", err), http.StatusBadRequest)
			return
		}

		params.SetCookie = &SetCookie

	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostGraduateFix(w, r, userId, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostGraduateSave operation middleware
func (siw *ServerInterfaceWrapper) PostGraduateSave(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameter("simple", false, "userId", chi.URLParam(r, "userId"), &userId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter userId: %s", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params PostGraduateSaveParams

	headers := r.Header

	// ------------- Optional header parameter "Set-Cookie" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Set-Cookie")]; found {
		var SetCookie SetCookie
		n := len(valueList)
		if n != 1 {
			http.Error(w, fmt.Sprintf("Expected one value for Set-Cookie, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "Set-Cookie", runtime.ParamLocationHeader, valueList[0], &SetCookie)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter Set-Cookie: %s", err), http.StatusBadRequest)
			return
		}

		params.SetCookie = &SetCookie

	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostGraduateSave(w, r, userId, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetHome operation middleware
func (siw *ServerInterfaceWrapper) GetHome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetHome(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostLabFix operation middleware
func (siw *ServerInterfaceWrapper) PostLabFix(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameter("simple", false, "userId", chi.URLParam(r, "userId"), &userId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter userId: %s", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params PostLabFixParams

	headers := r.Header

	// ------------- Optional header parameter "Set-Cookie" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Set-Cookie")]; found {
		var SetCookie SetCookie
		n := len(valueList)
		if n != 1 {
			http.Error(w, fmt.Sprintf("Expected one value for Set-Cookie, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "Set-Cookie", runtime.ParamLocationHeader, valueList[0], &SetCookie)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter Set-Cookie: %s", err), http.StatusBadRequest)
			return
		}

		params.SetCookie = &SetCookie

	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostLabFix(w, r, userId, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostLabSave operation middleware
func (siw *ServerInterfaceWrapper) PostLabSave(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameter("simple", false, "userId", chi.URLParam(r, "userId"), &userId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter userId: %s", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params PostLabSaveParams

	headers := r.Header

	// ------------- Optional header parameter "Set-Cookie" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Set-Cookie")]; found {
		var SetCookie SetCookie
		n := len(valueList)
		if n != 1 {
			http.Error(w, fmt.Sprintf("Expected one value for Set-Cookie, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "Set-Cookie", runtime.ParamLocationHeader, valueList[0], &SetCookie)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter Set-Cookie: %s", err), http.StatusBadRequest)
			return
		}

		params.SetCookie = &SetCookie

	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostLabSave(w, r, userId, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Login operation middleware
func (siw *ServerInterfaceWrapper) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Login(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Logout operation middleware
func (siw *ServerInterfaceWrapper) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params LogoutParams

	headers := r.Header

	// ------------- Optional header parameter "Set-Cookie" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Set-Cookie")]; found {
		var SetCookie SetCookie
		n := len(valueList)
		if n != 1 {
			http.Error(w, fmt.Sprintf("Expected one value for Set-Cookie, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "Set-Cookie", runtime.ParamLocationHeader, valueList[0], &SetCookie)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter Set-Cookie: %s", err), http.StatusBadRequest)
			return
		}

		params.SetCookie = &SetCookie

	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Logout(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL     string
	BaseRouter  chi.Router
	Middlewares []MiddlewareFunc
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/graduate/{userId}/fix", wrapper.PostGraduateFix)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/graduate/{userId}/save", wrapper.PostGraduateSave)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/home", wrapper.GetHome)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/lab/{userId}/fix", wrapper.PostLabFix)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/lab/{userId}/save", wrapper.PostLabSave)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/login", wrapper.Login)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/logout", wrapper.Logout)
	})

	return r
}
