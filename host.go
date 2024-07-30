package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"myplugin/pam_plugin" // Import the plugin interface package

	"github.com/hashicorp/go-plugin"
)

func main() {
	// Set up plugin handshake configuration
	handshakeConfig := plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "PLUGIN_MAGIC_COOKIE",
		MagicCookieValue: "PAM",
	}

	// Create a new client for the plugin
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"pam": &pam_plugin.PAMPlugin{},
		},
		Cmd: exec.Command("./plugin_main"), // Command to launch the plugin process
	})

	// Ensure the client is killed when main() exits
	defer client.Kill()

	// Connect to the plugin via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("pam")
	if err != nil {
		log.Fatal(err)
	}

	// Assert the raw plugin instance to the expected interface type
	pamPlugin, ok := raw.(pam_plugin.PAM)
	if !ok {
		log.Fatal("Plugin does not implement PAM interface")
	}

	// Use the DiscoverUsers method provided by the PAM interface
	users, err := pamPlugin.DiscoverUsers()
	if err != nil {
		log.Fatal(err)
	}
	str, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(str))

	// Use the RotatePassword method provided by the PAM interface
	err = pamPlugin.RotatePassword(pam_plugin.UserPasswordPair{Name: "user1", Password: "NewPassword1"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Password rotated for user1.")

	// Use the DiscoverUsers method provided by the PAM interface again
	users, err = pamPlugin.DiscoverUsers()
	if err != nil {
		log.Fatal(err)
	}
	str, err = json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(str))
}
