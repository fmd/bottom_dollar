class Background {
    private var stats(default, null):Dynamic;
    private var immigration_generation(default, null):Int;
    private var naturalized(default, null):Bool;
    private var religion(default, null):Religion;

    public function new(key) {
        generate_stats(key);
        cement_stats();
    }

    private function generate_stats(key):Void {
        var backgrounds = Loader.list("backgrounds/backgrounds.json");
        for (background in backgrounds) {
            if (key == background.name) {
                stats = background;
            }
        }
    }

    private function cement_stats():Void {

    }
}