<img src="https://codeberg.org/video-prize-ranch/rimgo/raw/branch/main/static/img/rimgo.svg" width="96" height="96" />

# rimgo
An alternative frontend for Imgur. Based on [rimgu](https://codeberg.org/3np/rimgu) and rewritten in Go.

<a href="https://www.gnu.org/licenses/agpl-3.0.en.html">
  <img alt="License: AGPLv3" src="https://shields.io/badge/License-AGPL%20v3-blue.svg">
</a>

It's read-only and works without JavaScript. Images and albums can be viewed without wasting resources from downloading and running tracking scripts. No sign-up nags.

## Features

- [x] URL-compatible with i.imgur.com - just replace the domain in the URL
- [x] Images and videos (gifv, mp4)
- [ ] Galleries with comments
- [x] Albums
- [ ] User page
- [ ] Tag page

This is currently very early stage software. Some things left to implement (contributions welcome!):

- [x] Streaming (currently media is downloaded in full in rimgu before it's returned)
- [ ] Localization/internationalization
- [x] Pretty CSS styling (responsive?)
- [ ] Support for other popular image sites
- [ ] Filtering and exploration on user/tags pages
- [ ] Responsive scaling of videos on user/tags pages
- [x] Logo

Things that are considered out of scope:

* Uploading, commenting, voting, etc - rimgo is read-only.
* Authentication, serving HTTPS, rate limiting, etc - Use a reverse proxy or load balancer like Caddy, Traefik, or NGINX.
* Anything requiring JavaScript or the client directly interacting with upstream servers.

## Instances

Open an issue to have your instance listed here!

| Website                                                                                                                                  | Country | Cloudflare |
|------------------------------------------------------------------------------------------------------------------------------------------|---------|------------|
| [i.bcow.xyz](https://i.bcow.xyz/)                                                                                                        | 🇨🇦 CA   |            |
| [rimgo.pussthecat.org](https://rimgo.pussthecat.org/)                                                                                    | 🇩🇪 DE   |            |
| [img.riverside.rocks](https://img.riverside.rocks) |     🇺🇸 US    |            |
| [l4d4owboqr6xcmd6lf64gbegel62kbudu3x3jnldz2mx6mhn3bsv3zyd.onion](http://l4d4owboqr6xcmd6lf64gbegel62kbudu3x3jnldz2mx6mhn3bsv3zyd.onion/) |         |            |
| [jx3dpcwedpzu2mh54obk5gvl64i2ln7pt5mrzd75s4jnndkqwzaim7ad.onion](http://jx3dpcwedpzu2mh54obk5gvl64i2ln7pt5mrzd75s4jnndkqwzaim7ad.onion) |     🇺🇸 US    |            |


## Install
Librarian can run on any platform Go compiles on.

### Docker
Install Docker and docker-compose, then clone this repository.
```
git clone https://codeberg.org/video-prize-ranch/rimgo
cd rimgo
```

Edit the `docker-compose.yml` file using your favorite text editor.
```
nvim docker-compose.yml
```

You can now run rimgo.
```
sudo docker-compose up -d
```

### Build from source

#### Requirements
* Go v1.16 or later

Clone the repository and `cd` into it.
```
git clone https://codeberg.org/video-prize-ranch/rimgo
cd rimgo
```

Build rimgo.
```
go build
```

Edit the config file using your preferred editor.
```
nvim config.yml
```

You can now run rimgo.
```
./rimgo
```


## Configuration

rimgo can be configured using environment variables or a config file.

### Environment variables

| Name                  | Default         |
|-----------------------|-----------------|
| RIMGU_PORT            | 3000            |
| RIMGU_HOST            | localhost       |
| RIMGU_ADDRESS         | 0.0.0.0         |
| RIMGU_IMGUR_CLIENT_ID | 546c25a59c58ad7 |

## Contributing

PRs are welcome!

This software is released under the AGPL 3.0 license. In short, this means that if you make any modifications to the code and then publish the result (e.g. by hosting the result on a web server), you must publicly distribute your changes and declare that they also use AGPL 3.0.