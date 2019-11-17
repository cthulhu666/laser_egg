import json
import os
import time

import requests

import dog
import mqtt


device_id = os.environ['LASEREGG_ID']
api_key = os.environ['LASEREGG_KEY']

outputs = [
    dog.send,
    mqtt.send
]


def get_measurement():
    rs = requests.get(f"https://api.origins-china.cn/v1/lasereggs/{device_id}?key={api_key}")
    data = json.loads(rs.content)
    return data['info.aqi']['ts'], data['info.aqi']['data']


def process():
    ts, measurement = get_measurement()
    print(ts, measurement)

    for out in outputs:
        try:
            out(ts, measurement)
        except RuntimeError as e:
            print(e)


while True:
    try:
        process()
    except (RuntimeError, IOError) as e:
        print(e)

    time.sleep(60*5)
