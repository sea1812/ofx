package routers

import (
	"QuoteServer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//用户登录，POST用户名和密码，返回连接授权码
	beego.Router("/login", &controllers.LoginController{})
	//订阅行情数据，使用POST方法，提交所有需要订阅的股票代码，服务器端与授权连接码一起保存到数据库，订阅时自动清除原有的订阅记录
	beego.Router("/sub", &controllers.SubController{})
	//下载代码表
	beego.Router("/codetable", &controllers.CodeTableController{})
	//下载交易日历
	beego.Router("/timetable", &controllers.TimeTableController{})
	//获取仿真环境列表
	beego.Router("/plans", &controllers.PlansController{})
	//下载指数TICK行情数据，按照订阅的数据记录返回下载地址，参数是日期索引，直接下载压缩的SNAP数据文件回去
	beego.Router("/indextick", &controllers.IndexTickController{})
	//下载个股TICK行情数据（包含盘口、逐笔成交），按照订阅的数据记录返回下载地址，参数是日期索引，直接下载压缩的SNAP、DW和TRADE数据文件回去
	beego.Router("/stocktick", &controllers.StockTickController{})
	//下载分钟线数据，按照订阅记录返回下载地址
	beego.Router("/minline", &controllers.MinutelineController{})
	//下载日线数据，按照订阅记录返回下载地址
	beego.Router("/dayline", &controllers.DaylineController{})
	//下载周线数据，按照订阅记录返回下载地址
	beego.Router("/weekline", &controllers.WeeklineController{})
	//下载月线数据，按照订阅记录返回下载地址
	beego.Router("/monthline", &controllers.MonthlineController{})
	//查询运行状态日志
	beego.Router("/status", &controllers.ClientStatusController{})
	//客户端上报运行状态
	beego.Router("/upstatus", &controllers.ClientUpstatusController{})
}
