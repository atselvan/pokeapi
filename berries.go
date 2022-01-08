package pokeapi

import (
	"fmt"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
)

const (
	berriesResource       = "/berry"
	berryFirmnessResource = "/berry-firmness"
	berryFlavoursResource = "/berry-flavor"
)

// BerriesClient represents the interface that needs to be implemented to perform actions on the berries endpoint.
type BerriesClient interface {
	Get(id ID) (*Berry, *errors.RestErr)
	GetFirmness(id ID) (*BerryFirmness, *errors.RestErr)
	GetFlavour(id ID) (*BerryFlavour, *errors.RestErr)
}

// berriesClient implements BerriesClient.
type berriesClient struct {
	client *Client
}

// Berry Berries are small fruits that can provide HP and status condition restoration, stat enhancement, and even damage
// negation when eaten by Pokémon.
type Berry struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	GrowthTime       int    `json:"growth_time"`
	MaxHarvest       int    `json:"max_harvest"`
	NaturalGiftPower int    `json:"natural_gift_power"`
	Size             int    `json:"size"`
	Smoothness       int    `json:"smoothness"`
	SoilDryness      int    `json:"soil_dryness"`
	Firmness         struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"firmness"`
	Flavors []struct {
		Potency int `json:"potency"`
		Flavor  struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"flavor"`
	} `json:"flavors"`
	Item struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"item"`
	NaturalGiftType struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"natural_gift_type"`
}

// BerryFirmness is the firmness of a berry. Berries can be soft or hard.
type BerryFirmness struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Berries []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"berries"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
}

// BerryFlavour determine whether a Pokémon will benefit or suffer from eating a berry based on their nature.
type BerryFlavour struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Berries []struct {
		Potency int `json:"potency"`
		Berry   struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"berry"`
	} `json:"berries"`
	ContestType struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"contest_type"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
}

// Get retrieves the data about a berry.
// id can be an id or the name of a berry.
func (bc *berriesClient) Get(id ID) (berry *Berry, restErr *errors.RestErr) {
	if restErr := id.Validate(); restErr != nil {
		return nil, restErr
	}
	berry = new(Berry)
	if restErr := bc.client.request(fmt.Sprintf("%s/%s", berriesResource, id.String()), berry); restErr != nil {
		return nil, restErr
	}
	return berry, restErr
}

// GetFirmness retrieves the data about a berry's firmness.
// id can be an id or the name of a berry.
func (bc *berriesClient) GetFirmness(id ID) (berryFirmness *BerryFirmness, restErr *errors.RestErr) {
	if restErr := id.Validate(); restErr != nil {
		return nil, restErr
	}
	berryFirmness = new(BerryFirmness)
	if restErr := bc.client.request(fmt.Sprintf("%s/%s", berryFirmnessResource, id.String()), berryFirmness); restErr != nil {
		return nil, restErr
	}
	return berryFirmness, restErr
}

// GetFlavour retrieves the data about a berry's flavour.
// id can be an id or the name of a berry.
func (bc *berriesClient) GetFlavour(id ID) (berryFlavour *BerryFlavour, restErr *errors.RestErr) {
	if restErr := id.Validate(); restErr != nil {
		return nil, restErr
	}
	berryFlavour = new(BerryFlavour)
	if restErr := bc.client.request(fmt.Sprintf("%s/%s", berryFlavoursResource, id.String()), berryFlavour); restErr != nil {
		return nil, restErr
	}
	return berryFlavour, restErr
}
