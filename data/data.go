package data

// les infos finales que l'on veut ajouter
type WPokemon struct {
	Id      int
	Type    string
	Name    string
	Hpoint  int
	Pattack int
	Smove   string
	Price   float64
	Count   int
}

// Extraction des infos d'un simple pokémon
type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
		} `json:"ability"`
	} `json:"abilities"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	Height int `json:"height"`
	ID     int `json:"id"`
	Moves  [1]struct {
		Move struct {
			Name string `json:"name"`
		} `json:"move"`
	} `json:"moves"`
	Name string `json:"name"`
	//Order   int    `json:"order"`
	Sprites struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}
