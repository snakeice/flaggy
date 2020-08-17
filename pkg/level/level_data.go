package level

import "github.com/snakeice/flaggy/pkg/properties"

type LevelData struct {
	Name       string
	Properties properties.GameObjectProps
	Objects    map[string]properties.GameObjectProps
}
