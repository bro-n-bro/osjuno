# osjuno - Juno impl for Osmosis network

Swaps and Pools data index.

If you have any problems with or questions about this software, please contact us through
a [GitHub issue](/../../issues). You are invited to
contribute [new features, fixes, or updates](/../../issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22), large or
small; We are always thrilled to receive pull requests, and do our best to process them as fast as we can.

## Postgres + Hasura + osjuno installation using docker-compose

Adjust settings using [.env file](./.env).`Run deploy script:

```bash
 docker-compose up --build
```

To delete all installation data run next command:

```bash
 docker-compose down -v
```

## Useful links

* [BDjuno repo](https://github.com/forbole/bdjuno)
* [Juno repo](https://github.com/RiccardoM/juno)