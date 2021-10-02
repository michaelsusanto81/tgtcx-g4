INSERT INTO toppers (toppers_name, toppers_type) VALUES('toppers 1', 'Bronze');
INSERT INTO toppers (toppers_name, toppers_type) VALUES('toppers 2', 'Silver');
INSERT INTO toppers (toppers_name, toppers_type) VALUES('toppers 3', 'Silver');
INSERT INTO toppers (toppers_name, toppers_type) VALUES('toppers 4', 'Gold');
INSERT INTO toppers (toppers_name, toppers_type) VALUES('toppers 5', 'Platinum');

INSERT INTO coupon (
    coupon_name,
    coupon_desc,
    coupon_duration,
    coupon_live,
    coupon_start_date,
    coupon_end_date,
    coupon_min_transaction,
    coupon_discount,
    coupon_max_discount_price,
    coupon_category
) VALUES (
    'Cashback',
    'Cashback Tokopedia hingga Rp50.000',
    3,
    TRUE,
    '2021-10-02',
    '2021-10-04',
    50000,
    10,
    50000,
    'Handphone & Tablet'
);

INSERT INTO coupon (
    coupon_name,
    coupon_desc,
    coupon_duration,
    coupon_live,
    coupon_start_date,
    coupon_end_date,
    coupon_min_transaction,
    coupon_discount,
    coupon_max_discount_price,
    coupon_category
) VALUES (
    'Cashback',
    'Cashback Tokopedia hingga Rp30.000',
    5,
    FALSE,
    '2021-10-02',
    '2021-10-04',
    500000,
    5,
    30000,
    'Angsuran Kredit'
);

INSERT INTO coupon (
    coupon_name,
    coupon_desc,
    coupon_duration,
    coupon_live,
    coupon_start_date,
    coupon_end_date,
    coupon_min_transaction,
    coupon_discount,
    coupon_max_discount_price,
    coupon_category
) VALUES (
    'Cashback',
    'Cashback Tokopedia hingga Rp15.000',
    3,
    TRUE,
    '2021-10-02',
    '2021-10-04',
    15000,
    5,
    15000,
    'Paket Data'
);

INSERT INTO coupon (
    coupon_name,
    coupon_desc,
    coupon_duration,
    coupon_live,
    coupon_start_date,
    coupon_end_date,
    coupon_min_transaction,
    coupon_discount,
    coupon_max_discount_price,
    coupon_category
) VALUES (
    'Cashback',
    'Cashback Tokopedia hingga Rp12.000',
    3,
    FALSE,
    '2021-10-02',
    '2021-10-04',
    20000,
    10,
    12000,
    'Listrik PLN'
);

INSERT INTO coupon (
    coupon_name,
    coupon_desc,
    coupon_duration,
    coupon_live,
    coupon_start_date,
    coupon_end_date,
    coupon_min_transaction,
    coupon_discount,
    coupon_max_discount_price,
    coupon_category
) VALUES (
    'Cashback',
    'Cashback Tokopedia hingga Rp20.000',
    3,
    FALSE,
    '2021-10-02',
    '2021-10-04',
    2000000,
    5,
    20000,
    'Pajak'
);

INSERT INTO coupon (
    coupon_name,
    coupon_desc,
    coupon_duration,
    coupon_live,
    coupon_start_date,
    coupon_end_date,
    coupon_min_transaction,
    coupon_discount,
    coupon_max_discount_price,
    coupon_category
) VALUES (
    'Cashback',
    'Cashback Tokopedia hingga Rp50.000',
    3,
    FALSE,
    '2021-10-02',
    '2021-10-04',
    0,
    5,
    50000,
    'Top-up & Tagihan'
);

INSERT INTO coupon_target_type (coupontarget_name, coupon_id) VALUES('Bronze',1);
INSERT INTO coupon_target_type (coupontarget_name, coupon_id) VALUES('Silver',1);
INSERT INTO coupon_target_type (coupontarget_name, coupon_id) VALUES('Silver',2);
INSERT INTO coupon_target_type (coupontarget_name, coupon_id) VALUES('Silver',3);
INSERT INTO coupon_target_type (coupontarget_name, coupon_id) VALUES('Gold',1);
INSERT INTO coupon_target_type (coupontarget_name, coupon_id) VALUES('Gold',2);
INSERT INTO coupon_target_type (coupontarget_name, coupon_id) VALUES('Gold',3);
INSERT INTO coupon_target_type (coupontarget_name, coupon_id) VALUES('Gold',4);
INSERT INTO coupon_target_type (coupontarget_name, coupon_id) VALUES('Gold',5);
INSERT INTO coupon_target_type (coupontarget_name, coupon_id) VALUES('Platinum',1);
INSERT INTO coupon_target_type (coupontarget_name, coupon_id) VALUES('Platinum',2);
INSERT INTO coupon_target_type (coupontarget_name, coupon_id) VALUES('Platinum',3);
INSERT INTO coupon_target_type (coupontarget_name, coupon_id) VALUES('Platinum',4);
INSERT INTO coupon_target_type (coupontarget_name, coupon_id) VALUES('Platinum',5);
INSERT INTO coupon_target_type (coupontarget_name, coupon_id) VALUES('Platinum',6);


INSERT INTO coupon_toppers (toppers_id, coupon_id) VALUES(1,1);
INSERT INTO coupon_toppers (toppers_id, coupon_id) VALUES(2,1);
INSERT INTO coupon_toppers (toppers_id, coupon_id) VALUES(2,2);
INSERT INTO coupon_toppers (toppers_id, coupon_id) VALUES(2,3);
INSERT INTO coupon_toppers (toppers_id, coupon_id) VALUES(3,1);
INSERT INTO coupon_toppers (toppers_id, coupon_id) VALUES(3,2);
INSERT INTO coupon_toppers (toppers_id, coupon_id) VALUES(3,3);
INSERT INTO coupon_toppers (toppers_id, coupon_id) VALUES(4,1);
INSERT INTO coupon_toppers (toppers_id, coupon_id) VALUES(4,2);
INSERT INTO coupon_toppers (toppers_id, coupon_id) VALUES(4,3);
INSERT INTO coupon_toppers (toppers_id, coupon_id) VALUES(4,4);
INSERT INTO coupon_toppers (toppers_id, coupon_id) VALUES(4,5);
INSERT INTO coupon_toppers (toppers_id, coupon_id) VALUES(5,1);
INSERT INTO coupon_toppers (toppers_id, coupon_id) VALUES(5,2);
INSERT INTO coupon_toppers (toppers_id, coupon_id) VALUES(5,3);
INSERT INTO coupon_toppers (toppers_id, coupon_id) VALUES(5,4);
INSERT INTO coupon_toppers (toppers_id, coupon_id) VALUES(5,5);
INSERT INTO coupon_toppers (toppers_id, coupon_id) VALUES(5,6);