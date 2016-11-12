#!/bin/bash
docker build -t create-machine-command .
docker run -it --rm --name create-machine-command create-machine-command create_machine $1
