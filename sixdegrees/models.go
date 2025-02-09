package sixdegrees

type ActorResults struct {
	Results []struct {
		ID                 int    `json:"id"`
		KnownForDepartment string `json:"known_for_department"`
	} `json:"results"`
}

type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title,omitempty"`
}

type MovieExtended struct {
	Movie         Movie
	OriginalMovie *MovieExtended
}

type MovieResults struct {
	Cast []Movie `json:"cast"`
}

type Cast struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CastResults struct {
	Cast []Cast `json:"cast"`
}
