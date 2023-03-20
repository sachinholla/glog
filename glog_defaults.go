//go:build glog_defaults

package glog

func init() {
	logging.toSyslog = true
}
