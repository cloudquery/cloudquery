policy "test_policy" {
  identifiers = ["id"]
  check "1" {
    query = "select 1 as id"
  }
}
