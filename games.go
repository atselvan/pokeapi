package pokeapi

const (
	generationsResource   = "generation"
	pokedexsResource      = "pokedex"
	versionsResource      = "version"
	versionGroupsResource = "version-group"
)

// Generation is a grouping of the Pokémon games that separates them based on the Pokémon they include.
// In each generation, a new set of Pokémon, Moves, Abilities and Types that did not exist in the previous
// generation are released.
type Generation struct {
	Id         int           `json:"id"`
	Name       string        `json:"name"`
	Abilities  []interface{} `json:"abilities"`
	MainRegion struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"main_region"`
	Moves []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"moves"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonSpecies []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"pokemon_species"`
	Types []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"types"`
	VersionGroups []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"version_groups"`
}

// Pokédex is a handheld electronic encyclopedia device; one which is capable of recording and retaining
// information of the various Pokémon in a given region except the national dex and some smaller
// dexes related to portion of a region.
type Pokédex struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	IsMainSeries bool   `json:"is_main_series"`
	Descriptions []struct {
		Description string `json:"description"`
		Language    struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEntries []struct {
		EntryNumber    int `json:"entry_number"`
		PokemonSpecies struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon_species"`
	} `json:"pokemon_entries"`
	Region struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"region"`
	VersionGroups []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"version_groups"`
}

// Version of the game. e.g., Red, Blue or Yellow.
type Version struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	VersionGroup struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"version_group"`
}

// VersionGroup categorize highly similar versions of the games.
type VersionGroup struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Order      int    `json:"order"`
	Generation struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"generation"`
	MoveLearnMethods []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"move_learn_methods"`
	Pokedexes []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"pokedexes"`
	Regions []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"regions"`
	Versions []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"versions"`
}
