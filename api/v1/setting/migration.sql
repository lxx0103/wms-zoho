/***
 *** Create Table s_shelves 货架表
***/
create table `s_shelves` (
    `id` int not null AUTO_INCREMENT,
    `code` varchar(64) NOT null DEFAULT '' COMMENT '货架编码',
    `level` tinyint not null DEFAULT 1 COMMENT '货架层数',
    `location` varchar(64) not null DEFAULT '' COMMENT '货架位置',
    `enabled` boolean not null DEFAULT 0 COMMENT '状态',
    `created` timestamp NOT null DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `created_by` varchar(64) not null DEFAULT '' COMMENT '创建人',
    `updated` timestamp not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `updated_by` varchar(64) not null DEFAULT '' COMMENT '更新人',
    PRIMARY KEY (`id`),
    UNIQUE KEY `code` (`code`)
);

/***
 *** Create Table s_location 货位表
***/
CREATE TABLE `s_locations` (
	`id` int not null AUTO_INCREMENT,
    `code` varchar(64) not null DEFAULT '' COMMENT '货位编码',
    `level` smallint not null DEFAULT 0 COMMENT '所在层',
    `shelf_id` int not null DEFAULT 0 COMMENT '货架ID',
    `sku` varchar(64) not null DEFAULT '' COMMENT '商品SKU',
    `unit` varchar(64) not null DEFAULT '' COMMENT '商品单位',
    `capacity` int not null DEFAULT 0 COMMENT '可容纳商品数量',
    `enabled` boolean not null DEFAULT 0 COMMENT '状态',
    `created` timestamp NOT null DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `created_by` varchar(64) not null DEFAULT '' COMMENT '创建人',
    `updated` timestamp not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `updated_by` varchar(64) not null DEFAULT '' COMMENT '更新人',
    PRIMARY KEY (`id`),
    UNIQUE KEY `code` (`code`)    
)

/***
 *** Create Table s_barcode 商品条码对应表
***/
CREATE TABLE `s_barcodes` (
	`id` int not null AUTO_INCREMENT,
    `code` varchar(64) not null DEFAULT '' COMMENT 'barcode编码',
    `sku` varchar(64) not null DEFAULT '' COMMENT '商品SKU',
    `unit` varchar(64) not null DEFAULT '' COMMENT '商品单位',
    `quantity` int not null DEFAULT 0 COMMENT '商品数量',
    `enabled` boolean not null DEFAULT 0 COMMENT '状态',
    `created` timestamp NOT null DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `created_by` varchar(64) not null DEFAULT '' COMMENT '创建人',
    `updated` timestamp not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `updated_by` varchar(64) not null DEFAULT '' COMMENT '更新人',
    PRIMARY KEY (`id`),
    UNIQUE KEY `code` (`code`)    
);
