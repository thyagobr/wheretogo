package clients

import (
	"encoding/json"
	"net/http"
	"net/url"
	"fmt"

	"github.com/thyagobr/wheretogo/internal/dtos"
)

type OpenMapsClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

type OpenMapsResponse struct {
	DisplayName string `json:"display_name"`
}

func NewOpenMapsClient() *OpenMapsClient {
	return &OpenMapsClient{
		BaseURL:    "https://nominatim.openstreetmap.org",
		HTTPClient: &http.Client{},
	}
}

func (client *OpenMapsClient) SearchAddress(req dtos.SearchAddressRequest) ([]OpenMapsResponse, error) {
	reqURL, err := url.Parse(client.BaseURL + "/search")
	if err != nil {
		return nil, err
	}

	// Join name city and country together in a string separated by comma
	queryParams := fmt.Sprintf("%s,%s,%s", req.Name, req.City, req.Country)
	q := reqURL.Query()
	q.Set("q", queryParams)
	q.Set("format", "json")
	q.Set("limit", fmt.Sprintf("%d", req.Limit))
	q.Set("addressdetails", "1")
	reqURL.RawQuery = q.Encode()

	httpReq, err := http.NewRequest(http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("User-Agent", "WhereToGoApp/1.0")
	httpReq.Header.Set("Accept", "application/json")

	resp, err := client.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	var openMapsResp []OpenMapsResponse
	err = json.NewDecoder(resp.Body).Decode(&openMapsResp)
	if err != nil {
		return nil, err
	}
	if len(openMapsResp) == 0 {
		return make([]OpenMapsResponse, 0), nil
	} else {
		return openMapsResp, nil
	}
}
