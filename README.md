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
    $ git push origin modlib/v1.0.0 ;# publish tag

Publish new version v2.0.0 for modlib.

    # Define modlib/v2/go.mod
    $ cat modlib/v2/go.mod
    module github.com/udhos/modhello/modlib/v2
    $ git add modlib/v2/go.mod
    $ git commit
    $ git tag modlib/v2.0.0
    $ git push --all
    $ git push origin modlib/v1.0.0 ;# publish tag

Create new module/app modhot depending on modlib v2.0.0.

    $ cat modhot/go.mod
    module github.com/udhos/modhello/modhot

    require github.com/udhos/modhello/modlib v2.0.0

