## Rime
Rime is simple json server.

## Installation

Download binary from [here](https://github.com/wataru0225/rime/releases/latest).
Or clone this repository and `go install`.

```
$ rime --dir JSON_DIR
2018/01/26 01:28:37 http://localhost:8080
2018/01/26 01:28:37 See browse -> http://localhost:8080/JSON_DIR/1.json
...
```

![rime_v_0_0_2](https://user-images.githubusercontent.com/7300913/35401843-05415f1c-023e-11e8-86d2-fb0585295763.gif)

## Options

```
Usage of rime:
-d, --dir string   Please select json files directory. (default "./")
    --noext        Please select if you need not to extenstion(.json).
-p, --port int     Please select if you want to use other port number except for 8080. (default 8080)
```

## Links

- [Goでjsonファイルを読み込んでHTTPで表示する](http://wataru0225.hateblo.jp/entry/2018/01/26/221512)
