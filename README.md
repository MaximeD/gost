# Gost

*The gist client in go*

Gost is a client for the gist API,
it allows you to create and browse the gists of [https://gist.github.com/](https://gist.github.com/).

# Installation

Assuming you have golang:

```
go get github.com/MaximeD/gost
```

# Usage

## Create gists

```
gost <files name>
```

You can supply as many files as you want here.


Full example:

```
gost -p -d "rocking the casbah" fun.go love_gist.rb
```

will create a private gist with description "rocking the casbah"
with two files `fun.go` and `love_gist.rb`


### Private gists

By default, your gists are public.
If you want a private one, just add `-p` to your command.
On first run, `gost` will create a prompt you for your credentials
and store a token in it's configuration file,
so that you will have to do this step only once.

### Description

To give your gist a description use `-d <description>`


## List gists

You can list the gist of a user with the following:

```
gost -l <username>
```

## Download a gist

```
gost -g <gist_id>
```

## Delete a gist

Given you have the id of one of your gist and the right to delete it, use `-D` flag:

```
gost -D <gist_id>
```
