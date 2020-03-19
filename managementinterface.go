package management

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

// managementinterfaceSpec holds information on all mapped endpoints
type managementinterfaceSpec struct {
	managementEndpoint endpoint
	endpoints          map[EndPointType]endpoint
}

// newManagementinterfaceSpec creates a new managementinterfaceSpec with a default management endpoint
func newManagementinterfaceSpec() *managementinterfaceSpec {
	mispec := &managementinterfaceSpec{
		endpoints: make(map[EndPointType]endpoint),
	}
	return mispec
}

func (mispec *managementinterfaceSpec) mapEndpoint(endpoint endpoint) {
	if endpoint.endpointid == Management {
		mispec.managementEndpoint = endpoint
	} else {
		mispec.endpoints[endpoint.endpointid] = endpoint
	}
}

func (mispec *managementinterfaceSpec) createManagementJSON(host string) ([]byte, error) {
	managementLinks := getLinkMapStructureForEndpoints(mispec.endpoints, host)
	managementLinksJSON, err := json.Marshal(managementLinks)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("Management template output:\n%s\n", string(managementLinksJSON))

	return managementLinksJSON, nil
}

func getLinkMapStructureForEndpoints(endpoints map[EndPointType]endpoint, host string) map[string]interface{} {
	endpointMap := make(map[string]interface{})
	for endpointid, endpoint := range endpoints {
		hrefMap := map[string]string{"href": endpoint.getEndpointURL(host)}
		endpointMap[string(endpointid)] = hrefMap
	}

	managementLinks := make(map[string]interface{})
	managementLinks["_links"] = endpointMap

	return managementLinks
}
