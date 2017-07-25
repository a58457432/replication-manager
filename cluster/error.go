// replication-manager - Replication Manager Monitoring and CLI for MariaDB and MySQL
// Authors: Guillaume Lefranc <guillaume@signal18.io>
//          Stephane Varoqui  <stephane.varoqui@mariadb.com>
// This source code is licensed under the GNU General Public License, version 3.
// Redistribution/Reuse of this code is permitted under the GNU v3 license, as
// an additional term, ALL code must carry the original Author(s) credit in comment form.
// See LICENSE in this directory for the integral text.
package cluster

var clusterError = map[string]string{
	"ERR00004": "Database %s access denied: %s",
	"ERR00005": "Could not get privileges for user %s@%s: %s",
	"ERR00006": "User must have REPLICATION CLIENT privilege",
	"ERR00007": "User must have REPLICATION SLAVE privilege",
	"ERR00008": "User must have SUPER privilege",
	"ERR00009": "User must have RELOAD privilege",
	"ERR00010": "Could not find a slave in topology",
	"ERR00011": "Found multiple masters in topology but not explicitely setup",
	"ERR00012": "Could not find a master in topology",
	"ERR00013": "Binary log disabled on slave: %s",
	"ERR00014": "Could not get binlog dump count on server %s: %s",
	"ERR00015": "Could not get get privileges for user %s on server %s: %s",
	"ERR00016": "Master is unreachable but slaves are replicating",
	"ERR00017": "Unable to fetch MaxScale monitoring information",
	"ERR00018": "Could not connect to MaxScale: %s",
	"ERR00019": "Could not get MaxScale maxadmin server list: %s",
	"ERR00020": "Could not get MaxScale maxinfo server list: %s",
	"ERR00021": "All cluster down in non-interactive mode",
	"ERR00022": "Running in passive mode",
	"ERR00023": "Constraint failed: state %s, conf.Interactive %t cluster.isMaxMasterFailedCountReach %t",
	"ERR00024": "Constraint failed: isExternalOk %t,isActiveArbitration %t,isBeetwenFailoverTimeTooShort %t ,isMaxClusterFailoverCountReach %t, isOneSlaveHeartbeatIncreasing %t, isMaxscaleSupectRunning %t",
	"ERR00025": "Could not get MaxScale maxinfo server list: %s",
	"ERR00026": "First node restarted is a slave, non-interactive mode",
	"ERR00027": "Number of cluster failovers exceeded",
	"ERR00028": "Slave %s can still communicate with the master",
	"ERR00029": "Time between failovers too short",
	"ERR00030": "Maxscale %s can still communicate with the master",
	"ERR00031": "External API can still communicate with the master",
	"ERR00032": "No candidates found in slaves list",
	"ERR00033": "Skip slave in election %s have no master log file, slave might have failed",
	"ERR00034": "Skip slave in election %s repl not electable for switchover",
	"ERR00035": "Skip slave in election %s multi-master and is already the master",
	"ERR00036": "Skip slave in election %s is relay",
	"ERR00037": "Skip slave in election %s in ignore list",
	"ERR00038": "Skip slave in election %s repl not electable for failover",
	"ERR00040": "Skip slave in election %s does not ping or has no binlogs",
	"ERR00041": "Skip slave in election %s has more than %d seconds of replication delay (%d)",
	"ERR00042": "Skip slave in election %s SQL Thread is stopped",
	"ERR00043": "Skip slave in election %s Semisync report unsynced",
	"ERR00044": "Can't connect OpenSVC collector %s",
	"WARN0023": "Failover number of master pings failure has been reached",
	"WARN0045": "Provision task is in queue",
	"WARN0046": "Provision task is working",
	"WARN0047": "Entreprise provision of MariaDB Sharding Cluster not yet implemented",
	"WARN0048": "No semisync settings on slave %s",
	"WARN0049": "No binlog format ROW on slave %s and flashback activated",
	"WARN0050": "No Heartbeat <= 1s on slave %s",
	"WARN0051": "No GTID replication on slave %s",
	"WARN0052": "No InnoDB durability on slave %s",
	"WARN0053": "No replication checksum on slave %s",
	"WARN0054": "No log of replication queries in slow query on slave %s",
	"WARN0055": "No replication queries in show processlist on slave %s",
	"WARN0056": "No compression of binlog on slave %s",
	"WARN0057": "No log-slave-updates on slave %s",
	"WARN0058": "No GTID strict mode on slave %s",
	"WARN0059": "No replication crash safe setting on slave %s",
	"WARN0060": "No semisync settings on master %s",
	"WARN0061": "No binlog format ROW on master %s and flashback activated",
	"WARN0062": "No Heartbeat <= 1s on master %s",
	"WARN0064": "No InnoDB durability on master %s",
	"WARN0065": "No replication checksum on master %s",
	"WARN0066": "No log of replication queries in slow query on master %s",
	"WARN0067": "No replication queries in show processlist on master %s",
	"WARN0068": "No compression of binlog on slave %s",
	"WARN0069": "No log-slave-updates on master %s",
	"WARN0070": "No GTID strict mode on master %s",
	"WARN0071": "No replication crash safe setting on master %s",
}
