package cmd

const (
	VERSION                  = "0.4.2"
	PROGNAME                 = "litt"
	CONFIG_FILE_NOT_FOUND    = 2
	CONFIG_FILE_NOT_LOADABLE = 4
	INVALID_CONFIG_FILE      = 8
	LISTED_FILE_NOT_FOUND    = 16
	LISTED_MONTAGE_NOT_FOUND = 32
)

var (
	config        = "./" + PROGNAME + ".json" // default config file
	configuration Configuration
)
