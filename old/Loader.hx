class Loader {
    public static function list(path):Array<Dynamic> {
    var contents = sys.io.File.getContent(path);
    var results:Array<Dynamic> = haxe.Json.parse(contents);
    return results;
}