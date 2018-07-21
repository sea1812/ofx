# encoding: UTF-8

from mods import octkeypass
from socket import *
import json
import config
import time

host = "192.168.0.64"
port = 1027
buffersize=1
addr = (host,port)

client = socket()
client.connect(addr)

print 'connected'

def test_000001():
    cmd = dict()
    cmd['cmd']=000001
    p = dict()
    p['username']='demo'
    p['password']='demo'
    cmd['params']=p
    cstr = json.dumps(cmd)
    cstr = octkeypass.encrypt(config.superkey, cstr)
    client.sendall("%s\r\n" % cstr)
    c= client.makefile().readline()
    print octkeypass.extract(c)

test_000001()
    
client.close()
