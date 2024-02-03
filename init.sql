CREATE DATABASE `ph` /*!40100 COLLATE 'utf8mb4_general_ci' */;

USE `ph`;

CREATE TABLE `user`
(
    `id`           INT(11) NOT NULL AUTO_INCREMENT,
    `phone_number` VARCHAR(191) NOT NULL COLLATE 'utf8mb4_general_ci',
    `password`     VARCHAR(191) NOT NULL COLLATE 'utf8mb4_general_ci',
    `created_at`   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`, `phone_number`) USING BTREE
) COLLATE = 'utf8mb4_general_ci'
    ENGINE = InnoDB
    AUTO_INCREMENT = 16
;


CREATE TABLE `product`
(
    `id`          INT(11) NOT NULL AUTO_INCREMENT,
    `created_at`  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `category`    VARCHAR(191) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
    `price`       BIGINT(20) NULL DEFAULT NULL,
    `cost`        BIGINT(20) NULL DEFAULT NULL,
    `name`        VARCHAR(191) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
    `description` VARCHAR(191) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
    `barcode`     VARCHAR(191) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
    `expired_at`  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `size`        VARCHAR(191) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
    PRIMARY KEY (`id`) USING BTREE
) COLLATE = 'utf8mb4_general_ci'
    ENGINE = InnoDB
    AUTO_INCREMENT = 42
;

DELIMITER //
DROP FUNCTION IF EXISTS FNC_CHOSUNG;
CREATE DEFINER=CURRENT_USER() FUNCTION `FNC_CHOSUNG`(str VARCHAR(191)) RETURNS VARCHAR(191) CHARSET utf8
BEGIN
     DECLARE rtrnStr VARCHAR(191);
     DECLARE cnt INT;
     DECLARE i INT;
     DECLARE j INT;
     DECLARE tmpStr VARCHAR(191);
     SET str = REPLACE(str,' ','');
     if str is null then
         return '';
end if;
     set cnt = ceil(length(str)/3);
     set i = 1;
     while i <= cnt DO
           set tmpStr = substring(str,i,1);
           set rtrnStr = concat(ifnull(rtrnStr,''),
            case when tmpStr rlike '^(ㄱ|ㄲ)' OR ( tmpStr >= '가' AND tmpStr < '나' ) then 'ㄱ'
                 when tmpStr rlike '^ㄴ' OR ( tmpStr >= '나' AND tmpStr < '다' ) then 'ㄴ'
                 when tmpStr rlike '^(ㄷ|ㄸ)' OR ( tmpStr >= '다' AND tmpStr < '라' ) then 'ㄷ'
                 when tmpStr rlike '^ㄹ' OR ( tmpStr >= '라' AND tmpStr < '마' ) then 'ㄹ'
                 when tmpStr rlike '^ㅁ' OR ( tmpStr >= '마' AND tmpStr < '바' ) then 'ㅁ'
                 when tmpStr rlike '^ㅂ' OR ( tmpStr >= '바' AND tmpStr < '사' ) then 'ㅂ'
                 when tmpStr rlike '^(ㅅ|ㅆ)' OR ( tmpStr >= '사' AND tmpStr < '아' ) then 'ㅅ'
                 when tmpStr rlike '^ㅇ' OR ( tmpStr >= '아' AND tmpStr < '자' ) then 'ㅇ'
                 when tmpStr rlike '^(ㅈ|ㅉ)' OR ( tmpStr >= '자' AND tmpStr < '차' ) then 'ㅈ'
                 when tmpStr rlike '^ㅊ' OR ( tmpStr >= '차' AND tmpStr < '카' ) then 'ㅊ'
                 when tmpStr rlike '^ㅋ' OR ( tmpStr >= '카' AND tmpStr < '타' ) then 'ㅋ'
                 when tmpStr rlike '^ㅌ' OR ( tmpStr >= '타' AND tmpStr < '파' ) then 'ㅌ'
                 when tmpStr rlike '^ㅍ' OR ( tmpStr >= '파' AND tmpStr < '하' ) then 'ㅍ'
                 when tmpStr rlike '^ㅎ' OR ( tmpStr >= '하' ) then 'ㅎ'
                 else ' ' end);
           set i=i+1;
end while;
RETURN rtrnStr;
END //
DELIMITER ;


INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('음료', 4000, 1500, '커피','ㅇㅇ', 121212, NOW(), 'large');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('음료', 4000, 1500, '디저트','ㄱㄹㄹ', 121212, NOW(), 'large');

INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('콜드브루 커피', 5800, 580, '나이트로 바닐라 크림','나바크 작은거', 100101, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('콜드브루 커피', 6300, 630, '나이트로 바닐라 크림','나바크 큰거', 100102, NOW(), 'large');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('콜드브루 커피', 19600, 1960, '시그니처 더 블랙 콜드 브루','시더블콜브', 100103, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('콜드브루 커피', 6000, 600, '돌체 콜드 브루','돌콜브 작은거', 100104, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('콜드브루 커피', 6500, 650, '돌체 콜드 브루','돌콜브 큰거', 100105, NOW(), 'large');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('콜드브루 커피', 5800, 580, '바닐라 크림 콜드 브루','바크콜브 작은거', 100106, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('콜드브루 커피', 6300, 630, '바닐라 크림 콜드 브루','바크콜브 큰거', 100107, NOW(), 'large');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('콜드브루 커피', 4900, 490, '콜드 브루','콜브 작은거', 100108, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('콜드브루 커피', 5400, 540, '콜드 브루','콜브 큰거', 100109, NOW(), 'large');

INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('피지오', 5900, 590, '쿨 라임 피지오','쿨라피 작은거', 100201, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('피지오', 6400, 640, '쿨 라임 피지오','쿨라피 큰거', 100202, NOW(), 'large');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('피지오', 5700, 570, '피치 딸기 피지오','피딸피 작은거', 100203, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('피지오', 6200, 620, '푸치 딸기 피지오','피딸피 큰거', 100204, NOW(), 'large');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('피지오', 5900, 590, '유자 패션 피지오','유패피 작은거', 100205, NOW(), 'small');

INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('리프레셔', 5900, 590, '망고 용과 레모네이드 스타벅스 리프레셔','망용레스리 작은거', 100301, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('리프레셔', 6400, 640, '망고 용과 레모네이드 스타벅스 리프레셔','망용레스리 큰거', 100302, NOW(), 'large');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('리프레셔', 5900, 590, '피플 드링크 위드 망고 용과 스타벅스 리프레셔','피드위망용스리 작은거', 100303, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('리프레셔', 6400, 640, '피플 드링크 위드 망고 용과 스타벅스 리프레셔','피드위망용스리 큰거', 100304, NOW(), 'large');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('리프레셔', 5900, 590, '딸기 아사이 레모네이드 스타벅스 리프레셔','딸아레스리 작은거', 100305, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('리프레셔', 6400, 640, '딸기 아사이 레모네이드 스타벅스 리프레셔','딸아레스리 큰거', 100306, NOW(), 'large');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('리프레셔', 5900, 590, '핑크 드링크 위드 딸기 아사이 스타벅스 리프레셔','핑드위딸아스리 작은거', 100307, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('리프레셔', 6400, 640, '핑크 드링크 위드 딸기 아사이 스타벅스 리프레셔','핑드위딸아스리 큰거', 100308, NOW(), 'large');

INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('블론드', 5900, 590, '블론드 바닐라 더블샷 마키아또','블바더마 큰거', 100401, NOW(), 'large');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('블론드', 5400, 540, '블론드 스타벅스 돌체 라떼','블스돌라 작은거', 100402, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('블론드', 5900, 590, '블론드 스타벅스 돌체 라떼','블스돌라 큰거', 100403, NOW(), 'large');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('블론드', 4500, 450, '블론드 카페라떼','블카 작은거', 100404, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('블론드', 5000, 500, '블론드 카페라떼','블카 큰거', 100405, NOW(), 'large');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('블론드', 4000, 400, '블론드 카페 아메리카노','블바더마 작은거', 100406, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('블론드', 4500, 450, '블론드 카페 아메리카노','블바더마 큰거', 100407, NOW(), 'large');

INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('에스프레소', 6300, 630, '토피 넛 라떼','토넛라 큰거', 100501, NOW(), 'large');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('에스프레소', 6100, 610, '더블 에스프레소 크림라떼','더에크 큰거', 100502, NOW(), 'large');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('에스프레소', 5400, 540, '바닐라 플랫 화이트','바플화 작은거', 100503, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('에스프레소', 5900, 590, '바닐라 플랫 화이트','바플화 큰거', 100504, NOW(), 'large');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('에스프레소', 5400, 540, '스타벅스 돌체 라떼','스돌라 작은거', 100505, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('에스프레소', 5900, 590, '스타벅스 돌체 라떼','스돌라 큰거', 100506, NOW(), 'large');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('에스프레소', 5000, 500, '카페 모카','카모 작은거', 100507, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('에스프레소', 5500, 550, '카페 모카','카모 큰거', 100508, NOW(), 'large');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('에스프레소', 4000, 400, '카페 아메리카노','카아 작은거', 100509, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('에스프레소', 4500, 450, '카페 아메리카노','카아 큰거', 100510, NOW(), 'large');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('에스프레소', 4500, 450, '카페 라떼','카라 작은거', 100511, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('에스프레소', 5000, 500, '카페 라떼','더에크 큰거', 100512, NOW(), 'large');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('에스프레소', 4500, 450, '카푸치노','카 작은거', 100513, NOW(), 'small');
INSERT INTO product(category, price, cost, NAME, DESCRIPTION, barcode, expired_at, size) VALUES('에스프레소', 5000, 500, '카푸치노','카 큰거', 100514, NOW(), 'large');
