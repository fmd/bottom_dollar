class Person {
    public var name(default, null):Name;
    public var background(default, null):Background;
    public var age(default, null):Int;
    public var gender(default, null):String;
    public var family(default, null):Family;
    public var dob(default, null):DateOfBirth;

    public function new(background, gender, family = null):Void {
        this.background = background;
        this.gender = gender;
        trace(this.background);

        if (family == null) {
            this.family = new Family(this);
        } else {
            this.family = family;
        }

        name = new Name(this.family.background, this.family.name, gender);
    }
}