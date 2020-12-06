-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: mariadb
-- Generation Time: Dec 06, 2020 at 12:04 AM
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
(13, 'ศราวุธ พิมสาย', 'Sarawuth Pimsai', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ', '0918861050', '2222222', '2020-12-06 04:53:58', '2020-12-06 07:04:03'),
(14, 'ศราวุธ พิมสาย', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ', '0957126336', '20215648', '2020-12-06 04:54:45', '2020-12-06 06:47:59'),
(15, 'ศราวุธ พิมสาย', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ', '0918861050', '555555', '2020-12-06 07:04:16', '2020-12-06 07:04:16');

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

-- --------------------------------------------------------

--
-- Table structure for table `job_groups`
--

CREATE TABLE `job_groups` (
  `job_group_id` int(11) NOT NULL,
  `job_group_label` int(11) NOT NULL,
  `job_type_id` int(11) NOT NULL,
  `job_group_created_at` datetime DEFAULT NULL,
  `job_group_updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

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
  `owner_id` int(11) NOT NULL,
  `project_created_at` datetime DEFAULT NULL,
  `project_updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

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
(171, '001', 'ศราวุธ พิมสาย', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ ถนนสายไหม', '0809789718', '2020-12-06 03:03:53', '2020-12-06 06:04:22'),
(172, '002', 'Sarawuth Pimsai', '79/141 หมู่บ้านพฤกษาวิลล์ 12เอ', '0918861050', '2020-12-06 03:04:01', '2020-12-06 06:02:35');

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
-- Indexes for table `jobs`
--
ALTER TABLE `jobs`
  ADD PRIMARY KEY (`job_id`),
  ADD KEY `job_type_id` (`job_type_id`,`job_group_id`);

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
  MODIFY `contractor_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `jobs`
--
ALTER TABLE `jobs`
  MODIFY `job_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `job_types`
--
ALTER TABLE `job_types`
  MODIFY `job_type_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `owners`
--
ALTER TABLE `owners`
  MODIFY `owner_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `projects`
--
ALTER TABLE `projects`
  MODIFY `project_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `user_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=173;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
