package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Set the base url and read the coupon codes
	baseUrl := "https://p-bifrostbackendnest.dsgapps.dk/api/vouchers-verify?tenantId=6&voucherCode="
	couponCodes, err := os.ReadFile("coupon-codes.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// String manipulation (I CANT)
	couponCodesString := string(couponCodes)
	couponCodesString = strings.ReplaceAll(couponCodesString, "\r", "")
	couponCodesList := strings.Split(couponCodesString, "\n")
	amountOfCodes := len(couponCodesList)

	for i, code := range couponCodesList {
		if i%1000 == 0 {
			fmt.Println(fmt.Sprintf("%d/%d checked", i, amountOfCodes))
		}

		resp, err := http.Get(baseUrl + code)
		if err != nil {
			log.Fatalln(err)
		}

		if resp.StatusCode >= 400 && resp.StatusCode <= 499 {
			continue
		}

		fmt.Println(fmt.Sprintf("Code: %s with status code %d", code, resp.StatusCode))
	}
}
