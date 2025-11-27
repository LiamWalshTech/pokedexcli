package pokeapi

// NamedResource is used for every object that only contains a name + URL
type NamedResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Language is the same as NamedResource (all languages follow this pattern)
type Language NamedResource

// EncounterMethod is the method used to encounter Pokémon (walk, surf, etc.)
type EncounterMethod NamedResource

// Version is the game version
type Version NamedResource

// Pokemon basic reference
type Pokemon NamedResource

// EncounterConditionValue (e.g. time-morning, radio-hoenn, …)
type EncounterConditionValue NamedResource

// EncounterDetail – one concrete way a Pokémon can be encountered
type EncounterDetail struct {
	Chance          int                       `json:"chance"`
	ConditionValues []EncounterConditionValue `json:"condition_values"`
	MaxLevel        int                       `json:"max_level"`
	MinLevel        int                       `json:"min_level"`
	Method          EncounterMethod           `json:"method"`
}

// VersionDetails – encounter information for a specific game version
type VersionDetails struct {
	EncounterDetails []EncounterDetail `json:"encounter_details"`
	MaxChance        int               `json:"max_chance"`
	Version          Version           `json:"version"`
}

// PokemonEncounter – all encounter data for a single Pokémon in this area
type PokemonEncounter struct {
	Pokemon        Pokemon          `json:"pokemon"`
	VersionDetails []VersionDetails `json:"version_details"`
}

// EncounterMethodRate – how often a certain encounter method appears in this area
type EncounterMethodRate struct {
	EncounterMethod EncounterMethod `json:"encounter_method"`
	VersionDetails  []struct {
		Rate    int     `json:"rate"`
		Version Version `json:"version"`
	} `json:"version_details"`
}

// NameEntry – multilingual name of the location area
type NameEntry struct {
	Language Language `json:"language"`
	Name     string   `json:"name"`
}

// RespShallowLocation – the complete object returned by the PokeAPI endpoint
type RespShallowLocation struct {
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	GameIndex            int                   `json:"game_index"`
	ID                   int                   `json:"id"`
	Location             NamedResource         `json:"location"`
	Name                 string                `json:"name"`
	Names                []NameEntry           `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}
