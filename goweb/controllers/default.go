package controllers

import (
	"QuoteServer/models"
	"encoding/json"
	//	"math"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}
type LoginController struct {
	beego.Controller
}
type SubController struct {
	beego.Controller
}
type CodeTableController struct {
	beego.Controller
}
type TimeTableController struct {
	beego.Controller
}
type IndexTickController struct {
	beego.Controller
}
type StockTickController struct {
	beego.Controller
}
type MinutelineController struct {
	beego.Controller
}
type DaylineController struct {
	beego.Controller
}
type WeeklineController struct {
	beego.Controller
}
type MonthlineController struct {
	beego.Controller
}
type ClientStatusController struct {
	beego.Controller
}
type ClientUpstatusController struct {
	beego.Controller
}
type PlansController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "localhost"
	c.Data["Email"] = ""
	c.TplName = "index.tpl"
}

func (c *LoginController) Get() {
	/*
	  处理用户登录，用户登录时传输的参数是用户名、密码字符串、登陆时间字符串（混淆用），用逗号间隔，都进行MD5加密
	  程序分割字符串后查找数据库，如果登陆合法，则从数据库中读userinfo
	  1、检查缓存是否存在，存在则从缓存中读取，如果不存在则从数据库中读取，并保存到缓存。读取用户信息，包括ID、Authcode等
	  2、记录登录信息到客户端状态数据库
	  3、将读取到的用户信息（loginName,AuthCode）返回给客户端
	*/
	//解析传递来的参数字符串
	var r models.ResponseMsg
	mtmp := c.GetString("a")
	mtmp2 := strings.Split(mtmp, ",")
	if len(mtmp2) == 3 {
		//获取登录名和密码的MD5字符串及MAC地址
		mlgoinnamemd5 := mtmp2[0]
		mpasswd := mtmp2[1]
		mmac := mtmp2[2]
		//检查缓存中是否存在对应的信息，缓存KEY是LoginNameMD5+mpasswd
		if models.Bm.IsExist(mlgoinnamemd5 + mpasswd + mmac) {
			//存在缓存，从缓存中读取数据
			mr := models.Bm.Get(mlgoinnamemd5 + mpasswd + mmac).(models.CachedUserinfo)
			//返回登陆成功数据
			r.Id = 200
			j, _ := json.Marshal(mr)
			r.Msg = string(j)
			c.Data["json"] = &r
			c.ServeJSON()
			//写入日志
			models.LogClientStatus(mr.Id, c.Ctx.Input.IP(), "登陆成功")
		} else {
			//从数据库中查询数据
			cid, cloginame, cauthcode, cmac := models.QueryCustomerId(mlgoinnamemd5, mpasswd)
			if cid != -1 {
				//找到记录，先判断MAC地址是否为空，如果为空表示从未绑定计算机，则更新数据库中的MAC记录，
				//如果不为空，则判断数据库中的MAC地址和提交的MAC地址是否一致
				var mloginstatus int
				mloginstatus = 0
				if cmac == "" {
					//MAC为空，更新数据库MAC记录允许登陆
					mloginstatus = 1
					//更新数据库记录
					models.UpdateClientMac(cid, mmac)
				} else {
					//MAC不为空，判断是否与提交的MAC一致
					if cmac != mmac {
						//数据库与提交的MAC不一致，返回锁定信息
						mloginstatus = 2
					} else {
						//数据库与提交的MAC一致，允许登陆
						mloginstatus = 1
					}
				}
				if mloginstatus == 1 {
					//登陆成功写入缓存
					var mr models.CachedUserinfo
					mr.Id = cid
					mr.Loginname = cloginame
					mr.Loginnamemd5 = mlgoinnamemd5
					mr.Authcode = cauthcode
					mr.Mac = cmac
					models.Bm.Put(mlgoinnamemd5+mpasswd+mmac, mr, 3600*time.Second)
					r.Id = 200
					j, _ := json.Marshal(mr)
					r.Msg = string(j)
					c.Data["json"] = &r
					c.ServeJSON()
					//写入日志
					models.LogClientStatus(mr.Id, c.Ctx.Input.IP(), "登陆成功")
				} else if mloginstatus == 2 {
					//锁定计算机
					var mr models.CachedUserinfo
					mr.Id = cid
					mr.Loginname = cloginame
					mr.Loginnamemd5 = mlgoinnamemd5
					mr.Authcode = "locked"
					mr.Mac = cmac
					r.Id = 200
					j, _ := json.Marshal(mr)
					r.Msg = string(j)
					c.Data["json"] = &r
					c.ServeJSON()
					//写入日志
					models.LogClientStatus(mr.Id, c.Ctx.Input.IP(), "登陆失败(MAC不符)")
				}
			} else {
				//没找到记录，登陆失败
				r.Id = 404
				r.Msg = ""
				c.Data["json"] = &r
				c.ServeJSON()
			}
		}
	}
}

