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
	"encoding/xml"
	"time"
)

type InXML struct {
	XMLName xml.Name `xml:"document"`
	Data    struct {
		XMLName xml.Name `xml:"data"`
		Rows    struct {
			XMLName xml.Name `xml:"rows"`
			Items   []Row    `xml:"row"`
		} `xml:"rows"`
	} `xml:"data"`
}

type Row struct {
	XMLName xml.Name `xml:"row"`
	ID      string   `xml:"SECID,attr"`
	Name    string   `xml:"SECNAME,attr"`
	Next    string   `xml:"NEXTCOUPON,attr"`
	Value   float64  `xml:"COUPONVALUE,attr"`
	Percent float64  `xml:"COUPONPERCENT,attr"`
	EndDate string   `xml:"MATDATE,attr"`
	Period  int      `xml:"COUPONPERIOD,attr"`
}

type Document struct {
	ID      string
	Name    string
	From    time.Time
	To      time.Time
	Percent float64
	Coupons []Coupon
}

func (d *Document) Add(date time.Time, sum float64) {
	d.Coupons = append(d.Coupons, Coupon{
		Date: date,
		Sum:  sum,
	})
}

func (d *Document) GetByDate(date time.Time) (float64, bool) {
	year, month, _ := date.Date()
	for _, coupon := range d.Coupons {
		cyear, cmonth, _ := coupon.Date.Date()
		if year == cyear && month == cmonth {
			return coupon.Sum, true
		}
	}
	return 0, false
}

type Coupon struct {
	Date time.Time
	Sum  float64
}
