package postgres

import (
	"MatchaServer/common"
	"MatchaServer/config"
	"MatchaServer/errors"
	"testing"
)

func TestPhotos(t *testing.T) {
	print(common.NO_COLOR)
	var (
		repo  = New()
		uid1  = 1000
		uid2  = 1002
		body1 = "qwerty" //[]byte("qwerty")
		body2 = "ytrewq" //[]byte("ytrewq")
		body3 = "asd"    //[]byte("asd")
	)
	conf, err := config.Create("../../config/")
	if err != nil {
		t.Errorf(common.RED_BG + "ERROR: Cannot get config file - " + err.Error() + common.NO_COLOR)
		return
	}
	err = repo.Connect(&conf.Sql)
	if err != nil {
		t.Errorf(common.RED_BG + "ERROR: Cannot connect to database - " + err.Error() + common.NO_COLOR)
		return
	}

	t.Run("Create Photo #1", func(t_ *testing.T) {
		_, err := repo.SetNewPhoto(uid1, body1)
		if err != nil {
			t_.Errorf(common.RED_BG + "Error: " + err.Error() + common.NO_COLOR)
			return
		}
		t_.Log(common.GREEN_BG + "Success" + common.NO_COLOR)
	})
	t.Run("Create Photo #2", func(t_ *testing.T) {
		_, err := repo.SetNewPhoto(uid1, body2)
		if err != nil {
			t_.Errorf(common.RED_BG + "Error: " + err.Error() + common.NO_COLOR)
			return
		}
		t_.Log(common.GREEN_BG + "Success" + common.NO_COLOR)
	})
	t.Run("Create Photo #3", func(t_ *testing.T) {
		_, err := repo.SetNewPhoto(uid2, body3)
		if err != nil {
			t_.Errorf(common.RED_BG + "Error: " + err.Error() + common.NO_COLOR)
			return
		}
		t_.Log(common.GREEN_BG + "Success" + common.NO_COLOR)
	})
	t.Run("Get photos by uid1 #1", func(t_ *testing.T) {
		photos, err := repo.GetPhotosByUid(uid1)
		if err != nil {
			t_.Errorf(common.RED_BG + "Error: " + err.Error() + common.NO_COLOR)
			return
		}
		if len(photos) != 2 {
			t_.Errorf(common.RED_BG+"Error: expected 2 photos, got %d"+common.NO_COLOR+"\n", len(photos))
			return
		}
		t_.Log(common.GREEN_BG + "Success" + common.NO_COLOR)
	})
	t.Run("Get photos by uid2 #1", func(t_ *testing.T) {
		photos, err := repo.GetPhotosByUid(uid2)
		if err != nil {
			t_.Errorf(common.RED_BG + "Error: " + err.Error() + common.NO_COLOR)
			return
		}
		if len(photos) != 1 {
			t_.Errorf(common.RED_BG+"Error: expected 1 photos, got %d"+common.NO_COLOR+"\n", len(photos))
			return
		}
		t_.Log(common.GREEN_BG + "Success" + common.NO_COLOR)
	})
	t.Run("Delete photo", func(t_ *testing.T) {
		var pid int
		photos, err := repo.GetPhotosByUid(uid1)
		if err != nil {
			t_.Errorf(common.RED_BG + "Error: " + err.Error() + common.NO_COLOR)
			return
		}
		for _, photo := range photos {
			pid = photo.Pid
		}
		err = repo.DeletePhoto(pid)
		if err != nil {
			t_.Errorf(common.RED_BG + "Error: " + err.Error() + common.NO_COLOR)
			return
		}
		t_.Log(common.GREEN_BG + "Success" + common.NO_COLOR)
	})
	t.Run("Get photos by uid1 #2", func(t_ *testing.T) {
		photos, err := repo.GetPhotosByUid(uid1)
		if err != nil {
			t_.Errorf(common.RED_BG + "Error: " + err.Error() + common.NO_COLOR)
			return
		}
		if len(photos) != 1 {
			t_.Errorf(common.RED_BG+"Error: expected 1 photos, got %d"+common.NO_COLOR+"\n", len(photos))
			return
		}
		t_.Log(common.GREEN_BG + "Success" + common.NO_COLOR)
	})
	t.Run("Invalid GetPhotoByPid", func(t_ *testing.T) {
		var pid int
		photos, err := repo.GetPhotosByUid(uid1)
		if err != nil {
			t_.Errorf(common.RED_BG + "Error: " + err.Error() + common.NO_COLOR)
			return
		}
		for _, photo := range photos {
			pid = photo.Pid
		}
		err = repo.DeletePhoto(pid)
		if err != nil {
			t_.Errorf(common.RED_BG + "Error: " + err.Error() + common.NO_COLOR)
			return
		}
		_, err = repo.GetPhotoByPid(pid)
		if errors.RecordNotFound.IsOverlapWithError(err) {
			t_.Log(common.GREEN_BG + "Success: there is `RecordNotFoundError` as it expected" + common.NO_COLOR)
		} else if err != nil {
			t_.Errorf(common.RED_BG + "Error: it should be RecordNotFound but it dont - " + err.Error() + common.NO_COLOR)
			return
		} else {
			t_.Errorf(common.RED_BG + "Error: it should be RecordNotFound but there is no error at all" + common.NO_COLOR)
			return
		}
	})
	t.Run("Valid GetPhotoByPid", func(t_ *testing.T) {
		var pid int
		photos, err := repo.GetPhotosByUid(uid2)
		if err != nil {
			t_.Errorf(common.RED_BG + "Error: " + err.Error() + common.NO_COLOR)
			return
		}
		for _, photo := range photos {
			pid = photo.Pid
		}
		_, err = repo.GetPhotoByPid(pid)
		if err != nil {
			t_.Errorf(common.RED_BG + "Error: " + err.Error() + common.NO_COLOR)
			return
		}
		err = repo.DeletePhoto(pid)
		if err != nil {
			t_.Errorf(common.RED_BG + "Error: " + err.Error() + common.NO_COLOR)
			return
		}
		t_.Log(common.GREEN_BG + "Success" + common.NO_COLOR)
	})
	print(common.YELLOW)
}
