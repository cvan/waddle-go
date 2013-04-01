waddle
======
Don't byte off more than you can chew.


Installation
------------

Go on and install `go`:

    brew install go

To make our lives easier we rely on [`autoenv`](https://github.com/kennethreitz/autoenv)
for managing our `go` environment variables. Go on and install that:

    brew install autoenv
    echo 'source /usr/local/Cellar/autoenv/*/activate.sh' >> ~/.zshrc
    . ~/.zshrc


heifer
------
[Heifer](https://github.com/potch/heifer) calculates your site's weight. Go and compile that:

    go build heifer && ./heifer
