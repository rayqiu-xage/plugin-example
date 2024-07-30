package main

import (
	"errors"
	"myplugin/pam_plugin" // Import the plugin interface package
	"time"

	"github.com/hashicorp/go-plugin"
)

var userMap map[string]pam_plugin.UserRecord

// Implementation of the Plugin interface.
type PAMUsers struct{}

func (p *PAMUsers) DiscoverUsers() ([]pam_plugin.UserRecord, error) {
	values := make([]pam_plugin.UserRecord, 0)
	for _, v := range userMap {
		values = append(values, v)
	}
	return values, nil
}

func (p *PAMUsers) RotatePassword(userPass pam_plugin.UserPasswordPair) error {
	u, ok := userMap[userPass.Name]
	if !ok {
		return errors.New("User does not exist")
	}
	u.Password = userPass.Password
	userMap[userPass.Name] = u
	return nil
}

// Handshake configuration.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "PLUGIN_MAGIC_COOKIE",
	MagicCookieValue: "PAM",
}

func main() {
	userMap = make(map[string]pam_plugin.UserRecord)
	userMap["user1"] = pam_plugin.UserRecord{Name: "user1", Password: "password1", CreatedAt: time.Now()}
	userMap["user2"] = pam_plugin.UserRecord{Name: "user2", Password: "password2", CreatedAt: time.Now()}
	pam := &PAMUsers{}

	var pluginMap = map[string]plugin.Plugin{
		"pam": &pam_plugin.PAMPlugin{Impl: pam},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		GRPCServer:      plugin.DefaultGRPCServer,
	})
}
