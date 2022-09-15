package dataStorage

type dataWriteable interface {
	Write(content string, append bool)
}
