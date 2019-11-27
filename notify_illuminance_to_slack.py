#!/usr/bin/env python3
# coding: utf-8

import smbus
import time
import requests, json
WEB_HOOK_URL = "https://hooks.slack.com/services/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"
message = u'defalut'

bus = smbus.SMBus(1)
bus.write_byte_data(0x13, 0x80, 0xFF)
bus.write_byte_data(0x13, 0x82, 0x00)
bus.write_byte_data(0x13, 0x84, 0x9D)
time.sleep(0.8)
data = bus.read_i2c_block_data(0x13, 0x85, 4)

luminance = data[0] * 256 + data[1]
proximity = data[2] * 256 + data[3]


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
