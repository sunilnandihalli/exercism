package twelve

import "strings"

// On the twelfth day of Christmas my true love gave to me:

func Verse(i int) string {
	gifts := []string{"twelve Drummers Drumming", "eleven Pipers Piping", "ten Lords-a-Leaping", "nine Ladies Dancing", "eight Maids-a-Milking", "seven Swans-a-Swimming", "six Geese-a-Laying", "five Gold Rings", "four Calling Birds", "three French Hens", "two Turtle Doves", "a Partridge in a Pear Tree"}
	day_ids := []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
	var gift_string string
	if i == 1 {
		gift_string = gifts[12-i]
	} else {
		gift_string = strings.Join(gifts[12-i:11], ", ") + ", and " + gifts[11]
	}

	return "On the " + day_ids[i-1] + " day of Christmas my true love gave to me: " + gift_string + "."

}

func Song() string {
	var verses [12]string
	for i := 1; i <= 12; i++ {
		verses[i-1] = Verse(i)
	}
	return strings.Join(verses[:], "\n")
}
