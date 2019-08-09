package cmd

const (
	VERSION                      = "0.4.9"
	PROGNAME                     = "litt"
	EMPTY_FILE                   = 2
	CONFIG_FILE_NOT_LOADABLE     = 4
	UNMARSHALING_FAILED          = 8
	LISTED_FILE_NOT_FOUND        = 16
	LISTED_MONTAGE_NOT_FOUND     = 32
	MARSHALING_FAILED            = 64
	WRITE_FILE_FAILED            = 128
	FILE_NOT_DEFINED             = 256
	MONTAGE_NOT_DEFINED          = 512
	INVALID_CONFIGURATION_FORMAT = 1024
	ROOT_EXECUTE_ERROR           = 2048
)

var (
	configuration Configuration
)
