# cqlc Example Application

This is an example application to demonstrate how `cqlc` can be used in a Go application. It is very basic and is not is any way an exhaustive showcase of the runtime API. It is just designed to show you how to use the `cqlc` tool with a real app.

`cqlc` is a source code generation tool that reads a schema from a given keyspace in Cassandra and produces strongly typed binding code for use in a Go application.

This example consists of:

* An example CQL schema for Cassandra;
* A Makefile that applies this schema to a local Cassandra instance and invokes `cqlc` to generate Go source code bindings for this schema;
* A simple main method that shows how to read/write to/from Cassandra using the generated bindings;

# Pre-requisites

* Go (only 1.2 has been tested, but other versions may work as well);
* Apache Cassandra (2.0.3 has been tested, but other versions will probably work too);
* `cqlc` needs to be available on your `$PATH`;
* (Optionally) `cqlsh` needs to be on your `$PATH` (if you use the supplied Makefile)

## Getting cqlc

To install `cqlc`:

    $ go get github.com/relops/cqlc

After installing this, make sure that the `bin` directory of your `$GOPATH` is in in your `$PATH`, so that the `cqlc` executable is available in your shell. 


# Generating the binding code

Running

    $ make schema

Will apply the schema from the `schema.cql` to a local instance of Cassandra. By default, the Makefile assumes that the Cassandra instance it should connect to is running on 127.0.0.1. The default keyspace is called `twitter_example` - this will be created if it not exist.

# Generating the binding code

By default, the generated code is written to `src/twitter/generated_schema.go` by running the default target from the Makefile:

    $ make default
      cqlc -i 127.0.0.1 -k twitter_example -p twitter -o src/twitter/generated_schema.go
      Reading schema from keyspace: twitter_example

# Running main.go

Once you have generated the binding code, make sure that it is on your local `$GOPATH`. For example, in the current directory, run:

    $ export GOPATH=$GOPATH:`pwd`

such that `pwd`/src/generated is on the `$GOPATH`.

Now you should be able to run:

    $ go run main.go 
    2014/01/16 18:06:29 Got tweets: [{Id:1 Name:tweeter}]
