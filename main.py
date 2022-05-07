from functools import partial
import click
import os
import shutil
import re

inject_old = re.compile(br"function validateString[^}]+}")
inject_new = br"function validateString(){};"


def check_prog(prog: str):
    if not os.access(prog, os.W_OK) or not os.access(os.path.join(os.path.dirname(prog),"../node_modules/"), os.W_OK):
        click.echo("Cannot write to program file: {}".format(prog), err=True)
        exit(1)


def backup_prog(prog: str):
    click.echo("Backuping main.node")
    if os.path.exists(prog + ".bak"):
        r = click.confirm("File already exists, overwrite?")
        if r:
            shutil.copyfile(prog, prog + ".bak")
    else:
        shutil.copyfile(prog, prog + ".bak")


def make_inject(path: str, matchobj: re.Match):
    inj = f"mod.require('{os.path.basename(path)}');"
    inj = inject_new + inj.encode()
    size = matchobj.end() - matchobj.start()
    if len(inj) > size:
        click.echo("Too long inject", err=True)
        exit(2)
    inj = inj.ljust(size, b" ")
    return inj


@click.command()
@click.option("-I", "--inject", default="inject/crack.js", help="Inject file", type=click.Path(exists=True))
@click.option("-B", "--from-bak", default=False, help="Use bak file as raw program", is_flag=True)
@click.argument("prog", default="main.node")
def main(prog: str, inject: str, from_bak: bool):
    if from_bak:
        _prog = prog + ".bak"
        if not os.path.exists(prog + ".bak"):
            shutil.copyfile(prog, prog + ".bak")
    else:
        _prog = prog
    
    check_prog(_prog)

    with open(_prog, "rb") as f:
        node = f.read()

    if not inject_old.search(node):
        click.echo(
            "Cannot find injection point in program file: {}".format(prog), err=True)
        exit(2)

    target = os.path.join(os.path.dirname(prog), "../node_modules/", os.path.basename(inject))
    shutil.copyfile(inject, target)

    if not from_bak:
        backup_prog(prog)

    click.echo(f"Injecting into {prog}")
    with open(prog, "wb") as f:
        f.write(inject_old.sub(partial(make_inject, target), node))
    click.echo("Done")


if __name__ == "__main__":
    main()
