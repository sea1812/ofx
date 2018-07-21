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

def test_compress():
    # 测试压缩和解压缩
    a = '''_price\": 0.0, \"exec_time\": 0, \"lower_limit_price\": 0.0, \"bid\": [0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0], \"ma_bid_price\": 0.0, \"warrant_upper_price\": 0.0, \"ask_qty\": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0], \"etf_buy_qty\": 0.0, \"pe_ratio2\": 0.0, \"cancel_sell_qty\": 0.0, \"cancel_sell_count\": 0, \"total_sell_count\": 0, \"qty\": 437300300, \"ticker\": \"399006\", \"num_bid_orders\": 0, \"ask\": [0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0], \"total_ask_qty\": 0, \"data_time\": 20180607102256000, \"settlement_price\": 0.0, \"num_ask_orders\": 0, \"etf_sell_money\": 0.0, \"cancel_sell_money\": 0.0, \"cancel_buy_qty\": 0.0, \"open_price\": 1747.93, \"duration_after_sell\": 0, \"turnover\": 6988980000.0}"], ["{\"etf_sell_count\": 0, \"trades_count\": 0, \"ma_ask_price\": 0.0, \"warrant_lower_price\": 0.0, \"pre_delta\": 0.0, \"etf_buy_money\": 0.0, \"cancel_buy_count\": 0, \"low_price\": 10360.96, \"ticker_status\": \"\", \"pre_close_price\": 10365.13, \"last_price\": 10381.67, \"is_market_closed\": \"\", \"upper_limit_price\": 0.0, \"total_buy_count\": 0, \"exchange_id\": 2, \"close_price\": 10381.67, \"total_bid_qty\": 0, \"iopv\": 0.0, \"yield_to_maturity\": 0.0, \"curr_delta\": 0.0, \"pre_open_interest\": 0.0, \"pe_ratio1\": 0.0, \"open_interest\": 0.0, \"cancel_buy_money\": 0.0, \"avg_price\": 0.0, \"duration_after_buy\": 0, \"total_warrant_exec_qty\": 0.0, \"bid_qty\": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0], \"ma_bond_bid_price\": 0.0, \"total_position\": 0.0, \"etf_sell_qty\": 0.0, \"high_price\": 10411.56, \"pre_settlement_price\": 0.0, \"etf_buy_count\": 0, \"ma_bond_ask_price\": 0.0, \"exec_time\": 0, \"lower_limit_price\": 0.0, \"bid\": [0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0], \"ma_bid_price\": 0.0, \"warrant_upper_price\": 0.0, \"ask_qty\": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0], \"etf_buy_qty\": 0.0, \"pe_ratio2\": 0.0, \"cancel_sell_qty\": 0.0, \"cancel_sell_count\": 0, \"total_sell_count\": 0, \"qty\": 5390183300, \"ticker\": \"399001\", \"num_bid_orders\": 0, \"ask\": [0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0], \"total_ask_qty\": 0, \"data_time\": 20180607102256000, \"settlement_price\": 0.0, \"num_ask_orders\": 0, \"etf_sell_money\": 0.0, \"cancel_sell_money\": 0.0, \"cancel_buy_qty\": 0.0, \"open_price\": 10397.65, \"duration_after_sell\": 0, \"turnover\": 72144530000.0}"], ["{\"etf_sell_count\": 0, \"trades_count\": 0, \"ma_ask_price\": 0.0, \"warrant_lower_price\": 0.0, \"pre_delta\": 0.0, \"etf_buy_money\": 0.0, \"cancel_buy_count\": 0, \"low_price\": 8.36, \"ticker_status\": \"T111    \", \"pre_close_price\": 8.36, \"last_price\": 8.4, \"is_market_closed\": \"\", \"upper_limit_price\": 0.0, \"total_buy_count\": 0, \"exchange_id\": 2, \"close_price\": 8.4, \"total_bid_qty\": 0, \"iopv\": 0.0, \"yield_to_maturity\": 0.0, \"curr_delta\": 0.0, \"pre_open_interest\": 0.0, \"pe_ratio1\": 0.0, \"open_interest\": 0.0, \"cancel_buy_money\": 0.0, \"avg_price\": 8.419421487603305, \"duration_after_buy\": 0, \"total_warrant_exec_qty\": 0.0, \"bid_qty\": [15800, 22100, 10200, 45800, 42000, 0, 0, 0, 0, 0], \"ma_bond_bid_price\": 0.0, \"total_position\": 0.0, \"etf_sel'''
    print u"原文------------------------"
    print a
    print u"压缩------------------------"
    b = octkeypass.compress(a)
    print b
    print u"解压缩----------------------"
    c = octkeypass.extract(b)
    print c

#test_100001()
#test_100002()
while True:
    test_100003()
    time.sleep(1)
    
#test_compress()

client.close()
