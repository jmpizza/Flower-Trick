package models

type Ability struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	IsMainSeries bool   `json:"is_main_series"`
	Generation   struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	EffectEntries []struct {
		Effect      string `json:"effect"`
		ShortEffect string `json:"short_effect"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"effect_entries"`
	EffectChanges []struct {
		VersionGroup struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version_group"`
		EffectEntries []struct {
			Effect   string `json:"effect"`
			Language struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"language"`
		} `json:"effect_entries"`
	} `json:"effect_changes"`
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
	Pokemon []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Pokemon  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon"`
}

type Characteristic struct {
	ID             int   `json:"id"`
	GeneModulo     int   `json:"gene_modulo"`
	PossibleValues []int `json:"possible_values"`
	HighestStat    struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"highest_stat"`
	Descriptions []struct {
		Description string `json:"description"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`
}

type EggGroup struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonSpecies []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon_species"`
}

type GrowthRates struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Formula      string `json:"formula"`
	Descriptions []struct {
		Description string `json:"description"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`
	Levels []struct {
		Level      int `json:"level"`
		Experience int `json:"experience"`
	} `json:"levels"`
	PokemonSpecies []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon_species"`
}

type Nature struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	DecreasedStat struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"decreased_stat"`
	IncreasedStat struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"increased_stat"`
	LikesFlavor struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"likes_flavor"`
	HatesFlavor struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"hates_flavor"`
	PokeathlonStatChanges []struct {
		MaxChange      int `json:"max_change"`
		PokeathlonStat struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokeathlon_stat"`
	} `json:"pokeathlon_stat_changes"`
	MoveBattleStylePreferences []struct {
		LowHpPreference  int `json:"low_hp_preference"`
		HighHpPreference int `json:"high_hp_preference"`
		MoveBattleStyle  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move_battle_style"`
	} `json:"move_battle_style_preferences"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
}

type PokeathlonStat struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	AffectingNatures struct {
		Increase []struct {
			MaxChange int `json:"max_change"`
			Nature    struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"nature"`
		} `json:"increase"`
		Decrease []struct {
			MaxChange int `json:"max_change"`
			Nature    struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"nature"`
		} `json:"decrease"`
	} `json:"affecting_natures"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
}

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt int `json:"level_learned_at"`
			VersionGroup   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			Order int `json:"order"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string `json:"back_default"`
		BackFemale       any    `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  any    `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      any    `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale any    `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string `json:"front_default"`
				FrontFemale  any    `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     string `json:"front_default"`
				FrontFemale      any    `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale any    `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
			Showdown struct {
				BackDefault      string `json:"back_default"`
				BackFemale       any    `json:"back_female"`
				BackShiny        string `json:"back_shiny"`
				BackShinyFemale  any    `json:"back_shiny_female"`
				FrontDefault     string `json:"front_default"`
				FrontFemale      any    `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale any    `json:"front_shiny_female"`
			} `json:"showdown"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault  string `json:"back_default"`
					BackGray     string `json:"back_gray"`
					FrontDefault string `json:"front_default"`
					FrontGray    string `json:"front_gray"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault  string `json:"back_default"`
					BackGray     string `json:"back_gray"`
					FrontDefault string `json:"front_default"`
					FrontGray    string `json:"front_gray"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"crystal"`
				Gold struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"gold"`
				Silver struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      string `json:"back_default"`
						BackFemale       any    `json:"back_female"`
						BackShiny        string `json:"back_shiny"`
						BackShinyFemale  any    `json:"back_shiny_female"`
						FrontDefault     string `json:"front_default"`
						FrontFemale      any    `json:"front_female"`
						FrontShiny       string `json:"front_shiny"`
						FrontShinyFemale any    `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Cries struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	PastTypes []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Types []struct {
			Slot int `json:"slot"`
			Type struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"type"`
		} `json:"types"`
	} `json:"past_types"`
	PastAbilities []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Abilities []struct {
			Ability  any  `json:"ability"`
			IsHidden bool `json:"is_hidden"`
			Slot     int  `json:"slot"`
		} `json:"abilities"`
	} `json:"past_abilities"`
}

