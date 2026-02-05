-- Munculkan data jumlah tipe kartu kredit terbanyak (Query)
SELECT
    credit_card_type,
    COUNT(*) AS total
FROM users
GROUP BY credit_card_type
ORDER BY total DESC
LIMIT 1;