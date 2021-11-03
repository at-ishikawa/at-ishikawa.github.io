-- Create an orchestrator user
CREATE DATABASE IF NOT EXISTS orchestrator;
CREATE USER 'orchestrator'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON `orchestrator`.* TO 'orchestrator'@'%';
