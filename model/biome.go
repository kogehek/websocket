package model

type Biome struct {
	Name        string
	EnemysCount int
	LockedCount int
}

var (
	defaultBiome = Biome{Name: "default", EnemysCount: 5, LockedCount: 4}
)

var Biomes []Biome = []Biome{
	defaultBiome,
	{
		Name: "forest", EnemysCount: 1, LockedCount: 8,
	},
}

func getBiom(name string) Biome {
	for _, b := range Biomes {
		if b.Name == name {
			return b
		}
	}
	return defaultBiome
}
