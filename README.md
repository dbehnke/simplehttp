# golang-gb-project-template

A template for starting a Go project using gb (http://getgb.io)

## Quickstart

1. Clone this repository
2. Execute: make

If successful, You should have a helloworld in the bin directory.

## Making it your own

Simply remove the sample helloworld (or keep it, it won't harm anything).

Add your project code under src and use "gb vendor" to vendor your dependencies.

The Makefile and use of make are optional. It is mostly for bootstrapping and does checking if go is installed, etc.  If you have gb installed already, you can simply do a "gb build".   If you are using make, make sure to modify Makefile to suit your needs.

I suggest reading the gb documentation at http://getgb.io as well.

## Activating Project Environment

If you are familiar with python virtual environments, this is similar.

There is a file named activate_env that can be sourced.

Activation

        $ source ./activate_env

Deactivate

        $ deactivate


## Making gocode happy

If you are using any IDEs / Text Editors that user gocode you must have lib-path configured.  Gocode won't see your vendored libraries or your source code in the src directory otherwise.  The activation script in the previous section will do this for you.   If you want to manually set the path use the following command.

                gocode set lib-path "$PWD/pkg/$GOOS-$GOARCH"
