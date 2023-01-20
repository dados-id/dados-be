DROP TRIGGER IF EXISTS tr_insert_update_overall_rating ON school_ratings;

DROP FUNCTION IF EXISTS update_overall_rating();

DROP TRIGGER IF EXISTS tr_insert_update_professor_field ON professor_ratings;

DROP FUNCTION IF EXISTS update_professor_field();

DROP TABLE IF EXISTS users CASCADE;

DROP TABLE IF EXISTS user_save_professors CASCADE;

DROP TABLE IF EXISTS professors CASCADE;

DROP TABLE IF EXISTS professor_course_associations CASCADE;

DROP TABLE IF EXISTS professor_ratings CASCADE;

DROP TABLE IF EXISTS professor_rating_tags CASCADE;

DROP TABLE IF EXISTS tags CASCADE;

DROP TABLE IF EXISTS courses CASCADE;

DROP TABLE IF EXISTS faculties CASCADE;

DROP TABLE IF EXISTS school_faculty_associations CASCADE;

DROP TABLE IF EXISTS schools CASCADE;

DROP TABLE IF EXISTS school_ratings CASCADE;

DROP TABLE IF EXISTS report_forms CASCADE;

DROP TABLE IF EXISTS correction_forms CASCADE;

DROP TYPE StatusRequest;
