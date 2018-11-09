package mlib

//音乐实体
type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

//音乐库
type MusicManager struct {
	musics []MusicEntry
}

//新建一个音乐库
func NewMusicManager() *MusicManager {
	return &MusicManager{
		make([]MusicEntry, 0),
	}
}

//音乐的size
func (this *MusicManager) Len() int {
	return len(this.musics)
}
