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
	"time"
)

func parse(row Row) (*Document, error) {
	date, err := time.Parse("2006-01-02", row.Next)
	if err != nil {
		return nil, err
	}
	endDate, err := time.Parse("2006-01-02", row.EndDate)
	if err != nil {
		return nil, err
	}
	d := &Document{
		ID:      row.ID,
		Name:    row.Name,
		From:    date,
		To:      endDate,
		Percent: row.Percent,
		Coupons: []Coupon{},
	}
	if date.Equal(endDate) {
		d.Add(endDate, row.Value)
	}
	for date.Before(endDate) {
		d.Add(date, row.Value)
		date = date.AddDate(0, 0, row.Period)
	}
	if !date.Equal(endDate) {
		d.Add(endDate, row.Value)
	}
	return d, nil
}