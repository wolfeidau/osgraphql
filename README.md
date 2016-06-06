# osgraphql

Simple service which uses graphql to query process and system information.

# usage

```
$ ./bin/osgraphql --help
usage: osgraphql [<flags>]

Flags:
      --help       Show context-sensitive help (also try --help-long and --help-man).
  -p, --port=9000  Listen port.

```

# query

Below is an example graphql query to retrieve cpu and partition information.

```
{
  cpus {
    cpu
    vendorId
    family
  }
  partitions {
    path
    fstype
    total
    free
    usedPercent
  }
}
```

Result.

```
{
  "data": {
    "cpus": [
      {
        "cpu": "0",
        "family": "6",
        "vendorId": "GenuineIntel"
      }
    ],
    "partitions": [
      {
        "free": "84828852224",
        "fstype": "hfs",
        "path": "/",
        "total": "249779191808",
        "usedPercent": 65
      },
      {
        "free": "0",
        "fstype": "devfs",
        "path": "/dev",
        "total": "354304",
        "usedPercent": 100
      },
      {
        "free": "0",
        "fstype": "autofs",
        "path": "/net",
        "total": "0",
        "usedPercent": null
      },
      {
        "free": "0",
        "fstype": "autofs",
        "path": "/home",
        "total": "0",
        "usedPercent": null
      }
    ]
  }
}
```

# references

* [graphql-go/graphql](https://github.com/graphql-go/graphql) An implementation of GraphQL for Go.

# disclaimer

This was an experimental thing I made while away for a weekend.

# license

Copyright (c) 2016 Mark Wolfe. Released under MIT license, see LICENSE.md file.
