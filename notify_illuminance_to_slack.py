#!/usr/bin/env python3
# coding: utf-8

import smbus
import time
import requests, json
import datetime
import jpholiday
WEB_HOOK_URL = "****"
message = u'defalut'

bus = smbus.SMBus(1)
bus.write_byte_data(0x13, 0x80, 0xFF)
bus.write_byte_data(0x13, 0x82, 0x00)
bus.write_byte_data(0x13, 0x84, 0x9D)
time.sleep(0.8)
data = bus.read_i2c_block_data(0x13, 0x85, 4)

luminance = data[0] * 256 + data[1]
proximity = data[2] * 256 + data[3]


today = datetime.date.today()
today_str = str(today).replace('-', '')
now_hour_int = datetime.datetime.now().hour

def isBizDay(DATE):
 Date = datetime.date(int(DATE[0:4]), int(DATE[4:6]), int(DATE[6:8]))
 if Date.weekday() >= 5 or jpholiday.is_holiday(Date):
  return 0
 else:		
  return 1

def sendMessage():
 if luminance >= 800:
  message = u'現在、オフィスに人がいます！早く帰りましょう！'	
 else:
  message = u'現在、オフィスに人はいません。'

 requests.post(WEB_HOOK_URL, data = json.dumps({
 'text': u'照度は、'+str(luminance)+u'です！'+message,  #通知内容
 'username': u'Bakira-Tech-Python-Bot',  #ユーザー名
 'icon_emoji': u':smile_cat:',  #アイコン
 'link_names': 1,  #名前をリンク化
 }))


if isBizDay(today_str) == 1:
 print("平日です")		
 if now_hour_int > 17:
  print("実行します")	
  sendMessage()
else: 	
 print("休日です")		
 if now_hour_int > 10:
  print("実行します")
  sendMessage()
