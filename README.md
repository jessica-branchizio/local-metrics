# local-metrics
 
This repo implements a prometheus exporter and an osClient using the gopsutil library to scrape system metrics and exposes them at port 9090  

It currently exposes the following metrics:  
- **CPU Load:** 1-minute, 5-minute, and 15-minute load averages, which indicate the average number of processes waiting to run
- **Total Memory:** Total physical memory available on the machine
- **Used Memory:** Amount of memory currently used by all processes
- **Available Memory:** Amount of memory currently available


## Future Metric Additions
**CPU Usage:** Percentage of CPU used by processes  

**Free Memory:** Amount of memory that is not being used  
**Buffer/Cache:** Memory used by the kernel for buffers or caching  

**Disk Utilization:** Percentage of disk space used  
**Disk I/O:** Read and write operations on the disk  
**Disk Throughput:** Speed of data transfer on disk  

**Network Traffic:** Amount of data sent and received over a network interface  
**Packet Statistics:** Number of packets sent and received, including errors and drops  

**Uptime:** How long the system has been running since the last reboot  

**Number of Processes:** Total number of active processes  
**Specific Process Monitoring:** Resources used by specific processes, such as memory and CPU usage  
