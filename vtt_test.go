package subtitles

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsVTT(t *testing.T) {
	expected := "WEBVTT\n" +
		"\n" +
		"00:00:04.630 --> 00:00:06.018\n" +
		"Go ninja!\n" +
		"\n" +
		"00:01:09.630 --> 00:01:11.005\n" +
		"No ninja!\n\n"

	in := Subtitle{[]Caption{{
		1,
		MakeTime(0, 0, 4, 630),
		MakeTime(0, 0, 6, 18),
		[]string{"Go ninja!"},
	}, {
		2,
		MakeTime(0, 1, 9, 630),
		MakeTime(0, 1, 11, 005),
		[]string{"No ninja!"},
	}}}

	assert.Equal(t, expected, in.AsVTT())
}

func ExampleNewFromSRT() {
	in := "1\n" +
		"00:00:04,630 --> 00:00:06,018\n" +
		"Go ninja!\n" +
		"\n" +
		"1\n" +
		"00:01:09,630 --> 00:01:11,005\n" +
		"No ninja!\n"

	res, _ := NewFromSRT(in)

	// Output: WEBVTT
	//
	// 00:00:04.630 --> 00:00:06.018
	// Go ninja!
	//
	// 00:01:09.630 --> 00:01:11.005
	// No ninja!
	fmt.Println(res.AsVTT())
}
