import json
import os
import paho.mqtt.client as mqtt


broker_url = "farmer.cloudmqtt.com"
broker_port = 12397


client = mqtt.Client()
client.username_pw_set(username=os.environ['MQTT_USERNAME'].strip(),
                       password=os.environ['MQTT_PASSWORD'].strip())
client.connect(broker_url, broker_port, keepalive=600)


def send(ts, measurement):
    client.publish(topic='laseregg', payload=json.dumps([ts, measurement]), qos=0, retain=False)
