
-- 教师表 唯一id 所教班级id 姓名 密码
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        'userId' bigint(20) not null,
                        `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
--学号 包含年级与班级
DROP TABLE IF EXISTS `stu`;
CREATE TABLE `stu` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `stuId` bigint(20) UNIQUE ,
                        `class_id` bigint(20) NOT NULL,
                        `stuName` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        `gender` tinyint(4) NOT NULL DEFAULT '0', -- 0 male 1 female
                        `birthday` timestamp NOT NULL default CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
--学号 包含年级与班级
DROP TABLE IF EXISTS `grades`;
CREATE TABLE `grades` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `stuId` bigint(20) UNIQUE ,
                        `height` tinyint(4) NOT NULL,
                        `weight` tinyint(4) NOT NULL,
                        `lungs`  INTEGER not null,
                        `50M` float not null, -- second 50m跑
                        `sittingForward`  tinyint(4) not null ,
                        `50M8` float , -- 50*8
                        `sit-upsPerMin` tinyint(4) not null,
                        `skipsPerMin` INTEGER not null ,
                        `standingLongJump` int not null ,
                        `800M` int,
                        `1000M` int,
                        `pullUp` tinyint(4),
                       PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

