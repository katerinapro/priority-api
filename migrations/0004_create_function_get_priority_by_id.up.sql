CREATE OR REPLACE FUNCTION lo.get_priority_by_id(p_id INT) RETURNS TABLE (
    id INT,
    title VARCHAR,
    description TEXT,
    created_date TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY 
    SELECT p.id, p.title, p.description, p.created_date
    FROM lo.priorities p
    WHERE p.id = p_id;
END;
$$ LANGUAGE plpgsql;
