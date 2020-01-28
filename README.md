# Android Applications Ethics Commission

To regulate dishonest and unethical practices by installed applications.

## server

Application configure defines `system events` and app specific action.

```
# app.ayml

events:
- bg:
    define:
    - instruction: deny
      extra:
        operation: SMS_READ
    - instruction: freeze
      extra:
        delay: 10s
    apply:
    - com.tencent.mm
    - com.gotokeep.keep.intl

apps:
  com.tencent.mm:
    fg:
    - instruction: allow
      extra:
        operation: SMS_READ
    bg:
    - instruction: ignore
      extra:
        operation: SMS_WRITE
    - super
```

Core configure shows the general options:

```
# aaec.yaml

pidfile: $HOME/aaec/aaec.pid
log_level: debug
log_filename: $HOME/aaec/aaec.log

unix_bind: $HOME/aaec/aaec.sock
```

Then we can run aaec server in Termux as a daemon.

```
aaecd --core-config aaec.yaml --app-config app.yaml
```

## client

`aaectl` is able to accept multiple event in a row and read event body from file using `@` prefix.

```
aaectl event create --pkg com.tencent.mm --type bg
```

Instead of send event, we can send instruction directly.

```
aaectl inst freeze --pkg com.tencent.mm --extra delay=10s
```

## installation

1. `GO111MODULE=on go get -u github.com/jschwinger23/aaec`
2. create Tasker profile for __App Changed Changed__`