func (c *SubController) Post() { //此处实际应用时需改成POST方法
	/*处理用户订阅股票，用户提交形式为字符串，形式为
	授权码|股票代码|所属市场（sh或sz）|是否指数（1或0）|是否基金（1或0）|
	处理步骤：
	1、先校验授权码
	2、如果授权码在缓存或数据库中存在，则解析订阅名单，清除缓存中的订阅记录，重新写入数据库和缓存
	3、否则返回未登陆信息
	*/
	var r models.ResponseMsg
	mtmp := c.GetString("a")
	mtmp2 := strings.Split(mtmp, "|")
	var mvalid bool
	mvalid = false
	var cid int
	cid = -1
	var mauth string
	//取授权码
	if (len(mtmp2)) > 1 {
		mauth = mtmp2[0]
		//校验授权码
		if models.Bm.IsExist(mauth) {
			//缓存中存在，校验成功，读出CID
			cid = models.Bm.Get(mauth).(int)
			mvalid = true
		} else {
			//缓存中不存在，查询数据库
			cid = models.QueryAuthcode(mauth)
			if cid != -1 {
				//数据库中存在，写入缓存
				mvalid = true
				models.Bm.Put(mauth, cid, 3600*time.Second)
			}
		}
	}
	if mvalid == true {
		//校验码通过验证，循环数组解析股票订阅信息
		is4 := (len(mtmp2) - 1) % 4
		if is4 == 0 {
			//数据长度是对的，开始循环
			mcount := (len(mtmp2) - 1) / 4
			var msubs []models.Clientsubs
			var msub models.Clientsubs
			for i := 1; i <= mcount; i++ {
				//向数组中添加纪录
				msub.Clientid = cid
				msub.Code = mtmp2[(i-1)*4+1]
				msub.Market = mtmp2[(i-1)*4+2]
				ii, _ := strconv.Atoi(mtmp2[(i-1)*4+3])
				msub.Isindex = ii
				ij, _ := strconv.Atoi(mtmp2[(i-1)*4+4])
				msub.Isfund = ij
				msubs = append(msubs, msub)
			}
			//添加完毕，将数据保存到数据库
			models.UpdateClientSubs(cid, msubs)
			//写入缓存
			if models.Bm.IsExist(mauth + "subs") {
				models.Bm.Delete(mauth + "subs")
			}
			models.Bm.Put(mauth+"subs", msubs, 3600*time.Second)
			//返回成功消息
			r.Id = 200
			r.Msg = ""
			c.Data["json"] = &r
			c.ServeJSON()
			//写入日志
			models.LogClientStatus(cid, c.Ctx.Input.IP(), "订阅股票成功")
		} else {
			//数据长度不对，打回去登陆失败
			r.Id = 404
			r.Msg = ""
			c.Data["json"] = &r
			c.ServeJSON()
		}

	} else {
		//校验码没有通过验证，返回未登陆信息
		r.Id = 404
		r.Msg = ""
		c.Data["json"] = &r
		c.ServeJSON()
	}
}

