class Main {
    static public function main():Void {
        var p = new Person(new Background("african"), "female");
        var n = new Pronouns(p.gender);

        trace(n.her_ob() + " friends call " + n.her_sub() + " " + p.name.nickname);
        trace(n.she() + " introduces " + n.herself() + " as " + p.name.full_informal);
        trace(n.she() + " signs " + n.her_ob() + " name as " + p.name.initials_with_last);
        trace(n.her_ob() + " full legal name is " + p.name.legal);
        trace(n.her_ob() + " parents' names are " + p.family.parent_a.name.first + " and " + p.family.parent_b.name.first + ".");
    }
}