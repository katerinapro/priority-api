CREATE OR REPLACE FUNCTION lo.delete_priority_by_id(p_id INT) RETURNS VOID AS $$
BEGIN
    DELETE FROM priorities
    WHERE id = p_id;
END;
$$ LANGUAGE plpgsql;
