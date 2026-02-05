-- Munculkan data country mana aja yang spend nya terbanyak (Query) - asumsi top 5 spender country
SELECT
    u.country,
    SUM(t.total_buy) AS total_spend
FROM user_transactions t
JOIN users u
    ON u.id = t.id_user
GROUP BY u.country
ORDER BY total_spend DESC
LIMIT 5;