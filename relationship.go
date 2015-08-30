package main

type RelationshipType int

const (
    PARENT RelationshipType = iota
    CHILD
    SIGNIFICANT
    SPOUSE
)

type Relationship struct {
    From *Person
    To   *Person
    Type RelationshipType
}