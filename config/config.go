// Package config parses command-line/environment/config file arguments and puts
// together the configuration of this instance, which is made available to other
// packages.
package config

import (
	"github.com/levenlabs/go-llog"
	"github.com/mediocregopher/lever"
)

var (
	ElasticSearchAddr string
	LuaInit           string
	LuaVMs            int
	LogLevel          string
)

func init() {
	l := lever.New("thumper", nil)
	l.Add(lever.Param{
		Name:        "--elasticsearch-addr",
		Description: "Address to find an elasticsearch instance on",
		Default:     "127.0.0.1:9200",
	})
	l.Add(lever.Param{
		Name:        "--lua-init",
		Description: "If set the given lua script will be executed at the initialization of every lua vm",
	})
	l.Add(lever.Param{
		Name:        "--lua-vms",
		Description: "How many lua vms should be used. Each vm is completely independent of the other, and requests are executed on whatever vm is available at that moment. Allows lua scripts to not all be blocked on the same os thread",
		Default:     "1",
	})
	l.Add(lever.Param{
		Name:        "--log-level",
		Description: "Adjust the log level. Valid options are: error, warn, info, debug",
		Default:     "info",
	})
	l.Parse()

	ElasticSearchAddr, _ = l.ParamStr("--elasticsearch-addr")
	LuaInit, _ = l.ParamStr("--lua-init")
	LuaVMs, _ = l.ParamInt("--lua-vms")
	LogLevel, _ = l.ParamStr("--log-level")
	llog.SetLevelFromString(LogLevel)
}
