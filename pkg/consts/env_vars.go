package consts

type EnvVar = string

var (
	// A comma-separated list of wildcard patterns of libs that will
	// be forced to be statically linked.
	EnvVarStaticLibsList = EnvVar("PKG_CONFIG_LIBS_FORCE_STATIC")

	// A comma-separated list of wildcard patterns of libs that will
	// be forced to be dynamically linked.
	EnvVarDynamicLibsList = EnvVar("PKG_CONFIG_LIBS_FORCE_DYNAMIC")
)
