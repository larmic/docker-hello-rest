# docker-hello-rest

[![Docker build and push](https://github.com/larmic/docker-hello-rest/actions/workflows/docker-build-and-push.yml/badge.svg)](https://github.com/larmic/docker-hello-rest/actions/workflows/docker-build-and-push.yml)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Tiny docker image provides an `/hello` endpoints supported on all architectures.

## Start

````shell
docker run -ti --rm -p 8080:8080 larmic/docker-hello-rest
````

## Example

```shell
$ curl localhost:8080/hello
```