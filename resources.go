package pokeapi

import (
	"fmt"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
)

const (
	defaultOffset     = 0
	defaultLimit      = 20
	queryParamsFormat = "?offset=%d&limit=%d"
)

// ResourcesClient describes the interface that needs to be implemented in order to fetch various resources from
// the pokeapi.
type ResourcesClient interface {
	Berries() (*[]Resource, *errors.RestErr)
	BerryFirmness() (*[]Resource, *errors.RestErr)
	BerryFlavours() (*[]Resource, *errors.RestErr)
	ContestTypes() (*[]Resource, *errors.RestErr)
	ContestEffects() (*[]Resource, *errors.RestErr)
	SuperContestEffects() (*[]Resource, *errors.RestErr)
	EncounterMethods() (*[]Resource, *errors.RestErr)
	EncounterConditions() (*[]Resource, *errors.RestErr)
	EncounterConditionValues() (*[]Resource, *errors.RestErr)
	EvolutionChains() (*[]Resource, *errors.RestErr)
	EvolutionTriggers() (*[]Resource, *errors.RestErr)
	Generations() (*[]Resource, *errors.RestErr)
	Pokedexes() (*[]Resource, *errors.RestErr)
	Version() (*[]Resource, *errors.RestErr)
	VersionGroups() (*[]Resource, *errors.RestErr)
	Abilities() (*[]Resource, *errors.RestErr)
	Characteristics() (*[]Resource, *errors.RestErr)
	EggGroups() (*[]Resource, *errors.RestErr)
	Genders() (*[]Resource, *errors.RestErr)
	GrowthRates() (*[]Resource, *errors.RestErr)
	Natures() (*[]Resource, *errors.RestErr)
	PokeathlonStats() (*[]Resource, *errors.RestErr)
	Pokemons() (*[]Resource, *errors.RestErr)
	PokemonColors() (*[]Resource, *errors.RestErr)
	PokemonForms() (*[]Resource, *errors.RestErr)
	PokemonHabitats() (*[]Resource, *errors.RestErr)
	PokemonShapes() (*[]Resource, *errors.RestErr)
	PokemonSpecies() (*[]Resource, *errors.RestErr)
	Stats() (*[]Resource, *errors.RestErr)
	Types() (*[]Resource, *errors.RestErr)
}

// resourcesClient implements ResourcesClient.
type resourcesClient struct {
	client *Client
}

// ResourceApiResp represents the response received by the pokeapi.
type ResourceApiResp struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Resource  `json:"results"`
}

// Resource is a pokeapi resource.
type Resource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// Berries retrieves a list of berries.
func (r *resourcesClient) Berries() (*[]Resource, *errors.RestErr) {
	return r.get(berriesResource)
}

// BerryFirmness retrieves a list of  berry firmness types.
func (r *resourcesClient) BerryFirmness() (*[]Resource, *errors.RestErr) {
	return r.get(berryFirmnessResource)
}

// BerryFlavours retrieves a list of berry flavours.
func (r *resourcesClient) BerryFlavours() (*[]Resource, *errors.RestErr) {
	return r.get(berryFlavoursResource)
}

// ContestTypes retrieves a list of contest types.
func (r *resourcesClient) ContestTypes() (*[]Resource, *errors.RestErr) {
	return r.get(contestTypesResource)
}

// ContestEffects retrieves a list of contest effects.
func (r *resourcesClient) ContestEffects() (*[]Resource, *errors.RestErr) {
	return r.get(contestEffectsResource)
}

// SuperContestEffects retrieves a list of super contest effects.
func (r *resourcesClient) SuperContestEffects() (*[]Resource, *errors.RestErr) {
	return r.get(superContestEffectsResource)
}

// EncounterMethods retrieves a list of encounter methods.
func (r *resourcesClient) EncounterMethods() (*[]Resource, *errors.RestErr) {
	return r.get(encounterMethodsResource)
}

// EncounterConditions retrieves a list of encounter conditions.
func (r *resourcesClient) EncounterConditions() (*[]Resource, *errors.RestErr) {
	return r.get(encounterConditionsResource)
}

// EncounterConditionValues retrieves a list of encounter condition values.
func (r *resourcesClient) EncounterConditionValues() (*[]Resource, *errors.RestErr) {
	return r.get(encounterConditionValuesResource)
}

// EvolutionChains retrieves a list of evolution chains.
func (r *resourcesClient) EvolutionChains() (*[]Resource, *errors.RestErr) {
	return r.get(evolutionChainsResource)
}

