package main

import (
    "testing"
)

func TestLeapYear01(t *testing.T) {
    expected := leapyear(2004)
    actual := true
    if expected != actual {
        t.Errorf("Error")
    }
}
func TestLeapYear02(t *testing.T) {
    expected := leapyear(2100)
    actual := false
    if expected != actual {
        t.Errorf("Error")
    }
}
func TestLeapYear03(t *testing.T) {
    expected := leapyear(2001)
    actual := false
    if expected != actual {
        t.Errorf("Error")
    }
}
func TestLeapYear04(t *testing.T) {
    expected := leapyear(2000)
    actual := true
    if expected != actual {
        t.Errorf("Error")
    }
}