# Octopus for XTP （简称OFX） 
## 简介
OFX是一套为对接XTP开发的前端服务软件，目的是简化客户端开发调试工作，降低工作量，同时为开发其他实用交易功能提供支持。
OFX主要希望实现以下功能：
* 客户端可以使用多种编程语言开发，使用标准的TCP连接，无需考虑XTP的SDK封装问题。
* 客户端使用统一格式的报文提交指令，而不是用具体的功能函数提交指令。
* 允许多个客户端使用同一个资金账号，摆脱官方SDK对于账号连接个数的限制。
* 内建高效的队列和内存数据管理机制，简化开发。
* 使用JSON进行数据交换，减少因官方频繁升级带来的结构体变化而修改程序的麻烦。
* 对上下行数据进行加密，提高安全性。
## 目前已经完成的工作
OFX已经完成原型开发，并经过了至少三个周期的测试，设计原理和基础框架可正常使用，基本达到了设计目标。目前OFX正处于快速迭代阶段，主要是扩充各项实用的交易功能。
## OFX应用实例
OFX目前已经应用到CSC定制交易终端软件的开发中。CSC交易终端是专门为CSC交易策略的一款定制软件，将交易策略与操作方法融为了一体。
