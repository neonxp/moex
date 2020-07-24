/*
Copyright © 2020 Alexander Kiryukhin <a.kiryukhin@mail.ru>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

const months = 48

func export(name string, documents []*Document) error {
	var sheet [][]string
	header := []string{"Облигации", "Тикер"}
	now := time.Now()
	for i := 0; i <= months; i++ {
		year, month, _ := now.Date()
		cellID := fmt.Sprintf("%02d.%d", month, year)
		header = append(header, cellID)
		now = now.AddDate(0, 1, 0)
	}
	sheet = append(sheet, header)
	for _, document := range documents {
		row := []string{document.Name, document.ID}
		now = time.Now()
		for i := 0; i <= months; i++ {
			if sum, ok := document.GetByDate(now); ok {
				row = append(row, strconv.FormatFloat(sum, 'f', 2, 64))
			} else {
				row = append(row, "")
			}
			now = now.AddDate(0, 1, 0)
		}
		sheet = append(sheet, row)
	}
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	if err := w.WriteAll(sheet); err != nil {
		return err
	}

	if err := w.Error(); err != nil {
		return err
	}
	return nil
}
