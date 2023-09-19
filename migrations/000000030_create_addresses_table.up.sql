CREATE TABLE IF NOT EXISTS `addresses` (
  `addrid` bigint unsigned NOT NULL AUTO_INCREMENT,
  `postalcode` varchar(7) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`addrid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci