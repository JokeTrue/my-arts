# HW03 Report

---
## In Idle State
Замеряем нагрузку Мастера в состоянии покоя
```
CPU Load: 0.01
Load Average: 0.04, 0.03, 0.00
Disk Usage: 2/25GB (11%)
Memory Usage: 728/981MB (74.21%)
```

## Before Async Replication
Замеряем нагрузку Мастера под нагрузкой до включения асинхронной репликации
```
CPU Load: 43.05
Load Average: 43.05, 11.29, 3.81
Disk Usage: 2/25GB (11%)
Memory Usage: 737/981MB (75.13%)
```

## Asynchronous Replication Setup
Настраиваем на Мастере и на Slave #1 асинхронную репликацию
```
Master
vim> /etc/mysql/mysql.conf.d/mysqld.cnf
- - -
server-id               = 1
log_bin                 = /var/log/mysql/mysql-bin.log
max_binlog_size         = 500M
binlog_do_db            = my_arts
- - -
mysql> create user rpl_user1@194.67.113.141 identified by '******';
mysql> grant replication slave on *.* to rpl_user1@194.67.113.141;


Slave 
/etc/mysql/mysql.conf.d/mysqld.cnf
- - -
server-id               = 2
read_only               = 1
log_bin                 = /var/log/mysql/mysql-bin.log
max_binlog_size         = 500M
sync_binlog             = 1
expire-logs-days        = 7
- - -

mysql> CHANGE MASTER TO MASTER_HOST='176.99.12.151',
-> MASTER_USER='rpl_user1',
-> MASTER_PASSWORD='******',
-> MASTER_LOG_FILE='mysql-bin.000002',
-> MASTER_LOG_POS=456;
```

## After Async Replication
Перенес все запросы чтения на Slave, все запросы на запись оставил на Мастере.  
Замеряем нагрузку Мастера под нагрузкой после включения асинхронной репликации.
Видно, что вся нагрузка на чтение перешла на Slave #1, поэтому показатели ниже соответствуют показателям в состоянии покоя.
```
CPU Load: 0.00
Load Average: 0.00, 0.00, 0.02
Disk Usage: 2/25GB (12%)
Memory Usage: 557/981MB (56.78%)
```

## Setting up Semi Sync Replication
Настраиваю полусинхронную репликацию в соответствии с документацией:
https://dev.mysql.com/doc/refman/8.0/en/replication-semisync-installation.html

Дополнительно настраиваю конфиги Мастера и Слейвов для работы:
* в режиме row-based
* с включенным gtid
* в режиме полусинхронной репликации
```
Master
/etc/mysql/mysql.conf.d/mysqld.cnf
- - -
...
binlog_format                = row
gtid-mode                    = ON
enforce-gtid-consistency     = ON
log-slave-updates            = ON
rpl_semi_sync_master_enabled = 1
- - -

Slave #1, Slave #2
/etc/mysql/mysql.conf.d/mysqld.cnf
- - -
...
binlog_format                = row
gtid-mode                    = ON
enforce-gtid-consistency     = ON
log-slave-updates            = OFF
rpl_semi_sync_slave_enabled  = 1
- - -
```

## Load Test
Нагрузку на запись буду проводить путем создания новых пользователей.  
Смотрим сколько было пользователей до непосредственной нагрузки.
```sql
SELECT COUNT(*) FROM users;
mysql> 1040305
```

Начинаю нагрузку путем добавления в таблицу пользователей пачками по 50 штук
```text
[2021/02/25 18:20:17] Processed 1 batch
...
[2021/02/25 18:20:33] Processed 646 batch
```

Убиваю Мастер примерно во время выполнения пачки #647, значит в базу должно было записаться 646 пачек по 50 пользователей = 36400 
```text
[2021/02/25 18:20:33] kill -9 on 647 batch
```

Промоутим Slave #1 в Мастера
```sql
STOP SLAVE;
STOP SLAVE IO_THREAD;
STOP REPLICA;
RESET MASTER;
```

Переключаем Slave #2 на Slave #1 в роли Мастера
```sql
STOP SLAVE;
STOP SLAVE IO_THREAD;
STOP REPLICA;
CHANGE MASTER TO MASTER_HOST='194.67.113.141';
START REPLICA;
START SLAVE;
```

Проверяем количество пользователей на Slave #2
```sql
SELECT COUNT(*) FROM users;
mysql> 1076705
```

Считаем количество пользователей, которое было потеряно во время тренировочного инцидента
```text
Стало       Добавили     Было
1076705  -  36400      = 1040305 

Потеря данных отсутствует.
```