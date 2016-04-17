package card

import (
	"github.com/asukakenji/clash-royale/arena"
	"github.com/asukakenji/clash-royale/attr"
	"github.com/asukakenji/clash-royale/lib"
	"github.com/asukakenji/clash-royale/rarity"
	"github.com/asukakenji/clash-royale/typ"
)

var spells = []*card{
	// --- Common Spells ---
	newCard(2000, map[attr.Attribute]interface{}{
		attr.Name:     "Arrows",
		attr.Arena:    arena.Arena0,
		attr.Rarity:   rarity.Common,
		attr.Type:     typ.Spell,
		attr.Desc:     `Arrows pepper a large area, damaging everyone hit. Reduced damage to Crown Towers.`,
		attr.Elixir:   3,
		attr.BaseADam: 115,
		attr.CTDam:    []int{46, 51, 56, 61, 67, 74, 81, 89, 98, 107, 118, 130},
		attr.Radius:   4,
	}),
	newCard(2050, map[attr.Attribute]interface{}{
		attr.Name:     "Zap",
		attr.Arena:    arena.Arena5,
		attr.Rarity:   rarity.Common,
		attr.Type:     typ.Spell,
		attr.Desc:     `Zaps enemies, briefly stunning them and dealing damage inside a small radius. Reduced damage to Crown Towers.`,
		attr.Elixir:   2,
		attr.BaseADam: 80,
		attr.CTDam:    []int{32, 36, 39, 43, 47, 52, 56, 62, 68, 75, 82, 90},
		attr.DurF:     1,
		attr.Radius:   2.5,
	}),

	// --- Rare Spells ---
	newCard(2100, map[attr.Attribute]interface{}{
		attr.Name:     "Fireball",
		attr.Arena:    arena.Arena0,
		attr.Rarity:   rarity.Rare,
		attr.Type:     typ.Spell,
		attr.Desc:     `Annnnnd... Fireball. Incinerates a small area, dealing high damage. Reduced damage to Crown Towers.`,
		attr.Elixir:   4,
		attr.BaseADam: 325,
		attr.CTDam:    []int{130, 143, 158, 173, 190, 208, 229, 251, 276, 303},
		attr.Radius:   2.5,
	}),
	newCard(2130, map[attr.Attribute]interface{}{
		attr.Name:     "Rocket",
		attr.Arena:    arena.Arena3,
		attr.Rarity:   rarity.Rare,
		attr.Type:     typ.Spell,
		attr.Desc:     `Deals high damage to a small area. Looks really awesome doing it. Reduced damage to Crown Towers.`,
		attr.Elixir:   6,
		attr.BaseADam: 700,
		attr.CTDam:    []int{280, 308, 339, 373, 409, 448, 493, 541, 594, 653},
		attr.Radius:   2,
	}),

	// --- Epic Spells ---
	newCard(2210, map[attr.Attribute]interface{}{
		attr.Name:    "Lightning",
		attr.Arena:   arena.Arena1,
		attr.Rarity:  rarity.Epic,
		attr.Type:    typ.Spell,
		attr.Desc:    `Bolts of lightning hit up to three enemy troops or buildings with the most hitpoints in the target area. Reduced damage to Crown Towers.`,
		attr.Elixir:  6,
		attr.Count:   3,
		attr.BaseDam: 650,
		attr.CTDam:   []int{260, 286, 315, 346, 380, 416, 458, 502},
		attr.DurF:    1.5,
		attr.Radius:  3.5,
	}),
	newCard(2211, map[attr.Attribute]interface{}{
		attr.Name:      "Goblin Barrel",
		attr.Arena:     arena.Arena1,
		attr.Rarity:    rarity.Epic,
		attr.Type:      typ.Spell,
		attr.Desc:      `Spawns three Goblins anywhere on the Arena. It's going to be a thrilling ride, boys!`,
		attr.Elixir:    4,
		attr.ADam:      lib.V50[0:8:8],
		attr.CTDam:     []int{20, 22, 24, 27, 30, 32, 36, 39},
		attr.Radius:    1.5,
		attr.BaseGobLV: 6,
		attr.GobCount:  3,
	}),
	newCard(2230, map[attr.Attribute]interface{}{
		attr.Name:   "Rage",
		attr.Arena:  arena.Arena3,
		attr.Rarity: rarity.Epic,
		attr.Type:   typ.Spell,
		attr.Desc:   `Increases troop movement and attack speed by 40%. Troop buildings and summoners deploy troop faster. Chaaaarge!`,
		attr.Elixir: 3,
		attr.DurU:   []float64{8, 8.5, 9, 9.5, 10, 10.5, 11, 11.5},
		attr.Radius: 5,
	}),
	newCard(2240, map[attr.Attribute]interface{}{
		attr.Name:   "Freeze",
		attr.Arena:  arena.Arena4,
		attr.Rarity: rarity.Epic,
		attr.Type:   typ.Spell,
		attr.Desc:   `Freezes troops and buildings, making them unable to move or attack. Everybody chill.`,
		attr.Elixir: 4,
		attr.DurU:   []float64{5, 5.3, 5.6, 5.9, 6.2, 6.5, 6.8, 7.1},
		attr.Radius: 3,
	}),
	newCard(2250, map[attr.Attribute]interface{}{
		attr.Name:     "Mirror",
		attr.Arena:    arena.Arena5,
		attr.Rarity:   rarity.Epic,
		attr.Type:     typ.Spell,
		attr.Desc:     `Mirrors your last card played for +1 Elixir`,
		attr.Elixir:   lib.X,
		attr.BaseMCLV: 5,
		attr.BaseMRLV: 3,
		attr.BaseMELV: 1,
		attr.MLLV:     []int{1, 1, 1, 1, 2, 3, 4, 5},
	}),
	newCard(2251, map[attr.Attribute]interface{}{
		attr.Name:   "Poison",
		attr.Arena:  arena.Arena5,
		attr.Rarity: rarity.Epic,
		attr.Type:   typ.Spell,
		attr.Desc:   `Covers the target area in a sticky toxin, damaging and slowing down troops and buildings. Remember: solvent abuse can kill!`,
		attr.Elixir: 4,
		attr.DPS:    []int{42, 46, 50, 55, 62, 67, 73, 81},
		attr.CTDPS:  []int{17, 19, 20, 22, 25, 27, 30, 33},
		attr.DurF:   10,
		attr.Radius: 3.5,
	}),
}