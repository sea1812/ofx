# encoding: UTF-8
# 测试交易前端

from mods import octkeypass
from socket import *
import json
import config
import time

host = "127.0.0.1"
port = 1026
buffersize=1
addr = (host,port)

client = socket()#AF_INET,SOCK_STREAM)
client.connect(addr)

print 'connected'

def test_200001():
    cmd = dict()
    cmd['cmd']=200001
    p = dict()
    p['code']='000001'
    p['market']='SZ'
    cmd['params']=p
    cstr = json.dumps(cmd)
    cstr = octkeypass.encrypt(config.superkey, cstr)
    print cstr
    client.sendall("%s\r\n" % cstr)
    c = client.makefile().readline()
    print octkeypass.extract(c)

def test_200002():
    cmd = dict()
    cmd['cmd']=200002
    p = dict()
    p['code']='000001'
    p['market']='SZ'
    cmd['params']=p
    cstr = json.dumps(cmd)
    cstr = octkeypass.encrypt(config.superkey, cstr)
    print cstr
    client.sendall("%s\r\n" % cstr)
    c = client.makefile().readline()
    print octkeypass.extract(c)
    
def test_200003():
    cmd = dict()
    cmd['cmd']=200003
    p = dict()
    cmd['params']=p
    cstr = json.dumps(cmd)
    cstr = octkeypass.encrypt(config.superkey, cstr)
    print cstr
    client.sendall("%s\r\n" % cstr)
    c = client.makefile().readline()
    print octkeypass.extract(c)

def test_200004():
    cmd = dict()
    cmd['cmd']=200004
    p = dict()
    cmd['reqid']=19
    cstr = json.dumps(cmd)
    cstr = octkeypass.encrypt(config.superkey, cstr)
    print cstr
    client.sendall("%s\r\n" % cstr)
    c = client.makefile().readline()
    print octkeypass.extract(c)

def test_200005():
    cmd = dict()
    cmd['cmd']=200005
    cmd['reqid']=6
    cstr = json.dumps(cmd)
    cstr = octkeypass.encrypt(config.superkey, cstr)
    print cstr
    client.sendall("%s\r\n" % cstr)
    c = client.makefile().readline()
    print octkeypass.extract(c)

def test_200006():
    cmd = dict()
    cmd['cmd']=200006
    cmd['reqid']=4
    cstr = json.dumps(cmd)
    cstr = octkeypass.encrypt(config.superkey, cstr)
    print cstr
    client.sendall("%s\r\n" % cstr)
    c = client.makefile().readline()
    print octkeypass.extract(c)

def test_210001():
    cmd = dict()
    cmd['cmd']=210001
    cmd['params']={}
    cstr = json.dumps(cmd)
    print cstr
    cstr = octkeypass.encrypt(config.superkey, cstr)
    print cstr
    client.sendall("%s\r\n" % cstr)
    c = client.makefile().readline()
    print octkeypass.extract(c)

def test_210002():
    cmd = dict()
    cmd['cmd']=210002
    cmd['params']={}
    cstr = json.dumps(cmd)
    print cstr
    cstr = octkeypass.encrypt(config.superkey, cstr)
    print cstr
    client.sendall("%s\r\n" % cstr)
    c = client.makefile().readline()
    print octkeypass.extract(c)
    
def test_210003():
    cmd = dict()
    cmd['cmd']=210003
    cmd['params']={}
    cstr = json.dumps(cmd)
    print cstr
    cstr = octkeypass.encrypt(config.superkey, cstr)
    print cstr
    client.sendall("%s\r\n" % cstr)
    c = client.makefile().readline()
    print octkeypass.extract(c)
    
def test_210004():
    cmd = dict()
    cmd['cmd']=210004
    cmd['params']={}
    cstr = json.dumps(cmd)
    print cstr
    cstr = octkeypass.encrypt(config.superkey, cstr)
    print cstr
    client.sendall("%s\r\n" % cstr)
    c = client.makefile().readline()
    print octkeypass.extract(c)

def test_210005():
    cmd = dict()
    cmd['cmd']=210005

    order = {}
    order['ticker'] = '000003'  # 平安银行
    order['market'] = 1         # 深圳A股
    order['price'] = 11.11
    order['quantity'] = 2100
    order['price_type'] = 1     # 限价单
    order['side'] = 1           # 买
    order['business_type'] = 0  # 普通股票业务
    order['order_client_id'] = 111 #自定义客户端号

    cmd['params'] = order
    print cmd['params'], type(cmd['params'])
    print order, type(order)
    cstr = json.dumps(cmd, ensure_ascii=True)
    #print cstr
    cstr = octkeypass.encrypt(config.superkey, cstr)
    #print cstr
    client.sendall("%s\r\n" % cstr)
    c = client.makefile().readline()
    print octkeypass.extract(c)

def test_210006():
    cmd = dict()
    cmd['cmd']=210006
    cmd['authcode']='66227638'
    cmd['params'] = 1 #ORDER_XTP_ID
    cstr = json.dumps(cmd, ensure_ascii=True)
    print cstr
    cstr = octkeypass.encrypt(config.superkey, cstr)
    #print cstr
    client.sendall("%s\r\n" % cstr)
    c = client.makefile().readline()
    print octkeypass.extract(c)

test_210006()