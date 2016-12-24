package basic

import (
	"os"
	logging "third/go-logging"

	"github.com/astaxie/beego/session"
)

//Log xx
var Log *logging.Logger
var GlobalSessions *session.Manager

//InitLogger xx
func InitLogger() (*logging.Logger, error) {
	var log = logging.MustGetLogger("")

	formatStr := "%{color}%{level:.4s}:%{time:2006-01-02 15:04:05.000}  %{shortfile}%{color:reset} %{message}"
	format := logging.MustStringFormatter(formatStr)
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend1Leveled := logging.NewBackendFormatter(backend1, format)
	logging.SetBackend(backend1Leveled)
	Log = log
	return log, nil
}

func init() {
	a := &session.ManagerConfig{}
	GlobalSessions, _ = session.NewManager("memory", a)
	go GlobalSessions.GC()
}
