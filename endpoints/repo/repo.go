package repo

import (
	"encoding/json"
	"errors"
	"os"
)

var storageFileName = "terraform-provider-sample.json"

type Repo map[string]interface{}

func New() Repo {
	repo := Repo{}
	repo.Load()
	return repo
}

func (me Repo) Put(id string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	me[id] = string(data)
	return me.Save()
}

func (me Repo) Get(id string, v interface{}) error {
	if storedJSON, ok := me[id]; ok {
		return json.Unmarshal([]byte(storedJSON.(string)), v)
	}
	return errors.New("404 not found")
}

func (me Repo) Delete(id string) {
	delete(me, id)
	me.Save()
}

func (me Repo) Save() error {
	var data []byte
	var err error

	if data, err = json.MarshalIndent(me, "", "  "); err != nil {
		return err
	}
	f, err := os.Create(storageFileName)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.Write(data); err != nil {
		return err
	}

	f.Sync()

	return nil
}

func (me Repo) Load() error {
	if _, err := os.Stat(storageFileName); errors.Is(err, os.ErrNotExist) {
		return nil
	}
	var data []byte
	var err error
	if data, err = os.ReadFile(storageFileName); err != nil {
		return err
	}
	if err = json.Unmarshal(data, &me); err != nil {
		return err
	}
	return nil
}
