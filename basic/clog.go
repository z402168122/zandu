package basic

import (
	logging "third/go-logging"

	"github.com/astaxie/beego/session"
)

//Log xx
var Log *logging.Logger
var GlobalSessions *session.Manager

func init() {
	a := &session.ManagerConfig{}
	GlobalSessions, _ = session.NewManager("memory", a)
	go GlobalSessions.GC()
}
