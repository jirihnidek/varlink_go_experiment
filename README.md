Experiment with varlink and Go
==============================

This is simple experiment with varlink and go that you can use on Linux.

Build
-----

Just run:

```
go build
```

Run and test
------------

First, run following service providing varlink interface on terminal 1:

```
sudo ./varlink_go_experiment
```

You can experiment with your varlink service on terminal 2.

Introspect interface:

```
sudo varlinkctl introspect /run/org.example.this org.example.this
```

should return this:

```
interface org.example.this

method Ping(
        in: string
) -> (
        out: string
)
```

You can call the `Pink()` method:

```
sudo varlinkctl call -j /run/org.example.this org.example.this.Ping '{"in": "hello"}'
```

and get echo response like this:

```json
{
        "out" : "hello"
}
```