package com

//#include "glua.h"

import(
    "C"
)

var OutstandingObjects int32
var LockCount int32
var ClassID REFIID
var ICF IClassFactory
const (
    S_OK = 0x00000000 //成功 
    S_FALSE = 0x00000001 //函数成功执行完成，但返回时出现错误 
    E_INVALIDARG = 0x80070057 //参数有错误 
    E_OUTOFMEMORY = 0x8007000E //内存申请错误 
    E_UNEXPECTED = 0x8000FFFF //未知的异常 
    E_NOTIMPL = 0x80004001 //未实现功能 
    E_FAIL = 0x80004005 //没有详细说明的错误。一般需要取得 Rich Error 错误信息(注1) 
    E_POINTER = 0x80004003 //无效的指针 
    E_HANDLE = 0x80070006 //无效的句柄 
    E_ABORT = 0x80004004 //终止操作 
    E_ACCESSDENIED = 0x80070005 //访问被拒绝 
    E_NOINTERFACE  = 0x80004002 //不支持接口 
)

type REFIID C.REFIID

//export GoGetClassObject
func GoGetClassObject( c,i,p uintptr)uint32{
    if C.int(C.isEqIid(,c)) == -1 {
        return E_NOINTERFACE
    }else{
        ICF
    }
    return S_OK
} 

//export DllCanUnloadNow
func DllCanUnloadNow(){

} 

//export DllRegisterServer
func DllRegisterServer(){

} 

//export DllUnregisterServer
func DllUnregisterServer(){

}

type IUnknown interface {
    QueryInterface() //查询接口
    AddRef()    //引用计数+1
    Release()   //引用计数-1
}

type IClassFactory struct {
         
     IUnknown
}

func CreateInstance(I IUnknown,guid string,ppv uintptr)

func LockServer(b bool)


type IDispatch  interface {
     GetTypeInfoCount(count uint*)
          
     GetTypeInfo(UINT iTInfo,  
            LCID lcid,  
            ITypeInfo **ppTInfo)  
          
     GetIDsOfNames(   
            REFIID riid,  
            LPOLESTR *rgszNames,  
            UINT cNames,  
            LCID lcid,  
            DISPID *rgDispId);  
          
     Invoke(DISPID dispIdMember,  
            REFIID riid,  
            LCID lcid,  
            WORD wFlags,  
            DISPPARAMS *pDispParams,  
            VARIANT *pVarResult,  
            EXCEPINFO *pExcepInfo,  
            UINT *puArgErr)
     IUnknown
}
/*
1、   CoGetClassObject
函数原型：
HRESULT CoGetClassObject(const CLSID& clsid, DWORD dwClsContext, COSERVERINFO* pServerInfo, const IID& iid, (void **)ppv);
参数说明：
         clsid: 指定COM类厂的CLSID标识符;
         dwClsContext: 指定组件类别(进程内、进程外组件等等);
         pServerInfo: DCOM留用;
         iid: 接口IClassFactory的标识符IID_IClassFactory;
         ppv: 存放类厂对象的接口指针;

2、   DllGetClassObject
函数定义：
HRESULT DllGetClassObject(const CLSID& clsid, const IID& iid, (void **)ppv)
{
           if(clsid == CLSID_Dictionary)
           {
                    CDictionaryFactory* pFactory = new CDictionaryFactory;
                    if(pFactory == NULL)
                    {
                             return E_OUTOFMEMORY;
                    }
                    HRESULT result = pFactory->QueryInterface(iid, ppv);
                    return result;
           }
           else
           {
                    Return CLASS_E_CLASSNOTAVAILABLE;
           }
}
参数说明：
clsid: 待创建对象的CLSID标识符;
iid: 指定接口IID;
ppv: 存放类厂对象的接口指针;
 
3、CoCreateInstance
函数定义：
HRESULT CoCrateInstance(const CLSID& clsid, IUnknown* pUnknownOuter, DWORD dwClsContext, const IID& iid, (void **)ppv)
{
           IClassFactory*   pCF;
           HRESULT          hr;
           hr = CoGetClassObject(clsid, dwClsContext, NULL, IID_IClassFactory, (void *)pCF);
           if(FAILED(hr))   return hr;
           hr = pCF->CreateInstance(pUnkOuter, iid, (void **)ppv);
           pCF->Release();
           return hr;
}
参数说明：
clsid: 待创建对象的CLSID标识符;
pUnknownOuter: 用于对象被聚合的情形;
dwClsContext: 指定组件类别(进程内、进程外组件等等);
iid: 指定接口IID;
ppv: 存放类厂对象的接口指针;
 
4、   CoCreateInstanceEx
函数定义：
HRESULT CoCreateInstaceEx(const CLSID& clsid, IUnknown* pUnknownOuter, DWORD dwClsContext, COSEVERINFO* pServerInfo, DWord dwCount, MULTI_QI* rgMultiQI);
参数说明：
clsid: 待创建对象的CLSID标识符;
pUnknownOuter: 用于对象被聚合的情形;
dwClsContext: 指定组件类别(进程内、进程外组件等等);
pServerInfo: 指定服务器信息;
dwCount: 与rgMultiQI指定一个结构数组，用于保存多个对象接口指针;
rgMultiQI: 同上;
 
5、   CreateInstance(类厂函数)
函数定义：
HRESULT CrateInstance(IUnknown *pUnknownOuter, const IID& iid, void **ppv)
{
CDictComp        *pObj;  
HRESULT hr = E_OUTOFMEMORY;
          *ppv = NULL;
if(pUnknownOuter != NULL)
{
return CLASS_E_NOAGGREGATION;
}
 
pObj = new CDictComp();
if(pObj == NULL)
{
return hr;  
}
 
hr = pObj->QueryInterface(iid, ppv);
if(hr != S_OK)
{
g_DictionaryNumber--;
delete pObj;
}
return hr; 
}
参数说明：
clsid: 待创建对象的CLSID标识符;
pUnknownOuter: 用于对象被聚合的情形;
dwClsContext: 指定组件类别(进程内、进程外组件等等);
iid: 指定接口IID;
ppv: 存放类厂对象的接口指针;
 
6、   IClassFactory
接口原型：
Class IClassFactory : public IUnknown
{
           Virtual HRESULT __stdcall CreateInstance(IUnknown* pUnknownOuter, const IID& iid, (void **)ppv) = 0;
           Virtual HRESULT __stdcall LockServer(BOOL bLock) = 0;
};
参数说明：
pUnknownOuter: 用于对象被聚合的情形;
iid: 指定接口IID;
ppv: 存放类厂对象的接口指针;
bLock: 加锁或解锁;
 
7、   创建对象三种方法
(1)、CoCreateInstance (最常用的方法)
(2)、CoGetClassObject (希望获取类厂对象(的函数))
(3)、CoCreateInstanceEx (创建远程对象或一次获取对象多个接口指针)

 */