# Octopus for XTP （简称OFX） 
## 简介
OFX是一套为对接XTP开发的前端服务软件，目的是简化客户端开发调试工作，降低工作量，同时为开发其他实用交易功能提供支持。
OFX主要希望实现以下功能：
* **使客户端不受XTP版本更新的影响，即使XTP推倒性升级，客户端仍能继续使用。**
* **允许客户端使用多种编程语言开发**，使用标准的TCP连接，无需考虑XTP的SDK封装问题。
* **客户端使用统一格式的报文提交指令**，而不是用具体的功能函数提交指令，**网关的升级不影响旧客户端使用**。
* **允许多个客户端使用同一个资金账号**，摆脱官方SDK对于账号连接个数的限制。
* 内建高效的队列和内存数据管理机制，使 **多个客户端可以共用网关里暂存的数据** ，并简化客户端开发。
* 使用JSON进行数据交换，**使各种编程语言的数据处理方式简单而统一**。
* 对网络上下行数据进行加密，**提高安全性和传输效率**。
## 不适合使用OFX的场景说明
* 以下场景或许不适合使用OFX：
    * 交易频率特别特别高，或者对指令传输延迟要求特别特别高的，这种场景最好是打入XTP内部或者直接对接交易所，即使是直接使用XTP官方提供的SDK也难免存在延迟。
    * 没有客户端软件开发能力的请绕行。OFX是为方便进行客制化客户端软件开发而设计的，没有客户端软件，什么也干不了。
    * 技术能力特别高超的黑科技科学家请移步。OFX不是黑科技实验场，为了追求实用目的，我们大量使用老套而稳定的技术，没有多少令人亢奋的新名词。
* 以下可能是使用OFX的理由：
    * 饱受XTP升级之苦。众所周知XTP每次升级带来的痛苦，OFX为减少痛苦而设计。
    * 饱受C++语言之苦。C++堪称最强编程语言，但未必所有人都喜欢或者擅长，而且高水平的C++程序员也不好招，OFX允许程序员使用各种编程语言，只要能实现TCP连接就行。
    * 希望多功能多客户端共用一个资金账号。例如，一个客户端用于人工交易，一个用于自动盯盘，一个用于算法分析等等，OFX可以让你这么用。每个客户端软件只实现一个功能，开发调试和维护都简单。
    * 需要使用多种语言实现客户端功能的。例如C#语言做人工操作界面，Python语言做技术分析（Pandas和TA-LIB），OFX可以让你这么做。每个客户端软件可以用不同语言实现，用最适合的语言做最适合的事情。
## 最近的主要更新
* 增加有关ETF清单和篮子信息的指令。
## 目前已经完成的工作
* OFX已经完成原型开发，并在XTP官方提供的测试环境下经过了至少30个交易日测试，可正常使用，基本达到了设计目标。
* 已有Python、Go、Delphi/Lazarus三种语言实现了客户端连接。
* CSC交易终端已经完成第一版开发，正在进行优化升级。
## 正在进行的工作
* 扩充报文指令，以更全面地支持XTP功能。
* 开发升级网关自动化运维功能。
* 开发升级网关数据处理功能。
## OFX应用情况
* OFX目前已经应用到CSC定制交易终端软件的开发中。CSC交易终端是专门为CSC交易策略的一款定制软件，将交易策略与操作方法融为了一体。目前该终端正在测试环境中进行测试。
* 已经有程序员用Python、Go、Delphi/Lazarus等编程语言实现的客户端成功与OFX服务对接。
## 其他
* OFX目前处于快速迭代当中，随时可能修改，请开发者务必注意！！！
* 大多数情况下OFX软件升级时，客户端程序不需做任何修改，如需修改，定会显著提醒！！！
* OFX是为特定的目标而产生的技术解决方案，追求的是实用而不是极致的技术指标，所以可能不适合那些极端的交易场景。由于无法理解或解释某些XTP特殊的情形，我们有时不得不通过技术方法做些折衷或妥协的处理。
* OFX只在XTP提供的测试环境中运行并测试，未在生产环境中测试过。
* 有关XTP的详细资料请访问其[官方网站](http://xtp.zts.com.cn)。
## 目录说明
* demo目录：测试和演示程序。
* document目录：有关文档和说明文件。
* octkeypass目录：OFX默认加密库文件。
* xtp目录：OFX开发时依照的XTP官方头文件。
