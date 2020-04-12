DELETE FROM article_tag WHERE article_id = (SELECT id FROM article WHERE slug = 'video-formats-codecs');
DELETE FROM tag WHERE name = 'video' OR name = 'web development';
DELETE FROM article WHERE slug = 'video-formats-codecs';