CREATE OR REPLACE FUNCTION update_priority_by_id(
    p_id INT,
    p_title VARCHAR,
    p_description TEXT
) RETURNS VOID AS $$
BEGIN
    UPDATE priorities
    SET title = p_title,
        description = p_description,
        created_date = CURRENT_TIMESTAMP
    WHERE id = p_id;
END;
$$ LANGUAGE plpgsql;
