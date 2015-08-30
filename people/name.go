package people

type Name struct {
	FirstSet  NameSet
    MiddleSet NameSet
	LastSet   NameSet
}

func NameFromBackgroundAndGender(g Gender, b *Background) *Name {
    n := &Name{}
    genderDir := g.NameDirectory()

    n.LastSet = lastName(genderDir, b)
    n.MiddleSet = middleName(genderDir, b)
    n.first = firstName(genderDir, b)

    return n
}

func lastName(genderDir string, b *Background) NameSet {
    return NameSetFromList(NameList, "last", genderDir, b.Name)
}

func middleName(genderDir string, b *Background) NameSet {
    if !b.HasMiddleName {
        return NameSet{""}
    }
    return NameSetFromList(NameList, "middle", genderDir, b.Name)
}

func firstName(genderDir string, b *Background) NameSet {
    return NameSetFromList(NameList, "first", genderDir, b.Name)
}

func First() {
    return FirstSet[0]
}

func Middle() {
    return MiddleSet[0]
}

func Last() {
    return LastSet[0]
}