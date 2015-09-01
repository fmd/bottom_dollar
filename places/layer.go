package places

type Layer [][]Tile

func NewExampleGroundLayer() Layer {
    return Layer{[]Tile{NONE, NONE,  NONE,  NONE,  NONE},
                 []Tile{NONE, FLOOR, FLOOR, FLOOR, NONE},
                 []Tile{NONE, FLOOR, FLOOR, FLOOR, NONE},
                 []Tile{NONE, FLOOR, FLOOR, FLOOR, NONE},
                 []Tile{NONE, NONE,  NONE,  NONE,  NONE}}
}

func NewExampleWallLayer() Layer {
    return Layer{[]Tile{WALL, WALL, WALL, WALL, WALL},
                 []Tile{WALL, NONE, NONE, NONE, WALL},
                 []Tile{WALL, NONE, NONE, NONE, WALL},
                 []Tile{WALL, NONE, NONE, NONE, WALL},
                 []Tile{WALL, WALL, WALL, WALL, WALL}}
}