func (c *CodeTableController) Get() { //此处实际应用时需改成POST方法
	/*处理下载代码表请求，请求参数只有授权码*/
	var r models.ResponseMsg
	mtmp := c.GetString("a")
	var mvalid bool
	mvalid = false
	var cid int
	cid = -1
	var mauth string
	mauth = mtmp
	if models.Bm.IsExist(mauth) {
		//缓存中存在，校验成功，读出CID
		cid = models.Bm.Get(mauth).(int)
		mvalid = true
	} else {
		//缓存中不存在，查询数据库
		cid = models.QueryAuthcode(mauth)
		if cid != -1 {
			//数据库中存在，写入缓存
			mvalid = true
			models.Bm.Put(mauth, cid, 3600*time.Second)
		}
	}
	if mvalid == true {
		//校验通过，查询全部股票代码记录，转换成JSON字符串
		mr := models.QueryCodetable()
		j, _ := json.Marshal(mr)
		r.Id = 200
		r.Msg = string(j)
		c.Data["json"] = &r
		c.ServeJSON()
		//写入日志
		models.LogClientStatus(cid, c.Ctx.Input.IP(), "下载代码表")
	} else {
		//校验不通过
		//校验码没有通过验证，返回未登陆信息
		r.Id = 404
		r.Msg = ""
		c.Data["json"] = &r
		c.ServeJSON()
	}
}

func (c *TimeTableController) Get() { //此处实际应用时需改成POST方法
	//处理下载交易日历表请求
	var r models.ResponseMsg
	mtmp := c.GetString("a")
	var mvalid bool
	mvalid = false
	var cid int
	cid = -1
	var mauth string
	mauth = mtmp
	if models.Bm.IsExist(mauth) {
		//缓存中存在，校验成功，读出CID
		cid = models.Bm.Get(mauth).(int)
		mvalid = true
	} else {
		//缓存中不存在，查询数据库
		cid = models.QueryAuthcode(mauth)
		if cid != -1 {
			//数据库中存在，写入缓存
			mvalid = true
			models.Bm.Put(mauth, cid, 3600*time.Second)
		}
	}
	if mvalid == true {
		//校验通过，查询全部交易日历记录，转换成JSON字符串
		mr := models.QueryTimetable()
		j, _ := json.Marshal(mr)
		r.Id = 200
		r.Msg = string(j)
		c.Data["json"] = &r
		c.ServeJSON()
		//写入日志
		models.LogClientStatus(cid, c.Ctx.Input.IP(), "下载交易日历")
	} else {
		//校验不通过
		//校验码没有通过验证，返回未登陆信息
		r.Id = 404
		r.Msg = ""
		c.Data["json"] = &r
		c.ServeJSON()
	}
}

func (c *IndexTickController) Post() {
	/*处理下载指数TICK文件请求
	请求参数为授权码|年月日，例如?a=daddada|20140102，返回用户订阅指数的下载地址列表
	*/
	var r models.ResponseMsg
	mtmp := c.GetString("a")
	mtmp2 := strings.Split(mtmp, "|")

	var mvalid bool
	mvalid = false
	var cid int
	cid = -1
	var mauth string
	var mdate string
	//取授权码
	if (len(mtmp2)) == 2 {
		mauth = mtmp2[0]
		//取日期值
		mdate = mtmp2[1]
		//校验授权码
		if models.Bm.IsExist(mauth) {
			//缓存中存在，校验成功，读出CID
			cid = models.Bm.Get(mauth).(int)
			mvalid = true
		} else {
			//缓存中不存在，查询数据库
			cid = models.QueryAuthcode(mauth)
			if cid != -1 {
				//数据库中存在，写入缓存
				mvalid = true
				models.Bm.Put(mauth, cid, 3600*time.Second)
			}
		}
	}
	if mvalid == true {
		//校验通过，查询订阅指数记录，生成下载地址，转换成JSON字符串
		var msubs []models.Clientsubs
		if models.Bm.IsExist(mauth + "subs") {
			//检查缓存中是否存在订阅记录
			msubs = models.Bm.Get(mauth + "subs").([]models.Clientsubs)
		} else {
			//缓存中不存在，从数据库中检索
			msubs = models.QueryClientSubs(cid)
			//写入缓存
			models.Bm.Put(mauth+"subs", msubs, 3600*time.Second)
		}
		//循环msubs，生成指数TICK文件下载地址
		var mr []models.Indextick_Download_Url
		mr = models.GenIndexTickUrls(mdate, msubs)
		j, _ := json.Marshal(mr)
		r.Id = 200
		r.Msg = string(j)
		c.Data["json"] = &r
		c.ServeJSON()
		//写入日志
		models.LogClientStatus(cid, c.Ctx.Input.IP(), "请求指数TICK")
	} else {
		//校验不通过
		//校验码没有通过验证，返回未登陆信息
		r.Id = 404
		r.Msg = ""
		c.Data["json"] = &r
		c.ServeJSON()
	}
}

