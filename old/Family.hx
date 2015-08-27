class Family {
    public var name(default, null):String;
    public var person(default, null):Person;
    public var background(default, null):Background;

    public var parent_a(default, null):Person;
    public var parent_b(default, null):Person;

    public function new(person):Void {
        this.person = person;
        this.background = person.background;
        this.name = new Name(person.background).last;

        parent_a = new Person(this.background, "male", this);
        parent_b = new Person(this.background, "female", this);
    }
}