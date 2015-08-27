class Background {
    public var name(default, null):String;
    public var is_naturalized(default, null):Bool;
    public var is_religious(default, null):Bool;

    public var religion(default, null):Religion;
    public var name_list(default, null):String;
    public var has_middle_name(default, null):Bool;

    public function all_backgrounds():Array<Dynamic> {
        var filename = "backgrounds/backgrounds.json";
        var contents = sys.io.File.getContent(filename);
        var results:Array<Dynamic> = haxe.Json.parse(contents).backgrounds;
        return results;
    }

    public function new(key):Void {
        for (background in all_backgrounds()) {
            if (key == background.name) {
                name = background.name;

                if (Random.int(1,100) <= background.naturalization_percent) {
                    is_naturalized = true;
                }

                religion = new Religion(Random.fromArray(background.religions));
                is_religious = false;
                if (Random.int(1,100) <= religion.religious_percent) {
                    is_religious = true;
                }
                name_list = Random.fromArray(all_name_choices());
                return;
            }
        }
    }

    public function all_name_choices():Array<String> {
        var choices = new Array<String>();
        if (!is_naturalized) {
            choices.push(name);
        } else {
            choices.push("european");
        }

        if (is_religious) {
            choices.push(religion.name);
        }

        return choices;
    }
}