-- Migrations for the project. These queries will be run automatically at the time when the containers are starting
-- The queries are written in such a way that they can be run multiple times without any issues

ALTER USER 'dbuser'@'%' IDENTIFIED WITH mysql_native_password BY 'dbpass';
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'dbpass';

-- User table to store the user details
CREATE TABLE IF NOT EXISTS `users`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `phone` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Admin table to store the admin details
CREATE TABLE IF NOT EXISTS `admins`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `name_password` (`name`, `password`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Loan application table to store the loan application details. Admin id is kept here to know which admin has approved/rejected the loan.
-- The status field is used to know the status of the loan application. Cancelled status is used when the user cancels the loan application.
CREATE TABLE IF NOT EXISTS `loan_applications` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `amount` DECIMAL(10,2) NOT NULL,
    `terms` SMALLINT UNSIGNED NOT NULL,
    `date` DATE NOT NULL,
    `status` TINYINT UNSIGNED NOT NULL COMMENT '1=>SUBMITTED, 2=>APPROVED, 3=>PAID, 4=>REJECTED, 5=>CANCELLED',
    `admin_id` BIGINT UNSIGNED DEFAULT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Loan repayments table to store the loan application's repayment details. The status field is used to know the status of the loan application.
-- Whenever an entry is created in loan_applications table, the respective term wise entries are created in this table.
-- Foreign key association has been intentionally left because we have seen a lot of issues with foreign key constraints in the past in my current organizations.
-- That association is currently being handled in the code through gorm - golang's ORM.
CREATE TABLE IF NOT EXISTS `loan_repayments` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `loan_application_id` BIGINT UNSIGNED NOT NULL,
    `installment_amount` DECIMAL(10,2) NOT NULL,
    `paid_amount` DECIMAL(10,2) NOT NULL,
    `due_date` DATE NOT NULL,
    `paid_date` DATE DEFAULT NULL,
    `status` TINYINT UNSIGNED NOT NULL COMMENT '1=>PENDING, 2=>PAID, 3=>SETTLED, 4=>CANCELLED',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `loan_application_id` (`loan_application_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;