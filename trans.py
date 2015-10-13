#!/usr/bin/python

import sys,os
import json, requests
from sys import argv

if(len(argv) != 3):
    print "Usage:\n\n python trans.py en-pl string"
    sys.exit()

_script, lngs, trans_string = argv

# GET SCRIPTDIR:
_script_dir = os.path.dirname(os.path.realpath(__file__))

# get APIKEY
APIKEY=open(_script_dir+'/APIKEY','r').read()

lngs = lngs.split('-')

URL = 'https://www.googleapis.com/language/translate/v2?key='+APIKEY+'&q='+trans_string +\
      '&source='+lngs[0]+'&target='+lngs[1]

data = json.loads(requests.get(URL).text)

print data['data']['translations'][0]['translatedText']
