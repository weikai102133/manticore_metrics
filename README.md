# manticore.conf 
=============

*  此文件配置jdbc连接，初始化jdbc需要该文件配置在工程的src/main/manticore.conf下

*  manticoreDriverName   manticore数据连接类型
*  username              数据库用户名                                    
*  password              数据库密码   
*  ip                    IP地址
*  port                  数据库端口号
*  databaseName          数据库名称

# 目前manticore.conf配置如下：

*   manticoreDriverName = mysql
*   username = root
*   password = root
*   ip = 127.0.0.1
*   port = 9306
*   databaseName = Manticore

# manticore-metrics 目前支持监控manticore以下指标
=============


*   Name: "manticore_metric_thread_num"
*   Help: "显示这个线程已经完成的任务总数"

*   Name: "manticore_thread_in_idle"
*   Help: "显示当前线程是否空闲 0为空闲 1为working"

*   Name: "manticore_thread_in_idle_num"
*   Help: "显示空闲线程总数 "

*   Name: "manticore_thread_in_working_num"
*   Help: "显示工作线程总数 "

*   Name: "manticore_thread_proto"
*   Help: "显示连接协议 0（sphinx）,1（mysql）,2（http）,3（ssl）,4（compressed ）,5（replication ）,6（combination ）,7线程空闲或其它协议",

*   Name: "manticore_thread_state"
*   Help: "显示线程状态 0（handshake），1（net_read），2（net_write），3（query），4（net_idle）,5(线程空闲)"

*   Name: "manticore_time_seconds"   
*   Help: "显示此线程当前任务执行时间 单位s"

*   Name: "manticore_index_type"
*   Help: "显示此索引类型 0（disk）, 1（rt）, 2（percolate）,3（template） ,  4（distributed），5（其它）"

*   Name: "manticore_indexed_documents_num"
*   Help: "显示此索引文档的数量"

*   Name: "manticore_indexed_documents_bytes"
*   Help: "显示此索引文档的大小"

*   Name: "manticore_indexed_ram_bytes"
*   Help: "显示常驻ram索引部分的总大小"

*   Name: "manticore_indexed_disk_bytes"
*   Help: "该索引的文件总大小"

*   Name: "manticore_indexed_disk_mapped_bytes"
*   Help: "该索引文件映射的总大小"

*   Name: "manticore_indexed_disk_mapped_cached_bytes"
*   Help: "该索引实际缓存在RAM中的文件映射的总大小"

*   Name: "manticore_disk_mapped_doclists_bytes"
*   Help: "显示文档列表映射部分大小"

*   Name: "manticore_indexed_disk_mapped_cached_doclists_bytes"
*   Help: "显示文档列表的缓存映射部分大小"

*   Name: "manticore_indexed_disk_mapped_hitlists_bytes"
*   Help: "显示命中列表的映射部分大小"

*   Name: "manticore_indexed_disk_mapped_cached_hitlists_bytes"
*   Help: "显示命中列表的缓存映射部分大小"

*   Name: "manticore_indexed_killed_documents_num"
*   Help: "显示该索引删除文档总数"

*   Name: "manticore_indexed_killed_rate"
*   Help: "显示该索引删除文档所占比例   %"

*   Name: "manticore_indexed_ram_chunk_bytes"
*   Help: "real-time or percolate index的RAM chunk大小"

*   Name: "manticore_indexed_ram_chunk_segments_count"
*   Help: "显示实时索引的磁盘块数"

*   Name: "manticore_indexed_disk_chunks_num"
*   Help: "显示实时索引的磁盘块数"

*   Name: "manticore_indexed_mem_limit_bytes"
*   Help: "显示索引的rt_mem_limit的实际值"

*   Name: "manticore_indexed_ram_bytes_retired_bytes"
*   Help: "显示该索引在RAM chunkst中垃圾的大小"

*   Name: "manticore_indexed_query_time_1min_num"
*   Help: "显示该索引最近一分钟内的查询次数统计"

*   Name: "manticore_indexed_query_time_1min_avg"
*   Help: "显示该索引最近一分钟内的查询执行次数均值 单位 : 次数/s"

*   Name: "manticore_indexed_index_found_rows_1min_avg"
*   Help: "显示该索引最近一分钟每次查询行数平均值"


