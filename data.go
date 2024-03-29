package main

import (
	"fmt"
	"os"
)

type data struct {
	file     *os.File
	pagesize int
}

func NewData(path string, pagesize int) (*data, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return &data{
		file:     file,
		pagesize: pagesize,
	}, nil
}

func (d *data) close() error {
	if d.file != nil {
		err := d.file.Close()
		if err != nil {
			return fmt.Errorf("close file error: %s", err)
		}
		d.file = nil
	}

	return nil
}
