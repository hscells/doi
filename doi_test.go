package doi

import (
	"log"
	"testing"
)

func TestParseDoiSuccessNumbers(t *testing.T) {
	_, err := Parse("10.1000/123456")
	if err != nil {
		log.Println(err)
		t.Fail()
	}
}

func TestParseDoiSuccessISSN(t *testing.T) {
	_, err := Parse("10.1038/issn.1476-4687")
	if err != nil {
		log.Println(err)
		t.Fail()
	}
}

func TestParseDoiFailIncorrectGeneral(t *testing.T) {
	_, err := Parse("11.1038/123456")
	if err == nil {
		t.Fail()
	}
}

func TestParseDoiFailNoRegistrantCode(t *testing.T) {
	_, err := Parse("10.1038")
	if err == nil {
		t.Fail()
	}
}

func TestParseDoiFailMissingRegistrantCode(t *testing.T) {
	_, err := Parse("10.1038/")
	if err == nil {
		t.Fail()
	}
}

func TestParseDoiFailMissingGeneral(t *testing.T) {
	_, err := Parse(".1038/123456")
	if err == nil {
		t.Fail()
	}
}

func TestParseDoiFailMissingDirectoryIndicator(t *testing.T) {
	_, err := Parse("10./123456")
	if err == nil {
		t.Fail()
	}
}

func TestIsValidSuccess(t *testing.T) {
	identifier := DigitalObjectIdentifier{
		General:            "10",
		DirectoryIndicator: "1038",
		RegistrantCode:     "213456",
	}
	if !identifier.IsValid() {
		t.Fail()
	}
}

func TestIsValidFailInvalidGeneral(t *testing.T) {
	identifier := DigitalObjectIdentifier{
		General:            "11",
		DirectoryIndicator: "1038",
		RegistrantCode:     "213456",
	}
	if identifier.IsValid() {
		t.Fail()
	}
}

func TestIsValidFailMissingGeneral(t *testing.T) {
	identifier := DigitalObjectIdentifier{
		DirectoryIndicator: "1038",
		RegistrantCode:     "213456",
	}
	if identifier.IsValid() {
		t.Fail()
	}
}

func TestIsValidFailMissingDirectoryIndicator(t *testing.T) {
	identifier := DigitalObjectIdentifier{
		General:        "10",
		RegistrantCode: "213456",
	}
	if identifier.IsValid() {
		t.Fail()
	}
}

func TestIsValidFailMissingRegistrantCode(t *testing.T) {
	identifier := DigitalObjectIdentifier{
		General:            "10",
		DirectoryIndicator: "1038",
	}
	if identifier.IsValid() {
		t.Fail()
	}
}

func TestToStringValidDoi(t *testing.T) {
	identifier := DigitalObjectIdentifier{
		General:            "10",
		DirectoryIndicator: "1038",
		RegistrantCode:     "213456",
	}
	if s, err := identifier.ToString(); err != nil {
		t.Fail()
		if s != "10.1038/213456" {
			t.Fail()
		}
	}
}

func TestToStringInvalidDoi(t *testing.T) {
	identifier := DigitalObjectIdentifier{
		General:            "11",
		DirectoryIndicator: "1038",
		RegistrantCode:     "213456",
	}
	if _, err := identifier.ToString(); err == nil {
		t.Fail()
	}
}
