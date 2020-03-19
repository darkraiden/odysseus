# Odysseus

[![Build Status](https://dev.azure.com/darkraiden/Odysseus/_apis/build/status/darkraiden.odysseus?branchName=master)](https://dev.azure.com/darkraiden/Odysseus/_build/latest?definitionId=2&branchName=master)

Odysseus is your friendly DNS pathfinder.

## Description

Odysseus is designed to update a list of [Cloudflare](https://cloudflare.com) _DNS A Records_ with your ISP Dynamic IP Address. The app reads the config from a `YAML` file, queries the Cloudflare API and if the content of each record has changed (i.e. the content of the DNS record != your IPS public IP), it'll go ahead and update it for you.

If you wrap this tool in a crontab, you might be all set to host your website/blog on a cluster of Raspberry Pis.

## How to use it

First things first, download the `odysseus` binary:

```bash
mdir -p /opt/odysseus && cd /opt/odysseus
wget https://github.com/darkraiden/odysseus/releases/download/v0.4/odysseus_<version>_<Linux|Darwin>_<i386|x86_64>.tar.gz
tar zxvf odysseus_<version>_<Linux|Darwin>_<i386|x86_64>.tar.gz
rm odysseus_<version>_<Linux|Darwin>_<i386|x86_64>.tar.gz
```

In the same directory where odysseus was downloaded, create a file called `cloudflare.yml`:

```bash
cat <<EOF > ./cloudflare.yml
cloudflare:
  zone_name: example.com
  email: user@example.com
  api_key: cloudflareapikeygoeshere
  records:
    - www.example.com
    - api.example.com
EOF
```

To establish a connection with Cloudflare, you need the `email` address you use to log in and the `api_key` which can be found in `My Profile` > `API Tokens` > `Global API Key`. Ensure that the rest of the details (`zone_name` and `records`) are correct too.

Now that your config is ready, simply type:

```bash
./odysseus
```

Alternatively, the config file can be stored elsewhere with a different name; if so, ensure to pass the appropriate flags when running the binary:

```bash
./odysseus -config-name someName.yml -config-path /path/to/config/file
```
