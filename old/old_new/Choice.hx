class Choice {
    public static function bool(percent):Bool {
        return Random.int(1,100) <= percent;
    }
}