policy "test_policy" {
  identifiers = ["id"]
  check "1" {
    query = "select 1 as id, 'test' as cq_reason"
  }
  policy "child_policy" {
    check "1" {
      query = "select 1 as id"
      reason = "test"
    }
  }
}
