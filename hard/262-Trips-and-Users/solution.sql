-- Execute Order!
--      1. From, Join
--          เอา client_id, driver_id ที่ table trips ไป join กับ table users เพื่อ filter banned user ออกไปก่อน

--      SELECT * FROM Trips AS trip
--      JOIN (
--          SELECT users_id FROM Users 
--          WHERE banned = 'No'
--      ) AS good_clients 
--      ON good_clients.users_id = trip.client_id
--      JOIN (
--          SELECT users_id FROM Users 
--          WHERE banned = 'No'
--      ) AS good_drivers
--      ON good_drivers.users_id = trip.driver_id

--      output 1
--      | id | client_id | driver_id | city_id | status              | request_at | users_id | users_id |
--      | -- | --------- | --------- | ------- | ------------------- | ---------- | -------- | -------- |
--      | 1  | 1         | 10        | 1       | completed           | 2013-10-01 | 1        | 10       |
--      | 3  | 3         | 12        | 6       | completed           | 2013-10-01 | 3        | 12       |
--      | 4  | 4         | 13        | 6       | cancelled_by_client | 2013-10-01 | 4        | 13       |
--      | 5  | 1         | 10        | 1       | completed           | 2013-10-02 | 1        | 10       |
--      | 7  | 3         | 12        | 6       | completed           | 2013-10-02 | 3        | 12       |
--      | 9  | 3         | 10        | 12      | completed           | 2013-10-03 | 3        | 10       |
--      | 10 | 4         | 13        | 12      | cancelled_by_driver | 2013-10-03 | 4        | 13       |


--      2. Where
--          อันนี้จะมีแค่เคสเดียวที่มีวันที่ 14 เกินมา
--      WHERE trip.request_at BETWEEN '2013-10-01' AND '2013-10-03'

--      3. Group by
--          group by request_date เนี่ยมันจะทำให้ result เหลือแค่

--      | id | client_id | driver_id | city_id | status              | request_at | users_id | users_id |
--      | -- | --------- | --------- | ------- | ------------------- | ---------- | -------- | -------- |
--      | 9  | 3         | 10        | 12      | completed           | 2013-10-03 | 3        | 10       |
--      | 5  | 1         | 10        | 1       | completed           | 2013-10-02 | 1        | 10       |
--      | 10 | 4         | 13        | 12      | cancelled_by_driver | 2013-10-03 | 4        | 13       |

--      4. SELECT
--          การคำนวน cancel rate คิดจาก `cancel record`(ที่ถูก group ตามวันที่) หารด้วย `record ทั้งหมด`(ที่ถูก group ตามวันที่)
--      ROUND(
--          SUM(trip.status IN ('cancelled_by_driver', 'cancelled_by_client')) 
--          / 
--          COUNT(*), 
--          2
--      ) AS 'Cancellation Rate'

--      | Day        | Cancellation Rate |
--      | ---------- | ----------------- |
--      | 2013-10-03 | 0.5               |
--      | 2013-10-02 | 0                 |
--      | 2013-10-01 | 0.33              |



SELECT
    trip.request_at AS Day,
    ROUND(
        SUM(trip.status IN ('cancelled_by_driver', 'cancelled_by_client')) 
        / 
        COUNT(*), 
        2
    ) AS 'Cancellation Rate'
FROM Trips AS trip

JOIN (
    SELECT users_id FROM Users 
    WHERE banned = 'No'
) AS good_clients 
ON good_clients.users_id = trip.client_id
JOIN (
    SELECT users_id FROM Users 
    WHERE banned = 'No'
) AS good_drivers
ON good_drivers.users_id = trip.driver_id

WHERE trip.request_at BETWEEN '2013-10-01' AND '2013-10-03'
GROUP BY trip.request_at;
