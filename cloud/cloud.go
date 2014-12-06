package cloud

import(
    "github.com/astaxie/beego/httplib"
    "strings"
    "crypto/md5"
    "crypto/tls"
    "fmt"
    "encoding/base64"
    "encoding/json"
    "time"
    "strconv"
)

type CloudSms struct {
    AccountSid string
    AccountToken string
    AppId        string
    SubAccountSid   string
    SubAccountToken string
    VoIPAccount     string
	VoIPPassword    string
	ServerIP        string
	ServerPort      string
	SoftVersion     string
	Batch           string
	BodyType        string `json`
	EnabeLog        bool
	Filename        string
}

type Respon struct {
    Ret int
    Msg string
}
type respon struct {
    statusCode int
    statusMsg  string
    templateSMS map[string]string
}
const (  
    base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"  
)

var coder = base64.NewEncoding(base64Table)


func NewCloud()*CloudSms{

    return &CloudSms{BodyType:"json"}
}


 /**
    * 设置主帐号
    * 
    * @param AccountSid 主帐号
    * @param AccountToken 主帐号Token
    */    

func (self *CloudSms )SetAccount(AccountSid,AccountToken string){
      self.AccountSid = AccountSid
      self.AccountToken = AccountToken  
}
    
   /**
    * 设置子帐号
    * 
    * @param SubAccountSid 子帐号
    * @param SubAccountToken 子帐号Token
    * @param VoIPAccount VoIP帐号
    * @param VoIPPassword VoIP密码
    */    
func (self *CloudSms )SetSubAccount( SubAccountSid, SubAccountToken, VoIPAccount, VoIPPassword string){
       self.SubAccountSid =  SubAccountSid;
       self.SubAccountToken =  SubAccountToken;
       self.VoIPAccount =  VoIPAccount;
       self.VoIPPassword =  VoIPPassword;     
}
    
    /**
    * 设置应用ID
    * 
    * @param AppId 应用ID
    */
func (self *CloudSms )SetAppId( AppId string){
        self.AppId =  AppId 
}
/**
* 发送模板短信
* @param to 短信接收彿手机号码集合,用英文逗号分开
* @param datas 内容数据
* @param  tempId 模板Id
*/       
func (self *CloudSms )SendTemplateSMS( to, datas, tempId string) Respon{
        //主帐号鉴权信息验证，对必选参数进行判空。
        var body string
         auth,b := self.accAuth()
        if b == false {
            return auth
        }
        // 拼接请求包体
        if( self.BodyType=="json"){
            body = "{templateSMS:{'to':'"+to+"','templateId':'"+tempId+"','appId':'"+self.AppId+"','param':'"+datas+"'}}";
           
        }
        fmt.Println(body)
        self.Batch = time.Now().Format("20060102150405")
        //self.Batch = fmt.Sprint(time.Now().Unix())
        // 大写的sig参数 
         sig :=  strings.ToUpper(Md5([]byte(self.AccountSid+self.AccountToken+fmt.Sprint(self.Batch))));
        // 生成请求URL        
         url :="https://"+self.ServerIP+"/"+self.SoftVersion+"/Accounts/"+self.AccountSid+"/Messages/templateSMS?sig="+sig
         
         req:=httplib.Post(url)
         req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
         // 生成包头  
         req.Header("Accept","application/"+self.BodyType)
         req.Header("Content-Type","application/"+self.BodyType+";charset=utf-8")
         
         // 生成授权：主帐户Id + 英文冒号 + 时间戳。
         authen := coder.EncodeToString([]byte(self.AccountSid+":"+fmt.Sprint(self.Batch)));
         req.Header("Authorization",authen)
         req.Body(body)
        // 发送请求
         var re map[string]interface{}
         result,_ :=  req.String();
         fmt.Println(result)
         //result := "{\"resp\":{\"respCode\":\"000000\",\"templateSMS\":{\"createDate\":\"20141205115214\",\"smsId\":\"3891fcd7973db2636e372ae65c159fa6\"}}}"
         json.Unmarshal([]byte(result), &re)

         resp := re["resp"].(map[string]interface{})
  
         if resp["respCode"] == "000000" {
            auth.Ret = 0
            //fmt.Println(re["templateSMS"].(type))
            auth.Msg = "发送成功！"
            //fmt.Sprint(re["templateSMS"].(map[string]interface{})["smsMessageSid"])
         } else {
            auth.Ret,_ = strconv.Atoi(resp["respCode"].(string))
            auth.Msg = "号码不合法"
         }
         //fmt.Println(re)
        return auth
}

