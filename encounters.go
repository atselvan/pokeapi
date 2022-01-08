package pokeapi

const (
	encounterMethodsResource         = "encounter-method"
	encounterConditionsResource      = "encounter-condition"
	encounterConditionValuesResource = "encounter-condition-value"
)

// EncounterMethod is a method by which the player might can encounter Pokémon in the wild, e.g., walking in tall grass.
type EncounterMethod struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
}

// EncounterCondition is a condition which affect what Pokémon might appear in the wild, e.g., day or night.
type EncounterCondition struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Values []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"values"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
}

// EncounterConditionValue is the states that an encounter condition can have, i.e., time of day can be either day or night.
type EncounterConditionValue struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Condition struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"condition"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
}
