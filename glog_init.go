package glog

import (
	"flag"
	"os"
	"strings"
)

func initFromArgv() {
	args := parseFlags(map[string]bool{
		"logtosyslog":     true,
		"logtostderr":     true,
		"alsologtostderr": true,
		"stderrthreshold": false,
		"logthreshold":    false,
		"v":               false,
		"vmodule":         false,
		"log_dir":         false,
	})

	for name, val := range args {
		if f := flag.Lookup(name); f != nil {
			f.Value.Set(val)
		}
	}
}

func parseFlags(flags map[string]bool) map[string]string {
	args := map[string]string{}
	for i, n := 1, len(os.Args); i < n; i++ {
		arg := os.Args[i]
		if arg == "--" {
			break
		}
		if !isOption(arg) {
			continue
		}
		arg = strings.TrimLeft(arg, "-")
		if j := strings.IndexByte(arg, '='); j > 0 {
			name, val := arg[:j], arg[j+1:]
			if _, ok := flags[name]; ok {
				args[name] = val
			}
		} else if boolFlag, ok := flags[arg]; ok {
			if boolFlag {
				args[arg] = "true"
			} else if i+1 < n && !isOption(os.Args[i+1]) {
				i++
				args[arg] = os.Args[i]
			}
		}
	}
	return args
}

func isOption(s string) bool {
	return len(s) != 0 && s[0] == '-'
}
