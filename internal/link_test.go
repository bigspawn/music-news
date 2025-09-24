package internal

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	itunes "github.com/bigspawn/go-itunes-api"
	"github.com/go-pkgz/lgr"
	"github.com/stretchr/testify/assert"
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
		{title: "Who Will Fix Me Now? - EP", want: "Who Will Fix Me Now?"},
		{title: "Who Will Fix Me Now - EP? - EP", want: "Who Will Fix Me Now - EP?"},
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

func Test_levenshteinDistance(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected int
	}{
		{
			name:     "identical strings",
			s1:       "test",
			s2:       "test",
			expected: 0,
		},
		{
			name:     "one character difference",
			s1:       "test",
			s2:       "tent",
			expected: 1,
		},
		{
			name:     "one character insertion",
			s1:       "test",
			s2:       "tests",
			expected: 1,
		},
		{
			name:     "one character deletion",
			s1:       "tests",
			s2:       "test",
			expected: 1,
		},
		{
			name:     "completely different strings",
			s1:       "kitten",
			s2:       "sitting",
			expected: 3,
		},
		{
			name:     "empty strings",
			s1:       "",
			s2:       "",
			expected: 0,
		},
		{
			name:     "one empty string",
			s1:       "music",
			s2:       "",
			expected: 5,
		},
		{
			name:     "case difference",
			s1:       "Test",
			s2:       "test",
			expected: 1,
		},
		{
			name:     "unicode characters",
			s1:       "привет",
			s2:       "привет мир",
			expected: 4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := levenshteinDistance(tc.s1, tc.s2)
			assert.Equal(t, tc.expected, result, "levenshtein distance calculation was incorrect")
		})
	}
}

func Test_findCollectionIDFromResultsByTitle(t *testing.T) {
	tests := []struct {
		name     string
		r        itunes.Result
		s        string
		expected string
		wantErr  bool
	}{
		{
			name: "found",
			r: itunes.Result{
				ArtistName:     "Slow Degrade",
				CollectionName: "Who Will Fix Me Now? - EP",
				ReleaseDate:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				CollectionId:   123,
				Kind:           itunes.KindAlbum,
			},
			s:        "Slow Degrade — Who Will Fix Me Now? [EP] (2024)",
			expected: "123",
			wantErr:  false,
		},
		{
			name: "UPFALL - ARTIFICIAL - EP",
			r: itunes.Result{
				ArtistName:     "UPFALL",
				CollectionName: "ARTIFICIAL - EP",
				ReleaseDate:    time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				CollectionId:   123,
				Kind:           itunes.KindAlbum,
			},
			s:        "Upfall - Artificial (EP) (2025)",
			expected: "123",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findCollectionIDFromResultsByTitle(lgr.Default(), []itunes.Result{tt.r}, tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("findCollectionIDFromResultsByTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("findCollectionIDFromResultsByTitle() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestProblematicTitlesFromLog(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		expected string
	}{
		{
			name:     "Cyrillic discography",
			title:    "Soap&Skin - дискография",
			expected: "soap&skin",
		},
		{
			name:     "Cyrillic discography with year",
			title:    "Lorna Shore - дискография",
			expected: "lorna shore",
		},
		{
			name:     "English album with year",
			title:    "The Project Hate MCMXCIX - Undivine Dethroning (2025)",
			expected: "the project hate mcmxcix - undivine dethroning",
		},
		{
			name:     "Cyrillic album with year",
			title:    "Mélancolia - andom.access.misery (2025)",
			expected: "mélancolia - andom.access.misery",
		},
		{
			name:     "Complex title with brackets and year",
			title:    "Wisdom In Chains x Evergreen Terrace - Wisdom In Chains / Evergreen Terrace [split EP] (2025)",
			expected: "wisdom in chains x evergreen terrace - wisdom in chains / evergreen terrace",
		},
		{
			name:     "Album with 2CD marking",
			title:    "Panic Lift - Split Pieces Stitched Together Again [2CD] (2025)",
			expected: "panic lift - split pieces stitched together again",
		},
		{
			name:     "Cyrillic album title with year",
			title:    "Multipass - Песни Осени [2CD] (2013)",
			expected: "multipass - песни осени",
		},
		{
			name:     "Album with special chars",
			title:    "NewTone - The World Сhanges (2025)",
			expected: "newtone - the world сhanges",
		},
		{
			name:     "Take It Down with Cyrillic",
			title:    "Take It Down - Культ (2025)",
			expected: "take it down - культ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := clearTitle(strings.ToLower(tt.title))
			if result != tt.expected {
				t.Errorf("clearTitle(%s) = %s, want %s", tt.title, result, tt.expected)
			}
		})
	}
}

func TestGenerateSearchVariants(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		expected []string
	}{
		{
			name:  "Simple artist - album format",
			title: "Artist Name - Album Name",
			expected: []string{
				"Artist Name - Album Name",
				"Artist Name",
			},
		},
		{
			name:  "Title with year",
			title: "Artist - Album (2025)",
			expected: []string{
				"Artist - Album (2025)",
				"Artist - Album",
				"Artist",
			},
		},
		{
			name:  "Title with discography",
			title: "Artist - дискография",
			expected: []string{
				"Artist - дискография",
				"Artist",
				"Artist",
			},
		},
		{
			name:  "Complex title",
			title: "Take It Down - Культ (2025)",
			expected: []string{
				"Take It Down - Культ (2025)",
				"Take It Down - Культ",
				"Take It Down",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateSearchVariants(tt.title)
			if len(result) != len(tt.expected) {
				t.Errorf("generateSearchVariants(%s) returned %d variants, want %d. Got: %v",
					tt.title, len(result), len(tt.expected), result)
				return
			}

			// Check that all expected variants are present
			for i, expected := range tt.expected {
				if i < len(result) && result[i] != expected {
					t.Errorf("generateSearchVariants(%s)[%d] = %s, want %s",
						tt.title, i, result[i], expected)
				}
			}
		})
	}
}
