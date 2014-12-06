package web

import (
	"drill/config"
	//"drill/logs"
	"crypto/tls"
	"drill/helper"
	"drill/lua"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"reflect"
	"strings"
	"time"
	//	"regexp"
)

type Server struct {
	c    *config.ConfigContainer
	Lua  *lua.State
	Log  *log.Logger
	l    net.Listener
	Name string
}

type Event struct {
	method      string
	handler     ControlInterface
	httpHandler http.Handler
}

var EventList map[string]Event = map[string]Event{}

//创建服务
func NewServer() *Server {
	s := &Server{}
	if c, err := config.NewConfig("ini", "conf/app.ini"); err != nil {
		panic(err.Error())
	} else {
		s.c = &c
	}

	s.Lua = lua.NewLua()
	s.Lua.Openlibs()
	s.Log = nil
	return s

}

func (self *Server) initServer() {

	if self.Log == nil {
		self.Log = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	}

}

//添加事件组

func (self *Server) addEvent(r string, method string, handler interface{}) {

	switch handler.(type) {
	case http.Handler:
		EventList[r] = Event{method: method, httpHandler: handler.(http.Handler)}
	case ControlInterface:

		EventList[r] = Event{method: method, handler: handler.(ControlInterface)}
	default:

		EventList[r] = Event{method: method, handler: handler.(ControlInterface), httpHandler: nil}
	}
}

// ServeHTTP is the interface method for Go's http server package

func (self *Server) ServeHTTP(c http.ResponseWriter, req *http.Request) {
	self.Process(c, req)
}

// Process invokes the routing system for server s
func (self *Server) Process(c http.ResponseWriter, req *http.Request) {
	reqPath := strings.Trim(req.URL.Path, "/")

	c.Header().Set("Server", self.Name)
	c.Header().Set("Date", helper.GMT(time.Now().UTC()))
	c.Header().Set("Content-Type", "text/html; charset=utf-8")

	var contrl string
	var method string

	if reqPath == "" {
		contrl = "default"
		method = "Index"

	} else {

		pathS := strings.Split(reqPath, "/")

		l := len(pathS)

		if l == 1 {
			contrl = pathS[0]
			method = "Index"

		}

		if l >= 2 {
			contrl = pathS[0]
			method = strings.ToUpper(pathS[1][0:1]) + pathS[1][1:len(pathS[1])]

		}
	}

	if _, ok := EventList[contrl]; ok {

		if EventList[contrl].httpHandler != nil {
			EventList[contrl].httpHandler.ServeHTTP(c, req)
			return
		} else {
			obj := EventList[contrl].handler
			obj.Lock()
			obj.SetRequest(req)
			obj.SetResponseWriter(c)
			self.Log.Printf(method)
			reflect.ValueOf(obj).MethodByName(method).Call(nil)
			obj.Flush()
			obj.Unlock()
			return
		}
	} else {
		c.WriteHeader(404)
		c.Write([]byte("Page Not Found"))
	}

	return

}

/*
func (self *Server) Websocket(route string, httpHandler websocket.Handler) {
	self.addEvent(route, "GET", httpHandler)
}
*/

// Run starts the web application and serves HTTP requests for s
func (self *Server) Run(addr string) {
	self.initServer()

	mux := http.NewServeMux()
	if true {
		mux.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		mux.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
		mux.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	}
	mux.Handle("/", self)

	self.Log.Printf("web.go serving %s\n", addr)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	self.l = l
	err = http.Serve(self.l, mux)
	self.l.Close()
}

// RunFcgi starts the web application and serves FastCGI requests for self.
func (self *Server) RunFcgi(addr string) {
	self.initServer()
	self.Log.Printf("web.go serving fcgi %s\n", addr)
	//self.listenAndServeFcgi(addr)
}

// RunScgi starts the web application and serves SCGI requests for self.
func (self *Server) RunScgi(addr string) {
	self.initServer()
	self.Log.Printf("web.go serving scgi %s\n", addr)
	//self.listenAndServeScgi(addr)
}

// RunTLS starts the web application and serves HTTPS requests for self.
func (self *Server) RunTLS(addr string, config *tls.Config) error {
	self.initServer()
	mux := http.NewServeMux()
	mux.Handle("/", self)
	l, err := tls.Listen("tcp", addr, config)
	if err != nil {
		log.Fatal("Listen:", err)
		return err
	}

	self.l = l
	return http.Serve(self.l, mux)
}

// Close stops server self.
func (self *Server) Close() {
	if self.l != nil {
		self.l.Close()
	}
}

func AddEvent(r string, method string, handler interface{}) {

	switch handler.(type) {
	case http.Handler:
		EventList[r] = Event{method: method, httpHandler: handler.(http.Handler)}
	case ControlInterface:

		EventList[r] = Event{method: method, handler: handler.(ControlInterface)}
	default:

		EventList[r] = Event{method: method, handler: handler.(ControlInterface), httpHandler: nil}
	}
}
