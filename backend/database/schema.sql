CREATE TABLE IF NOT EXISTS user (
    user_id SERIAL PRIMARY KEY,
    user_name TEXT,
    user_type TEXT
);

CREATE TABLE IF NOT EXISTS coupon (
    coupon_id SERIAL PRIMARY KEY,
    coupon_name TEXT,
    coupon_desc TEXT,
    coupon_duration INT,
    coupon_live BOOLEAN,
    coupon_start_date TEXT,
    coupon_end_date TEXT,
    coupon_min_transaction INT,
    coupon_discount INT,
    coupon_max_discount_price INT,
    coupon_category TEXT
);

CREATE TABLE IF NOT EXISTS coupon_user (
    couponuser_id SERIAL PRIMARY KEY,
    user_id INT,
    coupon_id INT,
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
            REFERENCES user(user_id)
            ON DELETE CASCADE
    CONSTRAINT fk_coupon
        FOREIGN KEY (coupon_id)
            REFERENCES coupon(coupon_id)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS coupon_target_type (
    coupontarget_id SERIAL PRIMARY KEY,
    coupontarget_name TEXT,
    coupon_id INT,
    CONSTRAINT fk_coupon
        FOREIGN KEY (coupon_id)
            REFERENCES coupon(coupon_id)
            ON DELETE CASCADE
);