# translator-cli #

I just want to be able to translate things in CLI. That's what this project is for :)

** Requirements **

Put your Google-Translate APIKEY into ~/.config/translator/APIKEY

More about [setting APIKEYs in Google Apps](https://support.google.com/cloud/answer/6158862?hl=en&ref_topic=6262490)

**Usage:** ```python trans.py <trans_langs> string```

Where **<trans_langs>** has format **sourcelang-destinationlang** e.g. **en-pl** when translating
from english to polish

**Example:** ```python trans.py en-pl core```

This project uses **google-translator API**:
- [developers' console](https://console.developers.google.com/project/)
- [API webpage](https://cloud.google.com/translate/v2)
