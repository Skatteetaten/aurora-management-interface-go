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

### Discovery endpoint - Management

  The management interface has a discovery endpoint that provides the endpoints for health and env. Use this to discover 
  the other endpoints.

* **URL**

  /management (default)

* **Method:**

  `GET`

*  **URL Params**

  None

* **Authorization**

  None

* **Success Response:**

  Returns a JSON structure of the configured management interface endpoints.

  * **Code:** 200 OK <br />
    **Content:** `{"_links":{"env":{"href":"http://localhost:8081/env"},"health":{"href":"http://localhost:8081/health"}}}`

* **Sample Call:**

```
  curl -H 'Content-Type: application/json' http://localhost:8081/management`
```

### Health endpoint

  The health endpoint provides a way to perform a health check. The returned JSON structure supports both deep and shallow 
  health checks. There is also a very shallow default health check implementation, returning simply a status of UP.

* **URL**

  /health (default)

* **Method:**

  `GET`

*  **URL Params**

  None

* **Authorization**

  None

* **Success Response:**

  Returns a simple JSON structure with status "UP"

  * **Code:** 200 OK <br />
    **Content:** `{"status":"UP"}`

* **The available status responses**

	- "UP" means the application is working properly. Http statuscode is 200, OK
	- "DOWN" means the application is down or (for deep health checks) that a critical 
	component has status "DOWN". Http statuscode is 503, Service Unavailable
	- "OUT_OF_SERVICE" means the application has been taken out of service.  Http statuscode is 503, Service Unavailable
	- "UNKNOWN" means that for some reason the status is unknown. Http statuscode is 200, OK
	- "OBSERVE" means that the application is working, but something is not quite the way it should. Examples may 
	be slow responses or (for deep health checks) that some some component has a status of "OBSERVE". 
	Http statuscode is 200, OK.

* **Sample Call:**

```
  curl -H 'Content-Type: application/json' http://localhost:8081/health`
```

### Env endpoint

  The env endpoint provides maps of environment variables.  There is a default implementation of this, listing the variables 
  system environment variables available to the application. There is also functionality for masking the 
  value of variables with secrets like access credentials.
  
* **URL**

  /env (default)

* **Method:**

  `GET`

*  **URL Params**

  None

* **Authorization**

  None

* **Success Response:**

  Returns a JSON structure with a map of environment variables

  * **Code:** 200 OK <br />
    **Content:** `{"activeProfiles":[],"propertySources":[{"name":"systemEnvironment","properties":{"DBUS_SESSION_BUS_ADDRESS":{"value":"unix:path"},"DESKTOP_SESSION":{"value":"gnome-flashback-metacity"},"DISPLAY":{"value":":0"},"GNOME_SESSION_XDG_SESSION_PATH":{"value":"/org/freedesktop/DisplayManager/Session0"},"IM_CONFIG_PHASE":{"value":"1"},"LANG":{"value":"en_US.UTF-8"},"LC_ADDRESS":{"value":"nb_NO.UTF-8"},"LC_IDENTIFICATION":{"value":"nb_NO.UTF-8"},"LC_MEASUREMENT":{"value":"nb_NO.UTF-8"},"LC_MONETARY":{"value":"nb_NO.UTF-8"},"LC_NAME":{"value":"nb_NO.UTF-8"},"LC_NUMERIC":{"value":"nb_NO.UTF-8"},"LC_PAPER":{"value":"nb_NO.UTF-8"},"LC_TELEPHONE":{"value":"nb_NO.UTF-8"},"LC_TIME":{"value":"nb_NO.UTF-8"},"QT_ACCESSIBILITY":{"value":"1"},"QT_IM_MODULE":{"value":"ibus"},"SHELL":{"value":"/bin/bash"},"SomeEnvKeyToMask":{"value":"***"}}}]}`

* **Sample Call:**

```
  curl -H 'Content-Type: application/json' http://localhost:8081/env`
```

