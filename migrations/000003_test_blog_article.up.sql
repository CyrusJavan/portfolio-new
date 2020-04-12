INSERT INTO author (name, image_url) VALUES ('Cyrus Javan', '/static/img/headshot.jpg');

INSERT INTO article (author_id, slug, title, content, snippet)
VALUES ((SELECT id FROM author WHERE name = 'Cyrus Javan'), 'test-article', 'Test Article: How to create a test article for your website',
'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,');

INSERT INTO tag (name) VALUES ('test');
INSERT INTO tag (name) VALUES ('also-test');

INSERT INTO article_tag (article_id, tag_id)
VALUES ((SELECT id FROM article WHERE slug = 'test-article'),
(SELECT id FROM tag WHERE name = 'test'));

INSERT INTO article_tag (article_id, tag_id)
VALUES ((SELECT id FROM article WHERE slug = 'test-article'),
(SELECT id FROM tag WHERE name = 'also-test'));