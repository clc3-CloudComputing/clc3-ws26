package main

import "testing"

func TestGetRoundedMinuteUp(t *testing.T) {

	res := getRoundedMinute(1, 30)
	if res != 2 {
		t.Errorf("getRoundedMinute(1, 30) = %d; wanted 2", res)
	}
}

func TestGetRoundedMinuteDown(t *testing.T) {

	res := getRoundedMinute(1, 15)
	if res != 1 {
		t.Errorf("getRoundedMinute(1, 15) = %d; wanted 1", res)
	}
}
