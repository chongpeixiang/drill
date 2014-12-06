package lua

/*
#include "glua.h"
#include <stdlib.h>
#cgo linux LDFLAGS: -lm
*/
import "C"

import "unsafe"
import "sync"
import "fmt"
import "reflect"

type State struct {
	S    *C.lua_State
	Lock *sync.Mutex
}

type GoFun func(*State) int
type Gofun_ struct {
	T uintptr
	F uintptr
}

type LuaState *C.lua_State

var reg map[uintptr]interface{} = make(map[uintptr]interface{})

/*
*state mainipulation
*
 */
func NewLua() *State {
	return &State{S: C.luaL_newstate(), Lock: new(sync.Mutex)}
}

func (self *State) Close() {
	C.lua_close(self.S)
}

func (self *State) NewThread() *State {
	s := C.lua_newthread(self.S)

	return &State{S: s, Lock: new(sync.Mutex)}
}

func (self *State) Version() float32 {
	return float32(*C.lua_version(self.S))
}

/*
*basic stack manipulation
 */

func (self *State) AbsIndex(idx int) int {
	return int(C.lua_absindex(self.S, C.int(idx)))
}

func (self *State) Gettop() int {
	return int(C.lua_gettop(self.S))
}

func (self *State) Settop(idx int) {
	C.lua_settop(self.S, C.int(idx))
}

func (self *State) Pushvalue(idx int) {
	C.lua_pushvalue(self.S, C.int(idx))
}

func (self *State) Remove(idx int) {
	C.lua_remove(self.S, C.int(idx))
}
func (self *State) Insert(idx int) {
	C.lua_insert(self.S, C.int(idx))
}

func (self *State) Replace(idx int) {
	C.lua_replace(self.S, C.int(idx))
}

func (self *State) Copy(src, dst int) {
	C.lua_copy(self.S, C.int(src), C.int(dst))
}

func (self *State) CheckStack(sz int) int {
	return int(C.lua_checkstack(self.S, C.int(sz)))
}

func Xmove(src, dst *State, n int) {
	C.lua_xmove(src.S, dst.S, C.int(n))
}

/**
*access function (stack -> Go)
**/

func (self *State) IsNumber(idx int) int {
	return int(C.lua_isnumber(self.S, C.int(idx)))
}

func (self *State) IsString(idx int) int {
	return int(C.lua_isstring(self.S, C.int(idx)))
}

func (self *State) IsCfunction(idx int) int {
	return int(C.lua_iscfunction(self.S, C.int(idx)))
}

func (self *State) IsUserdata(idx int) int {
	return int(C.lua_isuserdata(self.S, C.int(idx)))
}

func (self *State) Type(idx int) int {
	return int(C.lua_type(self.S, C.int(idx)))
}
func (self *State) TypeName(idx int) string {

	return C.GoString(C.lua_typename(self.S, C.int(idx)))
}

func (self *State) ToNumber(idx int) float64 {
	return float64(C.lua_tonumberx(self.S, C.int(idx), nil))
}
func (self *State) ToInteger(idx int) int {
	return int(C.lua_tointegerx(self.S, C.int(idx), nil))
}

func (self *State) ToUnsigned(idx int) uint {
	return uint(C.lua_tounsignedx(self.S, C.int(idx), nil))
}

func (self *State) ToBool(idx int) int {
	return int(C.lua_toboolean(self.S, C.int(idx)))
}

func (self *State) ToLString(idx int) string {
	return C.GoString(C.lua_tolstring(self.S, C.int(idx), nil))
}

func (self *State) Rawlen(idx int) int {
	return int(C.lua_rawlen(self.S, C.int(idx)))
}

func (self *State) ToUserdata(idx int) unsafe.Pointer {
	return unsafe.Pointer(C.lua_touserdata(self.S, C.int(idx)))
}

func (self *State) ToThread(idx int) *State {
	return &State{S: C.lua_tothread(self.S, C.int(idx))}
}

func (self *State) Openlibs() {
	C.luaL_openlibs(self.S)
}

func (self *State) Pushcfunction(fun interface{}) {

	if reflect.TypeOf(fun).Kind() != reflect.Func {
		panic("CallFunc: Can't call non-func")
	}
	//Tmp := unsafe.Pointer(&fun)
	Tmp := *(*Gofun_)(unsafe.Pointer(&fun))
	p := *(*uintptr)(unsafe.Pointer(Tmp.F))
	reg[p] = fun
	C.Glua_pushcfunction(self.S, C.int(p))
	//C.Glua_pushcfunction(self.S, C.int(uintptr(unsafe.Pointer(self))))
}

