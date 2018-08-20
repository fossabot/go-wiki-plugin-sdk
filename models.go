package module

import (
	"net/url"
	"encoding/gob"
)

type IGoWikiPlugin interface {
	Init()
	Name() string
	Version() string
	HandleRoute(string, HTTPRequest) string
}

type HTTPRequest struct {
	URL *url.URL
}

type GoWikiPluginConnector struct {
	Impl IGoWikiPlugin
}

func init() {
	gob.RegisterName("HTTPRequest", HTTPRequest{})
}