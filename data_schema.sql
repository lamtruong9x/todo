CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `full_name` varchar(255) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT (now())
);

CREATE TABLE `task` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `content` varchar(255) NOT NULL,
  `done` boolean NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `user_id` int
);

ALTER TABLE `task` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
