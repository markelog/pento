# Pento [![Build Status](https://travis-ci.org/markelog/pento.svg?branch=master)](https://travis-ci.org/markelog/pento)

> Time Tracker

## Intro

This project intendent to be looked at for inspiration reasons (mainly mine :-). It's not prepared for the deploy of any kind, only dev version is available

## Development

In order to check out this project you need to open to terminal sessions, one for backend and the other for frontend.

### Back
Requires installed `docker` stuff, `make` and `Go` language on your machine
```
$ cd back && cp .env.example .env && make dev
```

### Front
Only `Node.js` is required :), no tests or typescript – :(( I suck, haven't had the time

```bash
$ cd front && cp .env.example .env && npm i && npm run dev
```

### Then

And then you can check out the [`app`](http://localhost:3000)