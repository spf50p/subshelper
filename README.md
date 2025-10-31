# subshelper

Directory structure generation for subscriptions using the Caddy server. Example directory structure:

```
.work_dir
├── .headers.caddy
├── 099aaa26aacf19430366853c10ee7aa2
│   ├── .headers.caddy
│   ├── index.html
│   └── index.txt
```

Each directory contains files for a single user. Inside the directory:

- `index.html` allows opening a web page with the subscription in a browser;
- `index.txt` for all other clients (V2RayTun, curl, etc.), in `base64` format (`base64 -w 0 index.txt > index.txt.base64`);
- `.headers.caddy` contains user-specific `headers` (optional).

Example `Caddyfile` config:

```
s.subshelper.42 {
	encode zstd gzip
	root * /usr/share/caddy
	header -Server
	import /usr/share/caddy/.headers.caddy
	import /usr/share/caddy/*/.headers.caddy

	file_server {
		hide */.* */.headers.caddy */*.caddy */*.conf *~ *.bak *.swp
	}

	@badmethods {
		not method GET
	}
	respond @badmethods 405

	@html header Accept *text/html*
	handle @html {
		try_files {path}/index.html
		file_server
	}

	handle {
		header Content-Type "text/plain; charset=utf-8"
		try_files {path}/index.txt
		file_server
	}
}
```

Example `~/.subshelper.yaml` config:

```yaml
work_dir: .work_dir
subscription:
  title: SubsHelper                      # optional
  title_url_text: Subscription URL       # optional
  base_url: https://s.subshelper.42      # optional
  global_headers:                        # minimum one header is required
    profile-title: SubsHelper
    profile-update-interval: 10
    announce-url: https://subshelper.42/announce
    support-url: https://subshelper.42/support
    profile-web-page-url: https://subshelper.42/sh
  subs:
    - id: 099aaa26aacf19430366853c10ee7aa2
      links:
        - vless://418179d7-bb27-4f1f-98f4-cffba0878e91@gt.subshelper.42:443?flow=xtls-rprx-vision&type=tcp&security=tls&fp=firefox&alpn=http%2F1.1#🇳🇱 Netherlands
        - vless://418179d7-bb27-4f1f-98f4-cffba0878e91@ry.subshelper.42:443?flow=xtls-rprx-vision&type=tcp&security=tls&fp=firefox&alpn=http%2F1.1#🇺🇸 United States (East)
      headers:                           # optional
        announce: "Your ID: 323232"
        subscription-userinfo: upload=0;download=0;total=0;expire=1771070400
```

After generating the directory structure, you need to restart the Caddy server:

```sh
sudo systemctl reload caddy.service
```
