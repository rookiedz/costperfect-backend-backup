-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: mariadb
-- Generation Time: Jan 29, 2021 at 04:55 PM
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
  `contractor_acronym` varchar(20) NOT NULL,
  `contractor_address` varchar(256) DEFAULT NULL,
  `contractor_telephone` varchar(100) DEFAULT NULL,
  `contractor_fax` varchar(100) DEFAULT NULL,
  `contractor_created_at` datetime DEFAULT NULL,
  `contractor_updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `contractors`
--

INSERT INTO `contractors` (`contractor_id`, `contractor_name`, `contractor_name_eng`, `contractor_acronym`, `contractor_address`, `contractor_telephone`, `contractor_fax`, `contractor_created_at`, `contractor_updated_at`) VALUES
(28, 'บริษัท เอลเม็ควิศวกรรม จำกัด', 'Elmech engineering Co., Ltd.', 'ELM/EE(BD)', '801/236 Moo 8, T.Kukod, Lumlookkar, Pathumthani 12130', '025236106', '025319204', '2021-01-12 16:07:07', '2021-01-14 14:06:11'),
(29, 'บริษัท เวอธ อินโนเวชั่น จำกัด', 'Worth Innovation Co., Ltd.', 'WIN(LED)', '9-261 Moo 4 RamIntra Rd., Anudsawaree, Bangkean, Bangkok 10220, Thailand.', '029525330', '029525332', '2021-01-13 07:07:18', '2021-01-14 14:06:30'),
(31, ' บริษัท เอลเม็ควิศวกรรม จำกัด', ' Elmech engineering Co., Ltd.', ' ELM/EE(SD)', ' 801/236 Moo 8, T.Kukod, Lumlookkar, Pathumthani 12130', '025236106', '025319204', '2021-01-13 07:45:29', '2021-01-13 07:47:16');

-- --------------------------------------------------------

--
-- Table structure for table `contracts`
--

CREATE TABLE `contracts` (
  `contract_id` int(11) NOT NULL,
  `project_id` int(11) NOT NULL,
  `contractor_id` int(11) NOT NULL,
  `employer_id` int(11) NOT NULL,
  `job_id` int(11) NOT NULL,
  `contract_name` varchar(256) NOT NULL,
  `contract_no` varchar(256) NOT NULL,
  `contract_loi_no` varchar(256) NOT NULL,
  `contract_value` float(20,2) NOT NULL,
  `contract_tax` float(3,2) NOT NULL,
  `contract_tax_value` float(20,2) NOT NULL,
  `contract_net_value` float(20,2) NOT NULL,
  `contract_signing_date` date NOT NULL,
  `contract_begin_date` date NOT NULL,
  `contract_end_date` date NOT NULL,
  `contract_delivery_date` date NOT NULL,
  `contract_warranty_days` int(11) NOT NULL,
  `contract_payment_method` varchar(20) NOT NULL,
  `contract_payment_percentage` float(3,2) DEFAULT NULL,
  `contract_payment_amout` float(20,2) NOT NULL,
  `contract_payment_installments` int(11) DEFAULT NULL,
  `contract_advance_payment_method` varchar(20) NOT NULL,
  `contract_advance_payment_percentage` float(3,2) DEFAULT NULL,
  `contract_advance_payment_amout` float(20,2) NOT NULL,
  `contract_advance_payment_installments` int(11) DEFAULT NULL,
  `contract_deduct_method` varchar(20) NOT NULL,
  `contract_deduct_percentage` float(3,2) DEFAULT NULL,
  `contract_warranty_method` varchar(20) NOT NULL,
  `contract_warranty_percentage` float(3,2) DEFAULT NULL,
  `contract_performance_bond_percentage` float(3,2) DEFAULT NULL,
  `contract_retention_money_method` varchar(20) NOT NULL,
  `contract_retention_money_percentage` float(3,2) NOT NULL,
  `contract_note` text DEFAULT NULL,
  `contract_created_at` datetime DEFAULT NULL,
  `contract_updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

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
(7, 'Sarawuth Pimsai', 1, '2021-01-22 08:38:50', '2021-01-22 08:38:50'),
(8, 'Ketrat Na Songkhla', 2, '2021-01-22 08:39:35', '2021-01-22 08:39:35'),
(9, 'Sarawuth Pimsai', 2, '2021-01-22 08:39:35', '2021-01-22 08:39:35'),
(10, 'Kawinthida Pimsai', 2, '2021-01-22 08:39:35', '2021-01-22 08:39:35');

-- --------------------------------------------------------

--
-- Table structure for table `installments`
--

