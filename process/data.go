package process

// GameData Walkr game information
type GameData struct {
	Number       int
	Planet       string
	PlanetFile   string
	Satelite     string
	SateliteFile string
	Resource     string
}

// PlayerName Telegram player name mapping
type PlayerName struct {
	CodeName  string
	TgName    string
	WalkrName string
}
