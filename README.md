# Android Applications Ethics Commission

To regulate dishonest and unethical practices by installed applications.

## basic usage

```
apps:
  - name: com.tencent.mm:
    on_events:
      - kind: background
        callbacks:
          - command: pm disable com.tencent.mm
            delay: 10s
```

Then WeChat will be freezed in 10s after being put into background, unless it's put foreground within 10s.

## installation

```
GO111MODULE=on go get -u github.com/jschwinger23/aaec
```

then copy binary `aaec` under `/data/data/com.termux/files/home/.termux/tasker/`, and set Tasker profile for triggering __App Changed Event__ (Tasker>=5.8 required), and voila!