func (c *StockTickController) Post() {
	/*处理下载股票TICK文件请求
	请求参数为授权码|年月日，例如?a=daddada|20140102，返回用户订阅股票的下载地址列表
	*/
	var r models.ResponseMsg
	mtmp := c.GetString("a")
	mtmp2 := strings.Split(mtmp, "|")

	var mvalid bool
	mvalid = false
	var cid int
	cid = -1
	var mauth string
	var mdate string
	//取授权码
	if (len(mtmp2)) == 2 {
		mauth = mtmp2[0]
		//取日期值
		mdate = mtmp2[1]
		//校验授权码
		if models.Bm.IsExist(mauth) {
			//缓存中存在，校验成功，读出CID
			cid = models.Bm.Get(mauth).(int)
			mvalid = true
		} else {
			//缓存中不存在，查询数据库
			cid = models.QueryAuthcode(mauth)
			if cid != -1 {
				//数据库中存在，写入缓存
				mvalid = true
				models.Bm.Put(mauth, cid, 3600*time.Second)
			}
		}
	}
	if mvalid == true {
		//校验通过，查询订阅股票记录，生成下载地址，转换成JSON字符串
		var msubs []models.Clientsubs
		if models.Bm.IsExist(mauth + "subs") {
			//检查缓存中是否存在订阅记录
			msubs = models.Bm.Get(mauth + "subs").([]models.Clientsubs)
		} else {
			//缓存中不存在，从数据库中检索
			msubs = models.QueryClientSubs(cid)
			//写入缓存
			models.Bm.Put(mauth+"subs", msubs, 3600*time.Second)
		}
		//循环msubs，生成指数TICK文件下载地址
		var mr []models.Stocktick_Download_Url
		mr = models.GenStockTickUrls(mdate, msubs)
		j, _ := json.Marshal(mr)
		r.Id = 200
		r.Msg = string(j)
		c.Data["json"] = &r
		c.ServeJSON()
		//写入日志
		models.LogClientStatus(cid, c.Ctx.Input.IP(), "请求股票TICK")
	} else {
		//校验不通过
		//校验码没有通过验证，返回未登陆信息
		r.Id = 404
		r.Msg = ""
		c.Data["json"] = &r
		c.ServeJSON()
	}
}

