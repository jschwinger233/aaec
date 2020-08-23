import os
import time
import click
import signal
import atexit
import subprocess

import aaec.pidfile as pidfile
import aaec.subscribe as subscribe


@click.group(context_settings=dict(help_option_names=['-h', '--help']))
def cli():
    def term_handler(_, __):
        raise SystemError

    signal.signal(signal.SIGTERM, term_handler)

    subprocess.run('termux-wake-lock')
    atexit.register(lambda: subprocess.run('termux-wake-unlock'))


@cli.command()
@click.argument('package')
def bg(package: str):
    if not (delay := subscribe.check(package)):
        return

    pidfile.write(package)
    time.sleep(delay)
    subprocess.run(f'sudo pm disable {package}'.split())


@cli.command()
@click.argument('package')
def fg(package: str):
    if not subscribe.check(package):
        return

    pid = pidfile.read(package)
    if not pid:
        return

    os.kill(pid, 15)


@cli.command()
@click.argument('package')
def sub(package: str):
    subscribe.sub(package)


@cli.command()
@click.argument('package')
def unsub(package: str):
    subscribe.unsub(package)


@cli.command()
@click.argument('package')
@click.argument('incr', type=int)
def prolong(package: str, incr: int):
    if not subscribe.check(package):
        return

    subscribe.prolong(package, incr)
