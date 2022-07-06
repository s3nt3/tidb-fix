//go:build tools

package tools

import (
	_ "github.com/gin-gonic/gin"
	_ "github.com/gorilla/websocket"
	_ "golang.org/x/net/http2"
	_ "golang.org/x/text"
	_ "gopkg.in/yaml.v3"
)
