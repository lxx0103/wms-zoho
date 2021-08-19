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