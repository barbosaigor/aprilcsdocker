#!/usr/bin/env bash

docker run -d --name payment alpine tail -f /dev/null
docker run -d --name fees alpine tail -f /dev/null
docker run -d --name profile alpine tail -f /dev/null
docker run -d --name inventory alpine tail -f /dev/null
docker run -d --name shipping alpine tail -f /dev/null
docker run -d --name storefront alpine tail -f /dev/null