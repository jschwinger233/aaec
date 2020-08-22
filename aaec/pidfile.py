import os
import stat
import atexit

PIDFILE_DIR = os.getenv('PIDFILE_DIR', default=os.path.expanduser('~/aaec/'))


def write(ident: str):
    pidfile = _build_pidfile(ident)
    fd = os.open(
        pidfile, os.O_CREAT | os.O_EXCL | os.O_RDWR,
        stat.S_IRUSR | stat.S_IWUSR
    )
    atexit.register(lambda: os.remove(pidfile))
    os.write(fd, b'%d' % os.getpid())


def read(ident: str) -> int:
    pidfile = _build_pidfile(ident)
    with open(pidfile) as f:
        return int(f.read())


def _build_pidfile(ident: str) -> str:
    return os.path.join(PIDFILE_DIR, f'{ident}.pid')
