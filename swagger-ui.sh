#!/bin/bash

docker run -d -p 8081:8080 -e SWAGGER_JSON=/foo/swagger.json -v "$PWD":/foo swaggerapi/swagger-ui
