-- Solution Understanding!
-- เริ่มจาก WITH q1 as () <- อันนี้

-- WITH q1 as (
--     select 
--         *, 

--          row_number() over() as rrown, <- อันนี้ column สมมติไว้ Debug เฉยๆ

--         id - row_number() over() as id_diff
--     from stadium
--     where people >= 100
-- )
-- select * from q1

-- output:                                               and then add 
-- || rrown  ||   | id | visit_date | people | id_diff |
-- ||--------||   | -- | ---------- | ------ | ------- |
-- ||   1    ||   | 2  | 2017-01-02 | 109    | 1       | -> คือ rrown เนี่ยมันเป้นเลขของ sub table q1 ทีนี้ id_diff = id-rrown
-- ||   2    ||   | 3  | 2017-01-03 | 150    | 1       |
-- ||   3    ||   | 5  | 2017-01-05 | 145    | 2       | -> เช่น record นี้ มันคือการเอา 5-3 = 2
-- ||   4    ||   | 6  | 2017-01-06 | 1455   | 2       |
-- ||   5    ||   | 7  | 2017-01-07 | 199    | 2       |
-- ||   6    ||   | 8  | 2017-01-09 | 188    | 2       |


-- sub query
--      step 1
--          select id_diff 
--          from q1 

--          | id_diff |
--          | ------- |
--          | 1       |
--          | 1       |
--          | 2       |
--          | 2       |
--          | 2       |
--          | 2       |

--      step 2
--          select id_diff 
--          from q1 
--          group by id_diff 

--          | id_diff |
--          | ------- |
--          | 1       |
--          | 2       |

--      step 3
--          select id_diff 
--          from q1 
--          group by id_diff 
--          having count(*) >= 3    -> อันนี้ไม่มีไร  มันมาจากที่โจทย์ว่า display the records with three or more rows with consecutive id's

--          | id_diff |
--          | ------- |
--          | 2       |             -> มันเหลือแค่ id_diff 2 เพราะว่าไอเจ้า record ที่ id_diff = 1 มันมีแค่ 2 record 


-- ลองแทนค่าเพื่อให้เห็นภาพ 

-- WHERE id_diff IN (
--     SELECT id_diff 
--     FROM q1 
--     GROUP BY id_diff 
--     HAVING count(*) >= 3
-- )
-- จะกลายเป็น WHERE id_diff IN (1) หรือ WHERE id_diff IN (2) อยู่ที่ว่าจะได้ id_diff กลุ่มไหนมา

-- step สุดท้าย sort ด้วย visit date ตามโจทย์ต้องการ

-- key คือวิธีการหา consecutive id ด้วย id-row_number ครีเอทจัดๆ

WITH q1 AS (
    SELECT 
        *, 
        id - row_number() OVER() AS id_diff
    FROM stadium
    WHERE people >= 100
)

SELECT id, visit_date, people
FROM q1
WHERE id_diff IN (
    SELECT id_diff 
    FROM q1 
    GROUP BY id_diff 
    HAVING count(*) >= 3
)
ORDER BY visit_date