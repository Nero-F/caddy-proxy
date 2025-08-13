#!/bin/bash

docker-gen -config docker-gen.cfg
caddy start
