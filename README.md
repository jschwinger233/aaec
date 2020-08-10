# Android Applications Ethics Commission

To regulate dishonest and unethical practices by installed applications.

# Make it easy

## `aaec sub $package`

all subscribed packages would be freezed(`pm disable $package`) once being put background more than 60 seconds.

## `aaec unsub $package`

unsubscribe packages got alive after being backgrounded.

## `aaec bg $package`

notify specific package changes status from fg to bg.

## `aaec fg $package`

notify specific package changes status from bg to fg.
