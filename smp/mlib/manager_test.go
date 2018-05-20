package mlib


import (
	"testing"
)

func TestOps(t *testing.T){
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed.")
	}
	if mm.Len() != 0 {
		t.Error("NewMusicManaer failed, not empty.")
	}
	m0 := &MusicEntry{
		"1", "My Heat Will Go On.","Celion Dion",  "Http://qbox.me/24501234", MP3,
	}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("MusicManager.Add() failed.")
	}

	m := Find(m0.Nanme)
	if m == nil {
		t.Error("MusicManager.Find() failed.")
	}
	if m.Id != m0.Id || m.Artist != m0.Artist ||m.Nanme != m0.Nanme || 
	m.Source != m0.Source || m.Type != m0.Type {
		t.Error("MusicManager.Find() failed.", err)
	}

	m, err := mm.Get(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.Get() failed.", err)
	}

	m = mm.Remoeve(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove() failed.", err)
	}


}