/***
 *** Create Table i_items 商品表
***/
CREATE TABLE `i_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `sku` varchar(64) NOT NULL DEFAULT '' COMMENT '商品SKU',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名称',
  `zoho_item_id` varchar(64) NOT NULL DEFAULT '' COMMENT 'zoho ItemID',
  `unit` varchar(64) NOT NULL DEFAULT '' COMMENT '商品单位',
  `stock` int NOT NULL DEFAULT '0' COMMENT '商品库存',
  `stock_available` int NOT NULL DEFAULT '0' COMMENT '商品可用库存',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `zoho_item_id` (`zoho_item_id`),
  KEY `sku` (`sku`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
/***
 *** Create Table i_purchase_orders 采购订单表
***/
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
/***
 *** Create Table i_purchase_order_items 采购订单商品表
***/
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

/***
 *** Create Table i_transactions 采购订单商品表
***/
CREATE TABLE `i_transactions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `po_id` int NOT NULL DEFAULT '0' COMMENT '采购单ID',
  `po_number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '采购单编号',
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci