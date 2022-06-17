-- MySQL dump 10.13  Distrib 5.5.62, for Win64 (AMD64)
--
-- Host: 192.168.13.71    Database: wms
-- ------------------------------------------------------
-- Server version	8.0.21

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `i_adjustment_logs`
--

DROP TABLE IF EXISTS `i_adjustment_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `i_adjustment_logs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `adjustment_id` int NOT NULL DEFAULT '0',
  `transaction_id` int NOT NULL DEFAULT '0' COMMENT '货架位置',
  `quantity` int NOT NULL DEFAULT '0' COMMENT '商品总数量',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `i_adjustment_logs`
--

LOCK TABLES `i_adjustment_logs` WRITE;
/*!40000 ALTER TABLE `i_adjustment_logs` DISABLE KEYS */;
INSERT INTO `i_adjustment_logs` VALUES (1,1,2,1,'2021-12-30 08:50:15','andy2','2021-12-30 08:50:15','andy2');
/*!40000 ALTER TABLE `i_adjustment_logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `i_adjustments`
--

DROP TABLE IF EXISTS `i_adjustments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `i_adjustments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `sku` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商品SKU',
  `item_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `location_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0',
  `location_level` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `shelf_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '货架编码',
  `shelf_location` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '货架位置',
  `quantity` int NOT NULL DEFAULT '0' COMMENT '商品总数量',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `i_adjustments`
--

LOCK TABLES `i_adjustments` WRITE;
/*!40000 ALTER TABLE `i_adjustments` DISABLE KEYS */;
INSERT INTO `i_adjustments` VALUES (1,'106453893401','Container Round 4oz 100ml 50PCS*20SLV','BAY1AL1','G','BAY1A','B1A',-1,'aa','2021-12-30 08:50:15','andy2','2021-12-30 08:50:15','andy2');
/*!40000 ALTER TABLE `i_adjustments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `i_items`
--

DROP TABLE IF EXISTS `i_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `i_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `sku` varchar(64) NOT NULL DEFAULT '' COMMENT '商品SKU',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名称',
  `zoho_item_id` varchar(64) NOT NULL DEFAULT '' COMMENT 'zoho ItemID',
  `unit` varchar(64) NOT NULL DEFAULT '' COMMENT '商品单位',
  `stock` int NOT NULL DEFAULT '0' COMMENT '商品库存',
  `stock_available` int NOT NULL DEFAULT '0' COMMENT '商品可用库存',
  `stock_picking` int NOT NULL DEFAULT '0' COMMENT '捡货中库存',
  `stock_packing` int DEFAULT '0' COMMENT '打包中库存',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `zoho_item_id` (`zoho_item_id`),
  KEY `sku` (`sku`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1405 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `i_items`
--

LOCK TABLES `i_items` WRITE;
/*!40000 ALTER TABLE `i_items` DISABLE KEYS */;
INSERT INTO `i_items` VALUES (1,'106453893401','Container Round 4oz 100ml 50PCS*20SLV','8581000000118263','SLV',2,2,0,0,0,'2021-09-16 02:12:10','','2021-12-30 08:50:15','andy2'),(2,'106453893402','Container Round 8oz 100PCS*10SLV','8581000000118290','SLV',0,0,0,0,0,'2021-09-16 02:12:10','','2021-11-17 03:45:12','andy2'),(3,'106453893403','Container Round 10oz 100PCS*10SLV','8581000000118315','SLV',0,0,0,0,0,'2021-09-16 02:12:10','','2021-12-30 02:45:38','andy2'),(4,'106453893404','Container Round 70ml 100PCS*10SLV','8581000000118340','SLV',0,0,0,0,0,'2021-09-16 02:12:10','','2021-11-17 03:45:12','andy2'),(5,'106453893405','Container Round 150ml 50PCS*20SLV','8581000000118365','SLV',0,0,0,0,0,'2021-09-16 02:12:10','','2021-11-17 03:13:20','andy2'),(6,'106453893406','Container Round 440ml 50PCS*10SLV','8581000000118390','SLV',0,0,0,0,0,'2021-09-16 02:12:10','','2021-11-17 03:13:20','andy2'),(7,'106453893407','Container Round 220ml 50PCS*20SLV','8581000000118415','SLV',0,0,0,0,0,'2021-09-16 02:12:10','','2021-10-25 06:52:10','andy'),(8,'106453893408','Container Round 1750ml 50PCS*8SLV','8581000000118440','SLV',0,0,0,0,0,'2021-09-16 02:12:10','','2021-10-20 03:45:04',''),(9,'106453893409','Container Round 20oz (600ml) 50PCS*10SLV','8581000000118465','SLV',0,0,0,0,0,'2021-09-16 02:12:10','','2021-10-20 03:45:04',''),(10,'106453893410','Container Round 25oz (700ml) 50PCS*10SLV','8581000000118490','SLV',0,0,0,0,0,'2021-09-16 02:12:10','','2021-10-20 03:45:04',''),(11,'106453893411','Container Round 2oz 50PCS*20SLV','8581000000118515','SLV',0,0,0,0,0,'2021-09-16 02:12:10','','2021-10-20 03:45:04',''),(12,'106453893412','Container Round 3000ml 50PCS*2SLV','8581000000118542','SLV',0,0,0,0,0,'2021-09-16 02:12:10','','2021-10-20 03:45:04',''),(13,'106453893413','Container Round Tamper Evidence 300ml 30PCS*15SLV','8581000000118567','SLV',0,0,0,0,0,'2021-09-16 02:12:10','','2021-09-16 02:12:10',''),(14,'106453893414','Container Round Tamper Evidence 565ml 30PCS*15SLV','8581000000118576','SLV',0,0,0,0,0,'2021-09-16 02:12:10','','2021-09-16 02:12:10',''),(15,'106453893415','Container Round Tamper Evidence 870ml 25PCS*15SLV','8581000000118585','SLV',0,0,0,0,0,'2021-09-16 02:12:10','','2021-09-16 02:12:10',''),(16,'106453893417','NO Container Round 30oz 50PCS*10SLV','8581000000118596','SLV',0,0,0,0,0,'2021-09-16 02:12:10','','2021-10-20 03:45:04',''),(17,'107453893200','Container Rectangular 500ml Black 50PCS*10SLV','8581000000118621','SLV',0,0,0,0,0,'2021-09-16 02:12:10','','2021-10-20 03:45:04',''),(18,'107453893201','Container Rectangular 500ml 50PCS*10SLV','8581000000118646','SLV',0,0,0,0,0,'2021-09-16 02:12:10','','2021-10-20 03:45:04',''),(19,'107453893202','Container Rectangular 650ml 50PCS*10SLV','8581000000118671','SLV',0,0,0,0,0,'2021-09-16 02:12:10','','2021-10-20 03:45:04',''),(20,'107453893203','Container Rectangular 650ml 2 Compartments 50PCS*10SLV','8581000000118696','SLV',0,0,0,0,0,'2021-09-16 02:12:10','','2021-10-20 03:45:04','');
/*!40000 ALTER TABLE `i_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `i_packing_transactions`
--

DROP TABLE IF EXISTS `i_packing_transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `i_packing_transactions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `so_id` int NOT NULL DEFAULT '0' COMMENT '捡货单ID',
  `so_number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '捡货单编号',
  `item_name` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名称',
  `sku` varchar(64) NOT NULL DEFAULT '' COMMENT '商品SKU',
  `quantity` int NOT NULL DEFAULT '0' COMMENT '数量',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `i_packing_transactions`
--

LOCK TABLES `i_packing_transactions` WRITE;
/*!40000 ALTER TABLE `i_packing_transactions` DISABLE KEYS */;
/*!40000 ALTER TABLE `i_packing_transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `i_picking_order_details`
--

DROP TABLE IF EXISTS `i_picking_order_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `i_picking_order_details` (
  `id` int NOT NULL AUTO_INCREMENT,
  `picking_order_id` int NOT NULL DEFAULT '0',
  `shelf_location` varchar(64) NOT NULL DEFAULT '' COMMENT '货架位置',
  `shelf_code` varchar(64) NOT NULL DEFAULT '' COMMENT '货位编码',
  `location_level` varchar(64) NOT NULL DEFAULT '' COMMENT '货位所在层',
  `location_code` varchar(64) NOT NULL DEFAULT '' COMMENT '货位编码',
  `item_id` int NOT NULL DEFAULT '0',
  `sku` varchar(64) NOT NULL DEFAULT '',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名称',
  `quantity` int NOT NULL DEFAULT '0' COMMENT '商品总数量',
  `quantity_picked` int NOT NULL DEFAULT '0' COMMENT '已捡货数量',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  `zoho_item_id` varchar(64) NOT NULL DEFAULT '' COMMENT 'zoho商品ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `i_picking_order_details`
--

LOCK TABLES `i_picking_order_details` WRITE;
/*!40000 ALTER TABLE `i_picking_order_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `i_picking_order_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `i_picking_order_items`
--

DROP TABLE IF EXISTS `i_picking_order_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `i_picking_order_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `picking_order_id` int NOT NULL DEFAULT '0',
  `item_id` int NOT NULL DEFAULT '0',
  `sku` varchar(64) NOT NULL DEFAULT '',
  `zoho_item_id` varchar(64) NOT NULL DEFAULT '' COMMENT 'zoho商品ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名称',
  `quantity` int NOT NULL DEFAULT '0' COMMENT '商品总数量',
  `quantity_picked` int NOT NULL DEFAULT '0' COMMENT '已捡货数量',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `i_picking_order_items`
--

LOCK TABLES `i_picking_order_items` WRITE;
/*!40000 ALTER TABLE `i_picking_order_items` DISABLE KEYS */;
/*!40000 ALTER TABLE `i_picking_order_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `i_picking_order_logs`
--

DROP TABLE IF EXISTS `i_picking_order_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `i_picking_order_logs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `picking_order_id` int NOT NULL DEFAULT '0',
  `transaction_id` int NOT NULL DEFAULT '0' COMMENT '货架位置',
  `quantity` int NOT NULL DEFAULT '0' COMMENT '商品总数量',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `i_picking_order_logs`
--

LOCK TABLES `i_picking_order_logs` WRITE;
/*!40000 ALTER TABLE `i_picking_order_logs` DISABLE KEYS */;
/*!40000 ALTER TABLE `i_picking_order_logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `i_picking_orders`
--

DROP TABLE IF EXISTS `i_picking_orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `i_picking_orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '名称',
  `sales_orders` varchar(255) NOT NULL DEFAULT '' COMMENT '销售订单',
  `picking_date` date DEFAULT NULL COMMENT '捡货日期',
  `status` varchar(64) NOT NULL DEFAULT '' COMMENT '捡货单状态',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `i_picking_orders`
--

LOCK TABLES `i_picking_orders` WRITE;
/*!40000 ALTER TABLE `i_picking_orders` DISABLE KEYS */;
/*!40000 ALTER TABLE `i_picking_orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `i_picking_transactions`
--

DROP TABLE IF EXISTS `i_picking_transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `i_picking_transactions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `po_id` int NOT NULL DEFAULT '0' COMMENT '捡货单ID',
  `po_number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '捡货单编号',
  `item_name` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名称',
  `sku` varchar(64) NOT NULL DEFAULT '' COMMENT '商品SKU',
  `quantity` int NOT NULL DEFAULT '0' COMMENT '数量',
  `shelf_code` varchar(64) NOT NULL DEFAULT '' COMMENT '货架编码',
  `shelf_location` varchar(64) NOT NULL DEFAULT '' COMMENT '货架位置',
  `location_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '货位编码',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  `location_level` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `i_picking_transactions`
--

LOCK TABLES `i_picking_transactions` WRITE;
/*!40000 ALTER TABLE `i_picking_transactions` DISABLE KEYS */;
/*!40000 ALTER TABLE `i_picking_transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `i_purchase_order_items`
--

DROP TABLE IF EXISTS `i_purchase_order_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `i_purchase_order_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `po_id` int NOT NULL DEFAULT '0' COMMENT '采购订单ID',
  `item_id` int NOT NULL DEFAULT '0' COMMENT '商品ID',
  `sku` varchar(64) NOT NULL DEFAULT '' COMMENT 'SKU',
  `zoho_item_id` varchar(64) NOT NULL DEFAULT '' COMMENT 'zoho商品ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名称',
  `quantity` int NOT NULL DEFAULT '0' COMMENT '商品总数量',
  `quantity_received` int NOT NULL DEFAULT '0' COMMENT '已收取数量',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `i_purchase_order_items`
--

LOCK TABLES `i_purchase_order_items` WRITE;
/*!40000 ALTER TABLE `i_purchase_order_items` DISABLE KEYS */;
INSERT INTO `i_purchase_order_items` VALUES (1,1,1,'106453893401','8581000000118263','Container Round 4oz 100ml 50PCS*20SLV',30,3,1,'2021-10-20 09:29:34','Kenny','2021-12-30 08:48:06','andy2'),(2,1,2,'106453893402','8581000000118290','Container Round 8oz 100PCS*10SLV',30,0,1,'2021-10-20 09:29:34','Kenny','2021-11-17 03:45:12','andy2'),(3,2,2,'106453893402','8581000000118290','Container Round 8oz 100PCS*10SLV',30,0,1,'2021-10-20 09:29:34','Kenny','2021-11-17 03:45:12','andy2'),(4,2,3,'106453893403','8581000000118315','Container Round 10oz 100PCS*10SLV',30,0,1,'2021-10-20 09:29:34','Kenny','2021-11-17 03:45:12','andy2'),(5,2,4,'106453893404','8581000000118340','Container Round 70ml 100PCS*10SLV',30,0,1,'2021-10-20 09:29:34','Kenny','2021-11-17 03:45:12','andy2'),(6,3,3,'106453893403','8581000000118315','Container Round 10oz 100PCS*10SLV',40,0,1,'2021-10-20 09:29:34','Kenny','2021-11-17 03:45:12','andy2'),(7,4,4,'106453893404','8581000000118340','Container Round 70ml 100PCS*10SLV',30,0,1,'2021-10-20 09:29:34','Kenny','2021-10-26 01:42:50','andy'),(8,4,5,'106453893405','8581000000118365','Container Round 150ml 50PCS*20SLV',30,0,1,'2021-10-20 09:29:34','Kenny','2021-10-26 08:20:42','andy'),(9,4,6,'106453893406','8581000000118390','Container Round 440ml 50PCS*10SLV',30,0,1,'2021-10-20 09:29:34','Kenny','2021-10-26 01:42:50','andy'),(10,5,5,'106453893405','8581000000118365','Container Round 150ml 50PCS*20SLV',30,0,1,'2021-10-20 09:29:34','Kenny','2021-10-26 01:42:50','Kenny'),(11,5,6,'106453893406','8581000000118390','Container Round 440ml 50PCS*10SLV',30,0,1,'2021-10-20 09:29:34','Kenny','2021-10-26 01:42:50','Kenny'),(12,5,7,'106453893407','8581000000118415','Container Round 220ml 50PCS*20SLV',30,0,1,'2021-10-20 09:29:34','Kenny','2021-10-26 01:42:50','andy'),(13,6,6,'106453893406','8581000000118390','Container Round 440ml 50PCS*10SLV',30,0,1,'2021-10-20 09:29:34','Kenny','2021-10-26 01:42:50','Kenny'),(14,6,7,'106453893407','8581000000118415','Container Round 220ml 50PCS*20SLV',30,0,1,'2021-10-20 09:29:34','Kenny','2021-10-26 01:42:50','Kenny'),(15,6,8,'106453893408','8581000000118440','Container Round 1750ml 50PCS*8SLV',30,0,1,'2021-10-20 09:29:34','Kenny','2021-10-26 01:42:50','Kenny'),(16,7,7,'106453893407','8581000000118415','Container Round 220ml 50PCS*20SLV',3,0,1,'2021-10-20 09:29:34','Kenny','2021-10-20 09:29:34','Kenny'),(17,7,8,'106453893408','8581000000118440','Container Round 1750ml 50PCS*8SLV',4,0,1,'2021-10-20 09:29:34','Kenny','2021-10-20 09:29:34','Kenny'),(18,7,9,'106453893409','8581000000118465','Container Round 20oz (600ml) 50PCS*10SLV',5,0,1,'2021-10-20 09:29:34','Kenny','2021-10-20 09:29:34','Kenny'),(19,8,8,'106453893408','8581000000118440','Container Round 1750ml 50PCS*8SLV',3,0,1,'2021-10-20 09:29:34','Kenny','2021-10-20 09:29:34','Kenny'),(20,8,9,'106453893409','8581000000118465','Container Round 20oz (600ml) 50PCS*10SLV',4,0,1,'2021-10-20 09:29:34','Kenny','2021-10-20 09:29:34','Kenny'),(21,8,10,'106453893410','8581000000118490','Container Round 25oz (700ml) 50PCS*10SLV',5,0,1,'2021-10-20 09:29:34','Kenny','2021-10-20 09:29:34','Kenny'),(22,9,3,'106453893403','8581000000118315','Container Round 10oz 100PCS*10SLV',3,0,1,'2021-10-20 09:29:34','Kenny','2021-10-21 03:19:51','andy'),(23,9,5,'106453893405','8581000000118365','Container Round 150ml 50PCS*20SLV',4,0,1,'2021-10-20 09:29:34','Kenny','2021-10-21 03:19:51','andy'),(24,9,6,'106453893406','8581000000118390','Container Round 440ml 50PCS*10SLV',5,0,1,'2021-10-20 09:29:34','Kenny','2021-10-21 03:19:51','andy'),(25,10,8,'106453893408','8581000000118440','Container Round 1750ml 50PCS*8SLV',3,0,1,'2021-10-20 09:29:34','Kenny','2021-10-20 09:29:34','Kenny'),(26,10,9,'106453893409','8581000000118465','Container Round 20oz (600ml) 50PCS*10SLV',4,0,1,'2021-10-20 09:29:34','Kenny','2021-10-20 09:29:34','Kenny'),(27,10,10,'106453893410','8581000000118490','Container Round 25oz (700ml) 50PCS*10SLV',5,0,1,'2021-10-20 09:29:34','Kenny','2021-10-20 09:29:34','Kenny'),(30,0,0,'','','',0,0,0,'2021-10-26 01:42:50','','2021-10-26 01:42:50','');
/*!40000 ALTER TABLE `i_purchase_order_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `i_purchase_orders`
--

DROP TABLE IF EXISTS `i_purchase_orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `i_purchase_orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `zoho_po_id` varchar(64) NOT NULL DEFAULT '' COMMENT 'zoho 采购订单ID',
  `po_number` varchar(64) NOT NULL DEFAULT '' COMMENT '采购订单编码',
  `po_date` date DEFAULT NULL COMMENT '采购订单日期',
  `expected_delivery_date` date DEFAULT NULL COMMENT '预计到货日期',
  `reference_number` varchar(64) NOT NULL DEFAULT '' COMMENT '参考编码',
  `status` varchar(64) NOT NULL DEFAULT '' COMMENT '采购订单状态',
  `vendor_id` varchar(64) NOT NULL DEFAULT '' COMMENT '供应商ID',
  `vendor_name` varchar(64) NOT NULL DEFAULT '' COMMENT '供应商名称',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `po_number` (`po_number`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `i_purchase_orders`
--

LOCK TABLES `i_purchase_orders` WRITE;
/*!40000 ALTER TABLE `i_purchase_orders` DISABLE KEYS */;
INSERT INTO `i_purchase_orders` VALUES (1,'1','po-1','2021-10-01','2021-10-01','','CONFIRM','','LinYuan',1,'2021-10-20 09:29:33','Kenny','2021-11-17 03:45:12','andy2'),(2,'2','po-2','2021-10-02','2021-10-02','','CONFIRM','','LinYuan',1,'2021-10-20 09:29:33','Kenny','2021-11-12 09:28:27','Kenny'),(3,'3','po-3','2021-10-03','2021-10-03','','CONFIRM','','LinYuan',1,'2021-10-20 09:29:33','Kenny','2021-11-12 09:28:27','Kenny'),(4,'4','po-4','2021-10-04','2021-10-04','','CONFIRM','','LinYuan',1,'2021-10-20 09:29:33','Kenny','2021-11-12 09:28:27','Kenny'),(5,'5','po-5','2021-10-05','2021-10-05','','CONFIRM','','LinYuan',1,'2021-10-20 09:29:33','Kenny','2021-11-12 09:28:27','Kenny'),(6,'6','po-6','2021-10-06','2021-10-06','','CONFIRM','','LinYuan',1,'2021-10-20 09:29:33','Kenny','2021-11-12 09:28:27','Kenny'),(7,'7','po-7','2021-10-07','2021-10-07','','CONFIRM','','LinYuan',1,'2021-10-20 09:29:33','Kenny','2021-11-12 09:28:27','Kenny'),(8,'8','po-8','2021-10-08','2021-10-08','','CONFIRM','','LinYuan',1,'2021-10-20 09:29:33','Kenny','2021-11-12 09:28:27','Kenny'),(9,'9','po-9','2021-10-09','2021-10-09','','CONFIRM','','LinYuan',1,'2021-10-20 09:29:33','Kenny','2021-11-12 09:28:27','Kenny'),(10,'10','po-10','2021-10-10','2021-10-10','','CONFIRM','','LinYuan',1,'2021-10-20 09:29:33','Kenny','2021-11-12 09:28:27','Kenny');
/*!40000 ALTER TABLE `i_purchase_orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `i_sales_order_items`
--

DROP TABLE IF EXISTS `i_sales_order_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `i_sales_order_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `so_id` int NOT NULL DEFAULT '0' COMMENT '采购订单ID',
  `item_id` int NOT NULL DEFAULT '0' COMMENT '商品ID',
  `sku` varchar(64) NOT NULL DEFAULT '' COMMENT 'SKU',
  `zoho_item_id` varchar(64) NOT NULL DEFAULT '' COMMENT 'zoho商品ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名称',
  `quantity` int NOT NULL DEFAULT '0' COMMENT '商品总数量',
  `quantity_picked` int NOT NULL DEFAULT '0' COMMENT '已捡货数量',
  `quantity_packed` int NOT NULL DEFAULT '0' COMMENT '已打包数量',
  `quantity_shipped` int NOT NULL DEFAULT '0' COMMENT '已发货数量',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `i_sales_order_items`
--

LOCK TABLES `i_sales_order_items` WRITE;
/*!40000 ALTER TABLE `i_sales_order_items` DISABLE KEYS */;
INSERT INTO `i_sales_order_items` VALUES (1,1,1,'106453893401','8581000000118263','Container Round 4oz 100ml 50PCS*20SLV',3,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-11-17 03:45:12','andy2'),(2,1,2,'106453893402','8581000000118290','Container Round 8oz 100PCS*10SLV',4,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-11-17 03:45:12','andy2'),(3,1,3,'106453893403','8581000000118315','Container Round 10oz 100PCS*10SLV',5,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-11-17 03:45:12','andy2'),(4,2,2,'106453893402','8581000000118290','Container Round 8oz 100PCS*10SLV',3,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-11-17 03:45:12','andy2'),(5,2,3,'106453893403','8581000000118315','Container Round 10oz 100PCS*10SLV',4,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-11-17 03:45:12','andy2'),(6,2,4,'106453893404','8581000000118340','Container Round 70ml 100PCS*10SLV',5,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-11-17 03:45:12','andy2'),(7,3,3,'106453893403','8581000000118315','Container Round 10oz 100PCS*10SLV',3,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(8,3,4,'106453893404','8581000000118340','Container Round 70ml 100PCS*10SLV',4,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(9,3,5,'106453893405','8581000000118365','Container Round 150ml 50PCS*20SLV',5,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(10,4,5,'106453893405','8581000000118365','Container Round 150ml 50PCS*20SLV',3,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(11,4,6,'106453893406','8581000000118390','Container Round 440ml 50PCS*10SLV',4,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(12,4,7,'106453893407','8581000000118415','Container Round 220ml 50PCS*20SLV',5,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(13,5,6,'106453893406','8581000000118390','Container Round 440ml 50PCS*10SLV',3,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(14,5,7,'106453893407','8581000000118415','Container Round 220ml 50PCS*20SLV',4,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(15,5,8,'106453893408','8581000000118440','Container Round 1750ml 50PCS*8SLV',5,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(16,6,7,'106453893407','8581000000118415','Container Round 220ml 50PCS*20SLV',3,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(17,6,8,'106453893408','8581000000118440','Container Round 1750ml 50PCS*8SLV',4,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(18,6,9,'106453893409','8581000000118465','Container Round 20oz (600ml) 50PCS*10SLV',5,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(19,7,8,'106453893408','8581000000118440','Container Round 1750ml 50PCS*8SLV',3,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(20,7,9,'106453893409','8581000000118465','Container Round 20oz (600ml) 50PCS*10SLV',4,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(21,7,10,'106453893410','8581000000118490','Container Round 25oz (700ml) 50PCS*10SLV',5,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(22,8,1,'106453893401','8581000000118263','Container Round 4oz 100ml 50PCS*20SLV',3,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(23,8,3,'106453893403','8581000000118315','Container Round 10oz 100PCS*10SLV',4,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(24,8,5,'106453893405','8581000000118365','Container Round 150ml 50PCS*20SLV',5,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(25,9,2,'106453893402','8581000000118290','Container Round 8oz 100PCS*10SLV',3,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(26,9,4,'106453893404','8581000000118340','Container Round 70ml 100PCS*10SLV',4,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(27,9,6,'106453893406','8581000000118390','Container Round 440ml 50PCS*10SLV',5,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(28,10,3,'106453893403','8581000000118315','Container Round 10oz 100PCS*10SLV',3,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(29,10,5,'106453893405','8581000000118365','Container Round 150ml 50PCS*20SLV',4,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(30,10,6,'106453893406','8581000000118390','Container Round 440ml 50PCS*10SLV',5,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(31,11,1,'106453893401','8581000000118263','Container Round 4oz 100ml 50PCS*20SLV',3,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(32,11,2,'106453893402','8581000000118290','Container Round 8oz 100PCS*10SLV',4,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(33,11,5,'106453893405','8581000000118365','Container Round 150ml 50PCS*20SLV',5,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(34,12,4,'106453893404','8581000000118340','Container Round 70ml 100PCS*10SLV',3,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(35,12,8,'106453893408','8581000000118440','Container Round 1750ml 50PCS*8SLV',4,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(36,12,9,'106453893409','8581000000118465','Container Round 20oz (600ml) 50PCS*10SLV',5,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(37,13,6,'106453893406','8581000000118390','Container Round 440ml 50PCS*10SLV',3,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(38,13,7,'106453893407','8581000000118415','Container Round 220ml 50PCS*20SLV',4,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(39,13,8,'106453893408','8581000000118440','Container Round 1750ml 50PCS*8SLV',5,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(40,14,3,'106453893403','8581000000118315','Container Round 10oz 100PCS*10SLV',3,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(41,14,5,'106453893405','8581000000118365','Container Round 150ml 50PCS*20SLV',4,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(42,14,6,'106453893406','8581000000118390','Container Round 440ml 50PCS*10SLV',5,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(43,15,1,'106453893401','8581000000118263','Container Round 4oz 100ml 50PCS*20SLV',3,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(44,15,4,'106453893404','8581000000118340','Container Round 70ml 100PCS*10SLV',4,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny'),(45,15,10,'106453893410','8581000000118490','Container Round 25oz (700ml) 50PCS*10SLV',5,0,0,0,1,'2021-10-20 09:29:33','Kenny','2021-10-20 09:29:33','Kenny');
/*!40000 ALTER TABLE `i_sales_order_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `i_sales_order_pallets`
--

DROP TABLE IF EXISTS `i_sales_order_pallets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `i_sales_order_pallets` (
  `id` int NOT NULL AUTO_INCREMENT,
  `so_id` int NOT NULL DEFAULT '0' COMMENT '采购订单ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT 'pallet名称',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态1 shipped, 2 returned',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `i_sales_order_pallets`
--

LOCK TABLES `i_sales_order_pallets` WRITE;
/*!40000 ALTER TABLE `i_sales_order_pallets` DISABLE KEYS */;
INSERT INTO `i_sales_order_pallets` VALUES (2,1,'lewis',1,'2021-12-31 08:42:26','andy2','2021-12-31 09:00:55','andy2');
/*!40000 ALTER TABLE `i_sales_order_pallets` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `i_sales_orders`
--

DROP TABLE IF EXISTS `i_sales_orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `i_sales_orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `zoho_so_id` varchar(64) NOT NULL DEFAULT '' COMMENT 'zoho订单ID',
  `so_number` varchar(64) NOT NULL DEFAULT '' COMMENT '订单编号',
  `so_date` date DEFAULT NULL COMMENT '订单日期',
  `customer_id` varchar(64) NOT NULL DEFAULT '' COMMENT '供应商ID',
  `customer_name` varchar(64) NOT NULL DEFAULT '' COMMENT '供应商名称',
  `status` varchar(64) NOT NULL DEFAULT '' COMMENT '采购订单状态',
  `sales_name` varchar(64) NOT NULL DEFAULT '' COMMENT '销售人员名称',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  `has_pallet` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否带有pallet',
  `expected_shipment_date` date DEFAULT NULL COMMENT '预计到货日期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `i_sales_orders`
--

LOCK TABLES `i_sales_orders` WRITE;
/*!40000 ALTER TABLE `i_sales_orders` DISABLE KEYS */;
INSERT INTO `i_sales_orders` VALUES (1,'1','so-1','2021-10-01','','WangZiXin','CONFIRMED','LinYuan',1,'2021-10-20 09:29:33','Kenny','2022-06-01 09:04:02','andy2',1,'2021-10-01'),(2,'2','so-2','2021-10-02','','WangZiXin','CONFIRMED','LinYuan',1,'2021-10-20 09:29:33','Kenny','2022-06-01 09:04:02','andy2',0,'2021-10-01'),(3,'3','so-3','2021-10-03','','WangZiXin','CONFIRMED','LinYuan',1,'2021-10-20 09:29:33','Kenny','2022-06-01 09:04:02','andy',0,'2021-10-01'),(4,'4','so-4','2021-10-04','','WangZiXin','CONFIRMED','LinYuan',1,'2021-10-20 09:29:33','Kenny','2022-06-01 09:04:02','andy',0,'2021-10-01'),(5,'5','so-5','2021-10-05','','WangZiXin','CONFIRMED','LinYuan',1,'2021-10-20 09:29:33','Kenny','2022-06-01 09:04:02','Kenny',0,'2021-10-01'),(6,'6','so-6','2021-10-06','','WangZiXin','CONFIRMED','LinYuan',1,'2021-10-20 09:29:33','Kenny','2022-06-01 09:04:02','Kenny',0,'2021-10-01'),(7,'7','so-7','2021-10-07','','WangZiXin','CONFIRMED','LinYuan',1,'2021-10-20 09:29:33','Kenny','2022-06-01 09:04:02','Kenny',0,'2021-10-01'),(8,'8','so-8','2021-10-08','','WangZiXin','CONFIRMED','LinYuan',1,'2021-10-20 09:29:33','Kenny','2022-06-01 09:04:03','Kenny',0,'2021-10-01'),(9,'9','so-9','2021-10-09','','WangZiXin','CONFIRMED','LinYuan',1,'2021-10-20 09:29:33','Kenny','2022-06-01 09:04:03','Kenny',0,'2021-10-01'),(10,'10','so-10','2021-10-10','','WangZiXin','CONFIRMED','LinYuan',1,'2021-10-20 09:29:33','Kenny','2022-06-01 09:04:03','Kenny',0,'2021-10-01'),(11,'11','so-11','2021-10-11','','WangZiXin','CONFIRMED','LinYuan',1,'2021-10-20 09:29:33','Kenny','2022-06-01 09:04:03','Kenny',0,'2021-10-01'),(12,'12','so-12','2021-10-12','','WangZiXin','CONFIRMED','LinYuan',1,'2021-10-20 09:29:33','Kenny','2022-06-01 09:04:03','Kenny',0,'2021-10-01'),(13,'13','so-13','2021-10-13','','WangZiXin','CONFIRMED','LinYuan',1,'2021-10-20 09:29:33','Kenny','2022-06-01 09:04:03','Kenny',0,'2021-10-01'),(14,'14','so-14','2021-10-14','','WangZiXin','CONFIRMED','LinYuan',1,'2021-10-20 09:29:33','Kenny','2022-06-01 09:04:03','Kenny',0,'2021-10-01'),(15,'15','so-15','2021-10-15','','WangZiXin','CONFIRMED','LinYuan',1,'2021-10-20 09:29:33','Kenny','2022-06-01 09:04:03','Kenny',0,'2021-10-01');
/*!40000 ALTER TABLE `i_sales_orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `i_transactions`
--

DROP TABLE IF EXISTS `i_transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `i_transactions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `po_id` int NOT NULL DEFAULT '0' COMMENT '采购单ID',
  `po_number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '采购单编号',
  `item_name` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名称',
  `sku` varchar(64) NOT NULL DEFAULT '' COMMENT '商品SKU',
  `quantity` int NOT NULL DEFAULT '0' COMMENT '数量',
  `balance` int NOT NULL DEFAULT '0',
  `shelf_code` varchar(64) NOT NULL DEFAULT '' COMMENT '货架编码',
  `shelf_location` varchar(64) NOT NULL DEFAULT '' COMMENT '货架位置',
  `location_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '货位编码',
  `location_level` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `i_transactions`
--

LOCK TABLES `i_transactions` WRITE;
/*!40000 ALTER TABLE `i_transactions` DISABLE KEYS */;
INSERT INTO `i_transactions` VALUES (1,1,'po-1','Container Round 4oz 100ml 50PCS*20SLV','106453893401',1,0,'BAY1A','B1A','BAY1AL1','G',1,'2021-12-30 08:47:57','andy2','2021-12-30 08:48:45','andy2'),(2,1,'po-1','Container Round 4oz 100ml 50PCS*20SLV','106453893401',1,0,'BAY1A','B1A','BAY1AL1','G',1,'2021-12-30 08:48:03','andy2','2021-12-30 08:50:15','andy2'),(3,1,'po-1','Container Round 4oz 100ml 50PCS*20SLV','106453893401',1,1,'BAY1A','B1A','BAY1AL1','G',1,'2021-12-30 08:48:06','andy2','2021-12-30 08:48:06','andy2'),(4,1,'po-1','Container Round 4oz 100ml 50PCS*20SLV','106453893401',1,1,'BAY1A','B1A','BAY1AL2','2',1,'2021-12-30 08:47:57','andy2','2021-12-30 08:48:45','andy2');
/*!40000 ALTER TABLE `i_transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `i_transfer_transactions`
--

DROP TABLE IF EXISTS `i_transfer_transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `i_transfer_transactions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `from_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '来源货位',
  `to_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '目标货位',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名称',
  `sku` varchar(64) NOT NULL DEFAULT '' COMMENT '商品SKU',
  `quantity` int NOT NULL DEFAULT '0' COMMENT '数量',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `i_transfer_transactions`
--

LOCK TABLES `i_transfer_transactions` WRITE;
/*!40000 ALTER TABLE `i_transfer_transactions` DISABLE KEYS */;
INSERT INTO `i_transfer_transactions` VALUES (1,'BAY1AL1','BAY1AL2','','106453893401',1,1,'2021-12-30 08:48:45','andy2','2021-12-30 08:48:45','andy2');
/*!40000 ALTER TABLE `i_transfer_transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_barcodes`
--

DROP TABLE IF EXISTS `s_barcodes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `s_barcodes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(64) NOT NULL DEFAULT '' COMMENT 'barcode编码',
  `sku` varchar(64) NOT NULL DEFAULT '' COMMENT '商品SKU',
  `unit` varchar(64) NOT NULL DEFAULT '' COMMENT '商品单位',
  `quantity` int NOT NULL DEFAULT '0' COMMENT '商品数量',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_barcodes`
--

LOCK TABLES `s_barcodes` WRITE;
/*!40000 ALTER TABLE `s_barcodes` DISABLE KEYS */;
INSERT INTO `s_barcodes` VALUES (1,'5784587545','106453893401','SLV',1,1,'2021-10-20 06:29:27','andy','2021-10-20 08:09:50','andy'),(2,'4738927485','106453893402','SLV',1,1,'2021-10-20 07:02:26','andy','2021-10-20 08:09:45','andy'),(3,'4556354346','106453893403','SLV',1,1,'2021-10-20 07:02:44','andy','2021-10-20 07:02:44','andy'),(4,'4324252355','106453893404','SLV',1,1,'2021-10-20 08:07:40','andy','2021-10-20 08:07:40','andy'),(5,'3434343244','106453893405','SLV',1,1,'2021-10-20 08:09:23','andy','2021-10-20 08:09:23','andy'),(6,'5435423646','106453893406','SLV',1,1,'2021-10-20 08:09:37','andy','2021-10-20 08:09:37','andy'),(7,'5436365464','106453893407','SLV',1,1,'2021-10-20 08:10:06','andy','2021-10-20 08:10:43','andy'),(8,'4324324234','106453893408','SLV',1,1,'2021-10-20 08:10:24','andy','2021-10-20 08:10:48','andy'),(9,'5432543524','106453893409','SLV',1,1,'2021-10-20 08:10:35','andy','2021-10-20 08:10:51','andy'),(10,'12346546','106453893411','SLV',1,1,'2021-10-23 06:52:14','andy','2021-10-23 06:52:14','andy'),(11,'54645465','106453893403','SLV',1,1,'2021-10-23 07:19:48','andy','2021-10-23 07:19:48','andy');
/*!40000 ALTER TABLE `s_barcodes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_locations`
--

DROP TABLE IF EXISTS `s_locations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `s_locations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(64) NOT NULL DEFAULT '' COMMENT '货位编码',
  `level` varchar(64) NOT NULL DEFAULT '' COMMENT '所在层',
  `shelf_id` int NOT NULL DEFAULT '0' COMMENT '货架ID',
  `sku` varchar(64) NOT NULL DEFAULT '' COMMENT '商品SKU',
  `capacity` int NOT NULL COMMENT '可容纳商品',
  `alert` int DEFAULT '0',
  `quantity` int NOT NULL DEFAULT '0' COMMENT '已存放数量',
  `available` int NOT NULL DEFAULT '0' COMMENT '剩余空间',
  `can_pick` int NOT NULL DEFAULT '0',
  `unit` varchar(64) NOT NULL DEFAULT '' COMMENT '商品单位',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_locations`
--

LOCK TABLES `s_locations` WRITE;
/*!40000 ALTER TABLE `s_locations` DISABLE KEYS */;
INSERT INTO `s_locations` VALUES (1,'BAY1AL1','G',1,'106453893401',10,0,1,9,1,'SLV',1,'2021-10-20 06:28:31','andy','2021-12-30 08:50:15','andy2'),(2,'BAY1AL2','2',1,'106453893401',5,0,1,4,1,'SLV',1,'2021-10-20 06:47:48','andy','2021-12-30 08:48:45','andy2'),(3,'BAY1AL3a','3',1,'106453893401',5,0,0,5,0,'SLV',1,'2021-10-20 06:50:26','andy','2021-12-30 08:47:33','andy2'),(4,'BAY1AL3b','1',1,'106453893401',5,0,0,5,0,'SLV',1,'2021-10-20 07:29:51','andy','2021-12-30 08:47:33','andy2'),(5,'B1BL1','1',2,'106453893404',20,5,0,20,0,'SLV',1,'2021-10-20 08:12:48','andy','2021-12-30 06:54:59','andy2'),(6,'B2AL3a','3',3,'106453893404',20,0,0,20,0,'SLV',1,'2021-10-20 08:27:04','andy','2021-12-30 06:54:59','andy'),(7,'B2BL2','2',5,'106453893404',20,0,0,20,0,'SLV',1,'2021-10-20 08:28:15','andy','2021-12-30 06:54:59','andy'),(8,'B1BL2','2',2,'106453893404',50,20,0,50,0,'SLV',1,'2021-10-20 08:29:30','andy','2021-12-30 06:54:59','andy2'),(9,'B2AL1','1',3,'106453893405',20,0,0,20,0,'SLV',1,'2021-10-20 08:30:10','andy','2021-12-30 06:54:59','andy'),(10,'B2BL1a','1',5,'106453893405',20,5,0,20,0,'SLV',1,'2021-10-20 08:31:00','andy','2021-12-30 06:54:59','andy2'),(11,'B2AL2','2',3,'106453893406',20,5,0,20,0,'SLV',1,'2021-10-20 08:32:38','andy','2021-12-30 06:54:59','andy2'),(12,'B2BL3','3',5,'106453893407',20,0,0,20,0,'SLV',1,'2021-10-20 08:33:08','andy','2021-12-30 06:54:59','andy'),(13,'B1BL3a','3',2,'106453893408',20,5,0,20,0,'SLV',1,'2021-10-20 08:33:45','andy','2021-12-30 06:54:59','andy'),(14,'B1BL3b','3',2,'106453893409',20,2,0,20,0,'SLV',1,'2021-10-20 08:34:52','andy','2021-12-30 06:54:59','andy'),(15,'B2BL1b','1',5,'106453893409',30,0,0,30,0,'SLV',1,'2021-10-20 08:35:33','andy','2021-12-30 06:54:59','andy'),(16,'B2AL3b','3',3,'106453893403',4,0,0,4,0,'SLV',1,'2021-10-21 03:57:58','andy','2021-12-30 06:54:59','andy2'),(17,'fafd','1',1,'106453893403',1,1,0,1,0,'SLV',2,'2021-10-25 03:03:03','andy','2021-12-30 06:54:59','andy');
/*!40000 ALTER TABLE `s_locations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_shelves`
--

DROP TABLE IF EXISTS `s_shelves`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `s_shelves` (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(64) NOT NULL DEFAULT '' COMMENT '货架编码',
  `level` tinyint NOT NULL DEFAULT '1' COMMENT '货架层数',
  `location` varchar(64) NOT NULL DEFAULT '' COMMENT '货架位置',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_shelves`
--

LOCK TABLES `s_shelves` WRITE;
/*!40000 ALTER TABLE `s_shelves` DISABLE KEYS */;
INSERT INTO `s_shelves` VALUES (1,'BAY1A',3,'B1A',1,'2021-10-20 06:22:53','andy','2021-10-20 07:59:26','andy'),(2,'BAY1B',3,'B1B',1,'2021-10-20 08:02:56','andy','2021-10-20 08:02:56','andy'),(3,'BAY2A',3,'B2A',1,'2021-10-20 08:03:19','andy','2021-10-20 08:03:27','andy'),(5,'BAY2B',3,'B2B',1,'2021-10-20 08:04:50','andy','2021-10-20 08:26:29','andy');
/*!40000 ALTER TABLE `s_shelves` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_apis`
--

DROP TABLE IF EXISTS `user_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_apis` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL DEFAULT '',
  `route` varchar(128) NOT NULL DEFAULT '',
  `method` varchar(24) NOT NULL DEFAULT '',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `route` (`route`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_apis`
--

LOCK TABLES `user_apis` WRITE;
/*!40000 ALTER TABLE `user_apis` DISABLE KEYS */;
INSERT INTO `user_apis` VALUES (1,'GET/swagger/*any','/swagger/*any','GET',1,'2021-11-09 09:40:39','lewis','2021-11-09 09:41:12','lewis'),(2,'POST/signin','/signin','POST',1,'2021-11-09 09:40:39','lewis','2021-11-09 09:41:12','lewis'),(3,'POST/signup','/signup','POST',1,'2021-11-09 09:40:39','lewis','2021-11-09 09:41:12','lewis'),(4,'GET/users','/users','GET',1,'2021-11-09 09:40:39','lewis','2021-11-09 09:41:12','lewis'),(5,'GET/users/:id','/users/:id','GET',1,'2021-11-09 09:40:40','lewis','2021-11-09 09:41:12','lewis'),(6,'POST/users','/users','POST',1,'2021-11-09 09:40:40','lewis','2021-11-09 09:41:12','lewis'),(7,'GET/shelves','/shelves','GET',1,'2021-11-09 09:40:40','lewis','2021-11-09 09:41:12','lewis'),(8,'GET/shelves/:id','/shelves/:id','GET',1,'2021-11-09 09:40:40','lewis','2021-11-09 09:41:12','lewis'),(9,'PUT/shelves/:id','/shelves/:id','PUT',1,'2021-11-09 09:40:40','lewis','2021-11-09 09:41:12','lewis'),(10,'POST/shelves','/shelves','POST',1,'2021-11-09 09:40:40','lewis','2021-11-09 09:41:12','lewis'),(11,'GET/locations','/locations','GET',1,'2021-11-09 09:40:40','lewis','2021-11-09 09:41:12','lewis'),(12,'GET/locations/:id','/locations/:id','GET',1,'2021-11-09 09:40:40','lewis','2021-11-09 09:41:12','lewis'),(13,'PUT/locations/:id','/locations/:id','PUT',1,'2021-11-09 09:40:40','lewis','2021-11-09 09:41:12','lewis'),(14,'POST/locations','/locations','POST',1,'2021-11-09 09:40:40','lewis','2021-11-09 09:41:12','lewis'),(15,'GET/barcodes','/barcodes','GET',1,'2021-11-09 09:40:40','lewis','2021-11-09 09:41:12','lewis'),(16,'GET/barcodes/:id','/barcodes/:id','GET',1,'2021-11-09 09:40:40','lewis','2021-11-09 09:41:12','lewis'),(17,'PUT/barcodes/:id','/barcodes/:id','PUT',1,'2021-11-09 09:40:41','lewis','2021-11-09 09:41:12','lewis'),(18,'POST/barcodes','/barcodes','POST',1,'2021-11-09 09:40:41','lewis','2021-11-09 09:41:12','lewis'),(19,'POST/transfers','/transfers','POST',1,'2021-11-09 09:40:41','lewis','2021-11-09 09:41:12','lewis'),(20,'GET/transfers','/transfers','GET',1,'2021-11-09 09:40:41','lewis','2021-11-09 09:41:12','lewis'),(21,'GET/items','/items','GET',1,'2021-11-09 09:40:41','lewis','2021-11-09 09:41:12','lewis'),(22,'GET/items/:id','/items/:id','GET',1,'2021-11-09 09:40:41','lewis','2021-11-09 09:41:12','lewis'),(23,'GET/purchaseorders','/purchaseorders','GET',1,'2021-11-09 09:40:41','lewis','2021-11-09 09:41:12','lewis'),(24,'GET/purchaseorders/:id','/purchaseorders/:id','GET',1,'2021-11-09 09:40:41','lewis','2021-11-09 09:41:12','lewis'),(25,'POST/receives','/receives','POST',1,'2021-11-09 09:40:41','lewis','2021-11-09 09:41:12','lewis'),(26,'GET/receives','/receives','GET',1,'2021-11-09 09:40:41','lewis','2021-11-09 09:41:12','lewis'),(27,'GET/salesorders','/salesorders','GET',1,'2021-11-09 09:40:41','lewis','2021-11-09 09:41:12','lewis'),(28,'GET/salesorders/:id','/salesorders/:id','GET',1,'2021-11-09 09:40:41','lewis','2021-11-09 09:41:12','lewis'),(29,'GET/pickingorders','/pickingorders','GET',1,'2021-11-09 09:40:42','lewis','2021-11-09 09:41:12','lewis'),(30,'GET/pickingorders/:id','/pickingorders/:id','GET',1,'2021-11-09 09:40:42','lewis','2021-11-09 09:41:12','lewis'),(31,'POST/pickingorders','/pickingorders','POST',1,'2021-11-09 09:40:42','lewis','2021-11-09 09:41:12','lewis'),(32,'POST/pickings','/pickings','POST',1,'2021-11-09 09:40:42','lewis','2021-11-09 09:41:12','lewis'),(33,'POST/packings','/packings','POST',1,'2021-11-09 09:40:42','lewis','2021-11-09 09:41:12','lewis'),(34,'GET/roles','/roles','GET',1,'2021-11-09 09:40:42','lewis','2021-11-09 09:41:12','lewis'),(35,'GET/roles/:id','/roles/:id','GET',1,'2021-11-09 09:40:42','lewis','2021-11-09 09:41:12','lewis'),(36,'PUT/roles/:id','/roles/:id','PUT',1,'2021-11-09 09:40:42','lewis','2021-11-09 09:41:12','lewis'),(37,'POST/roles','/roles','POST',1,'2021-11-09 09:40:42','lewis','2021-11-09 09:41:12','lewis'),(38,'GET/apis','/apis','GET',1,'2021-11-09 09:40:42','lewis','2021-11-09 09:41:12','lewis'),(39,'GET/apis/:id','/apis/:id','GET',1,'2021-11-09 09:40:42','lewis','2021-11-09 09:41:12','lewis'),(40,'PUT/apis/:id','/apis/:id','PUT',1,'2021-11-09 09:40:43','lewis','2021-11-09 09:41:12','lewis'),(41,'POST/apis','/apis','POST',1,'2021-11-09 09:40:43','lewis','2021-11-09 09:41:12','lewis'),(42,'GET/menus','/menus','GET',1,'2021-11-09 09:40:43','lewis','2021-11-09 09:41:12','lewis'),(43,'GET/menus/:id','/menus/:id','GET',1,'2021-11-09 09:40:43','lewis','2021-11-09 09:41:12','lewis'),(44,'PUT/menus/:id','/menus/:id','PUT',1,'2021-11-09 09:40:43','lewis','2021-11-09 09:41:12','lewis'),(45,'POST/menus','/menus','POST',1,'2021-11-09 09:40:43','lewis','2021-11-09 09:41:12','lewis'),(46,'GET/rolemenus/:id','/rolemenus/:id','GET',1,'2021-11-09 09:40:43','lewis','2021-11-09 09:41:12','lewis'),(47,'POST/rolemenus/:id','/rolemenus/:id','POST',1,'2021-11-09 09:40:43','lewis','2021-11-09 09:41:12','lewis'),(48,'GET/menuapis/:id','/menuapis/:id','GET',1,'2021-11-09 09:40:43','lewis','2021-11-09 09:41:12','lewis'),(49,'POST/menuapis/:id','/menuapis/:id','POST',1,'2021-11-09 09:40:43','lewis','2021-11-09 09:41:12','lewis');
/*!40000 ALTER TABLE `user_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_auths`
--

DROP TABLE IF EXISTS `user_auths`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_auths` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL DEFAULT '0',
  `auth_type` tinyint NOT NULL,
  `identifier` varchar(255) NOT NULL,
  `credential` varchar(255) NOT NULL,
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `login` (`auth_type`,`identifier`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_auths`
--

LOCK TABLES `user_auths` WRITE;
/*!40000 ALTER TABLE `user_auths` DISABLE KEYS */;
INSERT INTO `user_auths` VALUES (4,138,1,'andy@test.com','$2a$14$LehW9Skw/m0ZUOPA3HgHSO48DYvtYiBrKTK5vxTJCfeAazFD.Q5Tm',1,'2021-08-23 00:21:32','system','2021-08-23 08:21:31','system'),(5,6133,1,'kenny@test.com','$2a$14$ymhhesfVRUxj2lwfCAx4gOGlH9s38XQPH7NM6ap0Z./48YVBh54d2',1,'2021-11-16 09:35:33','system','2021-11-16 09:35:33','system'),(6,6134,1,'spade@test.com','$2a$14$3aK77eRxJEsp9F2p71qjIubp2nYL01DMSoca2QW1bKg9/vbJDPk9y',1,'2021-11-16 09:35:42','system','2021-11-16 09:35:42','system');
/*!40000 ALTER TABLE `user_auths` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_menu_apis`
--

DROP TABLE IF EXISTS `user_menu_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_menu_apis` (
  `id` int NOT NULL AUTO_INCREMENT,
  `menu_id` int NOT NULL,
  `api_id` int NOT NULL,
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_menu_apis`
--

LOCK TABLES `user_menu_apis` WRITE;
/*!40000 ALTER TABLE `user_menu_apis` DISABLE KEYS */;
INSERT INTO `user_menu_apis` VALUES (1,1,1,2,'2021-11-08 09:11:11','andy','2021-11-08 09:49:07','andy'),(2,1,2,2,'2021-11-08 09:11:11','andy','2021-11-08 09:49:07','andy'),(3,1,3,2,'2021-11-08 09:11:11','andy','2021-11-08 09:49:07','andy'),(4,1,1,1,'2021-11-08 09:11:11','andy','2021-11-08 09:11:11','andy'),(5,1,2,1,'2021-11-08 09:11:11','andy','2021-11-08 09:11:11','andy'),(6,1,3,1,'2021-11-08 09:11:11','andy','2021-11-08 09:11:11','andy');
/*!40000 ALTER TABLE `user_menu_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_menus`
--

DROP TABLE IF EXISTS `user_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_menus` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL DEFAULT '',
  `action` varchar(64) NOT NULL DEFAULT '' COMMENT '图标',
  `title` varchar(64) NOT NULL DEFAULT '' COMMENT '标题',
  `path` varchar(128) NOT NULL DEFAULT '' COMMENT '路径',
  `component` varchar(255) NOT NULL DEFAULT '' COMMENT '组件',
  `is_hidden` tinyint NOT NULL DEFAULT '0' COMMENT '是否隐藏',
  `parent_id` int NOT NULL DEFAULT '0' COMMENT '父级ID',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_menus`
--

LOCK TABLES `user_menus` WRITE;
/*!40000 ALTER TABLE `user_menus` DISABLE KEYS */;
INSERT INTO `user_menus` VALUES (1,'dashboard','mdi-view-dashboard','Dashboard','/dashboard','dashboard',2,0,1,'2021-11-08 07:50:19','andy','2021-11-15 06:47:01','andy'),(2,'setting','mdi-semantic-web','Settings','','',2,0,1,'2021-11-09 11:23:36','andy','2021-11-15 06:47:01','andy'),(3,'shelf','','Bay','/shelf','shelf',2,2,1,'2021-11-09 11:26:24','andy','2021-11-15 06:44:20','andy'),(4,'goodsLocation','','GoodsLocation','/goodsLocation','shelf/goodsLocation',1,2,1,'2021-11-10 06:49:52','','2021-11-15 06:45:48',''),(5,'barcode','','Barcode','/barcode','shelf/barcode',2,2,1,'2021-11-10 06:50:42','','2021-11-15 06:46:35',''),(6,'items','','Items','/items','shelf/items',2,2,1,'2021-11-15 01:16:12','andy2','2021-11-15 06:47:30','andy2'),(7,'receivings','mdi-note-plus','Receiving','','',2,0,1,'2021-11-15 06:47:30','','2021-11-15 08:32:49','andy2'),(9,'receiving','','Purchase Order','/receiving','receiving',2,7,1,'2021-11-15 06:48:49','','2021-11-15 08:30:05','andy2'),(10,'purchaseScan','','Purchase Scan','/purchaseScan','receiving/purchaseScan',1,7,1,'2021-11-15 06:49:35','','2021-11-15 06:49:35',''),(11,'PickingMenu','mdi-table-arrow-right','Picking','','',2,0,1,'2021-11-15 06:50:05','','2021-11-15 06:56:07',''),(12,'sale','','Sales Order','/sale','picking/sale',2,11,1,'2021-11-15 06:50:44','','2021-11-15 06:50:44',''),(14,'picking','','Picking list','/picking','picking',2,11,1,'2021-11-15 06:51:20','','2021-11-15 06:56:11',''),(15,'Replenishment','mdi-table-large-plus','Replenishment','','',2,0,1,'2021-11-15 06:51:43','','2021-11-15 06:51:50',''),(16,'location','','Location List','/location','replenishment/location',2,15,1,'2021-11-15 06:52:18','','2021-11-15 06:52:18',''),(17,'transactions','','Transfer Record','/transactions','replenishment/transactions',2,15,1,'2021-11-15 06:52:42','','2021-11-16 09:55:35','andy2'),(18,'User Setting','mdi-account-cog','User Setting','','',2,0,1,'2021-11-15 06:53:13','','2021-11-15 06:56:46','andy2'),(19,'user','','User','/user','userSetting/user',2,18,1,'2021-11-15 06:53:39','','2021-11-15 06:53:39',''),(20,'role','','Role','/role','userSetting/role',2,18,1,'2021-11-15 06:54:24','','2021-11-15 06:54:24',''),(21,'menu','','Menu','/menu','userSetting/menu',2,18,1,'2021-11-15 06:54:24','','2021-11-15 06:54:24',''),(22,'adjustmentsRecord','','Adjustments Record','/adjustmentsRecord','replenishment/adjustmentsRecord',2,15,1,'2021-11-17 02:17:59','andy2','2021-11-17 02:23:52','andy2');
/*!40000 ALTER TABLE `user_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_profiles`
--

DROP TABLE IF EXISTS `user_profiles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_profiles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL COMMENT '姓名',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '电子邮箱',
  `role_id` int NOT NULL DEFAULT '0' COMMENT '角色ID',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6135 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_profiles`
--

LOCK TABLES `user_profiles` WRITE;
/*!40000 ALTER TABLE `user_profiles` DISABLE KEYS */;
INSERT INTO `user_profiles` VALUES (138,'andy2','andy@test.com',4,1,'2021-08-23 00:21:32','system','2021-11-17 08:39:59','andy2'),(6133,'kenny','kenny@test.com',1,1,'2021-11-16 09:35:33','system','2021-11-16 09:46:27','andy2'),(6134,'spade','spade@test.com',3,1,'2021-11-16 09:35:42','system','2021-11-17 08:41:34','andy2');
/*!40000 ALTER TABLE `user_profiles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_role_menus`
--

DROP TABLE IF EXISTS `user_role_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_role_menus` (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_id` int NOT NULL,
  `menu_id` int NOT NULL,
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_role_menus`
--

LOCK TABLES `user_role_menus` WRITE;
/*!40000 ALTER TABLE `user_role_menus` DISABLE KEYS */;
INSERT INTO `user_role_menus` VALUES (1,1,1,2,'2021-11-08 09:11:11','andy','2021-11-08 09:38:26','andy'),(2,1,2,2,'2021-11-08 09:11:11','andy','2021-11-08 09:38:26','andy'),(3,1,3,2,'2021-11-08 09:11:11','andy','2021-11-08 09:38:26','andy'),(4,1,1,2,'2021-11-08 09:11:11','andy','2021-11-08 09:38:26','andy'),(5,1,2,2,'2021-11-08 09:11:11','andy','2021-11-08 09:38:26','andy'),(6,1,3,2,'2021-11-08 09:11:11','andy','2021-11-08 09:38:26','andy'),(7,1,1,2,'2021-11-08 09:11:11','andy','2021-11-08 09:38:26','andy'),(8,1,2,2,'2021-11-08 09:11:11','andy','2021-11-08 09:38:26','andy'),(9,1,3,2,'2021-11-08 09:11:11','andy','2021-11-08 09:38:26','andy'),(10,1,1,2,'2021-11-08 09:11:11','andy','2021-11-08 09:38:59','andy'),(11,1,2,2,'2021-11-08 09:11:11','andy','2021-11-08 09:38:59','andy'),(12,1,3,2,'2021-11-08 09:11:11','andy','2021-11-08 09:38:59','andy'),(13,1,1,2,'2021-11-08 09:11:11','andy','2021-11-15 08:39:30','andy2'),(14,1,2,2,'2021-11-08 09:11:11','andy','2021-11-15 08:39:30','andy2'),(15,1,3,2,'2021-11-08 09:11:11','andy','2021-11-15 08:39:30','andy2'),(16,3,1,2,'2021-11-10 06:11:11','andy','2021-11-10 06:17:46','andy'),(17,3,2,2,'2021-11-10 06:11:11','andy','2021-11-10 06:17:46','andy'),(18,3,3,2,'2021-11-10 06:11:11','andy','2021-11-10 06:17:46','andy'),(19,3,1,1,'2021-11-10 06:11:11','andy','2021-11-10 06:11:11','andy'),(20,3,2,1,'2021-11-10 06:11:11','andy','2021-11-10 06:11:11','andy'),(21,3,3,1,'2021-11-10 06:11:11','andy','2021-11-10 06:11:11','andy'),(22,4,1,2,'2021-11-15 07:11:11','andy2','2021-11-17 02:19:11','andy2'),(23,4,2,2,'2021-11-15 07:11:11','andy2','2021-11-17 02:19:11','andy2'),(24,4,3,2,'2021-11-15 07:11:11','andy2','2021-11-17 02:19:11','andy2'),(25,4,4,2,'2021-11-15 07:11:11','andy2','2021-11-17 02:19:11','andy2'),(26,4,5,2,'2021-11-15 07:11:11','andy2','2021-11-17 02:19:11','andy2'),(27,4,6,2,'2021-11-15 07:11:11','andy2','2021-11-17 02:19:11','andy2'),(28,4,7,2,'2021-11-15 07:11:11','andy2','2021-11-17 02:19:11','andy2'),(29,4,9,2,'2021-11-15 07:11:11','andy2','2021-11-17 02:19:11','andy2'),(30,4,10,2,'2021-11-15 07:11:11','andy2','2021-11-17 02:19:11','andy2'),(31,4,11,2,'2021-11-15 07:11:11','andy2','2021-11-17 02:19:11','andy2'),(32,4,12,2,'2021-11-15 07:11:11','andy2','2021-11-17 02:19:11','andy2'),(33,4,14,2,'2021-11-15 07:11:11','andy2','2021-11-17 02:19:11','andy2'),(34,4,15,2,'2021-11-15 07:11:11','andy2','2021-11-17 02:19:11','andy2'),(35,4,16,2,'2021-11-15 07:11:11','andy2','2021-11-17 02:19:11','andy2'),(36,4,17,2,'2021-11-15 07:11:11','andy2','2021-11-17 02:19:11','andy2'),(37,4,18,2,'2021-11-15 07:11:11','andy2','2021-11-17 02:19:11','andy2'),(38,4,19,2,'2021-11-15 07:11:11','andy2','2021-11-17 02:19:11','andy2'),(39,4,20,2,'2021-11-15 07:11:11','andy2','2021-11-17 02:19:11','andy2'),(40,4,21,2,'2021-11-15 07:11:11','andy2','2021-11-17 02:19:11','andy2'),(41,1,1,1,'2021-11-15 08:11:11','andy2','2021-11-15 08:11:11','andy2'),(42,1,2,1,'2021-11-15 08:11:11','andy2','2021-11-15 08:11:11','andy2'),(43,1,3,1,'2021-11-15 08:11:11','andy2','2021-11-15 08:11:11','andy2'),(44,4,1,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2'),(45,4,2,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2'),(46,4,3,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2'),(47,4,4,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2'),(48,4,5,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2'),(49,4,6,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2'),(50,4,7,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2'),(51,4,9,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2'),(52,4,10,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2'),(53,4,11,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2'),(54,4,12,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2'),(55,4,14,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2'),(56,4,15,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2'),(57,4,16,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2'),(58,4,17,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2'),(59,4,18,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2'),(60,4,19,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2'),(61,4,20,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2'),(62,4,21,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2'),(63,4,22,1,'2021-11-17 02:11:11','andy2','2021-11-17 02:11:11','andy2');
/*!40000 ALTER TABLE `user_role_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_roles`
--

DROP TABLE IF EXISTS `user_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_roles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL DEFAULT '',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  `priority` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_roles`
--

LOCK TABLES `user_roles` WRITE;
/*!40000 ALTER TABLE `user_roles` DISABLE KEYS */;
INSERT INTO `user_roles` VALUES (1,'admin1',1,'2021-11-08 02:59:52','andy','2021-11-17 08:44:37','andy',0),(3,'admin2',1,'2021-11-09 08:31:43','andy','2021-11-15 08:37:44','andy2',0),(4,'Super admin',1,'2021-11-15 07:31:20','andy2','2021-11-17 08:44:37','andy2',99);
/*!40000 ALTER TABLE `user_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'wms'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-06-17 16:56:32