CREATE TABLE `installments` (
  `installment_id` int(11) NOT NULL,
  `installment_no` int(11) NOT NULL,
  `installment_value` int(11) NOT NULL,
  `installment_relations` varchar(20) NOT NULL,
  `contract_id` int(11) NOT NULL,
  `installment_created_at` datetime NOT NULL,
  `installment_updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

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
  `project_acronym` varchar(20) NOT NULL,
  `project_expand` varchar(20) NOT NULL,
  `project_created_at` datetime DEFAULT NULL,
  `project_updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `projects`
--

INSERT INTO `projects` (`project_id`, `project_name`, `project_owner_name`, `project_owner_name_eng`, `project_manager`, `project_acronym`, `project_expand`, `project_created_at`, `project_updated_at`) VALUES
(1, 'โครงการ CENTRAL PLAZA MAHACHAI', 'บริษัท เซ็นทรัลพัฒนา จำกัด (มหาชน)', 'Centra Pattana Public Company Limited', 'Trustry Project Management Co., Ltd.', '', '', '2020-12-12 07:06:59', '2021-01-22 08:38:50'),
(2, 'โครงการ CENTRAL PLAZA MAHACHAI', 'บริษัท เซ็นทรัลพัฒนา จำกัด (มหาชน)', 'Centra Pattana Public Company Limited', 'Trustry Project Management Co., Ltd.', '', '', '2020-12-12 07:13:35', '2021-01-22 08:39:35'),
(3, 'โครงการ CENTRAL PLAZA MAHACHAI', 'บริษัท เซ็นทรัลพัฒนา จำกัด (มหาชน)', 'Centra Pattana Public Company Limited', 'Trustry Project Management Co., Ltd.', '', '', '2020-12-12 07:17:44', '2020-12-12 07:17:44'),
(4, 'โครงการ CENTRAL PLAZA MAHACHAI', 'บริษัท เซ็นทรัลพัฒนา จำกัด (มหาชน)', 'Centra Pattana Public Company Limited', 'Trustry Project Management Co., Ltd.', '', '', '2020-12-12 07:18:18', '2021-01-21 11:28:54'),
(5, 'โครงการ CENTRAL PLAZA MAHACHAI', 'บริษัท เซ็นทรัลพัฒนา จำกัด (มหาชน)', 'Centra Pattana Public Company Limited', 'Trustry Project Management Co., Ltd.', '', '', '2020-12-12 07:22:47', '2021-01-21 11:14:58'),
(6, 'โครงการ CENTRAL PLAZA MAHACHAI', 'บริษัท เซ็นทรัลพัฒนา จำกัด (มหาชน)', 'Centra Pattana Public Company Limited', 'Trustry Project Management Co., Ltd.', '', '', '2020-12-12 07:23:06', '2020-12-12 07:23:06');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `user_id` int(10) UNSIGNED NOT NULL,
  `user_username` varchar(100) NOT NULL,
  `user_password` varchar(20) NOT NULL,
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

INSERT INTO `users` (`user_id`, `user_username`, `user_password`, `user_employee_id`, `user_fullname`, `user_address`, `user_telephone`, `user_created_at`, `user_updated_at`) VALUES
(15, 'rookiedz@gmail.com', 'iLove@0102635', '001', 'Sarawuth Pimsai', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ ถนนสายไหม แขวงสายไหม เขตสายไหม จังหวัดกรุงเทพมหานคร 10220', '0809789718', '2021-01-11 09:37:16', '2021-01-12 16:01:49');

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
-- Indexes for table `contracts`
--
ALTER TABLE `contracts`
  ADD PRIMARY KEY (`contract_id`),
  ADD KEY `idx_project_id` (`project_id`),
  ADD KEY `idx_contractor_id` (`contractor_id`),
  ADD KEY `idx_employer_id` (`employer_id`),
  ADD KEY `idx_contract_name` (`contract_name`),
  ADD KEY `idx_contract_no` (`contract_no`),
  ADD KEY `idx_contract_loi_no` (`contract_loi_no`),
  ADD KEY `idx_contract_signing_date` (`contract_signing_date`),
  ADD KEY `idx_contract_begin_date` (`contract_begin_date`),
  ADD KEY `idx_contract_end_date` (`contract_end_date`),
  ADD KEY `idx_contract_delivery_date` (`contract_delivery_date`),
  ADD KEY `idx_job_id` (`job_id`);

--
-- Indexes for table `employers`
--
ALTER TABLE `employers`
  ADD PRIMARY KEY (`employer_id`);

--
-- Indexes for table `installments`
--
ALTER TABLE `installments`
  ADD PRIMARY KEY (`installment_id`),
  ADD KEY `idx_installment_type_contract_id` (`installment_relations`,`contract_id`) USING BTREE,
  ADD KEY `idx_installment_no` (`installment_no`) USING BTREE;

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
  ADD PRIMARY KEY (`user_id`),
  ADD UNIQUE KEY `user_username` (`user_username`);

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
  MODIFY `contractor_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=32;

--
-- AUTO_INCREMENT for table `contracts`
--
ALTER TABLE `contracts`
  MODIFY `contract_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `employers`
--
ALTER TABLE `employers`
  MODIFY `employer_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `installments`
--
ALTER TABLE `installments`
  MODIFY `installment_id` int(11) NOT NULL AUTO_INCREMENT;

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
  MODIFY `job_type_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `owners`
--
ALTER TABLE `owners`
  MODIFY `owner_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `projects`
--
ALTER TABLE `projects`
  MODIFY `project_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `user_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
