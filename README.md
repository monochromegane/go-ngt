# go-ngt

NGT binding for golang.

[NGT](https://github.com/yahoojapan/NGT) is Neighborhood Graph and Tree for Indexing High-dimensional Data. This package provides golang bindings for the library.

## Golang code example

Create graph and tree.

```go
dim := 40
property, _ := ngt.NewNGTProperty(dim)
defer property.Free()

index, _ := ngt.CreateGraphAndTree("sample_database", property)
defer index.Close()

index.SaveIndex("sample_database")
```

Open and insert index.

```go
index, _ := ngt.OpenIndex("sample_database")
defer index.Close()

obj := make([]float64, dim)
for i := 0; i < dim; i++ {
	obj[i] = rand.Float64()
}
index.InsertIndex(obj)
index.CreateIndex(24) // number of threads
index.SaveIndex("sample_database")
```

Search index.

```go
query := make([]float64, dim)
for i := 0; i < dim; i++ {
	query[i] = rand.Float64()
}
results, _ := index.Search(query, 10)
```

## License

[Apache License Version 2.0](https://github.com/monochromegane/go-ngt/blob/master/LICENSE)

## Author

[monochromegane](https://github.com/monochromegane)
