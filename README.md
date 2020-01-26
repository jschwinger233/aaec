# Android Applications Ethics Commission

To regulate dishonest and unethical practices by installed applications.

## server

Application configure defines `system events` and app specific action.

```
# app.ayml

system_events:
- bg:
    define:
    - type: timer
      content:
        action: start
        duration: 10s
    - type: adb
      content:
        command: 'pm disable {{.Package}}'
    apply:
    - com.tencent.mm
    - com.gotokeep.keep.intl

apps:
  com.tencent.mm:
    fg:
    - type: adb
      content:
        command: 'appops set {{.Package}} READ_SMS allow'
    bg:
    - type: adb
      content:
        command: 'appops set {{.Package}} READ_SMS ignore'
    - super
```

Core configure shows the general options:

```
# aaec.yaml

pidfile: $HOME/aaec.pid
log_level: debug
log_filename: $HOME/aaec.log
```

Then we can run aaec server in Termux as a daemon.

```
aaecd --core-config aaec.yaml --app-config app.yaml
```

## client

`aaectl` is able to accept multiple event in a row and read event body from file using `@` prefix.

```
aaectl --event-type type1,type2,type3 --event-body '{}' --event-body '{}' --event-body @event.body
```

## installation

1. `GO111MODULE=on go get -u github.com/jschwinger23/aaec`
2. create Tasker profile for __App Changed Changed__`
