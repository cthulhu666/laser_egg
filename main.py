import json
import logging
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


class ProcessingError(RuntimeError):
    pass


class LaserEggApiError(ProcessingError):
    pass


def get_measurement():
    rs = requests.get(f"https://api.origins-china.cn/v1/lasereggs/{device_id}?key={api_key}")
    try:
        data = json.loads(rs.content)
    except json.decoder.JSONDecodeError as e:
        raise LaserEggApiError()
    return data['info.aqi']['ts'], data['info.aqi']['data']


def process():
    ts, measurement = get_measurement()
    logging.info(ts, measurement)

    for out in outputs:
        try:
            out(ts, measurement)
        except RuntimeError as e:
            logging.warning(e)


while True:
    try:
        process()
    except (RuntimeError, IOError, ProcessingError) as e:
        logging.warning(e)

    time.sleep(60*5)
