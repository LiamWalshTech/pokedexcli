package pokeapi

import "encoding/json"

// Reusable named resource (name + url)
type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Ability reference
type Ability struct {
	NamedAPIResource
}

// Version (game version)
//type Version struct {
//	NamedAPIResource
//}

// VersionGroup (e.g. "red-blue", "sword-shield")
type VersionGroup struct {
	NamedAPIResource
}

// Move reference
type Move struct {
	NamedAPIResource
}

// Type reference
type Type struct {
	NamedAPIResource
}

// Stat reference
type Stat struct {
	NamedAPIResource
}

// Item reference
type Item struct {
	NamedAPIResource
}

// PokemonSpecies reference
type PokemonSpecies struct {
	NamedAPIResource
}

// Generation reference
type Generation struct {
	NamedAPIResource
}

// PokemonAbility – ability with slot and hidden flag
type PokemonAbility struct {
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
	Ability  Ability `json:"ability"`
}

// GameIndex – index in a particular game version
type GameIndex struct {
	GameIndex int     `json:"game_index"`
	Version   Version `json:"version"`
}

// Cries – cry audio URLs
type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

// PokemonHeldItemVersion – version-specific held item details
type PokemonHeldItemVersion struct {
	Rarity  int     `json:"rarity"`
	Version Version `json:"version"`
}

// PokemonHeldItem – item that can be held + version details
type PokemonHeldItem struct {
	Item           Item                     `json:"item"`
	VersionDetails []PokemonHeldItemVersion `json:"version_details"`
}

// MoveLearnMethod reference
type MoveLearnMethod struct {
	NamedAPIResource
}

// VersionGroupDetail – how/when a move is learned in a version group
type VersionGroupDetail struct {
	LevelLearnedAt  int             `json:"level_learned_at"`
	MoveLearnMethod MoveLearnMethod `json:"move_learn_method"`
	VersionGroup    VersionGroup    `json:"version_group"`
}

// PokemonMoveVersion – move + learning details per version group
type PokemonMoveVersion struct {
	Move                Move                 `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

// PokemonMove – all ways a Pokémon can learn a move
type PokemonMove struct {
	Move                Move                 `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

// PokemonStat – base stat + effort points
type PokemonStat struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

// PokemonType – type with slot (1 = primary, 2 = secondary)
type PokemonType struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}

// PokemonForm reference
type PokemonForm struct {
	NamedAPIResource
}

// PastAbilityEntry – ability that existed in previous generations
type PastAbilityEntry struct {
	Generation Generation `json:"generation"`
	Abilities  []struct {
		Ability  *Ability `json:"ability,omitempty"` // can be null
		IsHidden bool     `json:"is_hidden"`
		Slot     int      `json:"slot"`
	} `json:"abilities"`
}

// Sprites – all sprite URLs (kept as raw map for simplicity; can be expanded if needed)
type PokemonSprites struct {
	BackDefault      string  `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        string  `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`

	Other struct {
		DreamWorld struct {
			FrontDefault string  `json:"front_default"`
			FrontFemale  *string `json:"front_female"`
		} `json:"dream_world"`
		Home struct {
			FrontDefault     string  `json:"front_default"`
			FrontFemale      *string `json:"front_female"`
			FrontShiny       string  `json:"front_shiny"`
			FrontShinyFemale *string `json:"front_shiny_female"`
		} `json:"home"`
		OfficialArtwork struct {
			FrontDefault string `json:"front_default"`
			FrontShiny   string `json:"front_shiny"`
		} `json:"official-artwork"`
		Showdown struct {
			BackDefault      string  `json:"back_default"`
			BackFemale       *string `json:"back_female"`
			BackShiny        string  `json:"back_shiny"`
			BackShinyFemale  *string `json:"back_shiny_female"`
			FrontDefault     string  `json:"front_default"`
			FrontFemale      *string `json:"front_female"`
			FrontShiny       string  `json:"front_shiny"`
			FrontShinyFemale *string `json:"front_shiny_female"`
		} `json:"showdown"`
	} `json:"other"`

	Versions json.RawMessage `json:"versions"` // huge nested object – keep as raw JSON or define fully
}

// RespShallowPokemon – the full Pokémon object
type RespShallowPokemon struct {
	ID                     int    `json:"id"`
	Name                   string `json:"name"`
	BaseExperience         int    `json:"base_experience"`
	Height                 int    `json:"height"` // in decimetres
	Weight                 int    `json:"weight"` // in hectograms
	IsDefault              bool   `json:"is_default"`
	Order                  int    `json:"order"`
	LocationAreaEncounters string `json:"location_area_encounters"`

	Abilities   []PokemonAbility  `json:"abilities"`
	Forms       []PokemonForm     `json:"forms"`
	GameIndices []GameIndex       `json:"game_indices"`
	HeldItems   []PokemonHeldItem `json:"held_items"`
	Moves       []PokemonMove     `json:"moves"`
	Species     PokemonSpecies    `json:"species"`
	Sprites     PokemonSprites    `json:"sprites"`
	Stats       []PokemonStat     `json:"stats"`
	Types       []PokemonType     `json:"types"`
	Cries       Cries             `json:"cries"`

	PastAbilities []PastAbilityEntry `json:"past_abilities"`
	PastTypes     []interface{}      `json:"past_types"` // usually empty
}
