package main

type piece struct {
	white bool
	piece int
}

const (
	none = iota
	pawn
	knight
	bishop
	rook
	queen
	king
)

type board [8][8]piece // a-z 1-8

func NewBoard() *board {
	var b board
	for c := 0; c < 8; c++ {
		for r := 0; r < 8; r++ {
			b[c][r] = piece{false, none}
		}
	}

	setPawns := func(r int, white bool) {
		for c := 0; c < 8; c++ {
			b[c][r] = piece{white, pawn}
		}
	}
	setPawns(1, true)
	setPawns(6, false)

	homeRow := func(r int, white bool) {
		b[0][r] = piece{white, rook}
		b[1][r] = piece{white, knight}
		b[2][r] = piece{white, bishop}
		b[3][r] = piece{white, queen}
		b[4][r] = piece{white, king}
		b[5][r] = piece{white, bishop}
		b[6][r] = piece{white, knight}
		b[7][r] = piece{white, rook}
	}
	homeRow(0, true)
	homeRow(7, false)

	return &b

}

func (b *board) String() string {
	var s string
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			switch b[c][r].piece {
			case none:
				if isWhite(c, r) {
					s += "*"
				} else {
					s += " "
				}
			case pawn:
				s += "P"
			case bishop:
				s += "B"
			case knight:
				s += "N"
			case rook:
				s += "R"
			case queen:
				s += "Q"
			case king:
				s += "K"
			}
		}
		s += "\n"
	}
	return s
}

func isWhite(c, r int) bool {
	return (c+r)%2 == 0
}
