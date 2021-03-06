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

/***
 *** Create Table i_picking_orders 捡货单表
***/
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci


/***
 *** Create Table i_picking_order_items 捡货单商品表
***/
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci


/***
 *** Create Table i_picking_order_details 捡货单详情表
***/
CREATE TABLE `i_picking_order_details` (
  `id` int NOT NULL AUTO_INCREMENT,
  `picking_order_id` int NOT NULL DEFAULT '0',
  `shelf_location` varchar(64) NOT NULL DEFAULT '' COMMENT '货架位置',
  `shelf_code` varchar(64) NOT NULL DEFAULT '' COMMENT '货位编码',
  `location_level` smallint NOT NULL DEFAULT '0' COMMENT '货位所在层',
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci



/***
 *** Create Table i_sales_orders 销售订单表
***/
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci


/***
 *** Create Table i_sales_order_items 销售订单商品表
***/
CREATE TABLE `i_sales_order_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `so_id` int NOT NULL DEFAULT '0' COMMENT '采购订单ID',
  `item_id` int NOT NULL DEFAULT '0' COMMENT '商品ID',
  `sku` varchar(64) NOT NULL DEFAULT '' COMMENT 'SKU',
  `zoho_item_id` varchar(64) NOT NULL DEFAULT '' COMMENT 'zoho商品ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名称',
  `quantity` int NOT NULL DEFAULT '0' COMMENT '商品总数量',
  `quantity_picked` int NOT NULL DEFAULT '0' COMMENT '已捡货数量',
  `quantity_packaged` int NOT NULL DEFAULT '0' COMMENT '已打包数量',
  `quantity_shipped` int NOT NULL DEFAULT '0' COMMENT '已发货数量',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci