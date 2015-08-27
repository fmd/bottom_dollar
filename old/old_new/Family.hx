class Family {
    private var parent_a(default, null):Person;
    private var parent_b(default, null):Person;

    public function new(person_a, person_b) {
        parent_a = person_a;
        parent_b = person_b;
    }
}