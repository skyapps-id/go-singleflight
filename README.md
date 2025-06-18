## Performance Comparison: With vs Without singleflight
| Metric                         | Without singleflight         | With singleflight           | Impact / Notes                      |
| ------------------------------ | ---------------------------- | --------------------------- | ----------------------------------- |
| **Concurrent Requests**        | Tested with high concurrency | Same test setup             | Equal load in both scenarios        |
| **Execution Time**             | Longer                       | Shorter                     | ✅ Reduced redundant processing      |
| **Duplicate Request Handling** | Multiple DB/API hits         | De-duplicated automatically | ✅ Prevents unnecessary workload     |
| **CPU/Memory Load**            | Higher                       | Lower                       | ✅ System resources used more wisely |
| **Average Response Time**      | Higher latency               | Lower latency               | ✅ Faster client experience          |
| **Scalability**                | Lower                        | Higher                      | ✅ More stable under load            |
