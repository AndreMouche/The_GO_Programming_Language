/**
 * Created with IntelliJ IDEA.
 * User: fun
 * Date: 13-8-20
 * Time: 上午11:31
 * To change this template use File | Settings | File Templates.
 */
package library

import (
	"testing"
)

func TestOps(t *testing.T) {
	mm := NewMusicManager();

	if mm == nil {
		t.Error("NewMusicManager failed.");
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManager failed,not empty");
	}

	m0 := &MusicEntry {
		"1","My Heart will go on","Celion Dion","Pop",
		"http://qbox.me" }
	mm.Add(m0);

	if mm.Len() != 1 {
		t.Error("MusicManager.Add() failed.");
	}

	m := mm.Find(m0.Name)

	if m == nil {
		t.Error("MusicManager.Find() failed.")
	}

	if m.Id != m0.Id || m.Artist != m0.Artist ||
	   m.Source != m0.Source {
		      t.Error("MusicManager.Find() failed, Found item mismatch.");
	}

	m,err := mm.Get(0);

	if m == nil {
		t.Error("MusicManager.Get() failed",err);
	}

	m = mm.Remove(0);
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove() failed.",err);
	}
}

