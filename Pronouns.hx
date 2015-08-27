class Pronouns {
    public var gender(default, null):String;

    public function new(gender) {
        this.gender = gender;
    }

    public function female():Bool {
        if (gender == "female") {
            return true;
        }

        return false;
    }

    // Notice with the female pronoun does not distinguish here like the male.
    // -- her/her vs his/him.
    // She had a dog. It was her dog. It belonged to her.
    // He had a dog. It was his dog. It belonged to him.
    //
    // Extra Note - You also don't say 'It was hers cat.' here.
    // She had a cat. It was hers.
    // He had a cat. It was his.

    public function herself():String {
        if (female()) {
            return "herself";
        }
        return "himself";
    }

    public function she():String {
        if (female()) {
            return "she";
        }
        return "he";
    }

    public function her_ob():String {
        if (female()) {
            return "her";
        }
        return "his";
    }

    public function her_sub():String {
        if (female()) {
            return "her";
        }
        return "him";
    }

    public function hers():String {
        if (female()) {
            return "hers";
        }
        return "his";
    }
}