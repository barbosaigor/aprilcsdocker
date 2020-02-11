#!/usr/bin/env bash

docker stop payment
docker rm payment
docker stop fees
docker rm fees
docker stop profile
docker rm profile
docker stop inventory
docker rm inventory
docker stop shipping
docker rm shipping
docker stop storefront
docker rm storefront