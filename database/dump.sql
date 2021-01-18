PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE `Program_list` (
	`ID`	TEXT NOT NULL UNIQUE,
	`ProgramID`	TEXT NOT NULL,
	PRIMARY KEY(`ID`),
	FOREIGN KEY(`ProgramID`) REFERENCES `Program`(`ID`)
);
CREATE TABLE `Perfomance` (
	`ID`	TEXT,
	`CPU`	TEXT,
	`RAM`	TEXT,
	`HDD`	TEXT,
	`Cluster`	INTEGER,
	PRIMARY KEY(`ID`)
);
CREATE TABLE `Fault_tolerance` (
	`ID`	INTEGER NOT NULL,
	`Commissioning_date`	TEXT,
	`Backup`	INTEGER DEFAULT 0,
	`Number_of_error`	INTEGER DEFAULT 0,
	`Cluster`	INTEGER,
	PRIMARY KEY(`ID`)
);
CREATE TABLE `Program` (
	`ID`	TEXT NOT NULL UNIQUE,
	`Manufacturer`	TEXT,
	`Name`	TEXT,
	`Version`	TEXT,
	`Install_on`	TEXT,
	PRIMARY KEY(`ID`)
);
CREATE TABLE `RAM` (
	`ID`	TEXT NOT NULL,
	`Manufacturer`	TEXT,
	`Size`	INTEGER,
	`Frequency`	INTEGER,
	`Serial_number`	TEXT,
	PRIMARY KEY(`ID`)
);
CREATE TABLE `HDD` (
	`ID`	TEXT NOT NULL,
	`Model`	TEXT,
	`Size`	INTEGER,
	`Type`	INTEGER DEFAULT 1,
	PRIMARY KEY(`ID`)
);
CREATE TABLE `Matherboard` (
	`ID`	TEXT NOT NULL,
	`Name`	TEXT,
	`Model`	TEXT,
	`Product`	TEXT,
	PRIMARY KEY(`ID`)
);
CREATE TABLE `CPU` (
	`ID`	TEXT NOT NULL,
	`Model`	TEXT,
	`Manufacturer`	TEXT,
	`Frequency`	INTEGER,
	`Number_cores`	INTEGER,
	`Number_threads`	INTEGER,
	PRIMARY KEY(`ID`)
);
CREATE TABLE `CPU_list` (
	`ID`	TEXT NOT NULL,
	`CPUID`	TEXT,
	FOREIGN KEY(`CPUID`) REFERENCES `CPU`(`ID`),
	PRIMARY KEY(`ID`)
);
CREATE TABLE `HDD_list` (
	`ID`	TEXT NOT NULL,
	`HDDID`	TEXT,
	PRIMARY KEY(`ID`),
	FOREIGN KEY(`HDDID`) REFERENCES `HDD`(`ID`)
);
CREATE TABLE `RAM_list` (
	`ID`	TEXT NOT NULL,
	`RAMID`	TEXT,
	FOREIGN KEY(`RAMID`) REFERENCES `RAM`(`ID`),
	PRIMARY KEY(`ID`)
);
CREATE TABLE `Hardware` (
	`ID`	TEXT NOT NULL,
	`CPU_listID`	TEXT,
	`MatherboardID`	TEXT,
	`RAM_listID`	TEXT,
	`HDD_listID`	TEXT,
	FOREIGN KEY(`HDD_listID`) REFERENCES `HDD_list`(`ID`),
	PRIMARY KEY(`ID`),
	FOREIGN KEY(`MatherboardID`) REFERENCES `Matherboard`(`ID`),
	FOREIGN KEY(`CPU_listID`) REFERENCES `CPU_list`(`ID`),
	FOREIGN KEY(`RAM_listID`) REFERENCES `RAM_list`(`ID`)
);
CREATE TABLE `Workstation` (
	`ID`	TEXT NOT NULL UNIQUE,
	`Name`	TEXT NOT NULL UNIQUE,
	`Comment`	TEXT,
	`Allow_analysis`	INTEGER NOT NULL DEFAULT 0,
	`HardwareID`	TEXT NOT NULL,
	`Program_listID`	TEXT NOT NULL,
	`PerfomanceID`	TEXT NOT NULL,
	`Fault_toleranceID`	TEXT NOT NULL,
	FOREIGN KEY(`Fault_toleranceID`) REFERENCES `Fault_tolerance`(`ID`),
	PRIMARY KEY(`ID`),
	FOREIGN KEY(`Program_listID`) REFERENCES `Program_list`(`ID`),
	FOREIGN KEY(`HardwareID`) REFERENCES `Hardware`(`ID`),
	FOREIGN KEY(`PerfomanceID`) REFERENCES `Perfomance`(`ID`)
);
CREATE TABLE `Upgrade_list_item` (
	`ID`	TEXT NOT NULL,
	`WorkstationID`	TEXT,
	`Description`	TEXT,
	FOREIGN KEY(`WorkstationID`) REFERENCES `Workstation`(`ID`),
	PRIMARY KEY(`ID`)
);
CREATE TABLE `Upgrade_workstation_list` (
	`ID`	TEXT NOT NULL,
	`Date`	TEXT,
	`Upgrade_list_itemID`	TEXT,
	FOREIGN KEY(`Upgrade_list_itemID`) REFERENCES `Upgrade_list_item`(`ID`),
	PRIMARY KEY(`ID`)
);
CREATE TABLE `Monitoring_result` (
	`ID`	TEXT NOT NULL,
	`Upgrade_workstation_listID`	TEXT,
	PRIMARY KEY(`ID`),
	FOREIGN KEY(`Upgrade_workstation_listID`) REFERENCES `Upgrade_workstation_list`(`ID`)
);
CREATE TABLE `Monitoring` (
	`ID`	TEXT NOT NULL,
	`Date`	TEXT,
	`ResultID`	TEXT,
	PRIMARY KEY(`ID`)
);
COMMIT;
