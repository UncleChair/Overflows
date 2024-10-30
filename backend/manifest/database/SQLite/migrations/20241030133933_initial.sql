-- Create "casbin_rule" table
CREATE TABLE `casbin_rule` (`id` integer NULL PRIMARY KEY AUTOINCREMENT, `p_type` varchar NULL, `v0` varchar NULL, `v1` varchar NULL, `v2` varchar NULL, `v3` varchar NULL, `v4` varchar NULL, `v5` varchar NULL, `v6` varchar NULL, `v7` varchar NULL);
-- Create index "casbin_rule_p_type_v0_v1_v2_v3_v4_v5_v6_v7" to table: "casbin_rule"
CREATE UNIQUE INDEX `casbin_rule_p_type_v0_v1_v2_v3_v4_v5_v6_v7` ON `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`);
-- Create "users" table
CREATE TABLE `users` (`uid` text NOT NULL, `avatar_url` text NOT NULL, `username` text NOT NULL, `email` text NOT NULL, `password` text NOT NULL, `created_at` datetime NULL DEFAULT (NULL), `updated_at` datetime NULL DEFAULT (NULL), `deleted_at` datetime NULL DEFAULT (NULL), `last_login` datetime NULL DEFAULT (NULL), `login_attempts` integer NULL DEFAULT 0, `lock` boolean NULL DEFAULT FALSE, `lock_at` datetime NULL DEFAULT (NULL), PRIMARY KEY (`uid`));
-- Create index "users_username" to table: "users"
CREATE UNIQUE INDEX `users_username` ON `users` (`username`);
-- Create index "users_email" to table: "users"
CREATE UNIQUE INDEX `users_email` ON `users` (`email`);
-- Create index "users_username_email" to table: "users"
CREATE UNIQUE INDEX `users_username_email` ON `users` (`username`, `email`);
-- Create index "idx_username" to table: "users"
CREATE INDEX `idx_username` ON `users` (`username`);
-- Create index "idx_email" to table: "users"
CREATE INDEX `idx_email` ON `users` (`email`);
