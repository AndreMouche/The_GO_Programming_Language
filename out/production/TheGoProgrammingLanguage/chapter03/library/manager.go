/**
 * Created with IntelliJ IDEA.
 * User: fun
 * Date: 13-8-20
 * Time: 上午10:44
 * To change this template use File | Settings | File Templates.
 */
package library

import (
    "fmt"
	"errors"
)


type MusicEntry struct  {
	Id string
	Name string
	Artist string
	Source string
	Type string
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry,0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Find (name string) *MusicEntry {
	 if len(m.musics) == 0 {
		 return nil;
	 }

	for _, m := range m.musics {    // may have bugs?
		if m.Name == name {
			return &m;
		}
	}
	return nil
}

func (m *MusicManager) Get (index int) (music *MusicEntry,err error)  {
	    if index < 0 || index >= m.Len() {
			return nil, errors.New("Index out of range.")
		}

	    return &m.musics[index],nil
}

func (m *MusicManager) Add (music *MusicEntry) {
	m.musics = append(m.musics, *music);
}


func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= m.Len() {
		return nil;
	}

	removeMusic := &m.musics[index];

	//从数组切片中删除
	if index < m.Len() - 1 { //从中间元素删除
	    m.musics = append(m.musics[:index],m.musics[index+1:]...);
	} else if index == 0 {                   //删除的是仅有的元素
		m.musics = make([]MusicEntry,0);
	} else {
		m.musics = m.musics[:index-1];
	}

	return removeMusic
}


func (m *MusicManager) RemoveByName(name string) *MusicEntry{
	 for id,music := range m.musics {
		 if music.Name == name {
			fmt.Println(id);
			return m.Remove(id)
		 }
	 }
	return nil
}
