package models

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
)

var (
	Bm cache.Cache
)

//用户表
type Customer struct {
	Id           int       `orm:"index"`
	Loginname    string    `orm:"index"`
	Loginnamemd5 string    `orm:"null;index"`
	Created      time.Time `orm:"null;auto_now_add;type(datetime)"`
	Passwd       string    `orm:"null;index"`
	Authcode     string    `orm:"null;index"`
	Enabled      int       `orm:"index;default(1)"`
	Company      string    `orm:"null"`
	Email        string    `orm:"null"`
	Tel          string    `orm:"null"`
	Mac          string    `orm:"null;index"`
}

//股票代码表
type Stockcodes struct {
	Id      int    `orm:"index"`
	Market  string `orm:"null;index"`
	Code    string `orm:"null;index"`
	Caption string `orm:"null"`
	Isindex int    `orm:"default(0)"`
	Isfund  int    `orm:"default(0)"`
}

//交易数据表
type Datatable struct {
	Id      int `orm:"index"`
	Stockid int `orm:"default(0);index"`
	Tyear   int `orm:"default(0);index"`
	Tmonth  int `orm:"default(0);index"`
	Tday    int `orm:"default(0);index"`
	Snap    int `orm:"default(1)"`
	Dw      int `orm:"default(1)"`
	Trade   int `orm:"default(1)"`
	Idx     int `orm:"default(0)"`
}

//交易日历表
type Timetable struct {
	Id      int `orm:"index"`
	Dateint int `orm:"index"`
	Tyear   int `orm:"index"`
	Tmonth  int `orm:"index"`
	Tday    int `orm:"index"`
}

//客户端运行状态表
type Clientstatus struct {
	Id         int       `orm:"index"`
	Clientid   int       `orm:"index;default(0)"`
	Ip         string    `orm:"null"`
	Reporttime time.Time `orm:"null;auto_now_add;type(datetime)"`
	Action     string    `orm:"null;type(text)"`
	Tag        int       `orm:"index;default(0)"`
}

//股票订阅信息表
type Clientsubs struct {
	Id       int    `orm:"index"`
	Clientid int    `orm:"index;default(0)"`
	Code     string `orm:"null;index"`
	Market   string `orm:"null;index"`
	Isindex  int    `orm:"default(0)"`
	Isfund   int    `orm:"default(0)"`
}

//作者表
type Authors struct {
	Id        int    `orm:"index"`
	Aname     string `orm:"index"`     //作者姓名
	AHomepage string `orm:"size(255)"` //作者主页地址
	AEmail    string `orm:"size(255)"` //邮箱
	Tel       string `orm:"size(255)"` //电话
}

//仿真环境分类表
type Planscategory struct {
	Id      int    `orm:"index"`
	Caption string `orm:"index"` //类别名称
}