func (c *MinutelineController) Post() {
	/*处理下载分钟线文件请求
	请求参数为授权码|年月日，例如?a=daddada|20140102，返回用户订阅股票的分钟线下载地址列表
	*/
	var r models.ResponseMsg
	mtmp := c.GetString("a")
	mtmp2 := strings.Split(mtmp, "|")

	var mvalid bool
	mvalid = false
	var cid int
	cid = -1
	var mauth string
	var mdate string
	//取授权码
	if (len(mtmp2)) == 2 {
		mauth = mtmp2[0]
		//取日期值
		mdate = mtmp2[1]
		//校验授权码
		if models.Bm.IsExist(mauth) {
			//缓存中存在，校验成功，读出CID
			cid = models.Bm.Get(mauth).(int)
			mvalid = true
		} else {
			//缓存中不存在，查询数据库
			cid = models.QueryAuthcode(mauth)
			if cid != -1 {
				//数据库中存在，写入缓存
				mvalid = true
				models.Bm.Put(mauth, cid, 3600*time.Second)
			}
		}
	}
	if mvalid == true {
		//校验通过，查询订阅指数记录，生成下载地址，转换成JSON字符串
		var msubs []models.Clientsubs
		if models.Bm.IsExist(mauth + "subs") {
			//检查缓存中是否存在订阅记录
			msubs = models.Bm.Get(mauth + "subs").([]models.Clientsubs)
		} else {
			//缓存中不存在，从数据库中检索
			msubs = models.QueryClientSubs(cid)
			//写入缓存
			models.Bm.Put(mauth+"subs", msubs, 3600*time.Second)
		}
		//循环msubs，生成指数TICK文件下载地址
		var mr []models.Minuteline_Download_Url
		mr = models.GenMinutelineUrls(mdate, msubs)
		j, _ := json.Marshal(mr)
		r.Id = 200
		r.Msg = string(j)
		c.Data["json"] = &r
		c.ServeJSON()
		//写入日志
		models.LogClientStatus(cid, c.Ctx.Input.IP(), "请求分钟线")
	} else {
		//校验不通过
		//校验码没有通过验证，返回未登陆信息
		r.Id = 404
		r.Msg = ""
		c.Data["json"] = &r
		c.ServeJSON()
	}
}

func (c *DaylineController) Post() {
	/*处理下载日线文件请求
	请求参数为授权码，例如?a=daddada，返回用户订阅股票和指数的下载地址列表
	*/
	var r models.ResponseMsg
	mtmp := c.GetString("a")

	var mvalid bool
	mvalid = false
	var cid int
	cid = -1
	var mauth string
	//取授权码
	mauth = mtmp
	//校验授权码
	if models.Bm.IsExist(mauth) {
		//缓存中存在，校验成功，读出CID
		cid = models.Bm.Get(mauth).(int)
		mvalid = true
	} else {
		//缓存中不存在，查询数据库
		cid = models.QueryAuthcode(mauth)
		if cid != -1 {
			//数据库中存在，写入缓存
			mvalid = true
			models.Bm.Put(mauth, cid, 3600*time.Second)
		}
	}
	if mvalid == true {
		//校验通过，查询订阅指数记录，生成下载地址，转换成JSON字符串
		var msubs []models.Clientsubs
		if models.Bm.IsExist(mauth + "subs") {
			//检查缓存中是否存在订阅记录
			msubs = models.Bm.Get(mauth + "subs").([]models.Clientsubs)
		} else {
			//缓存中不存在，从数据库中检索
			msubs = models.QueryClientSubs(cid)
			//写入缓存
			models.Bm.Put(mauth+"subs", msubs, 3600*time.Second)
		}
		//循环msubs，日线文件下载地址
		var mr []models.Dayline_Download_Url
		mr = models.GenDaylineUrls(msubs)
		j, _ := json.Marshal(mr)
		r.Id = 200
		r.Msg = string(j)
		c.Data["json"] = &r
		c.ServeJSON()
		//写入日志
		models.LogClientStatus(cid, c.Ctx.Input.IP(), "请求日线")
	} else {
		//校验不通过
		//校验码没有通过验证，返回未登陆信息
		r.Id = 404
		r.Msg = ""
		c.Data["json"] = &r
		c.ServeJSON()
	}
}

