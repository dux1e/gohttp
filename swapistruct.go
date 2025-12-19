package main

type SwapiResponse struct {
	Films     string `json:"films"`
	People    string `json:"people"`
	Planets   string `json:"planets"`
	Species   string `json:"species"`
	Vehicles  string `json:"vehicles"`
	Starships string `json:"starships"`
}

type SwapiPeople struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	EyeColor  string   `json:"eye_color"`
	BirthYear string   `json:"birth_year"`
	Gender    string   `json:"gender"`
	HomeWorld string   `json:"homeworld"`
	Films     []string `json:"films"`
	Species   []string `json:"species"`
	Vehicles  []string `json:"vehicles"`
	Starships []string `json:"starships"`
	URL       string   `json:"url"`
}
