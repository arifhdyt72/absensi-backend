-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jul 31, 2022 at 04:31 PM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 7.4.29

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `absensi`
--

-- --------------------------------------------------------

--
-- Table structure for table `attendance`
--

CREATE TABLE `attendance` (
  `ID` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `time_in` datetime DEFAULT NULL,
  `time_out` datetime DEFAULT NULL,
  `date_attendance` date DEFAULT NULL,
  `file` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `attendance`
--

INSERT INTO `attendance` (`ID`, `user_id`, `time_in`, `time_out`, `date_attendance`, `file`, `created_at`, `updated_at`) VALUES
(30, 1, '2022-07-31 10:19:59', '2022-07-31 20:20:04', '2022-07-25', '1659273600-1-download.jpg', '2022-07-31 20:20:00', '2022-07-31 20:20:04'),
(31, 1, '2022-07-31 10:20:40', '2022-07-31 20:20:47', '2022-07-31', '1659273640-1-download.jpg', '2022-07-31 20:20:40', '2022-07-31 20:20:47'),
(32, 3, '2022-07-31 20:31:36', '2022-07-31 20:31:39', '2022-07-31', '1659274297-3-download.jpg', '2022-07-31 20:31:37', '2022-07-31 20:31:39'),
(33, 7, '2022-07-31 21:13:06', '2022-07-31 21:13:08', '2022-07-31', '1659276787-7-download.jpg', '2022-07-31 21:13:07', '2022-07-31 21:13:09'),
(34, 8, '2022-07-31 21:14:04', '2022-07-31 21:14:07', '2022-07-31', '1659276844-8-download.jpg', '2022-07-31 21:14:04', '2022-07-31 21:14:07');

-- --------------------------------------------------------

--
-- Table structure for table `master_user`
--

CREATE TABLE `master_user` (
  `id` int(11) NOT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `fullname` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `master_user`
--

INSERT INTO `master_user` (`id`, `username`, `password`, `email`, `fullname`, `role`, `created_at`, `updated_at`) VALUES
(1, 'arif.hdyt', '$2a$04$O9E1S4D3S6avlqjT0SoFxOxAyUckKSRMT29pKdgUy3mtUYD8EIEL6', 'arif.hdyt@test.com', 'Arif Hidayat', 'admin', '2022-07-28 16:11:30', '2022-07-28 16:11:30'),
(2, 'arif.hdyt2', '$2a$04$kiUOiknQIiyMPxk6.qwvSOz7aoMwgKXMVheYARxoep..WVszceq0q', 'ariftest2@test.com', 'Arif Test 2', 'admin', '2022-07-29 20:51:18', '2022-07-29 20:51:18'),
(3, 'arif.test3', '$2a$04$mG7D4ymIPzw/ud54ZRmi4.acmb7kmV5/kN7U18wTwQHIO1S7G2TVC', 'ariftest3@test.com', 'Arif Test 3 Update 2', 'employee', '0000-00-00 00:00:00', '2022-07-31 21:12:06'),
(4, 'arif.test4', '$2a$04$cvtlQTr0KNzLneoaqICZOeL4Ruy9CyTGy1V1fcc9G/C7deAqejAwC', 'ariftest4@test.com', 'Arif Test 4 Update', 'admin', '0000-00-00 00:00:00', '2022-07-30 16:02:30'),
(7, 'test.admin', '$2a$04$wZWH24OcDT1TsmTg0APSrerDMqDGilwW2NDkdqOkiq7M9/739eDhi', 'test.admin@test.com', 'test.admin', 'admin', '2022-07-31 21:04:53', '2022-07-31 21:04:53'),
(8, 'test.employee', '$2a$04$zIe/elgSGHb8bYAz9cz5n.RYH17GYdbPcAeZYcZphN9rsMWRKwMqy', 'test.employee@test.com', 'test.employee', 'employee', '2022-07-31 21:05:40', '2022-07-31 21:05:40');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `attendance`
--
ALTER TABLE `attendance`
  ADD PRIMARY KEY (`ID`);

--
-- Indexes for table `master_user`
--
ALTER TABLE `master_user`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `attendance`
--
ALTER TABLE `attendance`
  MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=35;

--
-- AUTO_INCREMENT for table `master_user`
--
ALTER TABLE `master_user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
