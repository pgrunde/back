package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testbody string = `{| class="sortable"
! Number
! Date
! Alleged Shooter 
! Killed
! Wounded
! Location
! References
|-
|1
|1/1/2015
|Unknown
|0
|5
|Memphis, TN
|
{| class="mw-collapsible mw-collapsed wikitable"
!
|-
|<ref>http://www.wmcactionnews5.com/story/27743361/5-injured-in-new-years-day-party-bus-shooting</ref>
|}    
|-
|2
|1/2/2015
|Unknown
|1
|4
|Savannah, GA
|
{| class="mw-collapsible mw-collapsed wikitable"
!
|-
|<ref>http://savannahnow.com/crime/2015-01-02/police-1-dead-4-injured-overnight-savannah-shooting</ref>
|}
|-
|210
|7/28/2015
|Harley Wilder
|2
|2
|West Frankfort, IL
|
{| class="mw-collapsible mw-collapsed wikitable"
!
|-
|<ref>http://www.wsiltv.com/news/local/Emergency-Crews-Respond-to-Madison-Street-in-West-Frankfort-319177301.html</ref> <ref>http://thesouthern.com/news/the-latest-shooter-kills-one-man-self-injures-two-others/article_40b29eb9-8740-53fb-a377-a080d5b4dcb2.html</ref>
|}

|}
<references/>`

func TestShootingsFromTextArea(t *testing.T) {
	tn := Shooting{
		Number:         1,
		Date:           "1/1/2015",
		AllegedShooter: "Unknown",
		Killed:         0,
		Wounded:        5,
		City:           "Memphis",
		State:          "TN",
		References:     []string{"http://www.wmcactionnews5.com/story/27743361/5-injured-in-new-years-day-party-bus-shooting"},
	}
	ga := Shooting{
		Number:         2,
		Date:           "1/2/2015",
		AllegedShooter: "Unknown",
		Killed:         1,
		Wounded:        4,
		City:           "Savannah",
		State:          "GA",
		References:     []string{"http://savannahnow.com/crime/2015-01-02/police-1-dead-4-injured-overnight-savannah-shooting"},
	}
	il := Shooting{
		Number:         210,
		Date:           "7/28/2015",
		AllegedShooter: "Harley Wilder",
		Killed:         2,
		Wounded:        2,
		City:           "West Frankfort",
		State:          "IL",
		References:     []string{"http://www.wsiltv.com/news/local/Emergency-Crews-Respond-to-Madison-Street-in-West-Frankfort-319177301.html", "http://thesouthern.com/news/the-latest-shooter-kills-one-man-self-injures-two-others/article_40b29eb9-8740-53fb-a377-a080d5b4dcb2.html"},
	}
	expected := []Shooting{tn, ga, il}

	shootings, err := ShootingsFromTextArea(testbody)
	assert.Nil(t, err, "Valid text body should not error %s", err)

	require.Equal(t, 3, len(shootings), "The correct number of shootings should be returned")
	for i, result := range shootings {
		assert.Equal(t, expected[i].Number, result.Number)
		assert.Equal(t, expected[i].Date, result.Date)
		assert.Equal(t, expected[i].AllegedShooter, result.AllegedShooter)
		assert.Equal(t, expected[i].Killed, result.Killed)
		assert.Equal(t, expected[i].Wounded, result.Wounded)
		assert.Equal(t, expected[i].City, result.City)
		assert.Equal(t, expected[i].State, result.State)

		require.Equal(t, len(expected[i].References), len(result.References))
		for j, ref := range expected[i].References {
			assert.Equal(t, ref, result.References[j])
		}
	}
}