func (self *CloudSms )Call( to, code string) Respon{
        //主帐号鉴权信息验证，对必选参数进行判空。
        var body string
         auth,b := self.accAuth()
        if b == false {
            return auth
        }
        // 拼接请求包体
        if( self.BodyType=="json"){
            body = "{voiceCode:{'to':'"+to+"','appId':'"+self.AppId+"','verifyCode':'"+code+"'}}";
        }
        
        self.Batch = time.Now().Format("20060102150405")
        //self.Batch = fmt.Sprint(time.Now().Unix())
        // 大写的sig参数 
         sig :=  strings.ToUpper(Md5([]byte(self.AccountSid+self.AccountToken+fmt.Sprint(self.Batch))));
        // 生成请求URL        
         url :="https://"+self.ServerIP+"/"+self.SoftVersion+"/Accounts/"+self.AccountSid+"/Calls/voiceCode?sig="+sig
         
         req:=httplib.Post(url)
         req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
         // 生成包头  
         req.Header("Accept","application/"+self.BodyType)
         req.Header("Content-Type","application/"+self.BodyType+";charset=utf-8")
         
         // 生成授权：主帐户Id + 英文冒号 + 时间戳。
         authen := coder.EncodeToString([]byte(self.AccountSid+":"+fmt.Sprint(self.Batch)));
         req.Header("Authorization",authen)
         req.Body(body)
        // 发送请求
         var re map[string]interface{}
         result,_ :=  req.String();
         fmt.Println(result)
         //result := "{\"resp\":{\"respCode\":\"000000\",\"templateSMS\":{\"createDate\":\"20141205115214\",\"smsId\":\"3891fcd7973db2636e372ae65c159fa6\"}}}"
         json.Unmarshal([]byte(result), &re)

         resp := re["resp"].(map[string]interface{})
  
         if resp["respCode"] == "000000" {
            auth.Ret = 0
            //fmt.Println(re["templateSMS"].(type))
            auth.Msg = "发送成功！"
            //fmt.Sprint(re["templateSMS"].(map[string]interface{})["smsMessageSid"])
         } else {
            auth.Ret,_ = strconv.Atoi(resp["respCode"].(string))
            auth.Msg = "号码不合法"
         }
         //fmt.Println(re)
        return auth
}

func (self *CloudSms )Send( to, code string) Respon{

        req:=httplib.Post("http://119.37.197.110/sms.action")
        req.Param("u","zlqh")
        req.Param("p",Md5([]byte("123456")))
        req.Param("m",to)
        req.Param("c",code)
        req.Param("s","")
        req.Param("g","")
        result,_ :=  req.String()
        if result[0:1] != "0" {
            return Respon{Ret:1,Msg:"发送失败"}
        }
        return Respon{Ret:0,Msg:"发送成功！"}
}


 func (self *CloudSms )accAuth() (t Respon,b bool) {
       if self.ServerIP=="" {
            t.Ret = 172004
            t.Msg = "IP为空"
            b= false
            return 
        }
       if self.ServerPort=="" {
            t.Ret = 172005
            t.Msg = "端口错误（小于等于0）"
            b= false
            return 
        }
       if self.SoftVersion=="" {
            t.Ret = 172013
            t.Msg = "版本号为空"
            b= false
            return 
        }
       if self.AccountSid=="" {
            t.Ret = 172006
            t.Msg = "主帐号为空"
            b= false
            return 
        }
       if self.AccountToken=="" {
            t.Ret = 172007
            t.Msg = "主帐号令牌为空"
            b= false
            return 
        }
       if self.AppId=="" {
            t.Ret = 172012
            t.Msg = "应用ID为空"
            b= false
            return 
        }
        b = true
        return 

   }
  
  func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
  }
 