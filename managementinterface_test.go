package management

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewManagementinterfaceSpec(t *testing.T) {
	t.Run("Should create ManagementinterfaceSpec", func(t *testing.T) {

		managementinterfaceSpec := NewManagementinterfaceSpec()

		assert.NotNil(t, managementinterfaceSpec)
		assert.NotNil(t, managementinterfaceSpec.endpoints)
		assert.NotNil(t, managementinterfaceSpec.managementEndpoint)
	})
}

func TestMapEndpoint(t *testing.T) {
	t.Run("Should create map endpoints properly", func(t *testing.T) {
		managementinterfaceSpec := NewManagementinterfaceSpec()

		// Mapping of managementEndpoint
		managementEndpoint, err := newEndPoint(Management, "localhost", "8081", func(w http.ResponseWriter, r *http.Request) {})
		assert.Nil(t, err)
		managementinterfaceSpec.mapEndpoint(*managementEndpoint)
		assert.Equal(t, managementEndpoint.endpointid, managementinterfaceSpec.managementEndpoint.endpointid)
		assert.Equal(t, managementEndpoint.getEndpointURL(), managementinterfaceSpec.managementEndpoint.getEndpointURL())

		// Mapping of health and env endpoints
		healthEndpoint, err := newEndPoint(Health, "localhost", "8081", func(w http.ResponseWriter, r *http.Request) {})
		assert.Nil(t, err)
		managementinterfaceSpec.mapEndpoint(*healthEndpoint)
		assert.Equal(t, healthEndpoint.endpointid, managementinterfaceSpec.endpoints[Health].endpointid)
		assert.Equal(t, healthEndpoint.getEndpointURL(), managementinterfaceSpec.endpoints[Health].getEndpointURL())

		envEndpoint, err := newEndPoint(Env, "localhost", "8081", func(w http.ResponseWriter, r *http.Request) {})
		assert.Nil(t, err)
		managementinterfaceSpec.mapEndpoint(*envEndpoint)
		assert.Equal(t, envEndpoint.endpointid, managementinterfaceSpec.endpoints[Env].endpointid)
		assert.Equal(t, envEndpoint.getEndpointURL(), managementinterfaceSpec.endpoints[Env].getEndpointURL())

	})
}

func TestCreateManagementJSON(t *testing.T) {
	t.Run("Should create map endpoints properly", func(t *testing.T) {
		managementinterfaceSpec := NewManagementinterfaceSpec()
		managementEndpoint, _ := newEndPoint(Management, "localhost", "8081", func(w http.ResponseWriter, r *http.Request) {})
		managementinterfaceSpec.mapEndpoint(*managementEndpoint)
		healthEndpoint, _ := newEndPoint(Health, "localhost", "8081", func(w http.ResponseWriter, r *http.Request) {})
		managementinterfaceSpec.mapEndpoint(*healthEndpoint)
		envEndpoint, _ := newEndPoint(Env, "localhost", "8081", func(w http.ResponseWriter, r *http.Request) {})
		managementinterfaceSpec.mapEndpoint(*envEndpoint)

		managementJSON, err := managementinterfaceSpec.createManagementJSON()
		assert.Nil(t, err)
		assert.Equal(t, "{\"_links\":{\"env\":{\"href\":\"http://localhost:8081/env\"},\"health\":{\"href\":\"http://localhost:8081/health\"}}}", fmt.Sprintf("%s", managementJSON))
	})
}
