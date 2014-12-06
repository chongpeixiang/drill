package web

import(
    "os"
    "io/ioutil"
    //"log"
)

func include(L *lua.State)int{
    x := L.ToLString(1)
    file,err := os.Open(x)
    if err != nil {
        panic(err)
    }

    defer file.Close() 
    
    fd,err := ioutil.ReadAll(file)  
    
    L.PushString(string(fd)) 

    return 1
}

func parseHtml(L *lua.State)int{
    str := L.ToLString(1)
    return 1
}

func openTpl(L *lua.State) int{
    r := map[string]interface{}{"include":include,"test2":test2,"test3":test3}
    L.Newlib(r)
    return 1
}