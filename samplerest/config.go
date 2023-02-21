package samplerest

import (
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

// This uses go-cty: https://github.com/zclconf/go-cty/blob/main/docs/gocty.md#converting-to-and-from-structs
type SampleRESTConfig struct {
	Email       *string `cty:"email"`
	Password    *string `cty:"password"`
	OtherConfig *bool   `cty:"other_config"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"email":        {Type: schema.TypeString},
	"password":     {Type: schema.TypeString},
	"other_config": {Type: schema.TypeBool},
}

func ConfigInstance() interface{} {
	return &SampleRESTConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) SampleRESTConfig {
	if connection == nil || connection.Config == nil {
		return SampleRESTConfig{}
	}
	config, _ := connection.Config.(SampleRESTConfig)
	return config
}

func (c SampleRESTConfig) String() string {
	return fmt.Sprintf(
		"SampleRESTConfig{email=%s, password=*** (len=%d), other_config=%t}",
		*c.Email, len(*c.Password), *c.OtherConfig)
}
