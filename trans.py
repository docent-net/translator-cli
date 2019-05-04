#!/usr/bin/python

import os
import json
import requests

from sys import argv,exit

if(len(argv) < 3):
    print("Usage:\n\n\tChoose you languages: python trans.py <SRC_LNG>-<DST_LNG> <STRING>")
    print("\n\tYou can translate whole phrases: python trans.py <SRC_LNG>-<DST_LNG> \"some phrase to translate\"")
    print("\n\tFrom PL to EN: python trans.py pl <STRING>")
    print("\n\tFrom EN to PL: python trans.py en <STRING>\n")
    exit(1)

# let's define source and destination languages
default_src_lng = 'pl'
default_dst_lng = 'en'

_script, lngs, trans_string = argv

# using default language
if len(argv) == 2:
    if lngs == default_src_lng:
        src_lng = lngs
        dst_lng = default_dst_lng
    else:
        src_lng = default_src_lng
        dst_lng = lngs
# provided src-dst pair:
else:
    src_lng, dst_lng = lngs.split('-')

# get APIKEY
API_KEY_PATH='{}/.config/translator/APIKEY'.format(os.environ['HOME'])

if os.path.exists(API_KEY_PATH):
    with open(API_KEY_PATH, 'r') as api_key_file:
        API_KEY = api_key_file.read().replace('\n', '')
else:
    print('API key for Google translator not found in {}'.format(API_KEY_PATH))
    exit(1)

lngs = lngs.split('-')

URL = 'https://www.googleapis.com/language/translate/v2?key={}&q={}&source={}&target={}'.format(
    API_KEY, trans_string, src_lng, dst_lng
)

try:
    data = json.loads(requests.get(URL).text)
except ValueError as e:
    print("Couldn't translate")
    exit(1)

print(data['data']['translations'][0]['translatedText'])