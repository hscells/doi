package go_doi

import (
	"github.com/pkg/errors"
)

// Doi is a struct that contains the three fields of a doi:
// https://www.doi.org/doi_handbook/2_Numbering.html#2.2
type Doi struct {
	General            string
	DirectoryIndicator string
	RegistrantCode     string
}

// Parse takes a string as input and attempts to parse a valid Doi
// from it. The parsing is done to conform to the standard outlined in
// https://www.doi.org/doi_handbook/2_Numbering.html#2.2.
func Parse(d string) (Doi, error) {
	var general, directoryIndicator, registrantCode string
	state := 0

	for _, c := range d {
		if state == 0 && c == '.' {
			state++
			continue
		} else if state == 1 && c == '/' {
			state++
			continue
		}

		switch state {
		case 0:
			general += string(c)
		case 1:
			directoryIndicator += string(c)
		case 2:
			registrantCode += string(c)
		}
	}

	if general != "10" {
		return Doi{}, errors.New("DOI does not start with 10.")
	} else if len(directoryIndicator) == 0 || len(registrantCode) == 0 {
		return Doi{}, errors.New("Directory Indicator or Registrant Code was empty.")
	}

	return Doi{
		General:            general,
		DirectoryIndicator: directoryIndicator,
		RegistrantCode:     registrantCode,
	}, nil

}

// IsValid checks to see if a Doi is valid or not.
func (d Doi) IsValid() bool {
	if d.General != "10" {
		return false
	} else if len(d.DirectoryIndicator) == 0 || len(d.RegistrantCode) == 0 {
		return false
	}

	return true
}

// ToString creates a string representation of a Doi.
func (d Doi) ToString() (string, error) {
	if d.IsValid() {
		return d.General + "." + d.DirectoryIndicator + "/" + d.RegistrantCode, nil
	}
	return "", errors.New("DOI is invalid, not printable.")
}
