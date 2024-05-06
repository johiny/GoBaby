package models

type Routes struct {
	MAIN          string
	LOGS          string
	OPTIONS       string
	CLOCK         string
	RESTART_CYCLE string
	LOG_TABLE     string
	ERROR         string
	CLEAR_ERROR   string
	LOGIN         string
	REGISTER      string
	USERMENU      string
	CHANGEUSER    string
}

var RoutesInstance = Routes{
	MAIN:          "/",
	LOGS:          "/logs",
	OPTIONS:       "/options",
	CLOCK:         "/clock",
	RESTART_CYCLE: "/clock/restart-cycle",
	LOG_TABLE:     "/log-table",
	ERROR:         "/error",
	CLEAR_ERROR:   "/clear-error",
	LOGIN:         "/login",
	REGISTER:      "/register",
	USERMENU:      "/user-menu",
	CHANGEUSER:    "/change-user",
}
