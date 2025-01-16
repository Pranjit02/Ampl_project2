-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 16, 2025 at 09:55 PM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `test`
--

-- --------------------------------------------------------

--
-- Table structure for table `tasks`
--

CREATE TABLE `tasks` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `title` longtext DEFAULT NULL,
  `description` longtext DEFAULT NULL,
  `status` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `tasks`
--

INSERT INTO `tasks` (`id`, `title`, `description`, `status`) VALUES
(28, 'Task 1', 'Description 1', 'pending'),
(29, 'Task 2', 'Description 2', 'completed'),
(30, 'New Task', 'New Description', 'pending'),
(31, 'Task 1', 'Description 1', 'pending'),
(32, 'Task 2', 'Description 2', 'completed'),
(33, 'New Task', 'New Description', 'pending'),
(34, 'Task 1', 'Description 1', 'pending'),
(35, 'Task 2', 'Description 2', 'completed'),
(36, 'New Task', 'New Description', 'pending'),
(37, 'Task 1', 'Description 1', 'pending'),
(38, 'Task 2', 'Description 2', 'completed'),
(39, 'New Task', 'New Description', 'pending'),
(40, 'Task 1', 'Description 1', 'pending'),
(41, 'Task 2', 'Description 2', 'completed'),
(42, 'New Task', 'New Description', 'pending'),
(43, 'Task 1', 'Description 1', 'pending'),
(44, 'Task 2', 'Description 2', 'completed'),
(45, 'New Task', 'New Description', 'pending'),
(46, 'Task 1', 'Description 1', 'pending'),
(47, 'Task 2', 'Description 2', 'completed'),
(48, 'New Task', 'New Description', 'pending'),
(49, 'Task 1', 'Description 1', 'pending'),
(50, 'Task 2', 'Description 2', 'completed'),
(51, 'New Task', 'New Description', 'pending');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tasks`
--
ALTER TABLE `tasks`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `tasks`
--
ALTER TABLE `tasks`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=52;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
