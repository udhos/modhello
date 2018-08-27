# modhello

Recipe based on https://github.com/golang/go/wiki/Modules

    Repo:    modhello
    Module:  modhello/modlib     -- go.mod lives here
    Package: modhello/modlib/lib -- this is imported from go code

## Highlights

- Host multiple modules under single repo.
- Host multiple packages under single module.
- Modules depend on full modules (not individual packages).
- Go code imports packages (not modules).
- import "domain.com/repo/module/package" for package (v0 or v1).
- import "domain.com/repo/module/v2/package" for package v2.
- Two ways to release v2 or higher:
  1. Update the go.mod file to include a /v2 at the end of the module path.
  2. Alternatively, create a v2 directory with a new go.mod file defining the module path ending with /v2.
- Tag the repository with <modulename>/vX.X.X to publish new version for specific module.

## How to add intial Go Modules support for a repository

    git clone https://.../repo ;# must clone outside of GOPATH
    cd repo
    go mod init
    go mod tidy
    git add go.mod go.sum
    git commit
    git push

## Recipe

Define module modlib.

    $ cat modhello/modlib/go.mod
    module github.com/udhos/modhello/modlib

Define module modapp. Notice that modapp requires modlib v1.0.0.

    $ cat modhello/modapp/go.mod
    module github.com/udhos/modhello/modapp

    require github.com/udhos/modhello/modlib v1.0.0

Publish version v1.0.0 for modlib.

    $ cd modhello
    $ git tag modlib/v1.0.0
    $ git push origin modlib/v1.0.0 ;# publish tag

Build the app.

    $ cd modhello/modapp
    $ go install

Publish new version v2.0.0 for modlib.

    $ cd modhello
    # Define modlib/v2/go.mod
    $ cat modlib/v2/go.mod
    module github.com/udhos/modhello/modlib/v2
    $ git add modlib/v2/go.mod
    $ git commit
    $ git tag modlib/v2.0.0
    $ git push --all
    $ git push origin modlib/v2.0.0 ;# publish tag

Create new app modhot depending on modlib v2.0.0.

    $ cat modhello/modhot/go.mod
    module github.com/udhos/modhello/modhot

    require github.com/udhos/modhello/modlib/v2 v2.0.0

Build new app.

    $ cd modhello/modhot
    $ go install

## How to update Modules used by your application

### Non-major update

By default, Go will not update modules without being asked.

    $ go get -u                                ;# use the latest minor or patch releases
    $ go get -u=patch                          ;# use the latest patch releases (to 1.0.1 but not to 1.1.0)
    $ go get github.com/user/testmodule@v1.0.1 ;# update to a specific version

### Major update

The major version is for all intents and purposes a completely different package.

    import "github.com/user/testmodule/v2/pkg"

