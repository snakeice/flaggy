package sprite

import (
	"sync"
)

var spriteData map[string]*SpriteData = make(map[string]*SpriteData)
var spriteDataLock sync.RWMutex

func AddSprite(data *SpriteData) {
	spriteDataLock.Lock()
	defer spriteDataLock.Unlock()
	spriteData[data.Name] = data

}

func GetSprite(name string) *SpriteData {
	spriteDataLock.RLock()
	defer spriteDataLock.RUnlock()
	return spriteData[name]
}

func RemoveSprite(name string) {
	spriteDataLock.Lock()
	defer spriteDataLock.Unlock()
	delete(spriteData, name)
}
