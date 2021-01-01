-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: mariadb
-- Generation Time: Jan 01, 2021 at 06:33 PM
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
-- Table structure for table `contractors`
--

CREATE TABLE `contractors` (
  `contractor_id` int(11) NOT NULL,
  `contractor_name` varchar(256) NOT NULL,
  `contractor_name_eng` varchar(256) DEFAULT NULL,
  `contractor_address` varchar(256) DEFAULT NULL,
  `contractor_telephone` varchar(100) DEFAULT NULL,
  `contractor_fax` varchar(100) DEFAULT NULL,
  `contractor_created_at` datetime DEFAULT NULL,
  `contractor_updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `contractors`
--

INSERT INTO `contractors` (`contractor_id`, `contractor_name`, `contractor_name_eng`, `contractor_address`, `contractor_telephone`, `contractor_fax`, `contractor_created_at`, `contractor_updated_at`) VALUES
(18, 'ศราวุธ พิมสาย', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ', '0918861050', '0001', '2020-12-10 06:05:35', '2020-12-10 06:52:24'),
(19, 'ศราวุธ 002', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ', '0918861050', '222222', '2020-12-10 06:06:11', '2020-12-10 06:07:43'),
(20, 'ศราวุธ พิมสาย', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ', '0918861050', '', '2020-12-10 06:07:05', '2020-12-10 06:33:49'),
(21, 'sdfafa', 'fasfafa', 'fdsfsfaf', '21161611', '5454146546', '2020-12-10 06:20:42', '2020-12-10 06:28:21'),
(22, 'ศราวุธ พิมสาย', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ', '0918861050', '1111111', '2020-12-10 06:23:29', '2020-12-10 06:25:13'),
(23, 'Sarawuth Pimsai', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ', '0918861050', '5555555', '2020-12-10 06:23:37', '2020-12-10 06:52:15'),
(24, 'ศราวุธ พิมสาย', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ', '0918861050', '010101', '2020-12-10 06:29:43', '2020-12-10 06:33:46'),
(25, 'KATRAT NA-SONGKHLA', 'KATRAT NA-SONGKHLA', '79/141 หมู่บ้าน พฤกษาวิลล์ 12เอ ถนนสายไหม', '0918861050', '566666', '2020-12-10 06:29:54', '2020-12-10 06:55:25'),
(26, 'ศราวุธ พิมสาย', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ', '0918861050', '46545661', '2020-12-10 06:34:46', '2020-12-10 06:34:46');

-- --------------------------------------------------------

--
-- Table structure for table `employers`
--

CREATE TABLE `employers` (
  `employer_id` int(10) UNSIGNED NOT NULL,
  `employer_fullname` varchar(256) NOT NULL,
  `project_id` int(10) UNSIGNED NOT NULL,
  `employer_created_at` datetime DEFAULT NULL,
  `employer_updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `employers`
--

INSERT INTO `employers` (`employer_id`, `employer_fullname`, `project_id`, `employer_created_at`, `employer_updated_at`) VALUES
(1, 'คุณธวัชชัย วงศ์ศิริวรรณ', 5, '2020-12-12 07:22:47', '2020-12-12 07:22:47'),
(2, 'คุณศราวุธ พิมสาย', 5, '2020-12-12 07:22:47', '2020-12-12 07:22:47'),
(3, 'คุณเกศรัตน์ ณ สงขลา', 5, '2020-12-12 07:22:47', '2020-12-12 07:22:47'),
(4, 'คุณธวัชชัย วงศ์ศิริวรรณ', 6, '2020-12-12 07:23:06', '2020-12-12 07:23:06'),
(5, 'คุณศราวุธ พิมสาย', 6, '2020-12-12 07:23:06', '2020-12-12 07:23:06'),
(6, 'คุณเกศรัตน์ ณ สงขลา', 6, '2020-12-12 07:23:06', '2020-12-12 07:23:06'),
(7, 'คุณธวัชชัย วงศ์ศิริวรรณ', 7, '2020-12-17 09:35:37', '2020-12-17 09:35:37');

-- --------------------------------------------------------

--
-- Table structure for table `jobs`
--

CREATE TABLE `jobs` (
  `job_id` int(11) NOT NULL,
  `job_type_id` int(11) NOT NULL,
  `job_group_id` int(11) NOT NULL,
  `job_description` varchar(256) NOT NULL,
  `job_created_at` datetime DEFAULT NULL,
  `job_updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `jobs`
--

INSERT INTO `jobs` (`job_id`, `job_type_id`, `job_group_id`, `job_description`, `job_created_at`, `job_updated_at`) VALUES
(2, 1, 2, 'ออกแบบ ระบบงานฉนวน ลานจอดรถ', '2020-12-11 03:45:04', '2020-12-11 03:45:04'),
(3, 1, 1, 'งานออกแบบ ระบบหลังคาที่จอดรถ', '2020-12-11 03:46:52', '2020-12-11 03:46:52');

-- --------------------------------------------------------

--
-- Table structure for table `job_groups`
--

CREATE TABLE `job_groups` (
  `job_group_id` int(11) UNSIGNED NOT NULL,
  `job_group_label` varchar(256) NOT NULL,
  `job_type_id` int(11) NOT NULL,
  `job_group_created_at` datetime DEFAULT NULL,
  `job_group_updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `job_groups`
--

INSERT INTO `job_groups` (`job_group_id`, `job_group_label`, `job_type_id`, `job_group_created_at`, `job_group_updated_at`) VALUES
(1, 'งานหลังคา', 1, NULL, NULL),
(2, 'งานฉนวน', 1, NULL, NULL),
(3, 'งานผนัง', 1, NULL, NULL),
(4, 'งานประตู หน้าต่าง', 1, NULL, NULL),
(5, 'งานรั้ว', 1, NULL, NULL),
(6, 'งานที่จอดรถ', 1, NULL, NULL),
(9, 'งานระบบไฟฟ้า', 2, NULL, NULL),
(10, 'งานระบบประปาสุขาภิบาล', 2, NULL, NULL),
(13, 'งานระบบปรับอากาศ', 2, NULL, NULL),
(14, 'งานระบบถ่ายเทอากาศอัตโนมัติ', 2, NULL, NULL),
(15, 'งานระบบสื่อสาร', 2, NULL, NULL),
(16, 'งานถนน', 1, '2020-12-11 06:05:44', '2020-12-11 06:05:44'),
(17, 'งานถนนสองเส้น', 1, '2020-12-11 06:06:06', '2020-12-11 06:06:06');

-- --------------------------------------------------------

--
-- Table structure for table `job_types`
--

CREATE TABLE `job_types` (
  `job_type_id` int(11) NOT NULL,
  `job_type_label` varchar(256) NOT NULL,
  `job_type_created_at` datetime DEFAULT NULL,
  `job_type_updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `job_types`
--

INSERT INTO `job_types` (`job_type_id`, `job_type_label`, `job_type_created_at`, `job_type_updated_at`) VALUES
(1, 'งานโครงสร้าง S&A', NULL, NULL),
(2, 'งานระบบ M&E', NULL, NULL),
(3, 'งานออกแบบ', '2020-12-11 06:03:03', '2020-12-11 06:03:03'),
(4, 'งานติดตั้ง', '2020-12-11 06:03:23', '2020-12-11 06:03:23');

-- --------------------------------------------------------

--
-- Table structure for table `owners`
--

CREATE TABLE `owners` (
  `owner_id` int(11) NOT NULL,
  `owner_name` varchar(256) NOT NULL,
  `owner_name_eng` varchar(256) DEFAULT NULL,
  `owner_director` varchar(256) NOT NULL,
  `owner_created_at` datetime DEFAULT NULL,
  `owner_updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `projects`
--

CREATE TABLE `projects` (
  `project_id` int(11) NOT NULL,
  `project_name` varchar(256) NOT NULL,
  `project_owner_name` varchar(256) DEFAULT NULL,
  `project_owner_name_eng` varchar(256) DEFAULT NULL,
  `project_manager` varchar(256) DEFAULT NULL,
  `project_acronym` varchar(3) NOT NULL,
  `project_expand` varchar(3) NOT NULL,
  `project_created_at` datetime DEFAULT NULL,
  `project_updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `projects`
--

INSERT INTO `projects` (`project_id`, `project_name`, `project_owner_name`, `project_owner_name_eng`, `project_manager`, `project_acronym`, `project_expand`, `project_created_at`, `project_updated_at`) VALUES
(1, 'โครงการ CENTRAL PLAZA MAHACHAI', 'บริษัท เซ็นทรัลพัฒนา จำกัด (มหาชน)', 'Centra Pattana Public Company Limited', 'Trustry Project Management Co., Ltd.', '', '', '2020-12-12 07:06:59', '2020-12-12 07:06:59'),
(2, 'โครงการ CENTRAL PLAZA MAHACHAI', 'บริษัท เซ็นทรัลพัฒนา จำกัด (มหาชน)', 'Centra Pattana Public Company Limited', 'Trustry Project Management Co., Ltd.', '', '', '2020-12-12 07:13:35', '2020-12-12 07:13:35'),
(3, 'โครงการ CENTRAL PLAZA MAHACHAI', 'บริษัท เซ็นทรัลพัฒนา จำกัด (มหาชน)', 'Centra Pattana Public Company Limited', 'Trustry Project Management Co., Ltd.', '', '', '2020-12-12 07:17:44', '2020-12-12 07:17:44'),
(4, 'โครงการ CENTRAL PLAZA MAHACHAI', 'บริษัท เซ็นทรัลพัฒนา จำกัด (มหาชน)', 'Centra Pattana Public Company Limited', 'Trustry Project Management Co., Ltd.', '', '', '2020-12-12 07:18:18', '2020-12-12 07:18:18'),
(5, 'โครงการ CENTRAL PLAZA MAHACHAI', 'บริษัท เซ็นทรัลพัฒนา จำกัด (มหาชน)', 'Centra Pattana Public Company Limited', 'Trustry Project Management Co., Ltd.', '', '', '2020-12-12 07:22:47', '2020-12-12 07:22:47'),
(6, 'โครงการ CENTRAL PLAZA MAHACHAI', 'บริษัท เซ็นทรัลพัฒนา จำกัด (มหาชน)', 'Centra Pattana Public Company Limited', 'Trustry Project Management Co., Ltd.', '', '', '2020-12-12 07:23:06', '2020-12-12 07:23:06'),
(7, 'โครงการ CENTRAL PLAZA MAHACHAI', 'บริษัทเซ็นทรับพัฒนา จำกัด (มหาชน)', 'Central Pattana Public Company Limited', 'Trustry Project Management Co., Ltd.', '', '', '2020-12-17 09:35:37', '2020-12-17 09:35:37');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `user_id` int(10) UNSIGNED NOT NULL,
  `user_employee_id` varchar(20) DEFAULT NULL,
  `user_fullname` varchar(100) NOT NULL,
  `user_address` varchar(256) DEFAULT NULL,
  `user_telephone` varchar(20) DEFAULT NULL,
  `user_created_at` datetime DEFAULT NULL,
  `user_updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`user_id`, `user_employee_id`, `user_fullname`, `user_address`, `user_telephone`, `user_created_at`, `user_updated_at`) VALUES
(173, '001', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ', '0918861050', '2020-12-10 03:44:18', '2020-12-10 03:44:18'),
(174, '002', 'KATRAT NA-SONGKHLA', '79/141 หมู่บ้าน พฤกษาวิลล์ 12เอ ถนนสายไหม', '0918861050', '2020-12-10 03:44:59', '2020-12-10 03:44:59'),
(175, '003', 'Sarawuth Pimsai', '79/52 หมู่บ้านพฤกษาวิลล์ สายไหม ถนนสายไหม แขวงสายไหม เขตสายไหม กรุงเทพ 10220', '0918861050', '2020-12-10 03:45:11', '2020-12-10 05:02:40'),
(176, '004', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ', '0918861050', '2020-12-10 06:09:01', '2020-12-10 07:06:32'),
(177, '005', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ', '0918861050', '2020-12-10 06:11:41', '2020-12-10 07:06:39'),
(178, '006', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ', '0918861050', '2020-12-10 07:00:51', '2020-12-10 07:06:46'),
(179, '007', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ', '0918861050', '2020-12-10 07:01:04', '2020-12-10 07:06:53'),
(180, '008', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ', '0918861050', '2020-12-10 07:02:13', '2020-12-10 07:06:59'),
(181, '009', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ', '0918861050', '2020-12-10 07:04:15', '2020-12-10 07:07:05');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `authors`
--
ALTER TABLE `authors`
  ADD PRIMARY KEY (`author_id`);

--
-- Indexes for table `contractors`
--
ALTER TABLE `contractors`
  ADD PRIMARY KEY (`contractor_id`);

--
-- Indexes for table `employers`
--
ALTER TABLE `employers`
  ADD PRIMARY KEY (`employer_id`);

--
-- Indexes for table `jobs`
--
ALTER TABLE `jobs`
  ADD PRIMARY KEY (`job_id`),
  ADD KEY `job_type_id` (`job_type_id`,`job_group_id`);

--
-- Indexes for table `job_groups`
--
ALTER TABLE `job_groups`
  ADD PRIMARY KEY (`job_group_id`);

--
-- Indexes for table `job_types`
--
ALTER TABLE `job_types`
  ADD PRIMARY KEY (`job_type_id`);

--
-- Indexes for table `owners`
--
ALTER TABLE `owners`
  ADD PRIMARY KEY (`owner_id`);

--
-- Indexes for table `projects`
--
ALTER TABLE `projects`
  ADD PRIMARY KEY (`project_id`);

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
-- AUTO_INCREMENT for table `contractors`
--
ALTER TABLE `contractors`
  MODIFY `contractor_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=27;

--
-- AUTO_INCREMENT for table `employers`
--
ALTER TABLE `employers`
  MODIFY `employer_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `jobs`
--
ALTER TABLE `jobs`
  MODIFY `job_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `job_groups`
--
ALTER TABLE `job_groups`
  MODIFY `job_group_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT for table `job_types`
--
ALTER TABLE `job_types`
  MODIFY `job_type_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `owners`
--
ALTER TABLE `owners`
  MODIFY `owner_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `projects`
--
ALTER TABLE `projects`
  MODIFY `project_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `user_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=182;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
