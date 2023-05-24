CREATE PROCEDURE delete_threads()
LANGUAGE SQL
AS $$
DELETE FROM threads;
$$;