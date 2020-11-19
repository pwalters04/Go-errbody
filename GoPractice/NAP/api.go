package NAP

import (
	"fmt"
)

type API struct {
	BaseURL       string
	Resources     map[string]*RestResource
	DefaultRouter *CBRouter
	Client        *Client
}

func NewAPI(baseUrl string) *API {
	return &API{
		BaseURL:       baseUrl,
		Resources:     make(map[string]*RestResource),
		DefaultRouter: NewRouter(),
		Client:        NewClient(),
	}
}
func (api *API) SetAuth(auth Authentication) {
	api.Client.SetAuth(auth)
}
func (api *API) AddResource(name string, res *RestResource) {
	api.Resources[name] = res
}
func (api *API) Call(name string, params map[string]string) error {
	res, ok := api.Resources[name]

	if !ok {
		return fmt.Errorf("Resource does not exist: %s", name)
	}

	if err := api.Client.ProcessRequest(api.BaseURL, res, params); err != nil {
		return err
	}
	return nil

}
