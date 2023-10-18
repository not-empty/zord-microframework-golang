package config

type ArtefactsConfig struct {
	Path       string
	Replacers  []string
	Stub       string
	ArchConfig ArchConfig
}

type ArchConfig struct {
	Replacers []string
	Artefacts map[string]ArtefactsConfig
}
