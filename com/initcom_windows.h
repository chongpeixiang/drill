#include <stdio.h>
#include "_cgo_export.h"

typedef struct _GUID {
    unsigned long  Data1;
    unsigned short Data2;
    unsigned short Data3;
    unsigned char  Data4[8];
} GUID;

typedef struct _state
{
	void* byte;
	void* ptr;
	int   len;
} CState;

typedef GUID*	CLASSID;
typedef GUID*	POID;

typedef GoUint  HRESULT;
typedef GoUint  UINT;
typedef GoUint64 ULONG;

#define STDCALL   __stdcall
#define DEFINE_GUID(name, l, w1, w2, b1, b2, b3, b4, b5, b6, b7, b8) \
        GUID  name \
                = { l, w1, w2, { b1, b2,  b3,  b4,  b5,  b6,  b7,  b8 } }

#define CLASS(iface)    typedef struct iface {\
                                    const struct iface##Vtbl * lpVtbl; \
                                } iface; \
                                typedef const struct iface##Vtbl iface##Vtbl; \
                                const struct iface##Vtbl

#define THIS                    void* This

#define IsEqualGUID(rguid1, rguid2) (!memcmp(rguid1, rguid2, sizeof(GUID)))

#define STDMETHOD(method)       HRESULT (STDCALL * method)
#define STDMETHOD_(type,method) type (STDCALL * method)

#define IUnknown	STDMETHOD (QueryInterface)   (THIS, CLASSID, void **);\
	STDMETHOD_(ULONG, AddRef)    (THIS);\
	STDMETHOD_(ULONG, Release)   (THIS);\

CLASS(IClassFactory)
{
	IUnknown
	STDMETHOD (CreateInstance)   (THIS, POID, void **);
	STDMETHOD (LockServer)   (THIS,int);
	
};

CLASS(IDispatch)
{	
	IUnknown
	STDMETHOD (GetTypeInfoCount)   (THIS,UINT *pctinfo);
	STDMETHOD (GetTypeInfo)   (THIS,UINT iTInfo,LCID lcid,void** pp);
	STDMETHOD (GetIDsOfNames)   (THIS,UINT iTInfo,LCID lcid,void** pp);
	STDMETHOD (Invoke)   (THIS,UINT iTInfo,LCID lcid,void** pp);	
};
