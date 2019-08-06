package cmd

const (
	VERSION = "0.3.0"
	BULLET  = "⮞"
)

var (
	cfgFile       string
	jsonPath      string
	basedir       string
	configuration Configuration
	reference     string // name or id of used montage
	tag           bool
	nocontent     bool
	verbose       bool
)