func (c *WeeklineController) Get() {
	/*处理下载周线文件请求
	请求参数为授权码，例如?a=daddada，返回用户订阅股票和指数的下载地址列表
	*/
	var r models.ResponseMsg
	mtmp := c.GetString("a")

	var mvalid bool
	mvalid = false
	var cid int
	cid = -1
	var mauth string
	//取授权码
	mauth = mtmp
	//校验授权码
	if models.Bm.IsExist(mauth) {
		//缓存中存在，校验成功，读出CID
		cid = models.Bm.Get(mauth).(int)
		mvalid = true
	} else {
		//缓存中不存在，查询数据库
		cid = models.QueryAuthcode(mauth)
		if cid != -1 {
			//数据库中存在，写入缓存
			mvalid = true
			models.Bm.Put(mauth, cid, 3600*time.Second)
		}
	}
	if mvalid == true {
		//校验通过，查询订阅指数记录，生成下载地址，转换成JSON字符串
		var msubs []models.Clientsubs
		if models.Bm.IsExist(mauth + "subs") {
			//检查缓存中是否存在订阅记录
			msubs = models.Bm.Get(mauth + "subs").([]models.Clientsubs)
		} else {
			//缓存中不存在，从数据库中检索
			msubs = models.QueryClientSubs(cid)
			//写入缓存
			models.Bm.Put(mauth+"subs", msubs, 3600*time.Second)
		}
		//循环msubs，日线文件下载地址
		var mr []models.Indextick_Download_Url
		mr = models.GenWeeklineUrls(msubs)
		j, _ := json.Marshal(mr)
		r.Id = 200
		r.Msg = string(j)
		c.Data["json"] = &r
		c.ServeJSON()
		//写入日志
		models.LogClientStatus(cid, c.Ctx.Input.IP(), "请求周线")
	} else {
		//校验不通过
		//校验码没有通过验证，返回未登陆信息
		r.Id = 404
		r.Msg = ""
		c.Data["json"] = &r
		c.ServeJSON()
	}
}

func (c *MonthlineController) Get() {
	/*处理下载月线文件请求
	请求参数为授权码，例如?a=daddada，返回用户订阅股票和指数的下载地址列表
	*/
	var r models.ResponseMsg
	mtmp := c.GetString("a")

	var mvalid bool
	mvalid = false
	var cid int
	cid = -1
	var mauth string
	//取授权码
	mauth = mtmp
	//校验授权码
	if models.Bm.IsExist(mauth) {
		//缓存中存在，校验成功，读出CID
		cid = models.Bm.Get(mauth).(int)
		mvalid = true
	} else {
		//缓存中不存在，查询数据库
		cid = models.QueryAuthcode(mauth)
		if cid != -1 {
			//数据库中存在，写入缓存
			mvalid = true
			models.Bm.Put(mauth, cid, 3600*time.Second)
		}
	}
	if mvalid == true {
		//校验通过，查询订阅指数记录，生成下载地址，转换成JSON字符串
		var msubs []models.Clientsubs
		if models.Bm.IsExist(mauth + "subs") {
			//检查缓存中是否存在订阅记录
			msubs = models.Bm.Get(mauth + "subs").([]models.Clientsubs)
		} else {
			//缓存中不存在，从数据库中检索
			msubs = models.QueryClientSubs(cid)
			//写入缓存
			models.Bm.Put(mauth+"subs", msubs, 3600*time.Second)
		}
		//循环msubs，日线文件下载地址
		var mr []models.Indextick_Download_Url
		mr = models.GenMonthlineUrls(msubs)
		j, _ := json.Marshal(mr)
		r.Id = 200
		r.Msg = string(j)
		c.Data["json"] = &r
		c.ServeJSON()
		//写入日志
		models.LogClientStatus(cid, c.Ctx.Input.IP(), "请求月线")
	} else {
		//校验不通过
		//校验码没有通过验证，返回未登陆信息
		r.Id = 404
		r.Msg = ""
		c.Data["json"] = &r
		c.ServeJSON()
	}
}

func (c *ClientStatusController) Get() { //此处实际应用时需改成POST方法
	/*处理下载操作请求，请求参数只有授权码*/
	var r models.ResponseMsg
	mtmp := c.GetString("a")
	var mvalid bool
	mvalid = false
	var cid int
	cid = -1
	var mauth string
	mauth = mtmp
	if models.Bm.IsExist(mauth) {
		//缓存中存在，校验成功，读出CID
		cid = models.Bm.Get(mauth).(int)
		mvalid = true
	} else {
		//缓存中不存在，查询数据库
		cid = models.QueryAuthcode(mauth)
		if cid != -1 {
			//数据库中存在，写入缓存
			mvalid = true
			models.Bm.Put(mauth, cid, 3600*time.Second)
		}
	}
	if mvalid == true {
		//校验通过，查询全部股票代码记录，转换成JSON字符串
		mr := models.QueryClientStatus(cid)
		j, _ := json.Marshal(mr)
		r.Id = 200
		r.Msg = string(j)
		c.Data["json"] = &r
		c.ServeJSON()
		//写入日志
		models.LogClientStatus(cid, c.Ctx.Input.IP(), "查询日志")
	} else {
		//校验不通过
		//校验码没有通过验证，返回未登陆信息
		r.Id = 404
		r.Msg = ""
		c.Data["json"] = &r
		c.ServeJSON()
	}
}