func (self *State) Setglobal(key string) {
	str := C.CString(key)
	defer C.free(unsafe.Pointer(str))
	C.Glua_setglobal(self.S, str)
}

func (self *State) LoadFile(filename string) error {
	str := C.CString(filename)
	defer C.free(unsafe.Pointer(str))
	if flag := C.GluaL_loadfile(self.S, str); flag == 0 {
		return nil
	} else {
		fmt.Println(flag)
	}
	return fmt.Errorf("Lua --file onload fiald: %s\n", filename)
}

func (self *State) PCall(nargs, nresults, errfunc int) int {

	C.Glua_pcall(self.S, C.int(nargs), C.int(nresults), C.int(errfunc))
	return 0
}

func (self *State) Call(nargs, nresults int) int {

	C.Glua_call(self.S, C.int(nargs), C.int(nresults))
	return 0
}

func (self *State) Getglobal(key string) {
	str := C.CString(key)
	defer C.free(unsafe.Pointer(str))
	C.Glua_getglobal(self.S, str)
}

func (self *State) DoFile(filename string) bool {
	if self.LoadFile(filename) == nil {
		return self.PCall(0, 0, 0) == 0
	}
	return false

}

/*
*push function (Go ->stack)
*
 */
func (self *State) PushNil() {
	C.lua_pushnil(self.S)
}

func (self *State) PushNumber(n float64) {
	C.lua_pushnumber(self.S, C.lua_Number(n))
}

func (self *State) PushInteger(n int) {
	C.lua_pushinteger(self.S, C.lua_Integer(n))
}

func (self *State) PushUInteger(n uint) {
	C.lua_pushunsigned(self.S, C.lua_Unsigned(n))
}

func (self *State) PushString(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.lua_pushstring(self.S, cstr)
}

func (self *State) PushBool(b int) {
	C.lua_pushboolean(self.S, C.int(b))
}

func (self *State) PushLightUserdata() {
	//C.lua_pushlightuserdata()
}

func (self *State) PushThread() int {
	return int(C.lua_pushthread(self.S))
}

func (self *State) SetField(idx int, k string) {
	cstr := C.CString(k)
	defer C.free(unsafe.Pointer(cstr))
	C.lua_setfield(self.S, C.int(idx), cstr)
}

func (self *State) Checkint(index int) int {

	return int(C.GluaL_checkint(self.S, C.int(index)))
}

func (self *State) GetSubTable(idx int, fname string) int {
	cstr := C.CString(fname)
	defer C.free(unsafe.Pointer(cstr))
	return int(C.luaL_getsubtable(self.S, C.int(idx), cstr))
}

func (self *State) SetTable(idx int) {

	C.lua_settable(self.S, C.int(idx))

}

func (self *State) CreateTable(x, y int) {
	C.lua_createtable(self.S, C.int(x), C.int(y))
}

func (self *State) SetFuncs(r map[string]interface{}, nup int) {
	self.CheckStack(nup)
	for k, v := range r {
		for i := 0; i < nup; i++ {
			self.Pushvalue(-nup)
		}
		self.Pushcfunction(v)
		self.SetField(-(nup + 2), k)
	}

	self.Settop(-(nup) - 1)

}

func (self *State) Newlib(r map[string]interface{}) {
	self.CreateTable(0, len(r)-1)
	self.SetFuncs(r, 0)
}

func (self *State) Requiref(modname string, Gfun interface{}, glb int32) {
	self.Pushcfunction(Gfun)
	self.PushString(modname)
	self.Call(1, 1)
	self.GetSubTable(1, "_LOADED")
	self.Pushvalue(-2)
	self.SetField(-2, modname)
	self.Settop(-2)
	if glb > 0 {
		self.Pushvalue(-1)
		self.Setglobal(modname)
	}
}

//s LuaState,ptr uintptr
//export goCallback
func goCallback(l, p uintptr) int {
	L := &State{S: (*C.lua_State)(unsafe.Pointer(l))}
	f := reg[p].(func(*State) int)
	//fmt.Println(f,L)

	return f(L)
}

func Dump(args ...string) {
	arg := make([](*_Ctype_char), 0) //C语言char*指针创建切片
	l := len(args)
	for i, _ := range args {
		char := C.CString(args[i])
		defer C.free(unsafe.Pointer(char)) //释放内存
		strptr := (*_Ctype_char)(unsafe.Pointer(char))
		arg = append(arg, strptr) //将char*指针加入到arg切片
	}

	C.Gluac(C.int(l), (**_Ctype_char)(unsafe.Pointer(&arg[0]))) //即c语言的main(int argc,char**argv)
}
