package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPhotoAlbum(t *testing.T) {
	t.Run("new album", func(t *testing.T) {
		m := NewPhotoAlbum("ABC", "EFG")
		assert.Equal(t, "ABC", m.PhotoUID)
		assert.Equal(t, "EFG", m.AlbumUID)
	})
}

func TestPhotoAlbum_TableName(t *testing.T) {
	photoAlbum := &PhotoAlbum{}
	tableName := photoAlbum.TableName()

	assert.Equal(t, "photos_albums", tableName)
}

func TestFirstOrCreatePhotoAlbum(t *testing.T) {
	model := PhotoAlbumFixtures.Get("1", "pt9jtdre2lvl0yh7", "at9lxuqxpogaaba8")
	result := FirstOrCreatePhotoAlbum(&model)

	if result == nil {
		t.Fatal("result should not be nil")
	}

	if result.AlbumUID != model.AlbumUID {
		t.Errorf("AlbumUID should be the same: %s %s", result.AlbumUID, model.AlbumUID)
	}

	if result.PhotoUID != model.PhotoUID {
		t.Errorf("PhotoUID should be the same: %s %s", result.PhotoUID, model.PhotoUID)
	}
}
