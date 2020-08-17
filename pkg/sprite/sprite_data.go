package sprite

type SpriteGlyph struct {
	Display    string
	Foreground string
	Background string
}

type SpriteFrameData struct {
	Names []string
	Next  string
	Time  int
	Data  []string
}

type SpriteData struct {
	Name    string
	Width   int
	Height  int
	OriginX int
	OriginY int
	Layer   int
	Glyphs  map[string]SpriteGlyph
	Palette map[string]string
	Frames  []SpriteFrameData
}