// EvolutionTriggers retrieves a list of evolution triggers.
func (r *resourcesClient) EvolutionTriggers() (*[]Resource, *errors.RestErr) {
	return r.get(evolutionTriggersResource)
}

// Generations retrieves a list of pokemon generations.
func (r *resourcesClient) Generations() (*[]Resource, *errors.RestErr) {
	return r.get(generationsResource)
}

// Pokedexes retrieves a list of pokedexes.
func (r *resourcesClient) Pokedexes() (*[]Resource, *errors.RestErr) {
	return r.get(pokedexsResource)
}

// Version retrieves a list of versions.
func (r *resourcesClient) Version() (*[]Resource, *errors.RestErr) {
	return r.get(versionsResource)
}

// VersionGroups retrieves a list of version groups.
func (r *resourcesClient) VersionGroups() (*[]Resource, *errors.RestErr) {
	return r.get(versionGroupsResource)
}

// Abilities retrieves a list of pokemon abilities.
func (r *resourcesClient) Abilities() (*[]Resource, *errors.RestErr) {
	return r.get(abilitiesResource)
}

// Characteristics retrieves a list of pokemon characteristics.
func (r *resourcesClient) Characteristics() (*[]Resource, *errors.RestErr) {
	return r.get(characteristicsResource)
}

// EggGroups retrieves a list of egg groups.
func (r *resourcesClient) EggGroups() (*[]Resource, *errors.RestErr) {
	return r.get(eggGroupsResource)
}

// Genders retrieves a list of genders.
func (r *resourcesClient) Genders() (*[]Resource, *errors.RestErr) {
	return r.get(gendersResource)
}

// GrowthRates retrieves a list of growth rates.
func (r *resourcesClient) GrowthRates() (*[]Resource, *errors.RestErr) {
	return r.get(growthRatesResource)
}

// Natures retrieves a list of natures.
func (r *resourcesClient) Natures() (*[]Resource, *errors.RestErr) {
	return r.get(naturesResource)
}

// PokeathlonStats retrieves a list of pokeathlon stats.
func (r *resourcesClient) PokeathlonStats() (*[]Resource, *errors.RestErr) {
	return r.get(pokeathlonStatsResource)
}

// Pokemons retrieves a list of pokemon's.
func (r *resourcesClient) Pokemons() (*[]Resource, *errors.RestErr) {
	return r.get(pokemonsResource)
}

// PokemonColors retrieves a list of pokemon colors.
func (r *resourcesClient) PokemonColors() (*[]Resource, *errors.RestErr) {
	return r.get(pokemonColorsResource)
}

// PokemonForms retrieves a list of pokemon forms.
func (r *resourcesClient) PokemonForms() (*[]Resource, *errors.RestErr) {
	return r.get(pokemonFormsResource)
}

// PokemonHabitats retrieves a list of habitats.
func (r *resourcesClient) PokemonHabitats() (*[]Resource, *errors.RestErr) {
	return r.get(pokemonHabitatsResource)
}

// PokemonShapes retrieves a list of pokemon shapes.
func (r *resourcesClient) PokemonShapes() (*[]Resource, *errors.RestErr) {
	return r.get(pokemonShapesResource)
}

// PokemonSpecies retrieves a list of pokemon species.
func (r *resourcesClient) PokemonSpecies() (*[]Resource, *errors.RestErr) {
	return r.get(pokemonSpeciesResource)
}

// Stats retrieves a list of pokemon stats.
func (r *resourcesClient) Stats() (*[]Resource, *errors.RestErr) {
	return r.get(statsResource)
}

// Types retrieves a list of pokemon types.
func (r *resourcesClient) Types() (*[]Resource, *errors.RestErr) {
	return r.get(typesResource)
}

// get is a generic method that is used to retrieve various resources from the pokeapi.
// get returns a pointer to the list of resources or an errors.RestErr if the API call fails.
func (r *resourcesClient) get(resourcePath string) (*[]Resource, *errors.RestErr) {
	var (
		result  []Resource
		apiResp ResourceApiResp
		offset  = defaultOffset
		limit   = defaultLimit
	)

	for {
		if restErr := r.client.request(resourcePath+fmt.Sprintf(queryParamsFormat, offset, limit), &apiResp); restErr != nil {
			return nil, restErr
		}
		result = append(result, apiResp.Results...)

		if apiResp.Next == nil {
			break
		}
		offset = offset + limit
	}
	return &result, nil
}
