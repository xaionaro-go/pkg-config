package pkgconfig

type Config struct {
	CommandExecutor          CommandExecutor
	ForceStaticLinkPatterns  []string
	ForceDynamicLinkPatterns []string
}

type Option interface {
	apply(*Config)
}

type Options []Option

func (opts Options) apply(cfg *Config) {
	for _, opt := range opts {
		opt.apply(cfg)
	}
}

func (opts Options) Config() Config {
	cfg := Config{
		CommandExecutor: DefaultCommandExecutor,
	}
	opts.apply(&cfg)
	return cfg
}

type OptionForceStaticLinkPatterns []string

func (o OptionForceStaticLinkPatterns) apply(c *Config) {
	c.ForceStaticLinkPatterns = o
}

type OptionForceDynamicLinkPatterns []string

func (o OptionForceDynamicLinkPatterns) apply(c *Config) {
	c.ForceDynamicLinkPatterns = o
}

type OptionCommandExecutor struct{ CommandExecutor }

func (o OptionCommandExecutor) apply(c *Config) {
	c.CommandExecutor = o.CommandExecutor
}
