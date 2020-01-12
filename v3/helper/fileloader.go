package helper

import (
	"errors"
	"path/filepath"
	"sync"

	"github.com/davyxu/tabtoy/v3/report"
)

type FileGetter interface {
	GetFile(filename string) (TableFile, error)
}

type FileLoader struct {
	fileByName sync.Map
	inputFile  []string

	syncLoad bool

	UseGBKCSV bool
}

func (loader *FileLoader) AddFile(filename string) {

	loader.inputFile = append(loader.inputFile, filename)
}

func (loader *FileLoader) Commit() {

	var task sync.WaitGroup
	task.Add(len(loader.inputFile))

	for _, inputFileName := range loader.inputFile {

		go func(fileName string) {

			loader.fileByName.Store(fileName, loadFileByExt(fileName, loader.UseGBKCSV))

			task.Done()

		}(inputFileName)

	}

	task.Wait()

	loader.inputFile = loader.inputFile[0:0]
}

func loadFileByExt(filename string, useGBKCSV bool) interface{} {

	var tabFile TableFile
	switch filepath.Ext(filename) {
	case ".xlsx", ".xls", ".xlsm":

		tabFile = NewXlsxFile()

		err := tabFile.Load(filename)

		if err != nil {
			return err
		}

	case ".csv":
		tabFile = NewCSVFile()

		err := tabFile.Load(filename)

		if err != nil {
			return err
		}

		// 输入gbk, 内部utf8
		if useGBKCSV {
			tabFile.(*CSVFile).Transform(ConvGBKToUTF8)
		}

	default:
		report.Error("UnknownInputFileExtension", filename)
	}

	return tabFile
}

func (self *FileLoader) GetFile(filename string) (TableFile, error) {

	if self.syncLoad {

		result := loadFileByExt(filename, self.UseGBKCSV)
		if err, ok := result.(error); ok {
			return nil, err
		}

		return result.(TableFile), nil

	} else {
		if result, ok := self.fileByName.Load(filename); ok {

			if err, ok := result.(error); ok {
				return nil, err
			}

			return result.(TableFile), nil

		} else {
			return nil, errors.New("not found")
		}
	}

}

func NewFileLoader(syncLoad bool) *FileLoader {
	return &FileLoader{
		syncLoad: syncLoad,
	}
}
