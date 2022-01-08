package pokeapi

const (
	contestTypesResource        = "contest-type"
	contestEffectsResource      = "contest-effect"
	superContestEffectsResource = "super-contest-effect"
)

// ContestType is a category judges used to weigh a Pokémon's condition in Pokémon contests.
type ContestType struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	BerryFlavor struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"berry_flavor"`
	Names []struct {
		Name     string `json:"name"`
		Color    string `json:"color"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
}

// ContestEffect refer to the effects of moves when used in contests.
type ContestEffect struct {
	Id            int `json:"id"`
	Appeal        int `json:"appeal"`
	Jam           int `json:"jam"`
	EffectEntries []struct {
		Effect   string `json:"effect"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"effect_entries"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"flavor_text_entries"`
}

// SupperContestEffect refer to the effects of moves when used in super contests.
type SupperContestEffect struct {
	Id                int `json:"id"`
	Appeal            int `json:"appeal"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"flavor_text_entries"`
	Moves []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"moves"`
}
