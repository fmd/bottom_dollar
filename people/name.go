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
    n.FirstSet = firstName(genderDir, b)

    return n
}

func lastName(genderDir string, b *Background) NameSet {
    return NameSetFromList(NameList("last", genderDir, b.Name))
}

func middleName(genderDir string, b *Background) NameSet {
    if !b.HasMiddleName {
        return NameSet{""}
    }
    return NameSetFromList(NameList("middle", genderDir, b.Name))
}

func firstName(genderDir string, b *Background) NameSet {
    return NameSetFromList(NameList("first", genderDir, b.Name))
}

func (n *Name) First() string {
    return n.FirstSet[0]
}

func (n *Name) Middle() string {
    return n.MiddleSet[0]
}

func (n *Name) Last() string {
    return n.LastSet[0]
}

func (n *Name) Full() string {
    a := n.First() + " "
    if len(n.Middle()) > 0 {
        a = a + n.Middle() + " "
    }
    return a + n.Last()
}