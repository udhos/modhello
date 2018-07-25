# modhello

Define module modlib.

    $ cat modlib/go.mod
    module github.com/udhos/modhello/modlib

Define module modapp. Notice that modapp requires modlib v1.0.0.

    $ cat modapp/go.mod
    module github.com/udhos/modhello/modapp

    require github.com/udhos/modhello/modlib v1.0.0

Publish version v1.0.0 for modlib.

    $ git tag modlib/v1.0.0
    $ git tag
    modlib/v1.0.0
    $ git push origin modlib/v1.0.0 ;# publish tag

