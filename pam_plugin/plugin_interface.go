package pam_plugin

import (
	"net/rpc"
	"time"

	"github.com/hashicorp/go-plugin"
)

type UserRecord struct {
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserPasswordPair struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// PAM is the interface that PAM plugins need to implement.
type PAM interface {
	DiscoverUsers() ([]UserRecord, error)
	RotatePassword(userPass UserPasswordPair) error
}

// PluginRPC is the exported interface for RPC.
type PAMPluginRPC struct {
	client *rpc.Client
}

// PAM calls the plugin's DiscoverUsers method.
func (p *PAMPluginRPC) DiscoverUsers() ([]UserRecord, error) {
	var resp []UserRecord
	err := p.client.Call("Plugin.DiscoverUsers", new(interface{}), &resp)
	return resp, err
}

// PAM calls the plugin's RotatePassword method.
func (p *PAMPluginRPC) RotatePassword(userPass UserPasswordPair) error {
	err := p.client.Call("Plugin.RotatePassword", userPass, new(interface{}))
	return err
}

// PAMPluginRPCServer is the server-side implementation of Plugin.
type PAMPluginRPCServer struct {
	Impl PAM
}

func (s *PAMPluginRPCServer) DiscoverUsers(args interface{}, resp *[]UserRecord) error {
	result, err := s.Impl.DiscoverUsers()
	*resp = result
	return err
}

func (s *PAMPluginRPCServer) RotatePassword(userPass UserPasswordPair, resp *interface{}) error {
	return s.Impl.RotatePassword(userPass)
}

// PAMPlugin is the plugin implementation.
type PAMPlugin struct {
	Impl PAM
}

func (p *PAMPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &PAMPluginRPCServer{Impl: p.Impl}, nil
}

func (p *PAMPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &PAMPluginRPC{client: c}, nil
}
