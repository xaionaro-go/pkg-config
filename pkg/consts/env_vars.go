package consts

type EnvVar = string

var (
	// EnvVarStaticLibsList is the key of the environment variable that
	// contains a comma-separated list of wildcard patterns
	// of libs that will be forced to be statically linked.
	EnvVarStaticLibsList = EnvVar("PKG_CONFIG_LIBS_FORCE_STATIC")

	// EnvVarDynamicLibsList is the key of the environment variable that
	// contains a comma-separated list of wildcard patterns
	// of libs that will be forced to be dynamically linked.
	EnvVarDynamicLibsList = EnvVar("PKG_CONFIG_LIBS_FORCE_DYNAMIC")

	// EnvVarDynamicLibsList is the key of the environment variable that
	// contains a comma-separated list of wildcard patterns
	// of words to be deleted from the final output.
	EnvVarEraseList = EnvVar("PKG_CONFIG_ERASE")

	// EnvVarLogFile is the key of the environment variable that
	// contains the path to the file that will be used to dump logs to.
	EnvVarLogFile = EnvVar("PKG_CONFIG_WRAPPER_LOG")

	// EnvVarLogLevel is the key of the environment variable that
	// contains the logging level setting.
	EnvVarLogLevel = EnvVar("PKG_CONFIG_WRAPPER_LOG_LEVEL")
)
