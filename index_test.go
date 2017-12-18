package ngt

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCreateGraphAndTree(t *testing.T) {
	databaseDir := tempDatabaseDir("db")
	defer os.RemoveAll(databaseDir)

	property, err := NewNGTProperty(5)
	if err != nil {
		t.Errorf("NewNGTProperty should not return error, but return %v", err)
	}
	defer property.Free()

	index, err := CreateGraphAndTree(databaseDir, property)
	if err != nil {
		t.Errorf("CreateGraphAndTree should not return error, but return %v", err)
	}
	defer index.Close()
}

func TestOpenIndex(t *testing.T) {
	index, dir := testCreateGraphAndTree("db", 5)
	index.SaveIndex(dir)
	index.Close()
	defer os.RemoveAll(dir)

	index, err := OpenIndex(dir)
	if err != nil {
		t.Errorf("OpenIndex should not return error, but return %v", err)
	}
	defer index.Close()
}

func TestInsertAndRemoveIndex(t *testing.T) {
	dim := 5
	index, dir := testCreateGraphAndTree("db", dim)
	defer os.RemoveAll(dir)
	defer index.Close()

	rand.Seed(time.Now().UnixNano())
	features := make([]float64, dim)
	for i := 0; i < dim; i++ {
		features[i] = rand.Float64()
	}

	// Insert
	id, err := index.InsertIndex(features)
	if err != nil {
		t.Errorf("NGTIndex.InsertIndex should not return error, but return %v", err)
	}

	err = index.CreateIndex(1)
	if err != nil {
		t.Errorf("NGTIndex.CreateIndex should not return error, but return %v", err)
	}

	results, err := index.Search(features, 1, 0.1)
	if err != nil {
		t.Errorf("NGTIndex.Search should not return error, but return %v", err)
	}
	if cnt := len(results); cnt != 1 {
		t.Errorf("NGTIndex.Search should return %d results, but return %d results", 1, cnt)
	}
	if result := results[0].Id; result != id {
		t.Errorf("NGTIndex.Search should return ID: %d, but return %d", id, result)
	}

	// Remove
	err = index.RemoveIndex(id)
	if err != nil {
		t.Errorf("NGTIndex.RemoveIndex should not return error, but return %v", err)
	}

	results, err = index.Search(features, 1, 0.1)
	if err != nil {
		t.Errorf("NGTIndex.Search should not return error, but return %v", err)
	}
	if cnt := len(results); cnt != 0 {
		t.Errorf("NGTIndex.Search should return %d results, but return %d results", 0, cnt)
	}
}

func testCreateGraphAndTree(database string, dim int) (NGTIndex, string) {
	property, _ := NewNGTProperty(int32(dim))
	defer property.Free()
	databaseDir := tempDatabaseDir(database)
	index, _ := CreateGraphAndTree(databaseDir, property)
	return index, databaseDir
}

func tempDatabaseDir(database string) string {
	dir, _ := ioutil.TempDir("", "go-ngt-test")
	return filepath.Join(dir, database)
}
