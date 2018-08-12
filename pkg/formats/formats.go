package formats

type Formatter interface {
	Format([]byte) string
	Parse(string) []byte
}

var formatMap = map[uint16]map[uint8][]Formatter{
	/* Bitcoin Cash (Bitcoin address) */
	0: {
		/* P2PKH address */
		0: {
			NewBase58Formatter(0),
		},

		/* P2SH address */
		1: {
			NewBase58Formatter(5),
		},
	},

	/* Nano/RaiBlocks (Nano address) */
	1: {
		/* Standard address */
		0: {
			NewNanoFormatter("xrb_"),
			NewNanoFormatter("nano_"),
		},
	},
}

func Format(bs []byte, typ uint16, subType uint8) []string {
	var (
		subTypeMap map[uint8][]Formatter
		formatters []Formatter
		ok bool
	)

	if subTypeMap, ok = formatMap[typ]; !ok {
		return nil
	}

	if formatters, ok = subTypeMap[subType]; !ok {
		return nil
	}

	list := make([]string, 0, len(formatters))
	for _, formatter := range formatters {
		if str := formatter.Format(bs); len(str) > 0 {
			list = append(list, str)
		}
	}
	return list
}
