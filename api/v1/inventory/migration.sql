/***
 *** Create Table i_items 商品表
***/
CREATE TABLE `i_items` (
	`id` int not null AUTO_INCREMENT,
    `sku` varchar(64) not null DEFAULT '' COMMENT '商品SKU',
    `name` varchar(64) not null DEFAULT '' COMMENT '商品名称',
    `zoho_item_id` varchar(64) not null DEFAULT '' COMMENT 'zoho ItemID',
    `unit` varchar(64) not null DEFAULT '' COMMENT '商品单位',
    `stock` int not null DEFAULT 0 COMMENT '商品库存',
    `stock_available` int not null DEFAULT 0 COMMENT '商品可用库存',
    `enabled` boolean not null DEFAULT 0 COMMENT '状态',
    `created` timestamp NOT null DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `created_by` varchar(64) not null DEFAULT '' COMMENT '创建人',
    `updated` timestamp not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `updated_by` varchar(64) not null DEFAULT '' COMMENT '更新人',
    PRIMARY KEY (`id`),
    UNIQUE KEY `sku` (`sku`),
    UNIQUE KEY `zoho_item_id` (`zoho_item_id`)    
)
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
  UNIQUE KEY `zoho_po_id` (`zoho_po_id`),
  UNIQUE KEY `po_number` (`po_number`)
)
/***
 *** Create Table i_purchase_order_items 采购订单商品表
***/
CREATE TABLE `i_purchase_order_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `po_id` int NOT NULL DEFAULT '0' COMMENT '采购订单ID',
  `item_id` int NOT NULL DEFAULT '0' COMMENT '商品ID',
  `sku` varchar(64) not null DEFAULT '' COMMENT 'SKU',
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
)