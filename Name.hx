class Name {
    public var background(default, null):Background;
    public var gender(default, null):String;

    public var first(default, null):String;
    public var first_short(default, null):String;
    public var middle(default, null):String;
    public var middle_short(default, null):String;
    public var last(default, null):String;

    // Method aliases
    public var nickname(default, null):String;
    public var legal(get, null):String;
    public var goes_by_first(get, null):String;
    public var full_informal(get, null):String;
    public var initials(get, null):String; //First and middle
    public var initials_no_stop(get, null):String;
    public var initial_with_last(get, null):String;
    public var initials_with_last(get, null):String;

    public var last_initials(get, null):String; //First and last
    public var last_initials_no_stop(get, null):String;

    public function new(background, last = "", gender = "female"):Void {
        this.background = background;
        this.gender = gender;
        generate_name(last);
    }

    public function get_goes_by_first():String {
        if (first_short != null) {
            return first_short;
        }

        return first;
    }

    public function get_full_informal():String {
        return goes_by_first + " " + last;
    }

    public function get_legal():String {
        if (middle != null) {
            return first + " " + middle + " " + last;
        } else {
            return first + " " + last;
        }
    }

    //First and middle
    public function get_initials():String {
        if (middle != null) {
            return initial(first) + initial(middle);
        } else {
            return initial(first);
        }
    }

    public function get_initials_no_stop():String {
        if (middle != null) {
            return initial_no_stop(first) + initial_no_stop(middle);
        } else {
            return initial_no_stop(first);
        }
    }

    public function get_initial_with_last():String {
        return initial(first) + " " + last;
    }

    public function get_initials_with_last():String {
        if (middle != null) {
            return initial(first) + " " + initial(middle) + " " + last;
        } else {
            return initial(first) + " " + last;
        }
    }

    //First and last
    public function get_last_initials():String {
        return initial(first) + initial(last);
    }

    public function get_last_initials_no_stop():String {
        return initial_no_stop(first) + initial_no_stop(last);
    }

    /*** Methods ***/

    public function first_middle():String {
        return first_short + "-" + middle_short;
    }

    public function preferred_initials():Dynamic {
        var filename = "names/preferred_initials.json";
        var contents = sys.io.File.getContent(filename);
        return haxe.Json.parse(contents);
    }

    public function preferred_doubles():Dynamic {
        var filename = "names/preferred_double_names.json";
        var contents = sys.io.File.getContent(filename);
        return haxe.Json.parse(contents);
    }

    public function preferred_initials_kind():String {
        if (Lambda.has(preferred_initials().initials, initials_no_stop)) {
            return "middle";
        }

        if (Lambda.has(preferred_initials().initials, last_initials_no_stop)) {
            return "last";
        }

        return "none";
    }

    public function has_preferred_double_name():Bool {
        return Lambda.has(preferred_doubles().names, first_middle());
    }

    public function initial(which):String {
        return which.substr(0,1).toUpperCase() + ".";
    }

    public function initial_no_stop(which):String {
        return which.substr(0,1).toUpperCase();
    }

    public function parse_name_contents(which):Dynamic {
        var choices = "";

        if (which == "last") {
            choices = background.name;
        } else {
            choices = background.name_list;
        }

        var filename = "names/" + choices + "_" + gender + "_" + which + ".json";
        if (which == "last") {
            filename = "names/" + choices + "_" + which + ".json";
        }

        var contents = sys.io.File.getContent(filename);
        return haxe.Json.parse(contents);
    }

    public function random_name(which):Array<String> {
        return Random.fromArray(parse_name_contents(which).names);
    }

    public function generate_name(last = ""):Void {
        if (last.length > 0) {
            this.last = last;
        } else {
            this.last = random_name("last").shift();
        }

        var first_whole:Array<String> = random_name("first");
        first = first_whole.shift();
        while (first == last) {
            first_whole = random_name("first");
            first = first_whole.shift();
        }

        // Generates short version of first name. I.e. James -> Jim, Jimmy, Jay
        if (first_whole.length > 0) {
            first_short = Random.fromArray(first_whole);
        }

        if (background.has_middle_name) {
            var middle_whole:Array<String> = random_name("first");
            middle = middle_whole.shift();
            while (middle == first || middle == last) {
                middle_whole = random_name("first");
                middle = middle_whole.shift();
            }

            // Generates short version of first name. I.e. William -> Billy, Bill, Will
            if (middle_whole.length > 0) {
                middle_short = Random.fromArray(middle_whole);
            }
        }

        nickname = generate_nickname();
    }

    public function generate_nickname():String {

        // In 30% of cases, attempt to use a convenient initial-based nickname, like "JD" or "DC".
        if (Random.int(1,10) <= 3) {
            var preferred = preferred_initials_kind();
            if (preferred == "middle") {
                return initials_no_stop;
            } else if (preferred == "last") {
                return last_initials_no_stop;
            }
        }

        // In 20% of cases, attempt to use a nickname like "Billy-Bob" for "William Robert".
        if (Random.int(1,10) <= 2) {
            if (first_short != null && middle_short != null) {
                if (has_preferred_double_name()) {
                    return first_middle();
                }
            }
        }

        return goes_by_first;
    }
}