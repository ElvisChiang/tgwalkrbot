package process

// GameData Walkr game information
type GameData struct {
	Number        int
	Planet        string
	PlanetFile    string
	Satellite     string
	SatelliteFile string
	Resource      string
}

// PlayerName Telegram player name mapping
type PlayerName struct {
	CodeName  string
	TgName    string
	WalkrName string
}
