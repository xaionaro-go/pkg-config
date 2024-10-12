package pkgconfig

type Config struct {
	CommandExecutor          CommandExecutor
	ForceStaticLinkPatterns  Patterns
	ForceDynamicLinkPatterns Patterns
	ErasePatterns            Patterns
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

type OptionForceStaticLinkPatterns Patterns

func (opt OptionForceStaticLinkPatterns) apply(c *Config) {
	c.ForceStaticLinkPatterns = Patterns(opt)
}

type OptionForceDynamicLinkPatterns Patterns

func (opt OptionForceDynamicLinkPatterns) apply(c *Config) {
	c.ForceDynamicLinkPatterns = Patterns(opt)
}

type OptionCommandExecutor struct{ CommandExecutor }

func (o OptionCommandExecutor) apply(c *Config) {
	c.CommandExecutor = o.CommandExecutor
}

type OptionErasePatterns Patterns

func (opt OptionErasePatterns) apply(c *Config) {
	c.ErasePatterns = Patterns(opt)
}
