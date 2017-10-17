iap_curl
========

curl wrapper for making HTTP request to IAP-protected app in CLI more easier than curl

## Usage

```console
$ export GOOGLE_APPLICATION_CREDENTIALS="/path/to/service-account.json"
$ export IAP_CLIENT_ID="342624545358-asdfd8fas9df8sd7ga0sdguadfpvqp69.apps.googleusercontent.com"
$
$ iap_curl http://iap-protected.webapp.com
```

The option of iap_curl is fully compatible with curl one.

If you want to use [httpstat](https://github.com/b4b4r07/httpstat), please specify the `IAP_CURL_BIN` environment variable:

```console
$ export IAP_CURL_BIN="httpstat.sh"
$ iap_curl https://tellme.tokyo
Connected to 104.31.70.103:443

HTTP/2.0 200 OK
Server: cloudflare-nginx
Access-Control-Allow-Origin: *
Cache-Control: max-age=600
Cf-Ray: 3af48c40aa3694cf-NRT
Content-Type: text/html; charset=utf-8
Date: Tue, 17 Oct 2017 16:13:54 GMT
Expires: Tue, 17 Oct 2017 16:23:54 GMT
Last-Modified: Mon, 16 Oct 2017 04:33:46 GMT
Set-Cookie: __cfduid=db7e1d73f138bcb26e0d6a040e9f5df491508256834; expires=Wed, 17-Oct-18 16:13:54 GMT; path=/; domain=.tellme.tokyo; HttpOnly; Secure
Strict-Transport-Security: max-age=15552000; preload
X-Content-Type-Options: nosniff
X-Github-Request-Id: 2A8B:16E6:10351CA:186074E:59E62C3F

Body discarded

  DNS Lookup   TCP Connection   TLS Handshake   Server Processing   Content Transfer
[      2ms  |          57ms  |        320ms  |            303ms  |             0ms  ]
            |                |               |                   |                  |
   namelookup:2ms            |               |                   |                  |
                       connect:60ms          |                   |                  |
                                   pretransfer:381ms             |                  |
                                                     starttransfer:684ms            |
                                                                                total:684ms
```

## Installation

```
$ go get github.com/b4b4r07/iap_curl
```

## License

MIT

## Author

b4b4r07
