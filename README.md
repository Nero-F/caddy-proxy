# Caddy Proxy

A simple dynamic Reverse Proxy Container Ambassador
(heavily inspired by [jaswilder's nginx-proxy](https://github.com/nginx-proxy/nginx-proxy))

## Usage

To run it:

```bash
	docker run --name=caddy-proxy \
		--publish 80:80 \
		--publish 443:443 \
		--volume /var/run/docker.sock:/tmp/docker.sock:rw \
        --detach
        ghcr.io/nero-f/caddy-proxy
```

```bash
	docker run --name=my_app \
        --env VIRTUAL_HOST=foo.bar.com
        myapp_container
```

## Disclaimer

> [!Important]
> This project is a tool I use to manage my VPS-hosted applications.
> I built it to explore reverse proxies, dive deeper into Docker internals, and gain hands-on experience with the Single Node Pattern: Ambassador
> , as described in [Brendan Burns’ book Designing Distributed Systems](https://azure.microsoft.com/mediahandler/files/resourcefiles/designing-distributed-systems/Designing_Distributed_Systems.pdf#page=36&zoom=auto,-194,590).<br/>
> It is not intended for production use as a reverse proxy because this is a proof of concept demonstrating how a container Ambassador works, simplifying service management and network communication. For a robust, production-ready solution, I recommend using [nginx-proxy](https://github.com/nginx-proxy/nginx-proxy), which inspired this project. <br/>
> But if you want to contribute feel free to send a PR.
