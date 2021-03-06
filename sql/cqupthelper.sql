-- MySQL dump 10.13  Distrib 5.7.26, for Win64 (x86_64)
--
-- Host: localhost    Database: cqupt
-- ------------------------------------------------------
-- Server version	5.7.26

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `checkin`
--

DROP TABLE IF EXISTS `checkin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `checkin` (
  `mail` varchar(20) COLLATE utf8_unicode_ci NOT NULL,
  `time` datetime DEFAULT '2022-06-08 08:00:00',
  `offset` int(11) DEFAULT '0',
  `szdq` varchar(20) COLLATE utf8_unicode_ci DEFAULT '',
  `xxdz` varchar(50) COLLATE utf8_unicode_ci DEFAULT '',
  `locationBig` varchar(50) COLLATE utf8_unicode_ci DEFAULT '',
  `locationSmall` varchar(50) COLLATE utf8_unicode_ci DEFAULT '',
  `latitude` varchar(10) COLLATE utf8_unicode_ci DEFAULT '',
  `longitude` varchar(10) COLLATE utf8_unicode_ci DEFAULT '',
  `ywjcqzbl` varchar(3) COLLATE utf8_unicode_ci DEFAULT '',
  `ywjchblj` tinyint(1) DEFAULT '0',
  `xjzdywqzbl` tinyint(1) DEFAULT '0',
  `twsfzc` tinyint(1) DEFAULT '1',
  `ywytdzz` tinyint(1) DEFAULT '0',
  `jkmresult` varchar(2) COLLATE utf8_unicode_ci DEFAULT '',
  `beizhu` varchar(50) COLLATE utf8_unicode_ci DEFAULT '',
  `uid` varchar(35) COLLATE utf8_unicode_ci DEFAULT '',
  `enable` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`mail`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `checkinlog`
--

DROP TABLE IF EXISTS `checkinlog`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `checkinlog` (
  `mail` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `data` varchar(10) COLLATE utf8_unicode_ci DEFAULT NULL,
  `log` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL,
  `time` datetime DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `user_id` int(10) NOT NULL AUTO_INCREMENT,
  `mail` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `password` varchar(20) COLLATE utf8_unicode_ci DEFAULT 'password',
  `cqupt_user_id` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `cqupt_password` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `openid` varchar(28) COLLATE utf8_unicode_ci DEFAULT NULL,
  `name` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL,
  `xh` varchar(10) COLLATE utf8_unicode_ci DEFAULT NULL,
  `sex` varchar(1) COLLATE utf8_unicode_ci DEFAULT '鐢?,
  PRIMARY KEY (`user_id`)
) ENGINE=MyISAM AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-06-19 16:26:03
