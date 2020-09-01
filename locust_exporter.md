
# Sample :9646/metrics

``` Bash
# HELP go_gc_duration_seconds A summary of the GC invocation durations.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 1.5468e-05
go_gc_duration_seconds{quantile="0.25"} 5.5945e-05
go_gc_duration_seconds{quantile="0.5"} 8.238e-05
go_gc_duration_seconds{quantile="0.75"} 0.00016262
go_gc_duration_seconds{quantile="1"} 0.000387374
go_gc_duration_seconds_sum 0.062545795
go_gc_duration_seconds_count 626
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 11
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.13.8"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 8.14464e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 2.648051232e+09
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 1.495011e+06
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 9.040828e+06
# HELP go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.
# TYPE go_memstats_gc_cpu_fraction gauge
go_memstats_gc_cpu_fraction 2.3302348602264766e-05
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 2.38592e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 8.14464e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 5.6713216e+07
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 9.314304e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 11608
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 5.5484416e+07
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 6.602752e+07
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 1.58505847702685e+09
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 9.052436e+06
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 13888
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 58888
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 114688
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 9.938608e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 2.214165e+06
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 1.081344e+06
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 1.081344e+06
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 7.3335032e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 14
# HELP locust_errors The current number of errors.
# TYPE locust_errors gauge
locust_errors{error="\"HTTPError('404 Client Error: NOT FOUND for url: http://locust-master:8089/does_not_exist',)\"",method="GET",name="/does_not_exist"} 20099
# HELP locust_requests_avg_content_length 
# TYPE locust_requests_avg_content_length gauge
locust_requests_avg_content_length{method="",name="Aggregated"} 5887.493411956757
locust_requests_avg_content_length{method="GET",name="/"} 15045.000150157666
locust_requests_avg_content_length{method="GET",name="/does_not_exist"} 232
locust_requests_avg_content_length{method="GET",name="/stats/requests"} 2415.2776747682287
# HELP locust_requests_avg_response_time 
# TYPE locust_requests_avg_response_time gauge
locust_requests_avg_response_time{method="",name="Aggregated"} 8.325235602154045
locust_requests_avg_response_time{method="GET",name="/"} 8.417879713507075
locust_requests_avg_response_time{method="GET",name="/does_not_exist"} 8.222418855366953
locust_requests_avg_response_time{method="GET",name="/stats/requests"} 8.336038763915587
# HELP locust_requests_current_fail_per_sec 
# TYPE locust_requests_current_fail_per_sec gauge
locust_requests_current_fail_per_sec{method="",name="Aggregated"} 1.8
locust_requests_current_fail_per_sec{method="GET",name="/"} 0
locust_requests_current_fail_per_sec{method="GET",name="/does_not_exist"} 1.8
locust_requests_current_fail_per_sec{method="GET",name="/stats/requests"} 0
# HELP locust_requests_current_response_time_percentile_50 
# TYPE locust_requests_current_response_time_percentile_50 gauge
locust_requests_current_response_time_percentile_50 9
# HELP locust_requests_current_response_time_percentile_95 
# TYPE locust_requests_current_response_time_percentile_95 gauge
locust_requests_current_response_time_percentile_95 11
# HELP locust_requests_current_rps 
# TYPE locust_requests_current_rps gauge
locust_requests_current_rps{method="",name="Aggregated"} 3.3
locust_requests_current_rps{method="GET",name="/"} 0.8
locust_requests_current_rps{method="GET",name="/does_not_exist"} 1.8
locust_requests_current_rps{method="GET",name="/stats/requests"} 0.7
# HELP locust_requests_fail_ratio 
# TYPE locust_requests_fail_ratio gauge
locust_requests_fail_ratio 0.3347991937767561
# HELP locust_requests_max_response_time 
# TYPE locust_requests_max_response_time gauge
locust_requests_max_response_time{method="",name="Aggregated"} 50
locust_requests_max_response_time{method="GET",name="/"} 24
locust_requests_max_response_time{method="GET",name="/does_not_exist"} 25
locust_requests_max_response_time{method="GET",name="/stats/requests"} 50
# HELP locust_requests_median_response_time 
# TYPE locust_requests_median_response_time gauge
locust_requests_median_response_time{method="",name="Aggregated"} 9
locust_requests_median_response_time{method="GET",name="/"} 9
locust_requests_median_response_time{method="GET",name="/does_not_exist"} 9
locust_requests_median_response_time{method="GET",name="/stats/requests"} 9
# HELP locust_requests_min_response_time 
# TYPE locust_requests_min_response_time gauge
locust_requests_min_response_time{method="",name="Aggregated"} 2
locust_requests_min_response_time{method="GET",name="/"} 2
locust_requests_min_response_time{method="GET",name="/does_not_exist"} 3
locust_requests_min_response_time{method="GET",name="/stats/requests"} 3
# HELP locust_requests_num_failures 
# TYPE locust_requests_num_failures gauge
locust_requests_num_failures{method="",name="Aggregated"} 20099
locust_requests_num_failures{method="GET",name="/"} 0
locust_requests_num_failures{method="GET",name="/does_not_exist"} 20099
locust_requests_num_failures{method="GET",name="/stats/requests"} 0
# HELP locust_requests_num_requests 
# TYPE locust_requests_num_requests gauge
locust_requests_num_requests{method="",name="Aggregated"} 60033
locust_requests_num_requests{method="GET",name="/"} 19979
locust_requests_num_requests{method="GET",name="/does_not_exist"} 20099
locust_requests_num_requests{method="GET",name="/stats/requests"} 19955
# HELP locust_running The current state of the execution (0 = STOPPED 1 = HATCHING 2 = RUNNING,).
# TYPE locust_running gauge
locust_running 2
# HELP locust_worker_detail The current status of a worker with user count
# TYPE locust_worker_detail gauge
locust_worker_detail{id="84d64fa91fe7_f501b54dce7f4bb6a80db4e7b677ce51",state="running"} 10
# HELP locust_workers_count The current number of workers.
# TYPE locust_workers_count gauge
locust_workers_count 1
# HELP locust_workers_hatching_count The current number of hatching workers.
# TYPE locust_workers_hatching_count gauge
locust_workers_hatching_count 0
# HELP locust_workers_missing_count The current number of missing workers.
# TYPE locust_workers_missing_count gauge
locust_workers_missing_count 0
# HELP locust_workers_running_count The current number of running workers.
# TYPE locust_workers_running_count gauge
locust_workers_running_count 1
# HELP locust_up The current health status of the server (1 = UP, 0 = DOWN).
# TYPE locust_up gauge
locust_up 1
# HELP locust_users The current number of users.
# TYPE locust_users gauge
locust_users 10
# HELP locustexporter_build_info A metric with a constant '1' value labeled by version, revision, branch, and goversion from which locustexporter was built.
# TYPE locustexporter_build_info gauge
locustexporter_build_info{branch="",goversion="go1.13.8",revision="",version=""} 1
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 86.35
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 1.048576e+06
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 10
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 2.9392896e+07
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.58503749614e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 1.16871168e+08
# HELP process_virtual_memory_max_bytes Maximum amount of virtual memory available in bytes.
# TYPE process_virtual_memory_max_bytes gauge
process_virtual_memory_max_bytes -1
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 6771
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```