type PokemonColor struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonSpecies []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon_species"`
}

type PokemonForm struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Order        int    `json:"order"`
	FormOrder    int    `json:"form_order"`
	IsDefault    bool   `json:"is_default"`
	IsBattleOnly bool   `json:"is_battle_only"`
	IsMega       bool   `json:"is_mega"`
	FormName     string `json:"form_name"`
	Pokemon      struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
	Sprites struct {
		BackDefault      string `json:"back_default"`
		BackFemale       any    `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  any    `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      any    `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale any    `json:"front_shiny_female"`
	} `json:"sprites"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	VersionGroup struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version_group"`
}

type PokemonHabitad struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonSpecies []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon_species"`
}

type PokemonShape struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	AwesomeNames []struct {
		AwesomeName string `json:"awesome_name"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"awesome_names"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonSpecies []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon_species"`
}

type PokemonSpecies struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	Order                int    `json:"order"`
	GenderRate           int    `json:"gender_rate"`
	CaptureRate          int    `json:"capture_rate"`
	BaseHappiness        int    `json:"base_happiness"`
	IsBaby               bool   `json:"is_baby"`
	IsLegendary          bool   `json:"is_legendary"`
	IsMythical           bool   `json:"is_mythical"`
	HatchCounter         int    `json:"hatch_counter"`
	HasGenderDifferences bool   `json:"has_gender_differences"`
	FormsSwitchable      bool   `json:"forms_switchable"`
	GrowthRate           struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"growth_rate"`
	PokedexNumbers []struct {
		EntryNumber int `json:"entry_number"`
		Pokedex     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokedex"`
	} `json:"pokedex_numbers"`
	EggGroups []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"egg_groups"`
	Color struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"color"`
	Shape struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"shape"`
	EvolvesFromSpecies struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"evolves_from_species"`
	EvolutionChain struct {
		URL string `json:"url"`
	} `json:"evolution_chain"`
	Habitat    any `json:"habitat"`
	Generation struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Version struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"flavor_text_entries"`
	FormDescriptions []struct {
		Description string `json:"description"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"form_descriptions"`
	Genera []struct {
		Genus    string `json:"genus"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"genera"`
	Varieties []struct {
		IsDefault bool `json:"is_default"`
		Pokemon   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"varieties"`
}

type Stat struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	GameIndex      int    `json:"game_index"`
	IsBattleOnly   bool   `json:"is_battle_only"`
	AffectingMoves struct {
		Increase []struct {
			Change int `json:"change"`
			Move   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move"`
		} `json:"increase"`
		Decrease []struct {
			Change int `json:"change"`
			Move   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move"`
		} `json:"decrease"`
	} `json:"affecting_moves"`
	AffectingNatures struct {
		Increase []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"increase"`
		Decrease []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"decrease"`
	} `json:"affecting_natures"`
	Characteristics []struct {
		URL string `json:"url"`
	} `json:"characteristics"`
	MoveDamageClass struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"move_damage_class"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
}

type Type struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	DamageRelations struct {
		NoDamageTo []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"no_damage_to"`
		HalfDamageTo []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"half_damage_to"`
		DoubleDamageTo []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"double_damage_to"`
		NoDamageFrom []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"no_damage_from"`
		HalfDamageFrom []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"half_damage_from"`
		DoubleDamageFrom []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"double_damage_from"`
	} `json:"damage_relations"`
	PastDamageRelations []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		DamageRelations struct {
			NoDamageTo []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"no_damage_to"`
			HalfDamageTo []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"half_damage_to"`
			DoubleDamageTo []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"double_damage_to"`
			NoDamageFrom []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"no_damage_from"`
			HalfDamageFrom []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"half_damage_from"`
			DoubleDamageFrom []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"double_damage_from"`
		} `json:"damage_relations"`
	} `json:"past_damage_relations"`
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	Generation struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	MoveDamageClass struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"move_damage_class"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	Pokemon []struct {
		Slot    int `json:"slot"`
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon"`
	Moves []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"moves"`
}
