package drill

import (
	"drill/session"
	"drill/lua"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

const VERSION = "0.1.0"

type Server interface { 
	func Run()
	func Stop()
	func Restart()
	func Init()
}

var ServerList map[string]Server


func Start(Name string){ 
		iniconf, err := NewConfig("ini", "server.ini")

		if err != nil {
		    t.Fatal(err)
		}

		ServerList[Name].Run()
}
func Stop(){

}