//仿真环境表
type Plans struct {
	Id              int       `orm:"index"`
	Categoryid      int       `orm:"default(0);index"` //分类Id，参照Planscategory表
	Enabled         int       `orm:"index;default(1)"` //是否启用
	Createtime      time.Time `orm:"null;auto_now_add;type(datetime)"`
	Caption         string    `orm:"index"`              //标题名称
	Description     string    `orm:"null;size(4096)"`    //描述信息
	Authorid        int       `orm:"default(0);index"`   //作者Id，参照Authors表
	Cash            float32   `orm:"default(0)"`         //初始现金
	Positioncount   int       `orm:"default(0)"`         //持仓股票数
	Positions       string    `orm:"null;sizeof(10240)"` //持仓股票列表，json字符串
	Stockcount      int       `orm:"default(0)"`         //观察和交易股票个数
	Stocks          string    `orm:"null;sizeof(20480)"` //观察和交易股票列表，json字符串
	Tdays           int       `orm:"default(1)"`         //T+N交割天数
	Buyfeerate      float32   `orm:"default(0)"`         //买入费率
	Buyfeefix       float32   `òrm:"default(0)"`         //买入费每笔加收金额
	Buytaxrate      float32   `orm:"default(0)"`         //买入税率
	Sellfeerate     float32   `orm:"default(0)"`         //卖出费率
	Sellfeefix      float32   `orm:"default(0)"`         //卖出费每笔加收金额
	Selltaxrate     float32   `orm:"default0()"`         //卖出税率
	Fromdate        int       `orm:"default(0)"`         //起始日期
	Todate          int       `orm:"default(0)"`         //结束日期
	Ignoredw        int       `orm:"default(1)"`         //是否忽略盘口流动性，0=是 1=否
	Cancelrate      float32   `orm:"default(1)"`         //最大容许撤报比=撤单数/报单数，1表示允许全部撤销，0表示不允许撤单
	Orderdelay      int       `orm:"default(0)"`         //报单延迟周期数，0=不延迟，每周期按每TICK3秒计算
	Orderspersecond int       `orm:"default(0)"`         //每秒允许下单数，0表示不限制
	Dealrate        float32   `orm:"default(1)"`         //每次最多成交比例，1=全成
	Supportbacktest int       `orm:"default(1)"`         //是否支持回测
	Supportresearch int       `orm:"default(1)"`         //是否支持策略研究
	Supporttraining int       `orm:"default(1)"`         //是否支持仿真训练
	Sortindex       int       `orm:"default(0);index"`   //显示排序
}

//指数TICK文件下载地址表
type Indextick_Download_Url struct {
	Code    string
	Market  string
	Tdate   string
	Url     string
	Isindex int
}

//股票TICK文件下载地址表
type Stocktick_Download_Url struct {
	Code     string
	Market   string
	Tdate    string
	Snapurl  string
	Dwurl    string
	Tradeurl string
}

//日线下载地址表
type Dayline_Download_Url struct {
	Code    string
	Market  string
	Dayurl  string
	Weekurl string
	Isindex int
}

//股票或指数分钟线和日线下载地址表
type Minuteline_Download_Url struct {
	Code     string
	Market   string
	Tdate    string
	Min1url  string
	Min5url  string
	Min10url string
	Min30url string
	Min60url string
	Isindex  int
}

//返回信息结构
type ResponseMsg struct {
	Id  int
	Msg string //存储JSON数据
}

//缓存中保存的用户信息结构
type CachedUserinfo struct {
	Id           int
	Loginname    string
	Loginnamemd5 string
	Authcode     string
	Mac          string
}

func RegisterDB() {
	//注册 model
	orm.RegisterModel(
		new(Customer),
		new(Stockcodes),
		new(Timetable),
		new(Datatable),
		new(Clientstatus),
		new(Clientsubs),
		new(Plans),
		new(Planscategory),
		new(Authors))
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:phpcj@/octopus?charset=utf8")
}

func GenMd5(str string) string {
	//计算字符串MD5值
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}

func LogClientStatus(clientId int, ipaddress string, msg string) {
	//记录客户端状态信息
	o := orm.NewOrm()
	r := new(Clientstatus)
	r.Clientid = clientId
	r.Ip = ipaddress
	r.Action = msg
	o.Insert(r)
}

func QueryCustomerId(loginname, passwd string) (int, string, string, string) {
	//通过登录名和密码查询用户ID和授权码
	o := orm.NewOrm()
	var cus Customer
	err := o.QueryTable("customer").Filter("loginnamemd5", loginname).Filter("passwd", passwd).Filter("enabled", 1).One(&cus)
	if (err != orm.ErrNoRows) && (err != orm.ErrMultiRows) {
		return cus.Id, cus.Loginname, cus.Authcode, cus.Mac
	} else {
		return -1, "", "", ""
	}
}

func QueryAuthcode(authcode string) int {
	//通过授权码查找用户ID
	o := orm.NewOrm()
	var cus Customer
	err := o.QueryTable("customer").Filter("authcode", authcode).Filter("enabled", 1).One(&cus)
	if (err != orm.ErrNoRows) && (err != orm.ErrMultiRows) {
		return cus.Id
	} else {
		return -1
	}
}

