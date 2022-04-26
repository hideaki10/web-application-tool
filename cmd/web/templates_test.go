package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {

	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2022, 12, 17, 10, 0, 0, 0, time.UTC),
			want: "17 Dec 2022 at 10:00",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			// UTC -> JST 10:00 -> 19:00
			name: "JST",
			tm:   time.Date(2022, 12, 17, 10, 0, 0, 0, time.FixedZone("JST", -9*60*60)),
			want: "17 Dec 2022 at 19:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			if hd != tt.want {
				t.Errorf("want %q; got %q", tt.want, hd)
			}
		})
	}

	time := time.Date(2022, 4, 26, 10, 0, 0, 0, time.UTC)
	hm := humanDate(time)

	if hm != "26 Apr 2022 at 10:00" {
		t.Errorf("want %q; got %q", "26 Apr 2022 at 10:00", hm)
	}
}
