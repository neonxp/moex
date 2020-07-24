package main

import "log"

const (
	OFZ_URL  = `https://iss.moex.com/iss/engines/stock/markets/bonds/boards/TQOB/securities.xml?iss.meta=off&iss.only=securities&securities.columns=SECID,NEXTCOUPON,COUPONVALUE,SECNAME,COUPONPERCENT,MATDATE,COUPONPERIOD`
	CORP_URL = `https://iss.moex.com/iss/engines/stock/markets/bonds/boards/TQCB/securities.xml?iss.meta=off&iss.only=securities&securities.columns=SECID,NEXTCOUPON,COUPONVALUE,SECNAME,COUPONPERCENT,MATDATE,COUPONPERIOD`
)

func main() {
	d1, err := processFile(OFZ_URL)
	if err != nil {
		log.Fatal(err)
	}
	if err := export("офз.csv", d1); err != nil {
		log.Fatal(err)
	}
	d2, err := processFile(CORP_URL)
	if err != nil {
		log.Fatal(err)
	}
	if err := export("корп.csv", d2); err != nil {
		log.Fatal(err)
	}
}

func processFile(url string) ([]*Document, error) {
	in, err := download(url)
	if err != nil {
		return nil, err
	}
	var documents []*Document
	for _, item := range in.Data.Rows.Items {
		d, err := parse(item)
		if err != nil {
			log.Println(err)
			continue
		}
		documents = append(documents, d)
	}
	return documents, nil
}
