# fdjapi-lotto

FDJ API for the lotto

## TODO

* This package use a tempory http client which be use in production mode. But in the future, the client http will be move to gofast-pkg organization with a better implementation
and configuration.

* Replace fmt.Printf / ln by a logger (from gofast-pkg) or by an information struct like Warning of csvparser

* Move csvParser to gofast-pkg
