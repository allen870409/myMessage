CREATE TABLE `message` (
     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
     `content` varchar(200) NOT NULL,
     `created` datetime NOT NULL,
     PRIMARY KEY (`id`)
   ) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8
