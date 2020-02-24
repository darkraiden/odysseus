# Odysseus

Odysseus is your friendly DNS pathfinder.

## Description

Odysseus is designed to update a list of [Cloudflare](https://cloudflare.com) _DNS A Records_ with your ISP Dynamic IP Address. The app reads the config from a `yaml` file, queries the Cloudflare API and if the content of each record has changed (i.e. the content of the DNS record != your IPS public IP), it'll go ahead and update it for you.

If you wrap this tool in a crontab, you might be all set to host your website/blog on a cluster of Raspberry Pi's.

## How to use it

First things first, download the `odysseus` binary:

```bash
echo "Instructions on how to download the binary here"
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
