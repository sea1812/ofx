# encoding: UTF-8

from mods import octkeypass
from socket import *
import json
import config
import time

host = "192.168.0.64"
port = 1025
buffersize=1
addr = (host,port)

client = socket()
client.connect(addr)

print 'connected'

def test_100001():
    cmd = dict()
    cmd['cmd']=100001
    p = dict()
    p['code']='399006'
    p['market']='SZ'
    cmd['params']=p
    cstr = json.dumps(cmd)
    cstr = octkeypass.encrypt(config.superkey, cstr)
    print cstr
    client.sendall("%s\r\n" % cstr)
    c= client.makefile().readline()
    print octkeypass.extract(c)

def test_100002():
    cmd = dict()
    cmd['cmd']=100002
    p = dict()
    p['code']='399006'
    p['market']='SZ'
    cmd['params']=p
    cstr = json.dumps(cmd)
    cstr = octkeypass.encrypt(config.superkey, cstr)
    print cstr
    client.sendall("%s\r\n" % cstr)
    c = client.makefile().readline()
    print octkeypass.extract(c)

def test_100003():
    cmd = dict()
    cmd['cmd']=100003
    p = []
    r = dict()
    r['code']='399006'
    r['market']='SZ'
    p.append(r)

    r = dict()
    r['code']='399001'
    r['market']='SZ'
    p.append(r)

    r = dict()
    r['code']='300359'
    r['market']='SZ'
    p.append(r)

    cmd['params']=p
    cstr = json.dumps(cmd)
    cstr = octkeypass.encrypt(config.superkey, cstr)
    #print cstr
    client.sendall("%s\r\n" % cstr)
    c = client.makefile().readline()
    print octkeypass.extract(c)

def test_110001():
    cmd = dict()
    cmd['cmd']=110001
    p = []
    r = dict()
    r['code']='399006'
    r['market']='SZ'
    p.append(r)

    r = dict()
    r['code']='399001'
    r['market']='SZ'
    p.append(r)

    r = dict()
    r['code']='300359'
    r['market']='SZ'
    p.append(r)

    cmd['params']=p
    cstr = json.dumps(cmd)
    cstr = octkeypass.encrypt(config.superkey, cstr)
    print cstr
    client.sendall("%s\r\n" % cstr)
    c =  client.makefile().readline()
    print octkeypass.extract(c)

#test_100001()
#test_100002()
while True:
    test_100003()
    time.sleep(1)
    
#test_compress()

client.close()
