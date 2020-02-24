-- MariaDB dump 10.17  Distrib 10.4.11-MariaDB, for osx10.15 (x86_64)
--
-- Host: localhost    Database: rssfeeds
-- ------------------------------------------------------
-- Server version	10.4.10-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `feeds`
--

DROP TABLE IF EXISTS `feeds`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `feeds` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `source_id` int(11) DEFAULT NULL,
  `title` varchar(1000) DEFAULT NULL,
  `description` varchar(2000) DEFAULT NULL,
  `link` varchar(1000) DEFAULT NULL,
  `guid` text DEFAULT NULL,
  `pubDate` datetime DEFAULT NULL,
  `dateCreated` datetime DEFAULT NULL,
  `dateModified` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  FULLTEXT KEY `title` (`title`,`description`,`link`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `feeds`
--

LOCK TABLES `feeds` WRITE;
/*!40000 ALTER TABLE `feeds` DISABLE KEYS */;
/*!40000 ALTER TABLE `feeds` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sources`
--

DROP TABLE IF EXISTS `sources`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sources` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `publisher` varchar(50) DEFAULT NULL,
  `url` varchar(500) DEFAULT NULL,
  `topic` varchar(50) DEFAULT NULL,
  `description` varchar(200) DEFAULT NULL,
  `lastBuildDate` datetime DEFAULT NULL,
  `dateCreated` datetime DEFAULT NULL,
  `dateModified` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sources`
--

LOCK TABLES `sources` WRITE;
/*!40000 ALTER TABLE `sources` DISABLE KEYS */;
INSERT INTO `sources` VALUES (1,'cnn','http://rss.cnn.com/rss/edition_motorsport.rss','CNN.com - RSS Channel - Sport - Motorsport','CNN.com delivers up-to-the-minute news and information on the latest top stories, weather, entertainment, politics and more.','2020-02-21 11:24:00','2020-02-21 13:09:33','2020-02-21 10:09:33'),(2,'bbc','http://feeds.bbci.co.uk/news/video_and_audio/technology/rss.xml','BBC News - Technology','Technology','2020-02-21 11:22:00','2020-02-21 13:10:21','2020-02-21 10:10:21'),(3,'bbc','http://feeds.bbci.co.uk/news/business/rss.xml','BBC News - Business','Business','2020-02-21 11:22:00','2020-02-21 20:12:03','2020-02-21 17:12:03'),(4,'cnn','http://rss.cnn.com/rss/edition_travel.rss','CNN Travel','Travel','2020-02-21 11:22:00','2020-02-21 20:15:52','2020-02-21 17:15:52'),(5,'cnn','http://rss.cnn.com/rss/edition_world.rss','CNN.com - RSS Channel - World','World','2020-02-21 11:22:00','2020-02-21 20:53:29','2020-02-21 17:53:29'),(6,'cnn','http://rss.cnn.com/rss/edition_europe.rss','CNN.com - RSS Channel - Europe','Europe','2020-02-21 11:22:00','2020-02-21 20:56:40','2020-02-21 17:56:40'),(7,'bbc','http://feeds.bbci.co.uk/news/entertainment_and_arts/rss.xml','Entertainment & Arts','Entertainment & Arts','2020-02-21 11:22:00','2020-02-21 21:09:50','2020-02-21 18:09:50'),(8,'telegraph','https://www.telegraph.co.uk/sport/rss.xml','The Telegraph Sport','Sport','2020-02-22 08:22:00','2020-02-22 08:32:36','2020-02-22 05:34:49');
/*!40000 ALTER TABLE `sources` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-02-22 17:53:08

