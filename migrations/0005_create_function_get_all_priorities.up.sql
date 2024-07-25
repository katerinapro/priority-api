CREATE OR REPLACE FUNCTION lo.get_all_priorities() RETURNS TABLE (
    id INT,
    title VARCHAR,
    description TEXT,
    created_date TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY 
    SELECT p.id, p.title, p.description, p.created_date
    FROM lo.priorities p;
END;
$$ LANGUAGE plpgsql;