class Religion {
    public var name(default, null):String;
    public var plural(default, null):String;
    public var religious_percent(default, null):Int;

    public function all_religions():Array<Dynamic> {
        var filename = "backgrounds/religions.json";
        var contents = sys.io.File.getContent(filename);
        var results:Array<Dynamic> = haxe.Json.parse(contents).religions;
        return results;
    }

    public function new(key):Void {
        for (religion in all_religions()) {
            if (key == religion.name) {
                name = religion.name;
                plural = religion.plural;
                religious_percent = religion.religious_percent;
                return;
            }
        }
    }
}