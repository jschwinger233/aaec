import os
import pickle
import contextlib

PICKLE_FILENAME = os.getenv(
    'PICKLE_FILENAME', default=os.path.expanduser('~/aaec/sub.pkl')
)


def sub(package: str):
    with load_pickle() as subscribed:
        subscribed.add(package)


def unsub(package: str):
    with load_pickle() as subscribed:
        with contextlib.suppress(KeyError):
            subscribed.remove(package)


def check(package: str) -> bool:
    with load_pickle() as subscribed:
        return package in subscribed


@contextlib.contextmanager
def load_pickle():
    with open(PICKLE_FILENAME, 'r+b') as f:
        try:
            subscribed = pickle.load(f)
        except EOFError:
            subscribed = set()
        yield subscribed
        f.seek(0)
        pickle.dump(subscribed, f)
