-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: mariadb
-- Generation Time: Nov 07, 2020 at 08:21 PM
-- Server version: 10.5.6-MariaDB-1:10.5.6+maria~focal
-- PHP Version: 7.4.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `costperfect`
--

-- --------------------------------------------------------

--
-- Table structure for table `authors`
--

CREATE TABLE `authors` (
  `author_id` int(10) UNSIGNED NOT NULL,
  `author_username` varchar(20) NOT NULL,
  `author_password` varchar(256) NOT NULL,
  `author_salt` varchar(256) NOT NULL,
  `user_id` int(10) UNSIGNED NOT NULL,
  `author_deleted` tinyint(1) NOT NULL DEFAULT 0,
  `author_created_at` datetime NOT NULL,
  `author_updated_at` datetime NOT NULL,
  `author_deleted_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `user_id` int(10) UNSIGNED NOT NULL,
  `user_employee_id` varchar(20) DEFAULT NULL,
  `user_name` varchar(100) NOT NULL,
  `user_address` varchar(256) DEFAULT NULL,
  `user_telephone` varchar(20) DEFAULT NULL,
  `user_deleted` tinyint(1) NOT NULL DEFAULT 0,
  `user_created_at` datetime DEFAULT NULL,
  `user_updated_at` datetime DEFAULT NULL,
  `user_deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`user_id`, `user_employee_id`, `user_name`, `user_address`, `user_telephone`, `user_deleted`, `user_created_at`, `user_updated_at`, `user_deleted_at`) VALUES
(1, '001', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ ถนนสายไหม แขวงสายไหม เขตสายไหม จังหวัดกรุงเทพมหานคร 10220', '0918861050', 0, '2020-11-08 00:12:19', '2020-11-08 01:46:26', NULL),
(2, '002', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ ถนนสายไหม แขวงสายไหม เขตสายไหม จังหวัดกรุงเทพมหานคร 10220', '0918861050', 0, '2020-11-08 00:12:21', '2020-11-08 01:46:20', NULL),
(3, '003', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ ถนนสายไหม แขวงสายไหม เขตสายไหม จังหวัดกรุงเทพมหานคร 10220', '0918861050', 0, '2020-11-08 00:12:25', '2020-11-08 01:46:14', NULL),
(4, '004', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ ถนนสายไหม แขวงสายไหม เขตสายไหม จังหวัดกรุงเทพมหานคร 10220', '0918861050', 0, '2020-11-08 00:15:00', '2020-11-08 01:46:08', NULL),
(5, '005', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ ถนนสายไหม แขวงสายไหม เขตสายไหม จังหวัดกรุงเทพมหานคร 10220', '0918861050', 0, '2020-11-08 00:15:16', '2020-11-08 01:46:03', NULL),
(6, '006', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ ถนนสายไหม แขวงสายไหม เขตสายไหม จังหวัดกรุงเทพมหานคร 10220', '0918861050', 0, '2020-11-08 00:30:31', '2020-11-08 01:45:57', NULL),
(7, '007', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ ถนนสายไหม แขวงสายไหม เขตสายไหม จังหวัดกรุงเทพมหานคร 10220', '0918861050', 0, '2020-11-08 00:31:31', '2020-11-08 01:45:51', NULL),
(8, '008', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ ถนนสายไหม แขวงสายไหม เขตสายไหม จังหวัดกรุงเทพมหานคร 10220', '0918861050', 0, '2020-11-08 00:31:54', '2020-11-08 01:45:37', NULL),
(9, '009', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ ถนนสายไหม แขวงสายไหม เขตสายไหม จังหวัดกรุงเทพมหานคร 10220', '0918861050', 0, '2020-11-08 00:34:26', '2020-11-08 01:45:31', NULL),
(10, '010', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ ถนนสายไหม แขวงสายไหม เขตสายไหม จังหวัดกรุงเทพมหานคร 10220', '0918861050', 0, '2020-11-08 00:34:56', '2020-11-08 01:45:25', NULL),
(11, '011', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ ถนนสายไหม แขวงสายไหม เขตสายไหม จังหวัดกรุงเทพมหานคร 10220', '0918861050', 0, '2020-11-08 00:35:18', '2020-11-08 01:45:17', NULL),
(12, '012', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ ถนนสายไหม แขวงสายไหม เขตสายไหม จังหวัดกรุงเทพมหานคร 10220', '0918861050', 0, '2020-11-08 00:37:08', '2020-11-08 01:45:11', NULL),
(13, '013', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ ถนนสายไหม แขวงสายไหม เขตสายไหม จังหวัดกรุงเทพมหานคร 10220', '0809789718', 0, '2020-11-08 00:43:09', '2020-11-08 01:45:04', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `authors`
--
ALTER TABLE `authors`
  ADD PRIMARY KEY (`author_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`user_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `authors`
--
ALTER TABLE `authors`
  MODIFY `author_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `user_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
