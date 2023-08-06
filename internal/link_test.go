package internal

import (
	"context"
	"fmt"
	"testing"

	itunes "github.com/bigspawn/go-itunes-api"
	"github.com/stretchr/testify/require"
)

func Test_clearTitle(t *testing.T) {
	tests := []struct {
		title string
		want  string
	}{
		{title: "Bones - Gate of Night [EP] (2020)", want: "Bones - Gate of Night"},
		{title: "The Azimuth - Beholder (2020)", want: "The Azimuth - Beholder"},
		{title: "Hegemony - Enthroned by Persecution (2020)", want: "Hegemony - Enthroned by Persecution"},
		{title: "Goldmund Quartet - Travel Diaries (2020)", want: "Goldmund Quartet - Travel Diaries"},
		{title: "Primal Aggression - Sinister (2020)", want: "Primal Aggression - Sinister"},
		{title: "Sarah Longfield - Dusk (2020)", want: "Sarah Longfield - Dusk"},
		{title: "LOWLIVES - Bones (2020)", want: "LOWLIVES - Bones"},
		{title: "Lyka - MK I (2020)", want: "Lyka - MK I"},
		{title: "Nvision - She Said [Single] (2020)", want: "Nvision - She Said"},
		{title: "AliT - Festett Árnyak (2020)", want: "AliT - Festett Árnyak"},
		{title: "LOWLIVES - Bones [Single] (2020)", want: "LOWLIVES - Bones"},
		{title: "One Hundred Thousand - Taurus (Single) (2020)", want: "One Hundred Thousand - Taurus"},
		{title: "Car Seat Headrest - Making a Door Less Open (2020)", want: "Car Seat Headrest - Making a Door Less Open"},
		{title: "Plini - Birds / Surfers [Single] (2020)", want: "Plini - Birds / Surfers"},
		{title: "Eye of Nix - Concealing Waters [Single] (2020)", want: "Eye of Nix - Concealing Waters"},
		{title: "Mark Lanegan - Straight Songs of Sorrow (2020)", want: "Mark Lanegan - Straight Songs of Sorrow"},
		{title: "Beyond Unbroken - Falling Down + Heathens [Single] (2020)", want: "Beyond Unbroken - Falling Down + Heathens"},
		{title: "Her Bright Skies - IB4U [Single] (2020)", want: "Her Bright Skies - IB4U"},
		{title: "Kodaline - Saving Grace [Single] (2020)", want: "Kodaline - Saving Grace"},
		{title: "Bleed From Within - Night Crossing [Single] (2020)", want: "Bleed From Within - Night Crossing"},
		{title: "Suicide Silence - Overlord [Single] (2020)", want: "Suicide Silence - Overlord"},
		{title: "Breakthrough Even - Requiem [Single] (2020)", want: "Breakthrough Even - Requiem"},
		{title: "Boys of Fall - Overthinking [Single] (2020)", want: "Boys of Fall - Overthinking"},
		{title: "KITSUNE - GRIEF [Single] (2020)", want: "KITSUNE - GRIEF"},
		{title: "Two Year Break - #bitbright [Single] (2020)", want: "Two Year Break - #bitbright"},
		{title: "Next Time Mr. Fox - Basilisk [Single] (2020)", want: "Next Time Mr. Fox - Basilisk"},
		{title: "Poeta - Anxious Racing [Single] (2020)", want: "Poeta - Anxious Racing"},
		{title: "Off Road Minivan - YOU [Single] (2020)", want: "Off Road Minivan - YOU"},
		{title: "O'Brother - You and I (2020)", want: "O'Brother - You and I"},
		{title: "Tiberius - Mechanical Messiah [Single] (2020)", want: "Tiberius - Mechanical Messiah"},
		{title: "Trenchworm - Constant Death [Single] (2020)", want: "Trenchworm - Constant Death"},
		{title: "InSights - Insights [EP] (2020)", want: "InSights - Insights"},
		{title: "Octobers - Summer Waste EP (2020)", want: "Octobers - Summer Waste EP"},
		{title: "Octobers - Misfits EP (2015)", want: "Octobers - Misfits EP"},
		{title: "A Fall To Break - The Man in the Mask (2011)", want: "A Fall To Break - The Man in the Mask"},
		{title: "A Fall To Break - September Falls (2012)", want: "A Fall To Break - September Falls"},
		{title: "Outline In Color - Quarantine Sessions (2020)", want: "Outline In Color - Quarantine Sessions"},
		{title: "Smoke Signals - Volume One (Forsaken) [EP] (2020)", want: "Smoke Signals - Volume One (Forsaken)"},
		{title: "Biesy - Transsatanizm (2020)", want: "Biesy - Transsatanizm"},
		{title: "Want - Tilt [Single] (2020)", want: "Want - Tilt"},
		{title: "Dead Days - Black Summer [Single] (2020)", want: "Dead Days - Black Summer"},
		{title: "Color in the Clouds - In Melancholy [Single] (2020)", want: "Color in the Clouds - In Melancholy"},
		{title: "Dream Awake - Prosper [EP] (2020)", want: "Dream Awake - Prosper"},
		{title: "Beyond The Black - Wounded Healer [Single] (2020)", want: "Beyond The Black - Wounded Healer"},
		{title: "Regimen De Terror - Inherente Del Poder [EP] (2020)", want: "Regimen De Terror - Inherente Del Poder"},
		{title: "Within Temptation - Entertain You [Single] (2020)", want: "Within Temptation - Entertain You"},
		{title: "Vatic - Departure [Single] (2020)", want: "Vatic - Departure"},
		{title: "Glass Tides - Sew Your Mouth Shut [Single] (2020)", want: "Glass Tides - Sew Your Mouth Shut"},
		{title: "The Motion Below - Truth Hurts [Single] (2020)", want: "The Motion Below - Truth Hurts"},
		{title: "Of Colors (feat. Dennis Landt) - Bleak [Single] (2020)", want: "Of Colors (feat. Dennis Landt) - Bleak"},
		{title: "Relent - LOW [Single] (2020)", want: "Relent - LOW"},
		{title: "Fractures and Outlines - Kerosene (feat. Jericho Spencer-Champagne) [Single] (2020)", want: "Fractures and Outlines - Kerosene (feat. Jericho Spencer-Champagne)"},
		{title: "Chasing Apparitions - As Above, So Below [Single] (2020)", want: "Chasing Apparitions - As Above, So Below"},
		{title: "Butch Walker - American Love Story (2020)", want: "Butch Walker - American Love Story"},
		{title: "Blacklab - Abyss (2020)", want: "Blacklab - Abyss"},
		{title: "Combos - Steelo (2020)", want: "Combos - Steelo"},
		{title: "Omniarch - Omniarch (2020)", want: "Omniarch - Omniarch"},
		{title: "Foad - Returner (2020)", want: "Foad - Returner"},
		{title: "Trees Will Tell - Negative Results (2020)", want: "Trees Will Tell - Negative Results"},
		{title: "Agriculture - Agriculture (2023)", want: "Agriculture - Agriculture"},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			if got := clearTitle(tt.title); got != tt.want {
				t.Errorf("clearTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName(t *testing.T) {
	t.Skip()

	api, err := itunes.NewClient(itunes.ClientOption{})
	require.NoError(t, err)
	resp, err := api.Search(context.Background(), itunes.SearchRequest{
		Term:    "Agriculture - Agriculture",
		Country: itunes.US,
	})
	require.NoError(t, err)
	fmt.Println(resp)
}
