#include "initcom_windows.h"

int isEqGuid(GoUintptr p0,GoUintptr p1);

GoUintptr comGetGoUinptr(CState* L);
char* comGetString(CState* L,int len);
int comInt(CState* L);
void comfree(CState* L);

