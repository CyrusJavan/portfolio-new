-- For tracking page views
CREATE TABLE view_track(
   id serial PRIMARY KEY,
   page VARCHAR (4096) NOT NULL,
   view_type VARCHAR (128) NOT NULL,
   ip_address VARCHAR (128) NOT NULL,
   user_agent VARCHAR (4096) NOT NULL,
   created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_view_track_page ON view_track (
    page
);

CREATE INDEX idx_view_track_view_type ON view_track (
    view_type
);