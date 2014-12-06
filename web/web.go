package web

import (
	"bufio"
	"drill/lua"
	"net/http"
	"strings"
	"sync"
	//"reflect"
	//"os"
	//"path"
	//"path/filepath"
)

const VERSION = "0.1.0"

type Control struct {
	R     *http.Request
	W     http.ResponseWriter
	WBuf  *bufio.Writer
	Param map[string]interface{}
	Data  map[string]interface{}
	Lua   *lua.State
	Loc   *sync.Mutex
}

type ControlInterface interface {
	SetRequest(R *http.Request)
	SetResponseWriter(w http.ResponseWriter)
	Lock()
	Unlock()
	Flush()
}

func (self *Control) SetRequest(R *http.Request) {
	self.R = R
	self.Param = make(map[string]interface{})
	self.Param["_URL_"] = strings.Split(strings.Trim(R.URL.Path, "/"), "/")
	self.R.ParseForm()

	if len(self.R.Form) > 0 {
		for k, v := range self.R.Form {
			pn := len(v)
			if pn == 1 {
				self.Param[k] = v[0]
			} else if pn > 1 {
				self.Param[k] = v
			}
		}
	}

	if len(self.R.PostForm) > 0 {
		for k, v := range self.R.PostForm {
			pl := len(v)

			if pl == 1 {
				self.Param[k] = v[0]
			} else if pl > 1 {
				self.Param[k] = v
			}
		}
	}
}

func (self *Control) SetResponseWriter(w http.ResponseWriter) {
	self.W = w
	self.WBuf = bufio.NewWriter(self.W)
}

func (self *Control) Lock() {
	self.Loc.Lock()
}

func (self *Control) Unlock() {
	self.Loc.Unlock()
}

func (self *Control) Index() {
	self.W.WriteHeader(200)
	self.W.Write([]byte("hallo drill!"))
}

func (self *Control) Get() {
	self.writeString("this is get")
}

func (self *Control) Post() {
	self.writeString("this is post")
}

func (self *Control) Put() {
	self.writeString("this is put")
}

func (self *Control) Delete() {
	self.writeString("this is delete")
}

func (self *Control) Flush() {
	self.WBuf.Flush()
}

func (self *Control) Display() {
	self.Lua = lua.NewLua()
}

func (self *Control) Redirect(status int, url string) {
	self.W.Header().Set("Location", url)
	self.W.WriteHeader(status)
	self.writeString("Redirecting to:" + url)
}

func (self *Control) writeString(data string) {
	self.WBuf.Write([]byte(data))
}

func init() {
	AddEvent("default", "Get", &Control{Loc: new(sync.Mutex)})
}
