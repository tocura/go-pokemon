package enum

type GameVersion string

const (
	RedGameVersion        GameVersion = "RED"
	BlueGameVersion       GameVersion = "BLUE"
	YellowGameVersion     GameVersion = "YELLOW"
	GoldGameVersion       GameVersion = "GOLD"
	SilverGameVersion     GameVersion = "SILVER"
	CrystalGameVersion    GameVersion = "CRYSTAL"
	RubyGameVersion       GameVersion = "RUBY"
	SapphireGameVersion   GameVersion = "SAPPHIRE"
	EmeralGameVersion     GameVersion = "EMERALD"
	FireRedGameVersion    GameVersion = "FIRERED"
	LeafGreenGameVersion  GameVersion = "LEAFGREEN"
	DiamondGameVersion    GameVersion = "DIAMOND"
	PearlGameVersion      GameVersion = "PEARL"
	PlatinumGameVersion   GameVersion = "PLATINUM"
	HeartGoldGameVersion  GameVersion = "HEARTGOLD"
	SoulSilverGameVersion GameVersion = "SOULSILVER"
	BlackGameVersion      GameVersion = "BLACK"
	WhiteGameVersion      GameVersion = "WHITE"
)

func (gv GameVersion) ToString() string {
	return string(gv)
}