func UpdateClientMac(aid int, amac string) {
	//更新用户数据库的MAC记录
	o := orm.NewOrm()
	//var m Customer
	m := Customer{Id: aid}
	o.Read(&m)
	//m.Id = aid
	m.Mac = amac
	o.Update(&m, "Mac")
}

func UpdateClientSubs(aid int, avalue []Clientsubs) {
	//更新用户订阅数据库，先删除掉对应的ClientID的全部记录
	o := orm.NewOrm()
	var m Clientsubs
	m.Clientid = aid
	o.Delete(&m, "clientid")
	//插入avalue数据到数据库
	for i := 0; i < len(avalue); i++ {
		m = avalue[i]
		o.Insert(&m)
	}
}

func QueryClientSubs(cid int) []Clientsubs {
	//查询用户订阅的股票记录
	o := orm.NewOrm()
	var m []Clientsubs
	o.QueryTable("clientsubs").Filter("clientid", cid).All(&m)
	return m
}

func QueryCodetable() []Stockcodes {
	//查询股票代码表
	o := orm.NewOrm()
	var codes []Stockcodes
	o.QueryTable("stockcodes").All(&codes)
	return codes
}

func QueryTimetable() []Timetable {
	//查询交易日历表
	o := orm.NewOrm()
	var tb []Timetable
	o.QueryTable("timetable").All(&tb)
	return tb
}

func GenIndexTickUrls(Adate string, Asubs []Clientsubs) []Indextick_Download_Url {
	//根据传入的Asubs记录，生成指数下载地址
	var m []Indextick_Download_Url
	var msub Clientsubs
	for i := 0; i < len(Asubs); i++ {
		msub = Asubs[i]
		if msub.Isindex == 1 {
			var murl Indextick_Download_Url
			murl.Code = msub.Code
			murl.Market = msub.Market
			murl.Tdate = Adate
			murl.Url = beego.AppConfig.String("quotoserver") + Adate + "/" + msub.Market + "/" + msub.Code + ".index" //生成下载地址
			m = append(m, murl)
		}
	}
	return m
}

func GenStockTickUrls(Adate string, Asubs []Clientsubs) []Stocktick_Download_Url {
	//根据传入Asubs参数，生成股票Tick下载地址
	var m []Stocktick_Download_Url
	var msub Clientsubs
	for i := 0; i < len(Asubs); i++ {
		msub = Asubs[i]
		if msub.Isindex == 0 {
			var murl Stocktick_Download_Url
			murl.Code = msub.Code
			murl.Market = msub.Market
			murl.Tdate = Adate
			murl.Snapurl = beego.AppConfig.String("quotoserver") + Adate + "/" + msub.Market + "/" + msub.Code + ".snap"   //生成下载地址
			murl.Dwurl = beego.AppConfig.String("quotoserver") + Adate + "/" + msub.Market + "/" + msub.Code + ".dw"       //生成下载地址
			murl.Tradeurl = beego.AppConfig.String("quotoserver") + Adate + "/" + msub.Market + "/" + msub.Code + ".trade" //生成下载地址
			m = append(m, murl)
		}
	}
	return m
}

func GenMinutelineUrls(Adate string, Asubs []Clientsubs) []Minuteline_Download_Url {
	//根据传入的Asubs记录，生成指数下载地址
	var m []Minuteline_Download_Url
	var msub Clientsubs
	for i := 0; i < len(Asubs); i++ {
		msub = Asubs[i]
		var murl Minuteline_Download_Url
		murl.Code = msub.Code
		murl.Market = msub.Market
		murl.Tdate = Adate
		murl.Isindex = msub.Isindex
		murl.Min1url = beego.AppConfig.String("quotoserver") + Adate + "/" + msub.Market + "/" + msub.Code + ".1min"   //生成下载地址
		murl.Min5url = beego.AppConfig.String("quotoserver") + Adate + "/" + msub.Market + "/" + msub.Code + ".5min"   //生成下载地址
		murl.Min10url = beego.AppConfig.String("quotoserver") + Adate + "/" + msub.Market + "/" + msub.Code + ".10min" //生成下载地址
		murl.Min30url = beego.AppConfig.String("quotoserver") + Adate + "/" + msub.Market + "/" + msub.Code + ".30min" //生成下载地址
		murl.Min60url = beego.AppConfig.String("quotoserver") + Adate + "/" + msub.Market + "/" + msub.Code + ".60min" //生成下载地址
		m = append(m, murl)
	}
	return m
}

