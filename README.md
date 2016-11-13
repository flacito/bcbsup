# Background

Simple experiment in a few things:

1. my first time programming in golang
2. using golang to parse and use JSON (something I'm brewing for work)
3. wrapping it all in a Docker container so no golang runtime is needed

# Howto

You can build and run the Docker containers by hand if you want, but for your convenience these are included in build and run shell scripts.

- `build.sh`: creates the Docker image with the create machine command in it
- `bcbsup.sh machine create {JSON object}`: runs the Docker image with the required JSON object as the one and only argument.

The JSON you send must mach what you see here. You can put whatever you want for the values, but you must include HostName, MemorySizeGB, and CpuCount. For example,

```
./bcbsup.sh machine create '{"HostName":"hallo-tapco","MemorySizeGB":2,"CpuCount":2}'`
```
