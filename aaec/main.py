import os
import time
import click
import subprocess

import aaec.pidfile as pidfile
import aaec.subscribe as subscribe


@click.group(context_settings=dict(help_option_names=['-h', '--help']))
def cli():
    pass


@cli.command()
@click.argument('package')
def bg(package: str):
    if not subscribe.check(package):
        return

    pidfile.write(package)
    time.sleep(60)
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
