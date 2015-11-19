# Nested Service

This is a very simple HTTP service for me to play around with `Go` and `Consul`.  Some of my first Go code so please don't grade it ;)

## Installation

The `build.sh` script will pull down a golang container to compile the service.  `docker-compose` will handle creating the service container from the `Dockerfile`

```
$ ./build.sh
$ docker-compose up
```

## Usage

After you finish building you query it using the names of the service as the path.  Feel free to add/change services in the `docker-compose.yml`.  View the examples below to see how exactly it works.

## Contribution

Bug reports and pull requests are welcome on GitHub at https://github.com/zacheryph/nested-service

## Examples

```javascript
// curl http://local.docker:3000/slave/worker/terminate
{
  "hostname": "cf22dfad250e",
  "service_name": "master",
  "error": "",
  "messages": [
    {
      "hostname": "6d4901bcfeea",
      "service_name": "slave",
      "error": "",
      "messages": [
        {
          "hostname": "f778ff8568b6",
          "service_name": "worker",
          "error": "",
          "messages": [
            {
              "hostname": "",
              "service_name": "",
              "error": "Get http://terminate.service.consul:3000/: dial tcp: lookup terminate.service.consul on 172.17.0.1:53: server misbehaving",
              "messages": null
            }
          ]
        }
      ]
    }
  ]
}

// docker-compose scale slave=3 (or more)
// curl http://local.docker:3000/slave/slave/slave
{
  "hostname": "cf22dfad250e",
  "service_name": "master",
  "error": "",
  "messages": [
    {
      "hostname": "6d4901bcfeea",
      "service_name": "slave",
      "error": "",
      "messages": [
        {
          "hostname": "6d4901bcfeea",
          "service_name": "slave",
          "error": "",
          "messages": [
            {
              "hostname": "a33ea7fed264",
              "service_name": "slave",
              "error": "",
              "messages": []
            }
          ]
        }
      ]
    }
  ]
}
```
