/* Structure */
DROP TABLE IF EXISTS `records`;
CREATE TABLE `records` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`title` varchar(100),
	`content` varchar(2048),
	`type` varchar(20),
	`icon_url` varchar(256),
	`creator_id` INT NOT NULL,
	`created_at` DATETIME NOT NULL,
	`updated_at` DATETIME,
	`deleted_at` DATETIME,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` INT NOT NULL AUTO_INCREMENT,
	`email` varchar(100) NOT NULL,
	`password` varchar(64) NOT NULL,
	`name` varchar(64) NOT NULL,
	`created_at` DATETIME NOT NULL,
	`updated_at` DATETIME,
	`deleted_at` DATETIME,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

ALTER TABLE `records` ADD CONSTRAINT `records_fk0` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`);

/* Writting data (password: non-secure) */
LOCK TABLES `users` WRITE;
INSERT INTO `users` VALUES (1,'hiyali920@gmail.com','$2a$10$u8/65CcvmYJMV5l8w/8Or..h4pB.3m4KCdemBY788UCRr54GDjIfW','Salam Hiyali', NOW(), null, null);
UNLOCK TABLES;
