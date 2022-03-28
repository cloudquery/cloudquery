policy "example-policy" {
  title = "ExamplePolicy"

  configuration {
  }

  policy "example-policy-1" {
    title  = "Example Policy 1"

    check "1.1" {
      title = "Example Policy 1 Check 1"
      query = "SELECT 1;"
    }

  }
}
