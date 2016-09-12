package bank

// FIBanks is a list of finnish banks by flk
var FIBanks = []Bank{
	Bank{
		ID:   1,
		Name: "Nordea",
		Pad:  6,
	},
	Bank{
		ID:   2,
		Name: "Nordea",
		Pad:  6,
	},
	Bank{
		ID:   31,
		Name: "Handelsbanken",
		Pad:  6,
	},
	Bank{
		ID:   33,
		Name: "Skandinaviska Enskilda Banken (SEB)",
		Pad:  6,
	},
	Bank{
		ID:   34,
		Name: "Danske Bank",
		Pad:  6,
	},
	Bank{
		ID:   36,
		Name: "Tapiola Pankki (Tapiola)",
		Pad:  6,
	},
	Bank{
		ID:   37,
		Name: "DnB NOR Bank ASA (DnB NOR)",
		Pad:  6,
	},
	Bank{
		ID:   38,
		Name: "Swedbank",
		Pad:  6,
	},
	Bank{
		ID:   39,
		Name: "S-Pankki",
		Pad:  6,
	},
	Bank{
		ID:   4,
		Name: "Säästöpankit (Sp) ja paikallisosuuspankit (Pop) sekä Aktia",
		Pad:  7,
	},
	Bank{
		ID:   5,
		Name: "Osuuspankit (Op), OKO ja Okopankki",
		Pad:  7,
	},
	Bank{
		ID:   6,
		Name: "Ålandsbanken (ÅAB)",
		Pad:  6,
	},
	Bank{
		ID:   8,
		Name: "Sampo Pankki (Sampo)",
		Pad:  6,
	},
	Bank{
		ID:   405,
		Name: "Aktia Pankki",
		Pad:  6,
	},
	Bank{
		ID:   497,
		Name: "Aktia Pankki",
		Pad:  6,
	},
	Bank{
		ID:   470,
		Name: "POP Pankit",
		Pad:  6,
	},
	Bank{
		ID:   479,
		Name: "Bonum",
		Pad:  6,
	},
	Bank{
		ID:   713,
		Name: "Citibank",
		Pad:  6,
	},
	Bank{
		ID:   799,
		Name: "Holvi",
		Pad:  6,
	},
}

func FindBankByID(checksum int) *Bank {
	for _, b := range FIBanks {
		if checksum == b.ID {
			return &b
		}
	}
	return nil
}
