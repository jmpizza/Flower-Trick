package models

type Move struct {
	Accuracy      interface{} `json:"accuracy"`
	ContestCombos interface{} `json:"contest_combos"`
	ContestEffect interface{} `json:"contest_effect"`
	ContestType   interface{} `json:"contest_type"`
	DamageClass   struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"damage_class"`
	EffectChance      interface{}   `json:"effect_chance"`
	EffectChanges     []interface{} `json:"effect_changes"`
	EffectEntries     []interface{} `json:"effect_entries"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		VersionGroup struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version_group"`
	} `json:"flavor_text_entries"`
	Generation struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	ID               int `json:"id"`
	LearnedByPokemon []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"learned_by_pokemon"`
	Machines []interface{} `json:"machines"`
	Meta     interface{}   `json:"meta"`
	Name     string        `json:"name"`
	Names    []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PastValues         []interface{} `json:"past_values"`
	Power              int           `json:"power"`
	Pp                 int           `json:"pp"`
	Priority           int           `json:"priority"`
	StatChanges        []interface{} `json:"stat_changes"`
	SuperContestEffect interface{}   `json:"super_contest_effect"`
	Target             struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"target"`
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}