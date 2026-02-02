package config

type Config struct {
	Rows      int
	Cols      int
	NumAgents int
	NumPhones int
	NumWalls  int
}

var Default = Config{
	Rows:      10,
	Cols:      10,
	NumAgents: 2,
	NumPhones: 1,
	NumWalls:  2,
}
