# JSONenv

[![Build Status](https://travis-ci.org/assemblyline/json_env.svg?branch=master)](https://travis-ci.org/assemblyline/json_env)

## Instalation

`go install https://github.com/assemblyline/json_env`


## Usage
JSONenv writes to stdout so its output can be redirected to a file or somesuch.

```
$ env
SHELL=/bin/bash
EDITOR=vim
TERM=xterm
HOME=/Users/ed
GOPATH=/Users/ed/go
FOO=fooming_at_the_mouth
BAR=bar_snacks

$ JSON_ENV=FOO,BAR json_env
{"BAR":"bar_snacks","FOO":"fooming_at_the_mouth"}
```
## Licence

MIT see [LICENCE](https://github.com/assemblyline/json_env/blob/master/LICENCE)
