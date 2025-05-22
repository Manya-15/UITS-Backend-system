-- MySQL dump 10.13  Distrib 8.0.34, for Win64 (x86_64)
--
-- Host: localhost    Database: uits_final
-- ------------------------------------------------------
-- Server version	8.0.34

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `device`
--

DROP TABLE IF EXISTS `device`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `device` (
  `device_id` int NOT NULL AUTO_INCREMENT,
  `device_name` varchar(255) NOT NULL,
  `type_id` int DEFAULT NULL,
  `serial_no` varchar(255) NOT NULL,
  `model_no` varchar(255) NOT NULL,
  `purchase_date` date NOT NULL,
  `warranty_expiry` date DEFAULT NULL,
  `status` varchar(50) DEFAULT 'Active',
  `added_by` int DEFAULT NULL,
  `location_id` int DEFAULT '1',
  `added_on` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `replaced_device_id` int DEFAULT NULL,
  `status_flag` int DEFAULT '1',
  PRIMARY KEY (`device_id`),
  UNIQUE KEY `serial_no` (`serial_no`),
  KEY `type_id` (`type_id`),
  KEY `added_by` (`added_by`),
  KEY `location_id` (`location_id`),
  KEY `replaced_device_id` (`replaced_device_id`),
  KEY `fk_device_status` (`status_flag`),
  CONSTRAINT `device_ibfk_1` FOREIGN KEY (`type_id`) REFERENCES `device_type` (`type_id`),
  CONSTRAINT `device_ibfk_2` FOREIGN KEY (`added_by`) REFERENCES `user` (`user_id`),
  CONSTRAINT `device_ibfk_3` FOREIGN KEY (`location_id`) REFERENCES `location` (`location_id`),
  CONSTRAINT `device_ibfk_4` FOREIGN KEY (`replaced_device_id`) REFERENCES `device` (`device_id`),
  CONSTRAINT `fk_device_status` FOREIGN KEY (`status_flag`) REFERENCES `status_master` (`status_id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `device`
--

LOCK TABLES `device` WRITE;
/*!40000 ALTER TABLE `device` DISABLE KEYS */;
INSERT INTO `device` VALUES (1,'Dell Latitude 7420',1,'SRN001DL','Lat7420','2023-01-15','2026-01-14','Active',1,6,'2023-01-10 10:00:00',NULL,1),(2,'HP EliteDesk 800 G6',2,'SRN002HP','ED800G6','2022-11-20','2025-11-19','Active',1,7,'2022-11-15 11:00:00',NULL,1),(3,'Lenovo ThinkPad X1 Carbon (Old)',1,'SRN003LVO','TPX1C-G7','2020-05-10','2023-05-09','Replaced',2,9,'2020-05-01 12:00:00',NULL,3),(4,'HP LaserJet Pro M404dn',8,'SRN004HP','M404dn','2023-03-01','2025-02-28','Active',1,4,'2023-02-25 13:00:00',NULL,1),(5,'Cisco Catalyst 2960X',12,'SRN005CS','WS-C2960X-24PS-L','2022-08-01','2027-07-31','Active',4,6,'2022-07-28 14:00:00',NULL,1),(6,'Samsung 27\" Monitor',5,'SRN006SM','LS27A800','2023-02-10','2026-02-09','Active',2,7,'2023-02-05 15:00:00',NULL,1),(7,'Logitech K120 Keyboard',6,'SRN007LG','K120','2023-01-05','2024-01-04','Active',2,7,'2023-01-01 16:00:00',NULL,1),(8,'Apple iPhone 13',19,'SRN008APL','A2482','2023-09-01','2024-08-31','Maintenance',4,9,'2023-08-25 17:00:00',NULL,4),(9,'Synology DS920+ NAS',16,'SRN009SYN','DS920+','2022-06-15','2025-06-14','Active',1,6,'2022-06-10 10:30:00',NULL,1),(10,'Lenovo ThinkPad X1 Carbon (New)',1,'SRN010LNV','TPX1C-G10','2023-05-15','2026-05-14','Active',2,9,'2023-05-10 12:30:00',3,1),(11,'Epson EB-L510U Projector',21,'SRN011EPS','EB-L510U','2023-07-20','2026-07-19','Inactive',1,12,'2023-07-15 14:30:00',NULL,2),(12,'APC Smart-UPS 1500VA',18,'SRN012APC','SMT1500RMI2U','2021-12-01','2024-11-30','Decommissioned',4,6,'2021-11-25 09:00:00',NULL,5),(13,'HP Laptop',1,'HKRHURY234','KHFIRYI2345','2025-05-09','2025-05-14','Active',6,2,'2025-05-22 19:39:09',NULL,1);
/*!40000 ALTER TABLE `device` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `device_category`
--

DROP TABLE IF EXISTS `device_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `device_category` (
  `category_id` int NOT NULL AUTO_INCREMENT,
  `category_name` varchar(255) NOT NULL,
  `status` tinyint DEFAULT '1',
  PRIMARY KEY (`category_id`),
  UNIQUE KEY `category_name` (`category_name`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `device_category`
--

LOCK TABLES `device_category` WRITE;
/*!40000 ALTER TABLE `device_category` DISABLE KEYS */;
INSERT INTO `device_category` VALUES (1,'Computing Devices',1),(2,'Peripherals',1),(3,'Networking Equipment',1),(4,'Storage Solutions',1),(5,'Office Automation',1),(6,'Mobile Devices',1),(7,'AV Equipment',1);
/*!40000 ALTER TABLE `device_category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `device_specification`
--

DROP TABLE IF EXISTS `device_specification`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `device_specification` (
  `spec_id` int NOT NULL AUTO_INCREMENT,
  `device_id` int DEFAULT NULL,
  `spec_template_id` int DEFAULT NULL,
  `spec_master_id` int DEFAULT NULL,
  `status` tinyint DEFAULT '1',
  PRIMARY KEY (`spec_id`),
  UNIQUE KEY `idx_device_template_master` (`device_id`,`spec_template_id`,`spec_master_id`),
  KEY `spec_template_id` (`spec_template_id`),
  KEY `spec_master_id` (`spec_master_id`),
  CONSTRAINT `device_specification_ibfk_1` FOREIGN KEY (`device_id`) REFERENCES `device` (`device_id`),
  CONSTRAINT `device_specification_ibfk_2` FOREIGN KEY (`spec_template_id`) REFERENCES `specification_template` (`spec_template_id`),
  CONSTRAINT `device_specification_ibfk_3` FOREIGN KEY (`spec_master_id`) REFERENCES `specification_master` (`spec_master_id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `device_specification`
--

LOCK TABLES `device_specification` WRITE;
/*!40000 ALTER TABLE `device_specification` DISABLE KEYS */;
INSERT INTO `device_specification` VALUES (1,1,1,1,1),(2,1,2,5,1),(3,1,3,8,1),(4,1,4,10,1),(5,1,5,13,1),(6,2,6,2,1),(7,2,7,5,1),(8,2,8,8,1),(9,2,9,10,1),(10,4,17,15,1),(11,4,18,16,1),(12,4,19,19,1),(13,4,20,20,1),(14,5,21,22,1),(15,5,22,24,1),(16,5,23,26,1),(17,5,24,28,1),(18,10,1,3,1),(19,10,2,5,1),(20,10,3,8,1),(21,10,4,11,1),(22,10,5,13,1),(23,13,1,3,1),(24,13,4,11,1),(25,13,2,6,1),(26,13,5,13,1),(27,13,3,8,1);
/*!40000 ALTER TABLE `device_specification` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `device_type`
--

DROP TABLE IF EXISTS `device_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `device_type` (
  `type_id` int NOT NULL AUTO_INCREMENT,
  `category_id` int DEFAULT NULL,
  `type_name` varchar(255) NOT NULL,
  `status` tinyint DEFAULT '1',
  PRIMARY KEY (`type_id`),
  UNIQUE KEY `idx_category_type_name` (`category_id`,`type_name`),
  CONSTRAINT `device_type_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `device_category` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `device_type`
--

LOCK TABLES `device_type` WRITE;
/*!40000 ALTER TABLE `device_type` DISABLE KEYS */;
INSERT INTO `device_type` VALUES (1,1,'Laptop',1),(2,1,'Desktop PC',1),(3,1,'Server',1),(4,1,'All-in-One PC',1),(5,2,'Monitor',1),(6,2,'Keyboard',1),(7,2,'Mouse',1),(8,2,'Printer - Laser',1),(9,2,'Scanner',1),(10,2,'Webcam',1),(11,3,'Router',1),(12,3,'Switch',1),(13,3,'Access Point',1),(14,3,'Firewall',1),(15,4,'External HDD',1),(16,4,'NAS Unit',1),(17,4,'USB Flash Drive',1),(18,5,'UPS',1),(19,6,'Smartphone',1),(20,6,'Tablet',1),(21,7,'Projector',1),(22,7,'Smart Board',1);
/*!40000 ALTER TABLE `device_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `location`
--

DROP TABLE IF EXISTS `location`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `location` (
  `location_id` int NOT NULL AUTO_INCREMENT,
  `location_name` varchar(255) NOT NULL,
  `level_id` int DEFAULT NULL,
  `parent_location_id` int DEFAULT NULL,
  `status` tinyint DEFAULT '1',
  PRIMARY KEY (`location_id`),
  KEY `level_id` (`level_id`),
  KEY `parent_location_id` (`parent_location_id`),
  CONSTRAINT `location_ibfk_1` FOREIGN KEY (`level_id`) REFERENCES `location_level` (`level_id`),
  CONSTRAINT `location_ibfk_2` FOREIGN KEY (`parent_location_id`) REFERENCES `location` (`location_id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `location`
--

LOCK TABLES `location` WRITE;
/*!40000 ALTER TABLE `location` DISABLE KEYS */;
INSERT INTO `location` VALUES (1,'Tech University',1,NULL,1),(2,'Main Campus',2,1,1),(3,'Engineering Building',3,2,1),(4,'IT Department',6,3,1),(5,'First Floor - Engg',4,3,1),(6,'Room 101 - Server Room',5,5,1),(7,'Room 102 - Lab A',5,5,1),(8,'Second Floor - Engg',4,3,1),(9,'Room 201 - Faculty Office',5,8,1),(10,'Admin Building',3,2,1),(11,'Ground Floor - Admin',4,10,1),(12,'Room G05 - Reception',5,11,1),(13,'Library Building',3,2,1);
/*!40000 ALTER TABLE `location` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `location_abbreviation`
--

DROP TABLE IF EXISTS `location_abbreviation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `location_abbreviation` (
  `abbreviation_id` int NOT NULL AUTO_INCREMENT,
  `location_id` int DEFAULT NULL,
  `abbreviation` varchar(50) NOT NULL,
  `status` tinyint DEFAULT '1',
  PRIMARY KEY (`abbreviation_id`),
  UNIQUE KEY `abbreviation` (`abbreviation`),
  KEY `location_id` (`location_id`),
  CONSTRAINT `location_abbreviation_ibfk_1` FOREIGN KEY (`location_id`) REFERENCES `location` (`location_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `location_abbreviation`
--

LOCK TABLES `location_abbreviation` WRITE;
/*!40000 ALTER TABLE `location_abbreviation` DISABLE KEYS */;
INSERT INTO `location_abbreviation` VALUES (1,1,'TechU',1),(2,2,'MC',1),(3,3,'ENG-BLD',1),(4,4,'IT-DEPT',1),(5,6,'SRVR-101',1),(6,7,'LAB-A-102',1),(7,9,'FO-201',1),(8,10,'ADM-BLD',1);
/*!40000 ALTER TABLE `location_abbreviation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `location_level`
--

DROP TABLE IF EXISTS `location_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `location_level` (
  `level_id` int NOT NULL AUTO_INCREMENT,
  `level_name` varchar(50) NOT NULL,
  `status` tinyint DEFAULT '1',
  PRIMARY KEY (`level_id`),
  UNIQUE KEY `level_name` (`level_name`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `location_level`
--

LOCK TABLES `location_level` WRITE;
/*!40000 ALTER TABLE `location_level` DISABLE KEYS */;
INSERT INTO `location_level` VALUES (1,'University',1),(2,'Campus',1),(3,'Building',1),(4,'Floor',1),(5,'Room',1),(6,'Department',1);
/*!40000 ALTER TABLE `location_level` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ownership`
--

DROP TABLE IF EXISTS `ownership`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ownership` (
  `ownership_id` int NOT NULL AUTO_INCREMENT,
  `device_id` int DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  `start_datetime` datetime NOT NULL,
  `end_datetime` datetime DEFAULT NULL,
  `status` tinyint DEFAULT '1',
  PRIMARY KEY (`ownership_id`),
  KEY `device_id` (`device_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `ownership_ibfk_1` FOREIGN KEY (`device_id`) REFERENCES `device` (`device_id`),
  CONSTRAINT `ownership_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ownership`
--

LOCK TABLES `ownership` WRITE;
/*!40000 ALTER TABLE `ownership` DISABLE KEYS */;
INSERT INTO `ownership` VALUES (1,1,2,'2023-01-15 10:00:00',NULL,0),(2,2,3,'2022-11-20 11:00:00',NULL,1),(3,3,5,'2020-05-10 12:00:00','2023-05-14 23:59:59',0),(4,4,4,'2023-03-01 13:00:00',NULL,1),(5,5,1,'2022-08-01 14:00:00',NULL,1),(6,6,3,'2023-02-10 15:00:00','2023-08-10 09:00:00',0),(7,6,2,'2023-08-10 09:00:01',NULL,1),(8,8,5,'2023-09-01 17:00:00',NULL,1),(9,10,5,'2023-05-15 09:00:00',NULL,0),(10,1,6,'2025-05-24 01:10:00','2025-05-29 01:10:00',1),(11,10,6,'2025-05-24 01:10:00','2025-05-29 01:10:00',1),(12,13,6,'2025-05-24 01:10:00','2025-05-29 01:10:00',1);
/*!40000 ALTER TABLE `ownership` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `specification_master`
--

DROP TABLE IF EXISTS `specification_master`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `specification_master` (
  `spec_master_id` int NOT NULL AUTO_INCREMENT,
  `spec_value` varchar(255) NOT NULL,
  `status` tinyint DEFAULT '1',
  `spec_template_id` int NOT NULL,
  PRIMARY KEY (`spec_master_id`),
  UNIQUE KEY `idx_template_value` (`spec_template_id`,`spec_value`),
  CONSTRAINT `fk_spec_temp_id` FOREIGN KEY (`spec_template_id`) REFERENCES `specification_template` (`spec_template_id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `specification_master`
--

LOCK TABLES `specification_master` WRITE;
/*!40000 ALTER TABLE `specification_master` DISABLE KEYS */;
INSERT INTO `specification_master` VALUES (1,'Intel Core i5-1135G7',1,1),(2,'Intel Core i7-1260P',1,1),(3,'AMD Ryzen 7 5800U',1,1),(4,'8',1,2),(5,'16',1,2),(6,'32',1,2),(7,'256',1,3),(8,'512',1,3),(9,'1000',1,3),(10,'Windows 11 Pro',1,4),(11,'Ubuntu 22.04',1,4),(12,'macOS Monterey',1,4),(13,'14',1,5),(14,'15.6',1,5),(15,'Laser',1,17),(16,'Monochrome',1,18),(17,'Color',1,18),(18,'25',1,19),(19,'40',1,19),(20,'USB, Ethernet, WiFi',1,20),(21,'USB, WiFi',1,20),(22,'24',1,21),(23,'48',1,21),(24,'1',1,22),(25,'10',1,22),(26,'Yes',1,23),(27,'No',1,23),(28,'Yes',1,24),(29,'No',1,24),(30,'2',1,25),(31,'4',1,25),(32,'USB 3.0',1,26),(33,'USB-C',1,26),(34,'Apple',1,27),(35,'Samsung',1,27),(36,'iPhone 14',1,28),(37,'Galaxy S23',1,28),(38,'UniqueIMEI12345',1,29),(39,'UniqueIMEI67890',1,29),(40,'3500',1,30),(41,'5000',1,30),(42,'1920x1080',1,31),(43,'1280x720',1,31);
/*!40000 ALTER TABLE `specification_master` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `specification_template`
--

DROP TABLE IF EXISTS `specification_template`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `specification_template` (
  `spec_template_id` int NOT NULL AUTO_INCREMENT,
  `type_id` int DEFAULT NULL,
  `spec_name` varchar(255) NOT NULL,
  `status` tinyint DEFAULT '1',
  PRIMARY KEY (`spec_template_id`),
  UNIQUE KEY `idx_type_spec_name` (`type_id`,`spec_name`),
  CONSTRAINT `specification_template_ibfk_1` FOREIGN KEY (`type_id`) REFERENCES `device_type` (`type_id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `specification_template`
--

LOCK TABLES `specification_template` WRITE;
/*!40000 ALTER TABLE `specification_template` DISABLE KEYS */;
INSERT INTO `specification_template` VALUES (1,1,'CPU Model',1),(2,1,'RAM (GB)',1),(3,1,'Storage (GB)',1),(4,1,'OS',1),(5,1,'Screen Size (inch)',1),(6,2,'CPU Model',1),(7,2,'RAM (GB)',1),(8,2,'Storage (GB)',1),(9,2,'OS',1),(10,3,'CPU Model',1),(11,3,'RAM (GB)',1),(12,3,'Storage (TB)',1),(13,3,'RAID Level',1),(14,5,'Screen Size (inch)',1),(15,5,'Resolution',1),(16,5,'Panel Type',1),(17,8,'Print Technology',1),(18,8,'Color/Mono',1),(19,8,'PPM',1),(20,8,'Connectivity',1),(21,12,'Port Count',1),(22,12,'Speed (Gbps)',1),(23,12,'Managed',1),(24,12,'PoE Enabled',1),(25,15,'Capacity (TB)',1),(26,15,'Interface',1),(27,19,'Brand',1),(28,19,'Model',1),(29,19,'IMEI',1),(30,21,'Lumens',1),(31,21,'Resolution',1);
/*!40000 ALTER TABLE `specification_template` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `status_master`
--

DROP TABLE IF EXISTS `status_master`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `status_master` (
  `status_id` int NOT NULL AUTO_INCREMENT,
  `status_name` varchar(50) NOT NULL,
  PRIMARY KEY (`status_id`),
  UNIQUE KEY `status_name` (`status_name`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `status_master`
--

LOCK TABLES `status_master` WRITE;
/*!40000 ALTER TABLE `status_master` DISABLE KEYS */;
INSERT INTO `status_master` VALUES (1,'Active'),(5,'Decommissioned'),(2,'Inactive'),(4,'Maintenance'),(3,'Replaced');
/*!40000 ALTER TABLE `status_master` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` varchar(50) NOT NULL,
  `full_name` varchar(255) DEFAULT NULL,
  `designation` varchar(255) DEFAULT NULL,
  `username` varchar(100) DEFAULT NULL,
  `status` tinyint DEFAULT '1',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'admin@example.com','$2a$10$EJS3Jmkv1f6n2hA4kHfL9.HMw0Qh3TJ80u0M1xZf840DEu1VvG9u.','admin','Admin User','System Administrator','admin',1),(2,'john.doe@example.com','$2a$10$EJS3Jmkv1f6n2hA4kHfL9.HMw0Qh3TJ80u0M1xZf840DEu1VvG9u.','user','John Doe','IT Support Specialist','johndoe',1),(3,'alice.smith@example.com','$2a$10$EJS3Jmkv1f6n2hA4kHfL9.HMw0Qh3TJ80u0M1xZf840DEu1VvG9u.','user','Alice Smith','Lab Technician','alicesmith',1),(4,'bob.johnson@example.com','$2a$10$EJS3Jmkv1f6n2hA4kHfL9.HMw0Qh3TJ80u0M1xZf840DEu1VvG9u.','manager','Bob Johnson','IT Manager','bobjohnson',1),(5,'carol.white@example.com','$2a$10$EJS3Jmkv1f6n2hA4kHfL9.HMw0Qh3TJ80u0M1xZf840DEu1VvG9u.','user','Carol White','Faculty Staff','carolwhite',1),(6,'rohit@gmail.com','$2a$14$rihflGcuD22h4SvbmLFTvuj5J8p6w8pkciJyCgpRqj3h3HzwJ05f6','admin','Rohit Jain','Manager','rohitjain',1);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'uits_final'
--
/*!50003 DROP PROCEDURE IF EXISTS `sp_add_device` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `sp_add_device`(
    IN d_name VARCHAR(255),
    IN t_id INT,
    IN serial VARCHAR(255),
    IN model VARCHAR(255),
    IN purchase DATE,
    IN warranty DATE,
    IN added_by INT,
    IN location INT,
    IN replaced_id INT,
    OUT new_device_id INT
)
BEGIN
    INSERT INTO Device (
        device_name, type_id, serial_no, model_no, purchase_date,
        warranty_expiry, added_by, location_id, replaced_device_id
    )
    VALUES (
        d_name, t_id, serial, model, purchase, warranty,
        added_by, location, replaced_id
    );
    
    SET new_device_id = LAST_INSERT_ID();
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `sp_add_device_category` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `sp_add_device_category`(IN cat_name VARCHAR(255))
BEGIN
    INSERT INTO Device_Category (category_name) VALUES (cat_name);
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `sp_add_device_specification` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `sp_add_device_specification`(
    IN p_device_id INT,
    IN p_spec_template_id INT,
    IN p_spec_master_id INT
)
BEGIN
    INSERT INTO Device_Specification (device_id, spec_template_id, spec_master_id)
    VALUES (p_device_id, p_spec_template_id, p_spec_master_id);
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `sp_add_device_type` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `sp_add_device_type`(IN cat_id INT, IN type_name VARCHAR(255))
BEGIN
    INSERT INTO Device_Type (category_id, type_name) VALUES (cat_id, type_name);
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `sp_add_specification_master` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `sp_add_specification_master`(
    IN p_spec_value VARCHAR(255),
    IN p_spec_template_id INT,
    OUT p_spec_id INT
)
BEGIN
    INSERT INTO Specification_Master (spec_value, spec_template_id)
    VALUES (p_spec_value, p_spec_template_id);

    SET p_spec_id = LAST_INSERT_ID();
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `sp_add_specification_template` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `sp_add_specification_template`(
    IN p_type_id INT,
    IN p_spec_name VARCHAR(255),
    OUT p_template_id INT
)
BEGIN
    INSERT INTO Specification_Template (type_id, spec_name)
    VALUES (p_type_id, p_spec_name);

    SET p_template_id = LAST_INSERT_ID();
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-05-23  1:23:49
