package models

// sprites:
//   path: ./sprites
//   out: ./bin/data/sprites.bin

// levels:
//   path: ./levels
//   out: ./bin/data/levels.bin

// sounds:
//     path: ./sounds
//     out: ./bin/media/sounds

type SptiresSource struct {
	Path string
	Out  string
}

type LevelsSource struct {
	Path string
	Out  string
}

type SoundsSource struct {
	Path string
	Out  string
}

type SourceConfig struct {
	Sprites SptiresSource
	Levels  LevelsSource
	Sounds  SoundsSource
}
