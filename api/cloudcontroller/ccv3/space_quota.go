package ccv3

import (
	"bytes"
	"encoding/json"

	"code.cloudfoundry.org/cli/api/cloudcontroller"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/internal"
)

type SpaceQuota struct {
	Quota
	// OrgGUID is the unique ID of the owning organization
	OrgGUID string
	// SpaceGUIDs are the list of unique ID's of the associated spaces
	SpaceGUIDs []string
}

func (sq SpaceQuota) MarshalJSON() ([]byte, error) {
	appLimits := map[string]interface{}{
		"total_memory_in_mb":       sq.Apps.TotalMemory,
		"per_process_memory_in_mb": sq.Apps.InstanceMemory,
		"total_instances":          sq.Apps.TotalAppInstances,
	}

	serviceLimits := map[string]interface{}{
		"paid_services_allowed":   sq.Services.PaidServicePlans,
		"total_service_instances": sq.Services.TotalServiceInstances,
	}

	routeLimits := map[string]interface{}{
		"total_routes":         sq.Routes.TotalRoutes,
		"total_reserved_ports": sq.Routes.TotalReservedPorts,
	}

	relationships := map[string]interface{}{
		"organization": map[string]interface{}{
			"data": map[string]interface{}{
				"guid": sq.OrgGUID,
			},
		},
	}

	if len(sq.SpaceGUIDs) > 0 {
		spaceData := make([]map[string]interface{}, len(sq.SpaceGUIDs))
		for i, spaceGUID := range sq.SpaceGUIDs {
			spaceData[i] = map[string]interface{}{
				"guid": spaceGUID,
			}
		}

		relationships["spaces"] = map[string]interface{}{
			"data": spaceData,
		}
	}

	jsonMap := map[string]interface{}{
		"name":          sq.Name,
		"apps":          appLimits,
		"services":      serviceLimits,
		"routes":        routeLimits,
		"relationships": relationships,
	}

	return json.Marshal(jsonMap)
}

func (sq *SpaceQuota) UnmarshalJSON(data []byte) error {
	type alias SpaceQuota
	var defaultUnmarshalledSpaceQuota alias
	err := json.Unmarshal(data, &defaultUnmarshalledSpaceQuota)
	if err != nil {
		return err
	}

	*sq = SpaceQuota(defaultUnmarshalledSpaceQuota)

	type RemainingFieldsStruct struct {
		Relationships struct {
			Organization struct {
				Data struct {
					Guid string
				}
			}
			Spaces struct {
				Data []struct {
					Guid string
				}
			}
		}
	}

	var remainingFields RemainingFieldsStruct
	err = json.Unmarshal(data, &remainingFields)
	if err != nil {
		return err
	}

	sq.OrgGUID = remainingFields.Relationships.Organization.Data.Guid

	for _, spaceData := range remainingFields.Relationships.Spaces.Data {
		sq.SpaceGUIDs = append(sq.SpaceGUIDs, spaceData.Guid)
	}

	return nil
}

func (client Client) CreateSpaceQuota(spaceQuota SpaceQuota) (SpaceQuota, Warnings, error) {
	spaceQuotaBytes, err := json.Marshal(spaceQuota)

	if err != nil {
		return SpaceQuota{}, nil, err
	}

	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.PostSpaceQuotaRequest,
		Body:        bytes.NewReader(spaceQuotaBytes),
	})

	if err != nil {
		return SpaceQuota{}, nil, err
	}

	var createdSpaceQuota SpaceQuota
	response := cloudcontroller.Response{
		DecodeJSONResponseInto: &createdSpaceQuota,
	}

	err = client.connection.Make(request, &response)
	if err != nil {
		return SpaceQuota{}, response.Warnings, err
	}

	return createdSpaceQuota, response.Warnings, err
}

func (client *Client) GetSpaceQuotas(query ...Query) ([]SpaceQuota, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetSpaceQuotasRequest,
		Query:       query,
	})
	if err != nil {
		return nil, nil, err
	}

	var spaceQuotasList []SpaceQuota
	warnings, err := client.paginate(request, SpaceQuota{}, func(item interface{}) error {
		if spaceQuota, ok := item.(SpaceQuota); ok {
			spaceQuotasList = append(spaceQuotasList, spaceQuota)
		} else {
			return ccerror.UnknownObjectInListError{
				Expected:   SpaceQuota{},
				Unexpected: item,
			}
		}
		return nil
	})

	return spaceQuotasList, warnings, err
}