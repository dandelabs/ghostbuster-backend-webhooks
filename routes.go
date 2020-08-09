package main

import (
	//"net/http"

	"github.com/dandelabs/ghostbuster-backend-webhooks/api"
	"github.com/dandelabs/ghostbuster-backend-webhooks/router"

	"net/http"
)

// Constant regex and params
const (
	token         = "{token}"
	link          = "{link}"
	regexUUID     = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"
	regexEmaiil   = "{email:\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3}}"
	UUID          = "{uuid:" + regexUUID + "}"
	dependentUUID = "{uuid_dependent:" + regexUUID + "}"
	permission    = "{permission}"
)

//All routes should be added here
const (
	// VERSION = "v2"
	// we need home for ELB to stay alive
	HOME = "/"
	// /api/order
	postCreateOrder = "/api/order"
)

var routes = router.Routes{

	/****HOME****/
	router.Route{
		Name:        "HOME",
		Method:      "GET",
		Pattern:     HOME,
		NeedToken:   false,
		HandlerFunc: router.Home,
	},

	router.Route{
		"Create new order",
		"POST",
		postCreateOrder,
		false,
		func(w http.ResponseWriter, r *http.Request) {
			router.PostManager(w, r, api.CreateOrder)
		},
	},
}
