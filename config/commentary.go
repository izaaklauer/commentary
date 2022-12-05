package config

type Commentary struct {
HelloWorldMessage string `hcl:"hello_world_message,attr"`

// ... your config here
}

// DefaultCommentaryConfig returns default config values
func DefaultCommentaryConfig() Commentary {
	return Commentary{
	HelloWorldMessage:
		"hello from the default config",
	}
}
