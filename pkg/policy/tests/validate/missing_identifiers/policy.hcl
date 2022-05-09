policy "test_policy" {
  identifiers = ["id"]
  check "1" {
    query = "select 1 as other_value, 'test' as cq_reason"
  }
}
