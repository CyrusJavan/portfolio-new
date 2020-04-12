DELETE FROM article_tag WHERE article_id = (SELECT id FROM article WHERE slug = 'test-article');
DELETE FROM tag WHERE name = 'test' OR name = 'also-test';
DELETE FROM article WHERE slug = 'test-article';
DELETE FROM author WHERE name = 'Cyrus Javan';