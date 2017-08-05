package go_doi

import (
	"testing"
	"log"
)

func TestParseDoiSuccessNumbers(t *testing.T) {
	doiString := "10.1000/123456"
	_, err := Parse(doiString)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
}

func TestParseDoiSuccessISSN(t *testing.T) {
	doiString := "10.1038/issn.1476-4687"
	_, err := Parse(doiString)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
}

func TestParseDoiFailIncorrectGeneral(t *testing.T) {
	doiString := "11.1038/123456"
	_, err := Parse(doiString)
	if err == nil {
		t.Fail()
	}
}

func TestParseDoiFailNoRegistrantCode(t *testing.T) {
	doiString := "10.1038"
	_, err := Parse(doiString)
	if err == nil {
		t.Fail()
	}
}

func TestParseDoiFailMissingRegistrantCode(t *testing.T) {
	doiString := "10.1038/"
	_, err := Parse(doiString)
	if err == nil {
		t.Fail()
	}
}

func TestParseDoiFailMissingGeneral(t *testing.T) {
	doiString := ".1038/123456"
	_, err := Parse(doiString)
	if err == nil {
		t.Fail()
	}
}

func TestParseDoiFailMissingDirectoryIndicator(t *testing.T) {
	doiString := "10./123456"
	_, err := Parse(doiString)
	if err == nil {
		t.Fail()
	}
}

func TestIsValidSuccess(t *testing.T) {
	doi := Doi{
		General:            "10",
		DirectoryIndicator: "1038",
		RegistrantCode:     "213456",
	}
	if !doi.IsValid() {
		t.Fail()
	}
}

func TestIsValidFailInvalidGeneral(t *testing.T) {
	doi := Doi{
		General:            "11",
		DirectoryIndicator: "1038",
		RegistrantCode:     "213456",
	}
	if doi.IsValid() {
		t.Fail()
	}
}

func TestIsValidFailMissingGeneral(t *testing.T) {
	doi := Doi{
		DirectoryIndicator: "1038",
		RegistrantCode:     "213456",
	}
	if doi.IsValid() {
		t.Fail()
	}
}

func TestIsValidFailMissingDirectoryIndicator(t *testing.T) {
	doi := Doi{
		General:        "10",
		RegistrantCode: "213456",
	}
	if doi.IsValid() {
		t.Fail()
	}
}

func TestIsValidFailMissingRegistrantCode(t *testing.T) {
	doi := Doi{
		General:            "10",
		DirectoryIndicator: "1038",
	}
	if doi.IsValid() {
		t.Fail()
	}
}

func TestToStringValidDoi(t *testing.T) {
	doi := Doi{
		General:            "10",
		DirectoryIndicator: "1038",
		RegistrantCode:     "213456",
	}
	if s, err := doi.ToString(); err != nil {
		t.Fail()
		if s != "10.1038/213456" {
			t.Fail()
		}
	}
}

func TestToStringInvalidDoi(t *testing.T) {
	doi := Doi{
		General:            "11",
		DirectoryIndicator: "1038",
		RegistrantCode:     "213456",
	}
	if _, err := doi.ToString(); err == nil {
		t.Fail()
	}
}
