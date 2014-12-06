#include "com_windows.h"

int isEqGuid(GoUintptr p0,GoUintptr p1){
	
	return IsEqualGUID((CLASSID)p0,(CLASSID)p1);
}

GoUintptr comGetGoUinptr(CState* L)
{
	GoUintptr* ret;
	ret = L->ptr;
	L->ptr += sizeof(GoUintptr); 
	return *ret;
	/* code */
}

char* comGetString(CState* L,int len)
{ 
	char* str;
	str = mollac(sizeof(char)*len);
	memcpy(str,L->ptr,len)
	L->ptr += sizeof(char)*len;
	return str;
}

int comInt(CState* L)
{
	int* ret;
	ret = L->ptr;
	L->ptr += sizeof(int); 
	return *ret;
}

void comfree(CState* L){
	free(L->byte);
	free(L);
}

GoUintptr NewIClassFactory(){
	IClassFactory* Factory;
	Factory =  (IClassFactory*)malloc(sizeof(IClassFactory))
	IClassFactoryVtbl* FactoryVtbl;
	FactoryVtbl = (IClassFactoryVtbl*)malloc(sizeof(IClassFactoryVtbl))

	FactoryVtbl->QueryInterface = QueryInterface;
	FactoryVtbl->AddRef = AddRef;
	FactoryVtbl->Release = Release;

	FactoryVtbl->CreateInstance = CreateInstance;
	FactoryVtbl->LockServer = LockServer;

	return (GoUintptr)Factory;
}

GoUintptr NewIDispatch(){
	IDispatch* Idis;
	Idis =  (IDispatch*)malloc(sizeof(IDispatch))
	IDispatchVtbl* IdisVtbl;
	IdisVtbl = (DispatchVtbl*)malloc(sizeof(IDispatchVtbl))
	IdisVtbl->QueryInterface = QueryInterface;
	IdisVtbl->AddRef = AddRef;
	IdisVtbl->Release = Release;

	IdisVtbl->GetTypeInfoCount = GetTypeInfoCount;
	IdisVtbl->GetTypeInfo = GetTypeInfo;
	IdisVtbl->GetIDsOfNames = GetIDsOfNames;
	IdisVtbl->Invoke = Invoke;

	return (GoUintptr)Idis;
}


#pragma cgo_export_static DllGetClassObject
#pragma cgo_export_dynamic DllGetClassObject

static HRESULT DllGetClassObject(CLASSID cid,POID pid,void** p)
{
	HRESULT hr;

	hr = GoGetClassObject((GoUintptr) cid,(GoUintptr) pid,(GoUintptr) p);
	
	return hr;
}

#pragma cgo_export_static DllRegisterServer
#pragma cgo_export_dynamic DllRegisterServer
static HRESULT DllRegisterServer()
{
}

 static HRESULT STDCALL QueryInterface(void* this, REFIID vt, void **ppv)
{
	CState* a;
	a = (CState*)malloc(sizeof(CState));
	a->byte = (void*)malloc(sizeof(char)*8);
	a->ptr = a->byte;
	a->len = 8;
	memcpy(a->ptr,&vt,4);
	memcpy((a->ptr+4),&ppv,4);
	return goComback((GoUintptr) this,"QueryInterface",a);
}

static ULONG STDCALL AddRef(void* this)
{

	return goComback((GoUintptr) this,"AddRef",NULL);
}

static ULONG STDCALL Release(void* this)
{
	return goComback((GoUintptr) this,"Release",NULL);
}

static HRESULT STDCALL CreateInstance(void* this, REFIID vt, void **ppv)
{
	CState* a;
	a = (CState*)malloc(sizeof(CState));
	a->byte = (void*)malloc(sizeof(char)*8);
	a->ptr = a->byte;
	a->len = 8;
	memcpy(a->ptr,&vt,4);
	memcpy((a->ptr+4),&ppv,4);
 	return goComback((GoUintptr) this,"CreateInstance",a);
}

static HRESULT STDCALL LockServer(void* this,int flag)
{
	CState* a;
	a = (CState*)malloc(sizeof(CState));
	a->byte = (void*)malloc(sizeof(char)*4);
	a->ptr = a->byte;
	a->len = 4;
	memcpy(a->ptr,&flag,4);

 	return goComback((GoUintptr) this,"LockServer",a);
}

static HRESULT STDCALL GetTypeInfoCount(void* this,UINT* flag)
{
	CState* a;
	a = (CState*)malloc(sizeof(CState));
	a->byte = (void*)malloc(sizeof(char)*4);
	a->ptr = a->byte;
	a->len = 4;
	memcpy(a->ptr,&flag,4);
 	return goComback((GoUintptr) this,"GetTypeInfoCount"，a);
}

static HRESULT STDCALL GetTypeInfo(void* this,int flag)
{
 	return goComback((GoUintptr) this,"GetTypeInfo"，&a,16);
}

static HRESULT STDCALL GetIDsOfNames(void* this,int flag)
{
 	return goComback((GoUintptr) this,"GetIDsOfNames"，&a,16);
}

static HRESULT STDCALL Invoke(void* this,int flag)
{
 	return goComback((GoUintptr) this,"Invoke"，&a,16);
}






