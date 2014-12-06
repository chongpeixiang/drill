#include "glua.h"
#include "_cgo_export.h"

void Glua_pushcfunction (lua_State *L, int f)
{
	lua_pushgofunction(L,(lua_CFunction)f);
}

void Glua_setglobal (lua_State *L, const char *name)
{	
	lua_setglobal(L,name);

}

void Glua_getglobal (lua_State *L, const char *name)
{ 
	lua_getglobal(L,name);

}

int GluaL_checkint (lua_State *L, int narg)
{ 
	luaL_checkint(L,narg);

}

int GluaL_loadfile (lua_State *L,const char* f)
{
	return luaL_loadfile(L,f);
}
int Glua_pcall (lua_State *L, int nargs, int nresults, int errfunc)
{
	return lua_pcall(L,nargs,nresults,errfunc);
}

int Glua_call (lua_State *L, int nargs, int nresults)
{
	 lua_call(L,nargs,nresults);
	 return 1;
}

int Gojit(lua_State *L, int narg)
{ 
	return goCallback((unsigned int)L,narg);
	 
}