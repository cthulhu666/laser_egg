import datadog
import os

datadog.initialize(api_key=os.environ['DD_API_KEY'],
                   app_key=os.environ['DD_APP_KEY'])

stats = datadog.ThreadStats()
stats.start()


def send(_ts, measurement):
    stats.gauge('air.pm10', measurement['pm10'])
    stats.gauge('air.pm25', measurement['pm25'])
