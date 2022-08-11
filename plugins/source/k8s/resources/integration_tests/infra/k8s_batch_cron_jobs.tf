resource "kubernetes_cron_job" "batch_cron_jobs" {
  metadata {
    name = "cron-job${var.test_prefix}${var.test_suffix}"
  }
  spec {
    concurrency_policy            = "Replace"
    failed_jobs_history_limit     = 5
    schedule                      = "1 0 * * *"
    starting_deadline_seconds     = 10
    successful_jobs_history_limit = 10
    job_template {
      metadata {}
      spec {
        backoff_limit              = 2
        ttl_seconds_after_finished = 10
        template {
          metadata {}
          spec {
            container {
              name              = "hello"
              image_pull_policy = "Never"
              image             = "busybox"
              command           = ["/bin/sh", "-c", "date; echo Hello from the Kubernetes cluster"]
            }
          }
        }
      }
    }
  }
}