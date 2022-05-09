policy "test_policy" {
  identifiers = ["id"]
  check "1" {
    query = "select 1 as id, 'test' as cq_reason"
  }
  policy "child_policy" {
    identifiers = ["id2"]
    check "1" {
      query = "select 1 as id2, 'test' as cq_reason"
    }
  }
}
