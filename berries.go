package pokeapi

import (
	"github.com/privatesquare/bkst-go-utils/utils/errors"
)

const (
	berryApiPath = "/berry"
)

// BerriesClient represents the interface that needs to be implemented to perform actions on the berries endpoint.
type BerriesClient interface {
	Get(id ID) (*Berry, *errors.RestErr)
}

// berriesClient implements BerriesClient.
type berriesClient struct {
	client *Client
}

// Berry Berries are small fruits that can provide HP and status condition restoration, stat enhancement, and even damage
// negation when eaten by Pokémon.
type Berry struct {
	Firmness struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"firmness"`
	Flavors []struct {
		Flavor struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"flavor"`
		Potency int `json:"potency"`
	} `json:"flavors"`
	GrowthTime int `json:"growth_time"`
	Id         int `json:"id"`
	Item       struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"item"`
	MaxHarvest       int    `json:"max_harvest"`
	Name             string `json:"name"`
	NaturalGiftPower int    `json:"natural_gift_power"`
	NaturalGiftType  struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"natural_gift_type"`
	Size        int `json:"size"`
	Smoothness  int `json:"smoothness"`
	SoilDryness int `json:"soil_dryness"`
}

// Get retrieves the data about a berry that matches the id.
// id can be an id or the name of a berry.
// If id is nil then a list of berries will be returned. The size of the list depends on the
// offset and the limit set on the api client.
func (bc *berriesClient) Get(id ID) (berry *Berry, restErr *errors.RestErr) {
	if restErr := id.Validate(); restErr != nil {
		return nil, restErr
	}
	berry = new(Berry)
	if restErr := bc.client.request(berryApiPath+"/"+id.String(), berry); restErr != nil {
		return nil, restErr
	}
	return berry, restErr
}