func (c *ClientUpstatusController) Get() { //实际用的时候改成POST
	/*客户端上报运行状态接口，参数形式为授权码|TAG|信息字符串*/
	var r models.ResponseMsg
	mtmp := c.GetString("a")
	mtmp2 := strings.Split(mtmp, "|")

	var mvalid bool
	mvalid = false
	var cid int
	cid = -1
	var mauth string
	var mtag int
	var msg string

	//取授权码
	if (len(mtmp2)) == 3 {
		mauth = mtmp2[0]
		//取日期值
		mtag, _ = strconv.Atoi(mtmp2[1])
		msg = mtmp2[2]
		//校验授权码
		if models.Bm.IsExist(mauth) {
			//缓存中存在，校验成功，读出CID
			cid = models.Bm.Get(mauth).(int)
			mvalid = true
		} else {
			//缓存中不存在，查询数据库
			cid = models.QueryAuthcode(mauth)
			if cid != -1 {
				//数据库中存在，写入缓存
				mvalid = true
				models.Bm.Put(mauth, cid, 3600*time.Second)
			}
		}
	}
	if mvalid == true {
		//校验通过，保存上报状态
		models.InsertClientStatus(cid, mtag, c.Ctx.Input.IP(), msg)
		r.Id = 200
		r.Msg = ""
		c.Data["json"] = &r
		c.ServeJSON()
		//写入日志
		//models.LogClientStatus(cid, c.Ctx.Input.IP(), "上报状态")
	} else {
		//校验不通过
		//校验码没有通过验证，返回未登陆信息
		r.Id = 404
		r.Msg = ""
		c.Data["json"] = &r
		c.ServeJSON()
	}
}

func (c *PlansController) Get() {
	//用户请求全部模拟环境列表
	var r models.ResponseMsg
	mtmp := c.GetString("a")

	var mvalid bool
	mvalid = false
	var cid int
	cid = -1
	var mauth string
	//取授权码
	mauth = mtmp
	//校验授权码
	if models.Bm.IsExist(mauth) {
		//缓存中存在，校验成功，读出CID
		cid = models.Bm.Get(mauth).(int)
		mvalid = true
	} else {
		//缓存中不存在，查询数据库
		cid = models.QueryAuthcode(mauth)
		if cid != -1 {
			//数据库中存在，写入缓存
			mvalid = true
			models.Bm.Put(mauth, cid, 3600*time.Second)
		}
	}
	if mvalid == true {
		//校验通过，查询Plans表，转换成JSON字符串
		var mplans []models.Plans
		if models.Bm.IsExist("plans") {
			//检查缓存中是否存在订阅记录
			mplans = models.Bm.Get("plans").([]models.Plans)
		} else {
			//缓存中不存在，从数据库中检索
			mplans = models.QueryPlans()
			//写入缓存
			models.Bm.Put("plans", mplans, 3600*time.Second)
		}
		//循环plans
		j, _ := json.Marshal(mplans)
		r.Id = 200
		r.Msg = string(j)
		c.Data["json"] = &r
		c.ServeJSON()
		//写入日志
		models.LogClientStatus(cid, c.Ctx.Input.IP(), "请求全部仿真环境列表")
	} else {
		//校验不通过
		//校验码没有通过验证，返回未登陆信息
		r.Id = 404
		r.Msg = ""
		c.Data["json"] = &r
		c.ServeJSON()
	}
}
