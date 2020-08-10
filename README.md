# Android Applications Ethics Commission

To regulate dishonest and unethical practices by installed applications.

Introduction: https://jschwinger23.github.io/#2020-08-11%20Android%20Applications%20Ethics%20Commission

# Make it easy

## `aaec sub $package`

All subscribed packages would be freezed(`pm disable $package`) once being put background more than 60 seconds.

## `aaec unsub $package`

Unsubscribe packages got alive after being backgrounded.

## `aaec bg $package`

Notify specific package changes status from fg to bg.

## `aaec fg $package`

Notify specific package changes status from bg to fg.
