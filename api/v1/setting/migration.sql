/***
 *** Create Table s_shelves 货架表
***/
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

/***
 *** Create Table s_location 货位表
***/
CREATE TABLE `s_locations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(64) NOT NULL DEFAULT '' COMMENT '货位编码',
  `level` varchar(64) NOT NULL DEFAULT '' COMMENT '所在层',
  `shelf_id` int NOT NULL DEFAULT '0' COMMENT '货架ID',
  `sku` varchar(64) NOT NULL DEFAULT '' COMMENT '商品SKU',
  `capacity` int NOT NULL COMMENT '可容纳商品',
  `quantity` int NOT NULL DEFAULT '0' COMMENT '已存放数量',
  `available` int NOT NULL DEFAULT '0' COMMENT '剩余空间',
  `unit` varchar(64) NOT NULL DEFAULT '' COMMENT '商品单位',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` varchar(64) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

/***
 *** Create Table s_barcode 商品条码对应表
***/
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
