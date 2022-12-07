// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.1 DO NOT EDIT.
package oapi

// Graduate defines model for graduate.
type Graduate bool

// Labs defines model for labs.
type Labs struct {
	Name       *string    `json:"name,omitempty"`
	Professors *Professor `json:"professors,omitempty"`
}

// Professor defines model for professor.
type Professor struct {
	Name  *string `json:"name,omitempty"`
	Point *int    `json:"point,omitempty"`
}

// ログイン情報
type User struct {
	Password *string `json:"password,omitempty"`
	UserId   *string `json:"user_id,omitempty"`
}

// SetCookie defines model for Set-Cookie.
type SetCookie string

// UserId defines model for userId.
type UserId string

// OK defines model for OK.
type OK interface{}

// PostGraduateFixJSONBody defines parameters for PostGraduateFix.
type PostGraduateFixJSONBody Graduate

// PostGraduateFixParams defines parameters for PostGraduateFix.
type PostGraduateFixParams struct {

	// jwtの情報を入れる
	SetCookie *SetCookie `json:"Set-Cookie,omitempty"`
}

// PostGraduateSaveJSONBody defines parameters for PostGraduateSave.
type PostGraduateSaveJSONBody Graduate

// PostGraduateSaveParams defines parameters for PostGraduateSave.
type PostGraduateSaveParams struct {

	// jwtの情報を入れる
	SetCookie *SetCookie `json:"Set-Cookie,omitempty"`
}

// GetHomeJSONBody defines parameters for GetHome.
type GetHomeJSONBody Labs

// PostLabFixJSONBody defines parameters for PostLabFix.
type PostLabFixJSONBody Labs

// PostLabFixParams defines parameters for PostLabFix.
type PostLabFixParams struct {

	// jwtの情報を入れる
	SetCookie *SetCookie `json:"Set-Cookie,omitempty"`
}

// PostLabSaveJSONBody defines parameters for PostLabSave.
type PostLabSaveJSONBody Labs

// PostLabSaveParams defines parameters for PostLabSave.
type PostLabSaveParams struct {

	// jwtの情報を入れる
	SetCookie *SetCookie `json:"Set-Cookie,omitempty"`
}

// LoginJSONBody defines parameters for Login.
type LoginJSONBody User

// LogoutParams defines parameters for Logout.
type LogoutParams struct {

	// jwtの情報を入れる
	SetCookie *SetCookie `json:"Set-Cookie,omitempty"`
}

// PostGraduateFixJSONRequestBody defines body for PostGraduateFix for application/json ContentType.
type PostGraduateFixJSONRequestBody PostGraduateFixJSONBody

// PostGraduateSaveJSONRequestBody defines body for PostGraduateSave for application/json ContentType.
type PostGraduateSaveJSONRequestBody PostGraduateSaveJSONBody

// GetHomeJSONRequestBody defines body for GetHome for application/json ContentType.
type GetHomeJSONRequestBody GetHomeJSONBody

// PostLabFixJSONRequestBody defines body for PostLabFix for application/json ContentType.
type PostLabFixJSONRequestBody PostLabFixJSONBody

// PostLabSaveJSONRequestBody defines body for PostLabSave for application/json ContentType.
type PostLabSaveJSONRequestBody PostLabSaveJSONBody

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody LoginJSONBody