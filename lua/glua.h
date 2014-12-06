//#include "stdafx.h"
#include <stdio.h>
#include "lua.h"
#include "lualib.h"
#include "lauxlib.h"
/*
typedef void * (*lua_All) (void *ud,void *ptr,size_t osize,site_t nsize)
typedef int (*lua_CFunction) (lua_State *L);
typedef ptrdiff_t lua_Integer;
typedef const char * (*lua_Reader) (lua_State *L,void *data,size_t *size);
typedef struct lua_State lua_State;
typedef int (*lua_Writer) (lua_State *L,const void* p,size_t sz,void* ud);
*/
void Glua_pushcfunction (lua_State *L, int f);	
void Glua_setglobal (lua_State *L, const char *name);
void Glua_getglobal (lua_State *L, const char *name);
int GluaL_checkint (lua_State *L, int narg);
int GluaL_loadfile (lua_State *L,const char* f);
int Glua_pcall (lua_State *L, int nargs, int nresults, int errfunc);
int Glua_call (lua_State *L, int nargs, int nresults);
int Gluac(int argc, char* argv[]);