import os
import fcntl
import pickle
import contextlib

PICKLE_FILENAME = os.getenv(
    'PICKLE_FILENAME', default=os.path.expanduser('~/aaec/sub.pkl')
)
FLOCK_FILENAME = os.getenv(
    'FLOCK_FILENAME', default=os.path.expanduser('~/aaec/sub.flock')
)


def sub(package: str):
    with load_pickle() as subscribed:
        subscribed.add(package)
        subscribed[package] = 10


def unsub(package: str):
    with load_pickle() as subscribed:
        with contextlib.suppress(KeyError):
            del subscribed[package]


def check(package: str) -> int:
    with load_pickle() as subscribed:
        return subscribed[package]


def prolong(package: str, incr: int):
    with load_pickle() as subscribed:
        subscribed[package] += incr
        if subscribed[package] < 0:
            subscribed[package] = 0


@contextlib.contextmanager
def load_pickle():
    with open(FLOCK_FILENAME) as lock:
        fcntl.flock(lock.fileno(), fcntl.LOCK_EX)
        with open(PICKLE_FILENAME, 'r+b') as f:
            try:
                subscribed = pickle.load(f)
            except EOFError:
                subscribed = set()
            yield subscribed
            f.seek(0)
            pickle.dump(subscribed, f)
        fcntl.flock(lock.fileno(), fcntl.LOCK_UN)