func GenDaylineUrls(Asubs []Clientsubs) []Dayline_Download_Url {
	//根据传入Asub参数，生成日线文件下载地址
	var m []Dayline_Download_Url
	var msub Clientsubs
	for i := 0; i < len(Asubs); i++ {
		msub = Asubs[i]
		var murl Dayline_Download_Url
		murl.Code = msub.Code
		murl.Market = msub.Market
		murl.Isindex = msub.Isindex
		if murl.Isindex == 1 {
			murl.Dayurl = beego.AppConfig.String("quotoserver") + "day/" + msub.Market + "/" + msub.Code + ".iday"   //生成下载地址
			murl.Weekurl = beego.AppConfig.String("quotoserver") + "day/" + msub.Market + "/" + msub.Code + ".iweek" //生成下载地址
		} else {
			murl.Dayurl = beego.AppConfig.String("quotoserver") + "day/" + msub.Market + "/" + msub.Code + ".sday"   //生成下载地址
			murl.Weekurl = beego.AppConfig.String("quotoserver") + "day/" + msub.Market + "/" + msub.Code + ".sweek" //生成下载地址
		}
		m = append(m, murl)
	}
	return m
}

func GenWeeklineUrls(Asubs []Clientsubs) []Indextick_Download_Url {
	//根据传入Asub参数，生成周线文件下载地址
	var m []Indextick_Download_Url
	var msub Clientsubs
	for i := 0; i < len(Asubs); i++ {
		msub = Asubs[i]
		var murl Indextick_Download_Url
		murl.Code = msub.Code
		murl.Market = msub.Market
		murl.Tdate = ""
		murl.Url = beego.AppConfig.String("quotoserver") + msub.Market + "/week/" + msub.Code + ".week" //生成下载地址
		m = append(m, murl)
	}
	return m
}

func GenMonthlineUrls(Asubs []Clientsubs) []Indextick_Download_Url {
	//根据传入Asub参数，生成月线文件下载地址
	var m []Indextick_Download_Url
	var msub Clientsubs
	for i := 0; i < len(Asubs); i++ {
		msub = Asubs[i]
		var murl Indextick_Download_Url
		murl.Code = msub.Code
		murl.Market = msub.Market
		murl.Tdate = ""
		murl.Url = beego.AppConfig.String("quotoserver") + msub.Market + "/month/" + msub.Code + ".month" //生成下载地址
		m = append(m, murl)
	}
	return m
}

func QueryPlans() []Plans {
	//查询Plans仿真环境表
	var m []Plans
	o := orm.NewOrm()
	o.QueryTable("plans").Filter("enabled", 1).OrderBy("-sortindex", "id").All(&m)
	return m
}

func QueryClientStatus(cid int) []Clientstatus {
	//根据指定的cid返回操作日志
	o := orm.NewOrm()
	var m []Clientstatus
	o.QueryTable("clientstatus").Filter("clientid", cid).OrderBy("-id").All(&m)
	return m
}

func InsertClientStatus(cid int, tag int, ipaddress string, msg string) {
	//保存客户端上传状态
	o := orm.NewOrm()
	r := new(Clientstatus)
	r.Clientid = cid
	r.Ip = ipaddress
	r.Action = msg
	r.Tag = tag
	o.Insert(r)
}

func InitCache() {
	//初始化缓存对象
	Bm, _ = cache.NewCache("memory", `{"interval":60}`)
}
