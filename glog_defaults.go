// +build glog_defaults

package glog

func init() {
	logging.alsoToSyslog = true
}
