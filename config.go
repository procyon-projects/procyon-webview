package webview

import _ "github.com/procyon-projects/procyon-context"

type Configuration struct {
	Prefix string `yaml:"prefix" json:"prefix" default:"/views"`
	Suffix string `yaml:"suffix" json:"suffix" default:".gop"`
}

func newConfiguration() *Configuration {
	return &Configuration{}
}

func (configuration *Configuration) GetConfigurationPrefix() string {
	return "procyon.webview"
}
