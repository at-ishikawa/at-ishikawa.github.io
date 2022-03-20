-- Grant a root to dump data
-- CREATE USER 'root'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%';

-- Create a replication user
CREATE USER 'repl'@'%' IDENTIFIED BY 'password';
GRANT REPLICATION SLAVE ON *.* TO 'repl'@'%';

-- Create an orchestrator user for topology
CREATE USER 'orchestrator'@'%' IDENTIFIED BY 'password';
GRANT SUPER, PROCESS, REPLICATION SLAVE, RELOAD ON *.* TO 'orchestrator'@'%';
GRANT SELECT ON mysql.slave_master_info TO 'orchestrator'@'%';

FLUSH PRIVILEGES;
