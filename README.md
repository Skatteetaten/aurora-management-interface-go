# Aurora Management Interface Go

This library is a utility for fulfilling the need for a "management interface". The API is designed to be compatible with 
management software that use the Spring Boot actuator API for management and monitoring. It is a lightweight framework with 
some default implementations of endpoints, but it is also easily accepts customized endpoint implementations for applications.

## Supported Endpoints

Currently these endpoints are supported. The endpoint paths can be overridden for custom implementations. See API for details.

- **management**, default path: `/management` - provides a JSON map of the available endpoints
- **health**, default path: `/health` - performs health check(s) and returns a JSON of the status
- **env**, default path: `/env` - returns a JSON map of available environment variables
- **info**, default path: `/info` - can show links to infrastructure, dependencies and git/build information. No default 
implementation is available at this time. Needs custom implementation for the application.

## Getting started 

Have a look at the [example](_example/example.go) included in this repo.

## API

// TODO 

