ALTER TABLE pullreq_activities ADD COLUMN pullreq_activity_outdated BOOLEAN;
ALTER TABLE pullreq_activities ADD COLUMN pullreq_activity_code_comment_merge_base_sha TEXT;
ALTER TABLE pullreq_activities ADD COLUMN pullreq_activity_code_comment_source_sha TEXT;
ALTER TABLE pullreq_activities ADD COLUMN pullreq_activity_code_comment_path TEXT;
ALTER TABLE pullreq_activities ADD COLUMN pullreq_activity_code_comment_line_new INTEGER;
ALTER TABLE pullreq_activities ADD COLUMN pullreq_activity_code_comment_span_new INTEGER;
ALTER TABLE pullreq_activities ADD COLUMN pullreq_activity_code_comment_line_old INTEGER;
ALTER TABLE pullreq_activities ADD COLUMN pullreq_activity_code_comment_span_old INTEGER;