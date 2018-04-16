package gtin

import (
	"testing"
)

func TestCheckSum(t *testing.T) {
	if IsGTIN("2388060103489") != true {
		t.Errorf("Checksum should be correct")
	}
	if IsGTIN("2388060202977") != true {
		t.Errorf("Checksum should be correct")
	}
	if IsGTIN("2388060310856") != true {
		t.Errorf("Checksum should be correct")
	}
	if IsGTIN("97350053850043") != true {
		t.Errorf("Checksum should be correct")
	}
	if IsGTIN("012345678905") != true { // GTIN12
		t.Errorf("Checksum should be correct")
	}
	if IsGTIN("00012345678905") != true { // GTIN12 as GTIN14
		t.Errorf("Checksum should be correct")
	}
	if IsGTIN("0123456789012") != true { // GTIN13
		t.Errorf("Checksum should be correct")
	}
	if IsGTIN("00123456789012") != true { // GTIN13 as GTIN14
		t.Errorf("Checksum should be correct")
	}
	if IsGTIN("96385074") != true { // GTIN-8
		t.Errorf("Checksum should be correct")
	}
	if IsGTIN("00000096385074") != true { // GTIN-8 as GTIN-14
		t.Errorf("Checksum should be correct")
	}
	if IsGTIN("00012345600012") != true { // GTIN-14
		t.Errorf("Checksum should be correct")
	}
	if IsGTIN("12") != false { // wrong
		t.Errorf("Checksum should be wrong")
	}
	if IsGTIN("abc") != false { // wrong
		t.Errorf("Checksum should be wrong")
	}
	if IsGTIN("asgdhddfhbthb") != false { // wrong
		t.Errorf("Checksum should be wrong")
	}
	if IsGTIN("123475365ufdh") != false { // wrong
		t.Errorf("Checksum should be wrong")
	}
	if IsGTIN("000012345600012") != false { // wrong
		t.Errorf("Checksum should be wrong")
	}
	if IsGTIN("00012345600013") != false { // wrong
		t.Errorf("Checksum should be wrong")
	}
}
