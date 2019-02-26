# use Google maps Geolocation to find current location
import re
import subprocess
import requests

API_URL = 'https://www.googleapis.com/geolocation/v1/geolocate?key='
API_KEY = 'AIzaSyAeP2GuG2WLiT_XA6uMLCkQdvkW5Q9e0XE'  # replace with your own API key

# add wifi into request if available
payload = {'considerIP': 'true'}
try:
    cmdout = subprocess.check_output('netsh wlan show networks mode=bssid').decode(encoding='utf_8')
    macs = re.findall(r'(?:[0-9a-fA-F]:?){12}', cmdout)
    signals = re.findall(r'([\d]+)(%)', cmdout)
    channels = re.findall(r'(Channel[\W]+: )([\d]+)', cmdout)
    if not (len(macs) == len(signals) == len(channels)):
        raise
except:
    print("Could not retrieve WIFI Access Points...")
else:
    print('Using {0} Wifi Access Point{1}....'.format(len(macs), 's' if (len(macs) > 1) else ''))
    wifi_list = []
    for i in range(len(macs)):
        wifi_list.append({'macAddress': macs[i],
        'signalStrength': str(int(signals[i][0]) / 2 - 100),
        'channel': channels[i][1]
        })
    payload['wifiAccessPoints'] = wifi_list

# make request to google API
try:
    response = requests.post(API_URL+API_KEY, json=payload)
    data = response.json()
    if 'error' in data.keys():
        raise requests.RequestException('Error {0}: {1}'.format(data['error']['code'], data['error']['message']))
except Exception as e:
    print(e)
else:
    lat = data['location']['lat']
    lng = data['location']['lng']
    accuracy = data['accuracy']
    print('You are within {0}m of {1}N {2}E'.format(accuracy, lat, lng))
