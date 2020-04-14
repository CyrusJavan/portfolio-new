CREATE TABLE author(
    id serial PRIMARY KEY,
    name VARCHAR (128) NOT NULL,
    image_url VARCHAR (2048) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE article(
   id serial PRIMARY KEY,
   author_id INTEGER REFERENCES author(id) NOT NULL,
   slug VARCHAR (128) UNIQUE NOT NULL,
   title VARCHAR (1024) NOT NULL,
   content TEXT NOT NULL,
   snippet VARCHAR (1024),
   created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
   is_active BOOLEAN DEFAULT TRUE NOT NULL
);

CREATE TABLE tag(
    id serial PRIMARY KEY,
    name VARCHAR (128) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE article_tag(
    article_id INTEGER REFERENCES article(id),
    tag_id INTEGER REFERENCES tag(id),
    PRIMARY KEY (article_id, tag_id)
);