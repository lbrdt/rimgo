# rimgu: image host alternative frontend

## rimgu is an alternative frontend / proxy for imgur

It's read-only and works without browser JavaScript. Images and albums can be viewed without wasting resources and anyonymity from downloading and running tracking scripts. No sign-up nags.

It's lightweight and easy to configure.

Inspired by and (soon) integratable with:

* [searx](https://github.com/searx/searx)
* [teddit](https://codeberg.org/teddit/teddit)
* [Privacy Redirect](https://github.com/SimonBrazell/privacy-redirect)
* [nitter](https://github.com/zedeus/nitter)
* [bibliogram](https://sr.ht/~cadence/bibliogram/)

## Building

### Locally

Dependencies:

* node.js >= v16 (earlier most likely works fine)

```
$ npm install && npm run build
```

### Docker
```
$ docker build -t rimgu:latest .
```

## Running

### Locally

```
$ node dist/index.js
```

### Docker
```
$ docker run -p 8080:8080 -e -it RIMGU_ADDRESS=0.0.0.0 -e RIMGU_PORT=8080 rimgu:latest
```

## Configuration

Rimgu is configured via environment variables. See available variable in [src/config.ts](./src/config.ts).
