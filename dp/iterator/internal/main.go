package internal

import "fmt"

/*
Item
*/
type Music struct {
	Id     int64
	Title  string
	Artist string
}

func (m Music) String() string {
	return fmt.Sprintf("ID: %d, Title: %s, Artist: %s", m.Id, m.Title, m.Artist)
}

/*
Item aggregator
*/
type Musics struct {
	items    []*Music
	iterator *MusicsIterator
}

func NewMusics() *Musics {
	m := &Musics{
		iterator: &MusicsIterator{},
	}
	m.iterator.Musics = m
	return m
}

func (m *Musics) GetItemAt(idx int64) *Music {
	if m.GetSize() <= idx {
		return nil
	}
	return m.items[idx]
}

func (m *Musics) HasNext() bool {
	return m.iterator.HasNext()
}

func (m *Musics) Scan() *Music {
	return m.iterator.Scan()
}

func (m *Musics) Append(music *Music) {
	m.items = append(m.items, music)
}

func (m *Musics) GetSize() int64 {
	return int64(len(m.items))
}

/*
Iterator
*/
type MusicsIterator struct {
	Musics *Musics
	Idx    int64
}

func (m *MusicsIterator) HasNext() bool {
	if m.Idx < m.Musics.GetSize() {
		return true
	}
	return false
}

func (m *MusicsIterator) Scan() *Music {
	item := m.Musics.GetItemAt(m.Idx)
	m.Idx += 1
	return item
}
