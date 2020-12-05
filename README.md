# translator-cli #

I just want to be able to translate things in CLI. That's what this project is for :)

** Requirements **

1. You need a valid project in Google Cloud Platform
1. You need to create service account (w/at least "Cloud Translation API User" permissions) and export JSON auth file somewhere to your disk
1. Lastly, you need to create and fill configuration file in **~/.translator-cli.yml** - use **cli-config-template.yml** as template

**Usage:**

```translator-cli translate -s <source-lang> -d <destination-lang> text to translate```

**Example:**

```translator-cli translate -s en -d es Is it snowing now?```